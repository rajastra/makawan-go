package apimodels

type Request struct {
	CustomerName string `json:"customerName"`
	Items        []Item `json:"items"`
}

type Item struct {
	ID          uint   `json:"id"`
	Description string `json:"description"`
	ItemCode    string `json:"itemCode"`
	Price       int64  `json:"price"`
	Quantity    int64  `json:"quantity"`
}

type Response struct {
	Data         Request `json:"data"`
	DateTrans    string  `json:"dateTrans"`
	OrderID      string  `json:"orderID"`
	ResponseCode string  `json:"responseCode"`
	Status       string  `json:"status"`
	Total        int64   `json:"total"`
}
