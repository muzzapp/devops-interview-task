package handler

import (
	"context"
	"math/rand/v2"
	"time"

	"github.com/muzzapp/devops-interview-task/internal/muzz"
)

type Server struct {
	muzz.UnimplementedServiceServer
}

func (s Server) Echo(_ context.Context, req *muzz.EchoRequest) (*muzz.EchoResponse, error) {
	time.Sleep(time.Duration(rand.IntN(5000)) * time.Millisecond) // Added to simulate some delay for more interesting metrics
	return &muzz.EchoResponse{Message: req.GetMessage()}, nil
}
