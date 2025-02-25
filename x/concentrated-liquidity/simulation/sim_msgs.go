package simulation

import (
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	legacysimulationtype "github.com/cosmos/cosmos-sdk/types/simulation"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	appParams "github.com/osmosis-labs/osmosis/v19/app/params"
	osmosimtypes "github.com/osmosis-labs/osmosis/v19/simulation/simtypes"
	sdkrand "github.com/osmosis-labs/osmosis/v19/simulation/simtypes/random"
	clkeeper "github.com/osmosis-labs/osmosis/v19/x/concentrated-liquidity"
	clmodeltypes "github.com/osmosis-labs/osmosis/v19/x/concentrated-liquidity/model"
	cltypes "github.com/osmosis-labs/osmosis/v19/x/concentrated-liquidity/types"
	minttypes "github.com/osmosis-labs/osmosis/v19/x/mint/types"
)

var PoolCreationFee = sdk.NewInt64Coin("stake", 10_000_000)

func RandomMsgCreateConcentratedPool(k clkeeper.Keeper, sim *osmosimtypes.SimCtx, ctx sdk.Context) (*clmodeltypes.MsgCreateConcentratedPool, error) {
	poolCreator, coin0, coin1, tickSpacing, spreadFactor, err := RandomPreparePoolFunc(sim, ctx, k)
	if err != nil {
		return nil, err
	}

	// set permissionless pool creation to true
	defaultParams := cltypes.DefaultParams()
	defaultParams.IsPermissionlessPoolCreationEnabled = true
	k.SetParams(ctx, defaultParams)

	// make sure the denoms are valid authorized quote denoms
	authorizedQuoteDenoms := k.GetAuthorizedQuoteDenoms(ctx)
	authorizedQuoteDenoms = append(defaultParams.AuthorizedQuoteDenoms, coin1.Denom, coin0.Denom)
	k.SetAuthorizedQuoteDenoms(ctx, authorizedQuoteDenoms)

	denomMetaData := banktypes.Metadata{
		DenomUnits: []*banktypes.DenomUnit{{
			Denom:    appParams.BaseCoinUnit,
			Exponent: 6,
		}},
		Base: appParams.BaseCoinUnit,
	}

	sim.BankKeeper().SetDenomMetaData(ctx, denomMetaData)
	err = sim.BankKeeper().MintCoins(ctx, minttypes.ModuleName, sdk.NewCoins(sdk.NewCoin(appParams.BaseCoinUnit, sdk.NewInt(10000000))))
	if err != nil {
		return nil, err
	}

	return &clmodeltypes.MsgCreateConcentratedPool{
		Sender:       poolCreator.String(),
		Denom0:       coin0.Denom,
		Denom1:       coin1.Denom,
		TickSpacing:  tickSpacing,
		SpreadFactor: spreadFactor,
	}, nil
}

func RandMsgCreatePosition(k clkeeper.Keeper, sim *osmosimtypes.SimCtx, ctx sdk.Context) (*cltypes.MsgCreatePosition, error) {
	// get random pool
	clPool, poolDenoms, err := getRandCLPool(k, sim, ctx)
	if err != nil {
		return nil, err
	}

	positionCreator, tokens, lowerTick, upperTick, err := RandomPrepareCreatePositionFunc(sim, ctx, clPool, poolDenoms)
	if err != nil {
		return nil, err
	}

	token0Desired := tokens.AmountOf(clPool.GetToken0())
	token1Desired := tokens.AmountOf(clPool.GetToken1())

	tokenMinAmount0, tokenMinAmount1 := RandomMinAmount(sim, token0Desired, token1Desired)

	accountBalancePoolDenom0 := sim.BankKeeper().GetBalance(ctx, positionCreator, poolDenoms[0])
	accountBalancePoolDenom1 := sim.BankKeeper().GetBalance(ctx, positionCreator, poolDenoms[1])
	if accountBalancePoolDenom0.Amount.LT(tokens[0].Amount) || accountBalancePoolDenom1.Amount.LT(tokens[1].Amount) {
		return nil, fmt.Errorf("insufficient funds when creating a concentrated position")
	}

	return &cltypes.MsgCreatePosition{
		PoolId:          clPool.GetId(),
		Sender:          positionCreator.String(),
		LowerTick:       lowerTick,
		UpperTick:       upperTick,
		TokensProvided:  tokens,
		TokenMinAmount0: tokenMinAmount0,
		TokenMinAmount1: tokenMinAmount1,
	}, nil
}

