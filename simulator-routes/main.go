package main

import (
	"github.com/thainapires/app_delivery/infra/kafka"
	// route2 "github.com/thainapires/app_delivery/application/route"
	"github.com/joho/godotenv"
	"log"
)

func init(){
	err := godotenv.Load()
	if err != nil{
		log.Fatal("error loading .env file")
	}
}

func main(){

	producer := kafka.NewKafkaProducer()
	kafka.Publish("ola", "readtest", producer)

	for {
		_ = 1
	}

	/* route := route2.Route{
		ID: "1",
		ClientID: "1",
	}

	route.LoadPositions()
	stringjson, _ := route.ExportJsonPositions()
	fmt.Println(stringjson[0])
	 */
}