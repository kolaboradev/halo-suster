package response

type Web struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
