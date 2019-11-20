package main

import (
	"vocabulary/config"
	"vocabulary/controllers"
	"vocabulary/db"
	"log"
	"net/http"
	"flag"
	
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

	http.HandleFunc("/", controllers.GetCards)

	http.ListenAndServe(conf.Server.Addr, nil)
}