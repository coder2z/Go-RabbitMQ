package main

import "rabbitmq/RabbitMQ"

func main()  {
	imoocOne:=RabbitMQ.NewRabbitMQTopic("myxy99Topic","#")
	imoocOne.RecieveTopic()
}
