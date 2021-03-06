package invoker

import (
	"context"

	"github.com/kukgini/hlf-sdk-go-2/api"
	"github.com/hyperledger/fabric/msp"
	"github.com/hyperledger/fabric-protos-go/peer"
	"github.com/pkg/errors"
)

type invoker struct {
	core api.Core
}

func (i *invoker) Invoke(ctx context.Context, from msp.SigningIdentity, channel string, chaincode string, fn string, args [][]byte, transArgs api.TransArgs) (*peer.Response, api.ChaincodeTx, error) {
	return i.core.Channel(channel).Chaincode(chaincode).Invoke(fn).WithIdentity(from).ArgBytes(args).Transient(transArgs).Do(ctx)
}

func (i *invoker) Query(ctx context.Context, from msp.SigningIdentity, channel string, chaincode string, fn string, args [][]byte, transArgs api.TransArgs) (*peer.Response, error) {
	argSs := make([]string, 0)
	for _, arg := range args {
		argSs = append(argSs, string(arg))
	}

	if resp, err := i.core.Channel(channel).Chaincode(chaincode).Query(fn, argSs...).WithIdentity(from).Transient(transArgs).AsProposalResponse(ctx); err != nil {
		return nil, errors.Wrap(err, `failed to query chaincode`)
	} else {
		return resp.Response, nil
	}
}

func (i *invoker) Subscribe(ctx context.Context, from msp.SigningIdentity, channel, chaincode string) (api.EventCCSubscription, error) {
	return i.core.Channel(channel).Chaincode(chaincode).Subscribe(ctx)
}

func New(core api.Core) api.Invoker {
	return &invoker{core: core}
}
