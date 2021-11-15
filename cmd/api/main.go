package main

import (
	"fmt"

	"github.com/csolarz-ml/shop-directory/internal/app"
)

func main() {
	router := app.Start()

	if err := router.Run(); err != nil {
		fmt.Println("main error: " + err.Error())
	}
}
