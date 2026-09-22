package main

import (
	"fmt"
	"server/wfc"
)

func main() {
	fmt.Println("init")
	tileset, error := wfc.LoadTileSet("./data/tilesets/pirate.json")
	weights, error := wfc.LoadTileWeights("./data/tilesets/default_weigths.json")
	if error != nil {
		fmt.Println("f failed:", error)
	} else {
		filtered := wfc.FilterTerrainLayer(*tileset)
		generator := wfc.NewGenerator(filtered, *weights)
		grid, error := generator.Generate(50, 50, 123, 123)
		if error != nil {
			fmt.Println("f failed:", error)
		} else {
			fmt.Println("f worked:", grid.Cells)
		}

	}

}
