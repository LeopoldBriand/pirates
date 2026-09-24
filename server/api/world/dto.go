package world

type CreateWorldRequest struct {
	Height int    `json:"height" binding:"required,min=1"`
	Width  int    `json:"width" binding:"required,min=1"`
	Seed   uint64 `json:"seed" binding:"required"`
}

type WorldResponse struct {
	Grid   [][]string `json:"grid"`
	Height int        `json:"height"`
	Width  int        `json:"width"`
}
