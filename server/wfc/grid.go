package wfc

import (
	"math/rand/v2"
	"slices"
)

type Cell struct {
	Possibilities []int
}

type Position struct {
	X int
	Y int
}

func (c *Cell) Entropy() int {
	return len(c.Possibilities)
}

func (c *Cell) Collapse(
	weights map[int]float64,
	rng *rand.Rand,
) {
	tile := chooseWeighted(
		c.Possibilities,
		weights,
		rng,
	)

	c.Possibilities = []int{tile}
}

type Grid struct {
	Width  int
	Height int
	Cells  []Cell
}

func (g *Grid) Index(x, y int) int {
	return y*g.Width + x
}

func (g *Grid) Cell(x, y int) *Cell {
	return &g.Cells[g.Index(x, y)]
}

func (g *Grid) Neighbor(
	x, y int,
	direction Direction,
) (int, int, bool) {
	switch direction {
	case Up:
		if y == 0 {
			return 0, 0, false
		}
		return x, y - 1, true

	case Right:
		if x >= g.Width-1 {
			return 0, 0, false
		}
		return x + 1, y, true

	case Down:
		if y >= g.Height-1 {
			return 0, 0, false
		}
		return x, y + 1, true

	case Left:
		if x == 0 {
			return 0, 0, false
		}
		return x - 1, y, true
	}

	return 0, 0, false
}

func (g *Grid) LowestEntropy() (int, bool) {
	minEntropy := -1
	selected := -1

	for i, cell := range g.Cells {
		entropy := cell.Entropy()

		if entropy == 0 {
			return i, false
		}

		if entropy == 1 {
			continue
		}

		if minEntropy == -1 || entropy < minEntropy {
			minEntropy = entropy
			selected = i
		}
	}

	return selected, selected != -1
}

func (g *Grid) Propagate(
	rules *Rules,
	startX int,
	startY int,
) bool {
	queue := []Position{
		{X: startX, Y: startY},
	}

	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]

		cell := g.Cell(pos.X, pos.Y)

		for direction := Up; direction <= Left; direction++ {
			nx, ny, ok := g.Neighbor(
				pos.X,
				pos.Y,
				direction,
			)

			if !ok {
				continue
			}

			neighbor := g.Cell(nx, ny)

			changed := false

			var remaining []int

			for _, candidate := range neighbor.Possibilities {
				if HasSupport(
					candidate,
					cell,
					direction.Opposite(),
					rules,
				) {
					remaining = append(remaining, candidate)
				} else {
					changed = true
				}
			}

			if !changed {
				continue
			}

			neighbor.Possibilities = remaining

			// Contradiction
			if len(remaining) == 0 {
				return false
			}

			queue = append(queue, Position{
				X: nx,
				Y: ny,
			})
		}
	}

	return true
}

func (g *Grid) solve(
	rules Rules,
	weights map[int]float64,
	rng *rand.Rand,
) bool {
	for {
		index, ok := g.LowestEntropy()

		if !ok {
			return g.isSolved()
		}

		x := index % g.Width
		y := index / g.Width

		g.Cell(x, y).Collapse(
			weights,
			rng,
		)

		if !g.Propagate(
			&rules,
			x,
			y,
		) {
			return false
		}
	}
}

func (g *Grid) isSolved() bool {
	for _, cell := range g.Cells {
		if len(cell.Possibilities) != 1 {
			return false
		}
	}

	return true
}

func chooseWeighted(
	possibilities []int,
	weights map[int]float64,
	rng *rand.Rand,
) int {
	var total float64

	for _, tile := range possibilities {
		total += weights[tile]
	}

	r := rng.Float64() * total

	for _, tile := range possibilities {
		r -= weights[tile]

		if r < 0 {
			return tile
		}
	}

	return possibilities[len(possibilities)-1]
}

func HasSupport(
	candidate int,
	neighbor *Cell,
	direction Direction,
	rules *Rules,
) bool {
	for _, neighborTile := range neighbor.Possibilities {
		if slices.Contains(
			rules.Allowed[neighborTile][direction],
			candidate,
		) {
			return true
		}
	}

	return false
}
