package main

import (
	"github.com/dabkoa/golang-server/application"
)

func main() {

	svc := application.CreateApp("/resources/config.json")

	svc.Start()
}
