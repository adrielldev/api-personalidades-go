package main

import (
	"fmt"

	"github.com/adrielldev/api-rest-go/database"

	"github.com/adrielldev/api-rest-go/routes"
)

func main() {
	fmt.Println("Iniciando o servidor Rest com Go")
	database.ConectaComBancoDeDados()
	routes.HandleResquest()
}
