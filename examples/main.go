package main

import (
	"fmt"

	"github.com/mhelmetag/surfliner"
)

func main() {
	client, err := surfliner.DefaultClient()
	if err != nil {
		fmt.Printf("Error building Client: %v", err)
	}

	payload, err := client.ListAreas()
	if err != nil {
		fmt.Printf("Error fetching Areas: %v", err)
	}

	for i := range payload.Data {
		fmt.Printf("%+v\n", payload.Data[i])
	}
}
