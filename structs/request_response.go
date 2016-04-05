package structs

type Request struct {
	Lat  float64      `json:"lat" binding:"required"`
	Lng  float64      `json:"lng" binding:"required"`
	Dist float32      `json:"dist" binding:"required"`
}

type Response struct {
	S int `json:"s"`
}