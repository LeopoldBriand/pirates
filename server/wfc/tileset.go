package wfc

import (
	"encoding/json"
	"os"
	"slices"
)

type TileSet struct {
	Tiles []Tile `json:"tiles"`
}

type Tile struct {
	ID     string `json:"id"`
	Source string `json:"source"`
	Kind   string `json:"kind"`
	Edges  Edges  `json:"edges"`
	Layers []int  `json:"layers"`
}

type TileWeights struct {
	Weights map[string]float64 `json:"weights"`
}

type Edges struct {
	Top    string `json:"top"`
	Right  string `json:"right"`
	Bottom string `json:"bottom"`
	Left   string `json:"left"`
}

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

func (d Direction) Opposite() Direction {
	switch d {
	case Up:
		return Down
	case Right:
		return Left
	case Down:
		return Up
	case Left:
		return Right
	}

	return Up
}

func LoadTileSet(path string) (*TileSet, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var tileset TileSet

	if err := json.Unmarshal(data, &tileset); err != nil {
		return nil, err
	}

	return &tileset, nil
}

func LoadTileWeights(path string) (*TileWeights, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var weights TileWeights

	if err := json.Unmarshal(data, &weights); err != nil {
		return nil, err
	}

	return &weights, nil
}

func FilterTerrainLayer(tileSet TileSet) TileSet {
	var tiles []Tile

	for _, tile := range tileSet.Tiles {
		if slices.Contains(tile.Layers, 1) {
			tiles = append(tiles, tile)
		}
	}

	return TileSet{
		Tiles: tiles,
	}
}