func RandMsgWithdrawPosition(k clkeeper.Keeper, sim *osmosimtypes.SimCtx, ctx sdk.Context) (*cltypes.MsgWithdrawPosition, error) {
	rand := sim.GetRand()
	// get random pool
	clPool, _, err := getRandCLPool(k, sim, ctx)
	if err != nil {
		return nil, err
	}

	// Utilize the PoolId to PositionId mapping
	positionIds, err := k.GetAllPositionIdsForPoolId(ctx, cltypes.PositionPrefix, clPool.GetId())
	if err != nil {
		return nil, err
	}

	if len(positionIds) < 1 {
		return nil, fmt.Errorf("user does not have any positions")
	}

	randPositionId := positionIds[rand.Intn(len(positionIds))]
	position, err := k.GetPosition(ctx, randPositionId)
	if err != nil {
		return nil, err
	}

	withdrawAmount := sim.RandomDecAmount(position.Liquidity)
	if withdrawAmount.TruncateDec().IsNegative() {
		return nil, fmt.Errorf("Invalid withdraw amount")
	}

	if withdrawAmount.TruncateDec().GT(position.Liquidity) {
		return nil, fmt.Errorf("Insufficient funds from a concentrated position")
	}

	return &cltypes.MsgWithdrawPosition{
		PositionId:      position.PositionId,
		Sender:          position.GetAddress(),
		LiquidityAmount: withdrawAmount,
	}, nil
}

func RandMsgCollectSpreadRewards(k clkeeper.Keeper, sim *osmosimtypes.SimCtx, ctx sdk.Context) (*cltypes.MsgCollectSpreadRewards, error) {
	// get random pool
	clPool, poolDenoms, err := getRandCLPool(k, sim, ctx)
	if err != nil {
		return nil, err
	}

	poolI, err := k.GetPool(ctx, clPool.GetId())
	if err != nil {
		return nil, err
	}

	// positionCreator creates the position with pool denoms
	// get random user address with the pool denoms
	sender, _, senderExists := sim.SelAddrWithDenoms(ctx, poolDenoms)
	if !senderExists {
		return nil, fmt.Errorf("no sender with denoms %s exists", poolDenoms)
	}

	// get random user positions
	positions, err := k.GetUserPositions(ctx, sender.Address, clPool.GetId())
	if err != nil {
		return nil, fmt.Errorf("position does not exist")
	}

	var positionIds []uint64
	for _, position := range positions {
		positionIds = append(positionIds, position.PositionId)
	}

	// SwapOwner performs the swap from token0 to token1
	swapOwner, swapOwnerTokens, senderExists := sim.SelAddrWithDenoms(ctx, poolDenoms)
	if !senderExists {
		return nil, fmt.Errorf("no sender with denoms %s exists", poolDenoms)
	}

	// ensure that we always have 2 tokens
	// Note: tokens returns a random subset of poolDenoms, so  had to add this assertion
	if len(swapOwnerTokens) < 2 {
		return nil, fmt.Errorf("user does not have pool tokens")
	}

	// perform swap until token 1 runs out
	remainingSwapOwnerToken0Amt := swapOwnerTokens[0].Amount
	// we must use cacheCtx when calling a mutative method within a simulator method
	// otherwise, if this errors, it will partially commit and lead to unrealistic state
	cacheCtx, write := ctx.CacheContext()
	for remainingSwapOwnerToken0Amt.GT(sdk.ZeroInt()) {
		randToken0Amt := sim.RandPositiveInt(remainingSwapOwnerToken0Amt)

		if randToken0Amt.LTE(sdk.ZeroInt()) {
			return nil, fmt.Errorf("invalid amount to swap")
		}

		// perform swap from token0 to token1 until either token0 or token1 fund runs out
		_, err = k.SwapExactAmountIn(cacheCtx, swapOwner.Address, poolI, sdk.NewCoin(swapOwnerTokens[0].Denom, randToken0Amt), swapOwnerTokens[1].Denom, sdk.OneInt(), sdk.NewDecWithPrec(1, 2))
		if err != nil {
			return nil, err
		}

		remainingSwapOwnerToken0Amt = remainingSwapOwnerToken0Amt.Sub(randToken0Amt)
	}
	write()

	return &cltypes.MsgCollectSpreadRewards{
		Sender:      sender.Address.String(),
		PositionIds: positionIds,
	}, nil
}

