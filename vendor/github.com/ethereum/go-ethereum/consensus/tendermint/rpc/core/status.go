package core

import (
	ctypes "github.com/ethereum/go-ethereum/consensus/tendermint/rpc/core/types"
	//"github.com/ethereum/go-ethereum/consensus/tendermint/types"
)

func Status(context *RPCDataContext) (*ctypes.ResultStatus, error) {
	/*
	latestHeight := context.blockStore.Height()
	var (
		latestBlockMeta *types.BlockMeta
		latestBlockHash []byte
		latestAppHash   []byte
		latestBlockTime int64
	)
	if latestHeight != 0 {
		latestBlockMeta = context.blockStore.LoadBlockMeta(latestHeight)
		latestBlockHash = latestBlockMeta.BlockID.Hash
		latestAppHash = latestBlockMeta.Header.AppHash
		latestBlockTime = latestBlockMeta.Header.Time.UnixNano()
	}

	return &ctypes.ResultStatus{
		NodeInfo:          context.p2pSwitch.NodeInfo(),
		PubKey:            context.pubKey,
		LatestBlockHash:   latestBlockHash,
		LatestAppHash:     latestAppHash,
		LatestBlockHeight: latestHeight,
		LatestBlockTime:   latestBlockTime}, nil
	*/

	return nil, nil
}
