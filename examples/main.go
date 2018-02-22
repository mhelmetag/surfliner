package main

import (
	"fmt"

	"github.com/mhelmetag/surfliner"
)

func main() {
	client, err := surfliner.DefaultClient()
	if err != nil {
		fmt.Printf("Error building Client: %v", err)
		return
	}

	areas, err := client.Areas()
	if err != nil {
		fmt.Printf("Error fetching Areas: %v", err)
		return
	}

	fmt.Print("Fetching Areas...\n")
	for i := range areas {
		fmt.Printf("%+v\n", areas[i])
	}
	fmt.Print("\n")

	area := areas[0]
	regions, err := client.Regions(area.ID)
	if err != nil {
		fmt.Printf("Error fetching Regions: %v\n", err)
		return
	}

	fmt.Printf("Fetching Regions for Area '%s'...\n", area.Name)
	for i := range regions {
		fmt.Printf("%+v\n", regions[i])
	}
	fmt.Print("\n")

	region := regions[0]
	sregions, err := client.SubRegions(area.ID, region.ID)
	if err != nil {
		fmt.Printf("Error fetching Regions: %v\n", err)
		return
	}

	fmt.Printf("Fetching SubRegions for Region '%s'...\n", region.Name)
	for i := range sregions {
		fmt.Printf("%+v\n", sregions[i])
	}

	sregion, err := client.SubRegion("1", "4", "7")
	if err != nil {
		fmt.Printf("Error fetching Sub Region: %v\n", err)
		return
	}

	fmt.Printf("Fetching SubRegion\n")
	fmt.Printf("%+v\n", sregion)
}
