package world

import (
	"context"
	"server/wfc"
)

type WorldService struct {
	generator wfc.Generator
}

func (s *WorldService) CreateWorld(ctx context.Context, req CreateWorldRequest) ([][]string, error) {
	// TODO: Load new weights based on the requested world type
	grid, error := s.generator.Generate(req.Height, req.Width, req.Seed, req.Seed)
	IDGrid := s.generator.GetGridWithIDs(grid)
	if error != nil {
		return nil, error
	} else {
		return IDGrid, nil
	}

}

func NewService() WorldService {
	tileset, error := wfc.LoadTileSet("./data/tilesets/pirate.json")
	weights, error := wfc.LoadTileWeights("./data/tilesets/default_weigths.json")
	if error != nil {
		panic(error)
	} else {
		filtered := wfc.FilterTerrainLayer(*tileset)
		generator := wfc.NewGenerator(filtered, *weights)
		return WorldService{
			generator: generator,
		}
	}
}
