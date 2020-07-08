package main

import "rabbitmq/RabbitMQ"

func main()  {
	imoocOne:=RabbitMQ.NewRabbitMQTopic("myxy99Topic","myxy99.*.two")
	imoocOne.RecieveTopic()
}
