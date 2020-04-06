package config

import (
	"encoding/json"
	"io/ioutil"
	"log"

	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

var (
	cfg Config
)

func init() {

	log.SetOutput(&lumberjack.Logger{
		Filename:   "logs/app.log",
		MaxSize:    10, // megabytes
		MaxBackups: 7,
		MaxAge:     40,   //days
		Compress:   true, // disabled by default
	})

	log.Println("-------- * ------- Starting Logging -------- * -------")
}

// Parce config from config file
func Parce(path string) (*Config, error) {

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// Peek get config instance
func Peek() *Config {
	return &cfg
}

// Config is database server configuration
type Config struct {
	Server   server   `json:"server"`
	Database database `json:"database"`
}

type server struct {
	Addr string `json:"addr"`
	Name string `json:"name"`
}

type database struct {
	Addr   string `json:"addr"`
	User   string `json:"user"`
	Pass   string `json:"pass"`
	DBName string `json:"dbname"`
}
