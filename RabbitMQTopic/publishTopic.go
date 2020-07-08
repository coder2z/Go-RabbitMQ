package main

import (
	"rabbitmq/RabbitMQ"
	"strconv"
	"time"
	"fmt"
)

func main()  {
	mqOne:=RabbitMQ.NewRabbitMQTopic("myxy99Topic","myxy99.topic.one")
	mqTwo:=RabbitMQ.NewRabbitMQTopic("myxy99Topic","myxy99.topic.two")
	for i := 0; i <= 10; i++ {
		mqOne.PublishTopic("Hello myxy99 topic one!" + strconv.Itoa(i))
		mqTwo.PublishTopic("Hello myxy99 topic Two!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
	
}
