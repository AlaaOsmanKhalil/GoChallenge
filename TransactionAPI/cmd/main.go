package main

import (
	"TransactionAPI/config"
	"TransactionAPI/internal/adapters/api"
	"TransactionAPI/internal/adapters/db"
	"TransactionAPI/internal/adapters/stream"
	"TransactionAPI/internal/repositories/transaction"
	"TransactionAPI/internal/services/transactionsvc"
	"TransactionAPI/resources"
	"github.com/go-playground/validator/v10"
)

func main() {
	log, closer := resources.NewLogger()
	defer closer()

	configs := config.LoadConfig(log)

	httpServer := api.NewHttpServer(log, configs.Server) //

	validate := validator.New()

	conn := db.NewDatabaseConnection(log, configs.Database)

	transactionRepo := transaction.NewDatabaseRepository(log, conn)

	transactionSvc := transactionsvc.NewDefaultService(log, transactionRepo)

	api.NewTransactionController(httpServer, validate, transactionSvc)

	//routes.TransactionRoute(router)

	stream.InitializeKafka(configs)
	go stream.Consume(log)

	httpServer.Start()
}
