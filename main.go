package main

import (
	"log"

	"github.com/Aaryan-palit/week2-GL1-CipherSchools/database"
	"github.com/Aaryan-palit/week2-GL1-CipherSchools/routers"
)

func main() {
	database.Setup()                    //establish the dtatbase connection
	engine := routers.Router()          //get the customized engine
	err := engine.Run("127.0.0.1:8080") // start the engine
	if err != nil {
		log.Fatal(err)
	}
}
