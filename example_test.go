package amclient_test

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"go.artefactual.dev/amclient"
)

// Server information can be retrieved explicitly when no other API response
// is available.
func ExampleClient_ServerInfo() {
	client := amclient.NewClient(
		http.DefaultClient,
		"https://archivematica.example.com/",
		"api-user",
		"api-key",
	)

	info, _, err := client.ServerInfo(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	if info.Version.AtLeast(1, 19, 0) {
		fmt.Println("Archivematica supports idempotent package creation")
	}
}

func ExampleResponse_ServerInfo() {
	client := amclient.NewClient(
		http.DefaultClient,
		"https://archivematica.example.com/",
		"api-user",
		"api-key",
	)

	_, response, err := client.ProcessingConfig.List(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	info, err := response.ServerInfo()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Archivematica %s (pipeline %s)\n", info.Version, info.ID)
}
