package peer

import (
	fd "github.com/hyperledger/fabric/protos/discovery"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/s7techlab/hlf-sdk-go/api"
	"github.com/s7techlab/hlf-sdk-go/api/config"
	"github.com/s7techlab/hlf-sdk-go/util"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type discovery struct {
	pool api.PeerPool
	cli  fd.DiscoveryClient
	conn *grpc.ClientConn
}

type discoveryOpts struct {
	MspId    string                  `yaml:"msp_id"`
	CertPath string                  `yaml:"cert_path"`
	KeyPath  string                  `yaml:"key_path"`
	Peer     config.ConnectionConfig `yaml:"peer"`
}

func (d *discovery) Initialize(options config.DiscoveryConfigOpts, pool api.PeerPool, log *zap.Logger) (api.DiscoveryProvider, error) {
	var opts discoveryOpts
	var di discovery
	if err := mapstructure.Decode(options, &opts); err != nil {
		return nil, errors.Wrap(err, `failed to decode params`)
	}

	if grpcOpts, err := util.NewGRPCOptionsFromConfig(opts.Peer, log); err != nil {
		return nil, errors.Wrap(err, `failed `)
	} else {
		if di.conn, err = grpc.Dial(opts.Peer.Host, grpcOpts...); err != nil {
			return nil, errors.Wrap(err, `failed to connect peer`)
		}
	}

	di.cli = fd.NewDiscoveryClient(di.conn)

	return &di, nil
}

func (d *discovery) Channels() ([]api.DiscoveryChannel, error) {
	panic("implement me")
}

func (d *discovery) Chaincode(channelName string, ccName string) (*api.DiscoveryChaincode, error) {
	panic("implement me")
}

func (d *discovery) Chaincodes(channelName string) ([]api.DiscoveryChaincode, error) {
	panic("implement me")
}
