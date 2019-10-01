package main

import (
	"flag"
	"github.com/kanatmg/http-rest-api/internal/app/apiserver"
	"log"
)

var (
	CONFIGPATH = "configs/apiserver.toml" //путь в конфиг файлу
)

func init() {
	flag.StringVar(&CONFIGPATH,"config",CONFIGPATH,"Путь к конфиг файлу")
	
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	
	server := apiserver.New(config)
	if error := server.Start(); error !=nil {
		log.Fatal(error)
	}

}
