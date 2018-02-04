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

	areas, err := client.ListAreas()
	if err != nil {
		fmt.Printf("Error fetching Areas: %v", err)
	}

	fmt.Print("Fetching Areas...\n")
	for i := range areas {
		fmt.Printf("%+v\n", areas[i])
	}
	fmt.Print("\n")

	area := areas[0]
	regions, err := client.ListRegions(area.ID)
	if err != nil {
		fmt.Printf("Error fetching Regions: %v\n", err)
	}

	fmt.Printf("Fetching Regions for Area '%s'...\n", area.Name)
	for i := range regions {
		fmt.Printf("%+v\n", regions[i])
	}
	fmt.Print("\n")

	region := regions[0]
	sregions, err := client.ListSubRegions(area.ID, region.ID)
	if err != nil {
		fmt.Printf("Error fetching Regions: %v\n", err)
	}

	fmt.Printf("Fetching SubRegions for Region '%s'...\n", region.Name)
	for i := range sregions {
		fmt.Printf("%+v\n", sregions[i])
	}
}
