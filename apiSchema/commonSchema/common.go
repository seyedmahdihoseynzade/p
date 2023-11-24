package commonSchema

type Response struct {
	Status string      `json:"status" msgpack:"status"`
	Error  Error       `json:"error" msgpack:"error"`
	Data   interface{} `json:"data" msgpack:"data"`
}

type Error struct {
	Code    string `json:"code" msgpack:"code"`
	Message string `json:"message" msgpack:"message"`
}
