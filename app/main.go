package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"app/models"
	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/olahol/melody.v1"
)

type ReceiveData struct {
	UserId string `json:"userId"`
	Text   string `json:"text"`
}

func init() {
	connectionString := os.Getenv("MONGODB_CONNECTION_STRING")
	if len(connectionString) == 0 {
		connectionString = "mongodb://root:password@mongodb:27017"
	}

	err := mgm.SetDefaultConfig(nil, "app", options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.Println("Websocket App start.")

	router := gin.Default()
	m := melody.New()

	rg := router.Group("/")
	rg.GET("/", func(ctx *gin.Context) {
		http.ServeFile(ctx.Writer, ctx.Request, "index.html")
	})

	rg.GET("/ws", func(ctx *gin.Context) {
		m.HandleRequest(ctx.Writer, ctx.Request)
	})

	m.HandleMessage(func(s *melody.Session, data []byte) {
		var receiveData ReceiveData
		err := json.Unmarshal(data, &receiveData)
		if err != nil {
			log.Fatal(err)
		}
		doc := models.NewMessage(receiveData.UserId, 1, receiveData.Text)
		err2 := mgm.Coll(doc).Create(doc)
		if err2 != nil {
			log.Fatal(err2)
		} else {
			m.Broadcast(data)
		}
	})

	m.HandleConnect(func(s *melody.Session) {
		log.Printf("websocket connection open. [session: %#v]\n", s)
	})

	m.HandleDisconnect(func(s *melody.Session) {
		log.Printf("websocket connection close. [session: %#v]\n", s)
	})

	// Listen and server on 0.0.0.0:8989
	router.Run(":8989")

	fmt.Println("Websocket App End.")
}
