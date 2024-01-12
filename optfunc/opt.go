package optfunc

// main class
type House struct {
	WallMaterial string
	GarageArea   float64
}

type optFunc func(h *House)

func WithWallMaterial(wallMaterial string) optFunc {
	return func(h *House) {
		h.WallMaterial = wallMaterial
	}
}

func WithGarageArea(garageArea float64) optFunc {
	return func(h *House) {
		h.GarageArea = garageArea
	}
}

func NewHouse(options ...optFunc) *House {
	h := &House{}
	for _, o := range options {
		o(h)
	}
	return h
}
