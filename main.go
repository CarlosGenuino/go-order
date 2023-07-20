package main

import (
	"database/sql"
	"fmt"

	"github.com/carlosgenuino/fieldservice/internal/infra/database"
	"github.com/carlosgenuino/fieldservice/internal/usecase"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")

	if err != nil {
		panic(err)
	}

	orderRepository := database.NewOrderRepository(db)
	uc := usecase.NewUseCaseCalculateFinalPrice(orderRepository)
	input := &usecase.OrderInput{
		Id:    "123",
		Price: 10.0,
		Tax:   1.0,
	}
	output, err := uc.Execute(*input)

	if err != nil {
		panic(err)
	}

	fmt.Println(&output)

}
