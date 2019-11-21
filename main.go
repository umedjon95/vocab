package main

import (
	"flag"
	"log"
	"net/http"
	"vocab/config"
	"vocab/controllers"
	"vocab/db"

	_ "github.com/go-sql-driver/mysql"
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

	http.HandleFunc("/", controllers.GetPage)
	http.HandleFunc("/page", controllers.GetCard)

	http.HandleFunc("/all", controllers.GetCards)

	http.ListenAndServe(conf.Server.Addr, nil)
}
