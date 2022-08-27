package server

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"github.com/faruan/codebank/infrastructure/grpc/pb"
	"github.com/faruan/codebank/infrastructure/grpc/service"
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

func (g GRPCServer) Serve() {
	lis, err := net.Listen("tcp","0.0.0.0:50052")
	if err != nil {
		log.Fatalf("couldn't listen tpc port")
	}
	transactionService := service.NewTransactionService()
	transactionService.ProcessTransactionUseCase = g.ProcessTransactionUseCase
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterPaymentServiceServer(grpcServer, transactionService)
	grpcServer.Serve(lis)
}