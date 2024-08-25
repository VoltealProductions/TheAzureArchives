package main

import "github.com/VoltealProductions/TheAzureArcchives/cmd/api"

func main() {
	server := api.NewApiServer(":3030", nil)
	server.Run()
}
