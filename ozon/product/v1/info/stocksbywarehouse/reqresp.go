package stocksbywarehouse

type FBSRequest struct {
	SKU    []string `json:"sku"`
	FBSSKU []string `json:"fbs_sku"`
}

type FBSResponseResult struct {
	SKU           int64  `json:"sku"`
	FBSSKU        int64  `json:"fbs_sku"`
	Present       int64  `json:"present"`
	ProductID     int64  `json:"product_id"`
	Reserved      int64  `json:"reserved"`
	WarehouseID   int64  `json:"warehouse_id"`
	WarehouseName string `json:"warehouse_name"`
}

type FBSResponse struct {
	Result []FBSResponseResult `json:"result"`
}
