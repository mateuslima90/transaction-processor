package api

import (

	"github.com/labstack/echo/v4"
	"github.com/mateuslima90/transaction-processor/entity"
	"github.com/mateuslima90/transaction-processor/usecase/process_transaction"
	"net/http"
)

type WebServer struct {
	Repository entity.TransactionRepository
}

func NewServer() * WebServer{
	return &WebServer{}
}

func (w WebServer) Serve() {
	e := echo.New()
	e.POST("/transaction", w.process)
	e.Logger.Fatal(e.Start(":8585"))
}

func (w WebServer) process(c echo.Context) error {
	transactionDto := &process_transaction.TransactionDtoInput{}
	c.Bind(transactionDto)
	usecase := process_transaction.NewProcessTransaction(w.Repository)
	output, _ := usecase.Execute(*transactionDto)
	return c.JSON(http.StatusOK, output)
}