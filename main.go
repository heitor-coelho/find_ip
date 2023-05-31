package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	aplicacao := app.Gerar()

	aplicacao.Run(os.Args)

	if err := aplicacao.Run(os.Args); err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

}
