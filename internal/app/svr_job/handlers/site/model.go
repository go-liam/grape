package site

type Model struct {
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Data    ModelData `json:"data"`
}

type ModelData struct {
	Status int8        `json:"status"`
	List   interface{} `json:"list"`
}
