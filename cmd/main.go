package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joaokbmartins/app-go-kafka-mysql/infra/kafka"
	repository2 "github.com/joaokbmartins/app-go-kafka-mysql/infra/repository"
	usecase2 "github.com/joaokbmartins/app-go-kafka-mysql/usecase"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/goapp_db")
	if err != nil {
		log.Fatalln(err)
	}

	repository := repository2.CourseMySqlRepository{Db: db}
	usecase := usecase2.CreateCourse{Repository: repository}

	var msgChan = make(chan *ckafka.Message)
	configMapConsumer := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"group.id":          "appgo",
	}
	topics := []string{"courses"}
	consumer := kafka.NewConsumer(configMapConsumer, topics)

	go consumer.Consume(msgChan)

	for msg := range msgChan {
		var input usecase2.CreateCourseInputDTO
		json.Unmarshal(msg.Value, &input)

		output, err := usecase.Execute(input)
		if err != nil {
			fmt.Println("Error: ", err)
		} else {
			fmt.Println("OUTPUT: ", output)
		}
	}

}
