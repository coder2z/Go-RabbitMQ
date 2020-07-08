package main

import (
	"fmt"
	"rabbitmq/RabbitMQ"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("test")
	rabbitmq.PublishSimple("Hello myxy99!")
	fmt.Println("发送成功！")
}
