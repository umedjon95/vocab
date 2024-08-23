package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/umedjon95/vocab/config"
	"github.com/umedjon95/vocab/controllers"
	"github.com/umedjon95/vocab/db"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

func main() {

	cfgFilePath := flag.String("config", "config/config.json", "path of the config file")
	conf, err := config.Parce(*cfgFilePath)
	if err != nil {
		log.Fatalln("Parcing config file err:", err)
	}

	err = db.Connect()
	if err != nil {
		log.Fatalln("Connecting to db err:", err)
	}
	defer db.Close()

	router := httprouter.New()

	//ping project;
	router.GET("/ping", controllers.Ping)

	//get all words;
	router.GET("/word", controllers.GetCards)

	//get 4 random cards;
	router.GET("/randcard", controllers.RandCard)

	//get 4 random cards;
	router.PUT("/randcard", controllers.CheckCard)

	//get a word;
	router.GET("/word/:key", controllers.GetCard)

	//add a word;
	router.POST("/word", controllers.InsertCard)

	//edit a word;
	router.PUT("/word", controllers.EditCard)

	//delete the word;
	router.DELETE("/word/:key", controllers.DeleteCard)

	http.ListenAndServe(conf.Server.Addr, router)
}