func RandMsgCollectIncentives(k clkeeper.Keeper, sim *osmosimtypes.SimCtx, ctx sdk.Context) (*cltypes.MsgCollectIncentives, error) {
	// get random pool
	clPool, poolDenoms, err := getRandCLPool(k, sim, ctx)
	if err != nil {
		return nil, err
	}

	// positionCreator creates the position with pool denoms
	// get random user address with the pool denoms
	sender, _, senderExists := sim.SelAddrWithDenoms(ctx, poolDenoms)
	if !senderExists {
		return nil, fmt.Errorf("no sender with denoms %s exists", poolDenoms)
	}

	// get random users position
	positions, err := k.GetUserPositions(ctx, sender.Address, clPool.GetId())
	if err != nil {
		return nil, fmt.Errorf("position does not exist")
	}

	if len(positions) < 1 {
		return nil, fmt.Errorf("user does not have any positions")
	}

	randPositionId := positions[sim.RandIntBetween(0, len(positions))]

	// PositionCreator collects the incentives from random position
	return &cltypes.MsgCollectIncentives{
		PositionIds: []uint64{randPositionId.PositionId},
		Sender:      sender.Address.String(),
	}, nil
}

// createPoolRestriction creates specific restriction for the creation of a pool.
func createPoolRestriction(sim *osmosimtypes.SimCtx, ctx sdk.Context) osmosimtypes.SimAccountConstraint {
	return func(acc legacysimulationtype.Account) bool {
		accCoins := sim.BankKeeper().SpendableCoins(ctx, acc.Address)
		hasTwoCoins := len(accCoins) >= 3
		hasPoolCreationFee := accCoins.AmountOf(PoolCreationFee.Denom).GT(PoolCreationFee.Amount)
		return hasTwoCoins && hasPoolCreationFee
	}
}

// getRandCLPool gets a concentrated liquidity pool with its pool denoms.
func getRandCLPool(k clkeeper.Keeper, sim *osmosimtypes.SimCtx, ctx sdk.Context) (cltypes.ConcentratedPoolExtension, []string, error) {
	rand := sim.GetRand()

	// get all pools
	clPools, err := k.GetPools(ctx)
	if err != nil {
		return nil, nil, err
	}

	numPools := len(clPools)
	if numPools == 0 {
		return nil, nil, fmt.Errorf("no pools created")
	}

	randPool := clPools[rand.Intn(numPools)]
	randClPool, ok := randPool.(cltypes.ConcentratedPoolExtension)
	if !ok {
		return nil, nil, fmt.Errorf("pool is not concentrated liquidity pool")
	}

	poolDenoms := []string{randClPool.GetToken0(), randClPool.GetToken1()}

	return randClPool, poolDenoms, err
}

// getRandomTickPositions returns random lowerTick and upperTick divisible by tickSpacing value.
func getRandomTickPositions(sim *osmosimtypes.SimCtx, minTick, maxTick int64, tickSpacing uint64) (int64, int64, error) {
	lowerTick, err := RandomTickDivisibility(sim, minTick, maxTick, tickSpacing)
	if err != nil {
		return 0, 0, err
	}

	if lowerTick == -1 {
		return 0, 0, fmt.Errorf("random lowertick divisible by tickSpacing not found")
	}

	upperTick, err := RandomTickDivisibility(sim, lowerTick, maxTick, tickSpacing)
	if err != nil {
		return 0, 0, err
	}

	if upperTick == -1 {
		return 0, 0, fmt.Errorf("random lowertick divisible by tickSpacing not found")
	}

	if lowerTick == upperTick {
		return 0, 0, fmt.Errorf("lower tick and upper tick cannot be the same")
	}

	return lowerTick, upperTick, nil
}

func RandomMinAmount(sim *osmosimtypes.SimCtx, token0Desired, token1Desired sdk.Int) (sdk.Int, sdk.Int) {
	rand := sim.GetRand()
	percent := sdk.NewDec(int64(sdkrand.RandIntBetween(rand, 0, 100) / 100))
	minAmount0 := sdk.NewDecFromInt(token0Desired).Mul(percent).TruncateInt()
	minAmount1 := sdk.NewDecFromInt(token1Desired).Mul(percent).TruncateInt()
	return minAmount0, minAmount1
}

