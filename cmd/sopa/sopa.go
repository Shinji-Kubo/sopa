package main

import "github.com/Shinji-Kubo/sopa/handler"

func main() {
	r := handler.New()

	r.Run(":8080")
}
