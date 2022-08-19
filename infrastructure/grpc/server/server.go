package server

import (
	"log"
	"net"
	"github.com/faruan/codebank/usecase"
)

type GRPCServer struct {
	ProcessTransactionUseCase usecase.UseCaseTransaction
}

func NewGRPCServer() GRPCServer {
	return GRPCServer{}
}

func (s GRPCServer) Serve() {
	lis, err := net.Listen("tcp","0.0.0.0:50052")
	if err != nil {
		log.Fatalf("couldn't listen tpc port")
	}
	// transactionService := service.NewTransactionService()
	
}