// RandomTickDivisibility calculates a random number between minTick - maxTick (inclusive) that is divisible by tickSpacing
func RandomTickDivisibility(sim *osmosimtypes.SimCtx, minTick int64, maxTick int64, tickSpacing uint64) (int64, error) {
	rand := sim.GetRand()

	// Generate a random number in the range [minTick, maxTick]
	randomNumber := rand.Int63n(maxTick-minTick+1) + minTick

	// Find the next multiple of x that is greater than or equal to the random number
	nextMultiple := ((randomNumber + int64(tickSpacing) - 1) / int64(tickSpacing)) * int64(tickSpacing)

	// If the next multiple is within the range [a, b], return it
	if nextMultiple >= minTick && nextMultiple <= maxTick {
		return nextMultiple, nil
	}

	// If the next multiple is not within the range [a, b], return -1
	return int64(-1), nil
}

func RandomPreparePoolFunc(sim *osmosimtypes.SimCtx, ctx sdk.Context, k clkeeper.Keeper) (sdk.AccAddress, sdk.Coin, sdk.Coin, uint64, sdk.Dec, error) {
	rand := sim.GetRand()

	authorizedTickSpacing := cltypes.AuthorizedTickSpacing
	authorizedSpreadFactor := cltypes.AuthorizedSpreadFactors

	// find an address with two or more distinct denoms in their wallet
	sender, senderExists := sim.RandomSimAccountWithConstraint(createPoolRestriction(sim, ctx))
	if !senderExists {
		return nil, sdk.Coin{}, sdk.Coin{}, 0, sdk.Dec{}, fmt.Errorf("no sender with two different denoms & pool creation fee exists")
	}

	// get random 3 coins, use 2 to create pool and 1 for fees (stake denom).
	poolCoins, ok := sim.GetRandSubsetOfKDenoms(ctx, sender, 3)
	if !ok {
		return nil, sdk.Coin{}, sdk.Coin{}, 0, sdk.Dec{}, fmt.Errorf("provided sender with requested number of denoms does not exist")
	}

	// check if the sender has sufficient amount for fees
	if poolCoins.Add(PoolCreationFee).IsAnyGT(sim.BankKeeper().SpendableCoins(ctx, sender.Address)) {
		return nil, sdk.Coin{}, sdk.Coin{}, 0, sdk.Dec{}, errors.New("chose an account / creation amount that didn't pass fee limit")
	}

	if poolCoins[0].Denom == sdk.DefaultBondDenom || poolCoins[1].Denom == sdk.DefaultBondDenom {
		return nil, sdk.Coin{}, sdk.Coin{}, 0, sdk.Dec{}, fmt.Errorf("poolCoins contains denom stake which contains invalid metadata")
	}

	coin0 := poolCoins[0]
	coin1 := poolCoins[1]
	tickSpacing := authorizedTickSpacing[rand.Intn(len(authorizedTickSpacing))]
	spreadFactor := authorizedSpreadFactor[rand.Intn(len(authorizedSpreadFactor))]

	return sender.Address, coin0, coin1, tickSpacing, spreadFactor, nil
}

func RandomPrepareCreatePositionFunc(sim *osmosimtypes.SimCtx, ctx sdk.Context, clPool cltypes.ConcentratedPoolExtension, poolDenoms []string) (sdk.AccAddress, sdk.Coins, int64, int64, error) {
	// make sure that the position creator has the poolTokens
	positionCreator, tokens, senderExists := sim.SelAddrWithDenoms(ctx, poolDenoms)
	if !senderExists {
		return nil, sdk.Coins{}, 0, 0, fmt.Errorf("no sender with denoms %s exists", poolDenoms)
	}

	// ensure that we always have 2 tokens
	// Note: tokens returns a random subset of poolDenoms, so  had to add this assertion
	if len(tokens) < 2 {
		return nil, sdk.Coins{}, 0, 0, fmt.Errorf("user does not have pool tokens")
	}

	//  Retrieve minTick and maxTick from kprecision factor
	minTick, maxTick := cltypes.MinInitializedTick, cltypes.MaxTick

	// Randomize lowerTick and upperTick from max values to create position
	lowerTick, upperTick, err := getRandomTickPositions(sim, minTick, maxTick, clPool.GetTickSpacing())
	if err != nil {
		return nil, sdk.Coins{}, 0, 0, err
	}

	return positionCreator.Address, tokens, lowerTick, upperTick, nil
}
