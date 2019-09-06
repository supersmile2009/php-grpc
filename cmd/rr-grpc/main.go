package main

import (
	"github.com/spiral/php-grpc"
	"github.com/spiral/roadrunner/service/limit"
	"github.com/spiral/roadrunner/service/metrics"
	"github.com/spiral/roadrunner/service/rpc"

	rr "github.com/spiral/roadrunner/cmd/rr/cmd"

	// grpc specific commands
	_ "github.com/spiral/php-grpc/cmd/rr-grpc/grpc"
)

func main() {
	rr.Container.Register(rpc.ID, &rpc.Service{})
	rr.Container.Register(grpc.ID, &grpc.Service{})

	rr.Container.Register(metrics.ID, &metrics.Service{})
	rr.Container.Register(limit.ID, &limit.Service{})

	rr.Execute()
}
