package main

import (
	"api/database"
	"api/server"
	"fmt"
)

func main() {
	fmt.Println("Conexao com o banco")
	database.ConectaBanco()

	s := server.NewSever()
	s.Run()
}
