package health

type ResponseCheckIdCard struct {
	Status string `json:"status"`
	Data   struct {
		Status bool   `json:"status"`
		Code   string `json:"code"`
		Desc   string `json:"desc"`
	} `json:"data"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}