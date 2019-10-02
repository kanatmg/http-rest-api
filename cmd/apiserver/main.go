package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/kanatmg/http-rest-api/internal/app/apiserver"
	"log"
)

var (
	configPath string //путь в конфиг файлу
)

func init() {
	flag.StringVar(&configPath,"config","configs/apiserver.toml","Путь к конфиг файлу")
	
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if(err != nil){
		log.Fatal(err)
	}

	server := apiserver.New(config)
	if error := server.Start(); error !=nil {
		log.Fatal(error)
	}

}
