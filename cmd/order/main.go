package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/carlosgenuino/fieldservice/internal/infra/database"
	"github.com/carlosgenuino/fieldservice/internal/usecase"
	"github.com/carlosgenuino/fieldservice/pkg/rabbitmq"
	_ "github.com/mattn/go-sqlite3"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	orderRepository := database.NewOrderRepository(db)
	uc := usecase.NewUseCaseCalculateFinalPrice(orderRepository)
	ch, err := rabbitmq.OpenChannel()

	if err != nil {
		panic(err)
	}

	defer ch.Close()
	msgRabbitMQChannel := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, msgRabbitMQChannel)
	rabbitmqWorker(msgRabbitMQChannel, uc)
	// input := &usecase.OrderInput{
	// 	Id:    "123",
	// 	Price: 10.0,
	// 	Tax:   1.0,
	// }
	// output, err := uc.Execute(*input)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(&output)
}

func rabbitmqWorker(msgChan chan amqp.Delivery, uc *usecase.CalculateFinalPrice) {
	fmt.Println("Starting rabbitmq")

	for msg := range msgChan {
		var input usecase.OrderInput

		err := json.Unmarshal(msg.Body, &input)

		if err != nil {
			panic(err)
		}

		output, err := uc.Execute(input)

		if err != nil {
			panic(err)
		}

		msg.Ack(false)
		fmt.Println("Mensagem Processada e Salva no Banco:", output)
	}
}
