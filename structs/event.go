package structs

type Event struct {
	Id       int `json:"id" binding:"required"`
	Img_id   int `json:"img_id" binding:"required"`
	Message  string `json:"message" binding:"required"`
	//[latitude, longitude]
	Position [] float64 `json:"position" binding:"required"`
	Username string `json:"username" binding:"required"`
	Status   int `json:"status" binding:"required"`
}
