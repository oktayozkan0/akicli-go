package main

import (
	"fmt"
	"log"

	"github.com/oktayozkan0/akicli-go/api"
	"github.com/oktayozkan0/akicli-go/client"
)

func main() {
	a, err := api.NewAPI(
		client.WithToken(""),
		client.WithBaseURL(""),
	)
	if err != nil {
		log.Fatal(err)
	}
	apps, err := a.GetApps()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(apps.Results)
}
