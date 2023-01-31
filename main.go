package main

import (
	"net/http"

	routers "github.com/aldiramdan/go-crud-mysql/routers"
	"github.com/joho/godotenv"
)

func main() {

	envMap, mapErr := godotenv.Read(".env")
	if mapErr != nil {
		panic(mapErr)
	}

	routers.Routers()

	http.ListenAndServe(envMap["PORT"], nil)

}
