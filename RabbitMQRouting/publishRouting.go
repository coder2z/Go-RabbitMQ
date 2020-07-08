package main

import (
	"rabbitmq/RabbitMQ"
	"strconv"
	"time"
	"fmt"
)

func main()  {
	mqOne:=RabbitMQ.NewRabbitMQRouting("myxy99","myxy99_one")
	mqTwo:=RabbitMQ.NewRabbitMQRouting("myxy99","myxy99_two")
	for i := 0; i <= 10; i++ {
		mqOne.PublishRouting("Hello myxy99 one!" + strconv.Itoa(i))
		mqTwo.PublishRouting("Hello myxy99 Two!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
	
}
