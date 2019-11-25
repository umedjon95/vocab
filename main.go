package main

import (
	"flag"
	"log"
	"net/http"
	"vocab/config"
	"vocab/controllers"
	"vocab/db"

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
	// 1. get all words;
	router.GET("/word", controllers.GetCards)
	// 2. get a word;
	router.GET("/word/:key", controllers.GetCard)
	// 3. add a word;
	router.POST("/word", controllers.InsertCard)
	// 4. edit a word;
	router.POST("/edit", controllers.EditCard)
	// 5. delete the word;
	router.DELETE("/delete", controllers.DeleteCard)

	http.ListenAndServe(conf.Server.Addr, router)
}
