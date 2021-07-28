package main

import "go-web/internal/auth"

func main() {
	app := auth.NewApp("../../config", "auth", "yml")
	app.Run()
}
