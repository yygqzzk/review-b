package data

import (
	"context"

	pb "github.com/yygqzzk/review-b/api/review/v1"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/yygqzzk/review-b/internal/conf"

	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/hashicorp/consul/api"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewDiscovery, NewReviewServiceClient, NewData, NewBusinessRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	reviewClient pb.ReviewClient
	log          *log.Helper
}

// NewData .
func NewData(c *conf.Data, reviewClient pb.ReviewClient, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		reviewClient: reviewClient,
		log:          log.NewHelper(logger),
	}, cleanup, nil
}

func NewDiscovery(reg *conf.Registry) registry.Discovery {
	config := api.DefaultConfig()
	config.Address = reg.Consul.Address
	config.Scheme = reg.Consul.Scheme
	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}
	return consul.New(client)
}

func NewReviewServiceClient(reg *conf.Registry, discovery registry.Discovery) pb.ReviewClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///"+reg.Service.Name),
		grpc.WithDiscovery(discovery),
		grpc.WithMiddleware(
			recovery.Recovery(),
			validate.Validator(),
		),
	)
	if err != nil {
		panic(err)
	}
	return pb.NewReviewClient(conn)
}
