package main

import (
	"github.com/hugoing97/twittor/bd"
	"github.com/hugoing97/twittor/handlers"
	"log"
)

func main() {
	if bd.ChequeoConection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
