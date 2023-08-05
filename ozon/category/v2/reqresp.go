package v2

type TreeRequest struct {
	CategoryId int64    `json:"category_id"`
	Language   Language `json:"language"`
}

type TreeResponseResult struct {
	CategoryId int64                `json:"category_id"`
	Children   []TreeResponseResult `json:"children"`
	Title      string               `json:"title"`
}

type TreeResponse struct {
	Result []TreeResponseResult `json:"result"`
}
