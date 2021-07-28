package main

import "go-web/internal/admin"

func main() {
	app := admin.NewApp("../../config", "admin", "yml")
	app.Run()
}
