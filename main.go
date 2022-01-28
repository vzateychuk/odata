package main

import (
	"context"
	"fmt"
	"github.com/vez/odata/handler"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

// var ctx = context.TODO()

func init() {
	log.Println("INIT")
}

func main() {
	client := connectToMongo("mongodb://root:example@localhost:27017")
	col := client.Database("meta").Collection("metadata")
	/*
		col.InsertOne(ctx, meta.NewMeta("metaRU", "RU"))
		col.InsertOne(ctx, meta.NewMeta("metaUSA", "USA"))
		col.InsertOne(ctx, meta.NewMeta("metaEN", "EN"))
		col.InsertOne(ctx, meta.NewMeta("metaFR", "FR"))
	*/

	// Объявляем http роутеры, связывая Path и обработчики
	h := handler.NewHandlerInstance(col)

	r := mux.NewRouter()
	r.HandleFunc("/", h.List).Methods("GET")
	r.HandleFunc("/metas", h.List).Methods("GET")

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", r)
}

func connectToMongo(uri string) *mongo.Client {

	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	//  убедимся, что ваш сервер MongoDB был обнаружен и подключен для успешного использования метода Ping
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
