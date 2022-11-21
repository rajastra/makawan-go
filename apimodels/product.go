package apimodels

type RequestProduct struct {
	ID          uint   `json:"id"`
	ProductID   string `json:"productID"`
	ProductName string `json:"productName"`
	Price       int64  `json:"price"`
	Description string `json:"description"`
}

type ResponseProduct struct {
	Data         RequestProduct `json:"data"`
	ProductID    string         `json:"productID"`
	ResponseCode string         `json:"responseCode"`
	Status       string         `json:"status"`
	Total        int64          `json:"total"`
}