package wfc

import (
	"fmt"
	"math/rand/v2"
)

type Generator struct {
	tilset  TileSet
	weights TileWeights
	rules   Rules
}

func (g *Generator) CreateGrid(width, height int) *Grid {
	cells := make([]Cell, width*height)

	allTiles := make([]int, len(g.tilset.Tiles))

	for i := range len(g.tilset.Tiles) {
		allTiles[i] = i
	}

	for i := range cells {
		cells[i] = Cell{
			Possibilities: append([]int(nil), allTiles...),
		}
	}

	return &Grid{
		Width:  width,
		Height: height,
		Cells:  cells,
	}
}

func (g *Generator) mapWeights() (map[int]float64, error) {
	weights := make(map[int]float64, len(g.tilset.Tiles))

	for i, tile := range g.tilset.Tiles {
		weight, ok := g.weights.Weights[tile.ID]

		if !ok {
			return nil, fmt.Errorf(
				"missing weight for tile %q",
				tile.ID,
			)
		}

		if weight < 0 {
			return nil, fmt.Errorf(
				"negative weight for tile %q",
				tile.ID,
			)
		}

		weights[i] = weight
	}

	return weights, nil
}

func (g *Generator) Generate(
	width int,
	height int,
	seed1 uint64,
	seed2 uint64,
) (*Grid, error) {
	rng := rand.New(rand.NewPCG(seed1, seed2))
	mappedWeights, error := g.mapWeights()
	if error != nil {
		fmt.Println("f failed:", error)
	}
	for attempt := 0; attempt < 100; attempt++ {
		grid := g.CreateGrid(width, height)

		if grid.solve(g.rules, mappedWeights, rng) {
			return grid, nil
		}
	}

	return nil, fmt.Errorf(
		"failed to generate map after 100 attempts",
	)
}

func (g *Generator) GetGridWithIDs(grid *Grid) [][]string {
	IDGrid := make([][]string, grid.Height)
	for y := 0; y < grid.Height; y++ {
		IDGrid[y] = make([]string, grid.Width)
		for x := 0; x < grid.Width; x++ {
			cell := grid.Cell(x, y)
			if len(cell.Possibilities) == 1 {
				tileIndex := cell.Possibilities[0]
				tileID := g.tilset.Tiles[tileIndex].ID
				IDGrid[y][x] = tileID
			} else {
				IDGrid[y][x] = "?"
			}
		}
	}
	return IDGrid
}

// Check if 2 tiles are compatible
func Compatible(a Tile, b Tile, direction Direction) bool {
	switch direction {
	case Up:
		return a.Edges.Top == b.Edges.Bottom

	case Right:
		return a.Edges.Right == b.Edges.Left

	case Down:
		return a.Edges.Bottom == b.Edges.Top

	case Left:
		return a.Edges.Left == b.Edges.Right
	}

	return false
}

// Precompile tileset rules
type Rules struct {
	// [tile][direction] = tiles that can be next to it
	Allowed [][][]int
}

func buildRules(tiles []Tile) *Rules {
	allowed := make([][][]int, len(tiles))

	for i := range tiles {
		allowed[i] = make([][]int, 4)

		for d := 0; d < 4; d++ {
			direction := Direction(d)

			for j := range tiles {
				if Compatible(tiles[i], tiles[j], direction) {
					allowed[i][d] = append(allowed[i][d], j)
				}
			}
		}
	}

	return &Rules{
		Allowed: allowed,
	}
}

func NewGenerator(set TileSet, weights TileWeights) Generator {
	rules := buildRules(set.Tiles)
	return Generator{
		tilset:  set,
		weights: weights,
		rules:   *rules,
	}
}
