package main

import (
	"context"
	"log"
	"net/http"

	"github.com/m1heng/oapi-codegen-reproduce/client"
)

func Try() {
	hc := http.Client{}
	c, err := client.NewClientWithResponses("http://localhost:1234", client.WithHTTPClient(&hc))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := c.PromptsGetWithResponse(context.Background(), "das", nil)
	if err != nil {
		log.Fatal(err)
	}

	resp.JSON200.AsPrompt0()
}
