package main

import "rabbitmq/RabbitMQ"

func main()  {
	mqOne:=RabbitMQ.NewRabbitMQRouting("myxy99","myxy99_two")
	mqOne.RecieveRouting()
}
