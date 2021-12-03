package main

import (
	"database/sql"
	"github.com/mateuslima90/transaction-processor/adapter/api"
	"github.com/mateuslima90/transaction-processor/adapter/grpc/pb"
	"github.com/mateuslima90/transaction-processor/adapter/grpc/service"
	"github.com/mateuslima90/transaction-processor/adapter/repository"
	"github.com/mateuslima90/transaction-processor/entity"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {

	db, err := sql.Open("sqlite3", "test.db")

	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewTransactionRepositoryDb(db)
	webserver := api.NewServer()
	webserver.Repository = repo
	go webserver.Serve()
	startGRPCServer(repo)
}

func startGRPCServer(repo entity.TransactionRepository) {
	listen, err := net.Listen("tcp", "localhost:50052")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	service := service.NewProcessService()
	service.Repository = repo
	pb.RegisterTransactionServiceServer(grpcServer, service)
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatal(err)
	}
}