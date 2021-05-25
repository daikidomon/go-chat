package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"unsafe"

	"app/models"
	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/olahol/melody.v1"
)

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

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		doc := models.NewMessage("user_id", 1, *(*string)(unsafe.Pointer(&msg)))
		err := mgm.Coll(doc).Create(doc)
		if err != nil {
			log.Fatal(err)
		} else {
			m.Broadcast(msg)
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