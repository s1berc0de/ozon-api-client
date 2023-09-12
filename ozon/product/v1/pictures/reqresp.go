package pictures

type ImportRequest struct {
	ColorImage string   `json:"color_image"`
	Images     []string `json:"images"`
	Images360  []string `json:"images360"`
	ProductID  int64    `json:"product_id"`
}

type ImportResponseResultPicture struct {
	Is360     bool                             `json:"is_360"`
	IsColor   bool                             `json:"is_color"`
	IsPrimary bool                             `json:"is_primary"`
	ProductID int64                            `json:"product_id"`
	State     ImportResponseResultPictureState `json:"state"`
	URL       string                           `json:"url"`
}

type ImportResponseResult struct {
	Pictures []ImportResponseResultPicture `json:"pictures"`
}

type ImportResponse struct {
	Result ImportResponseResult `json:"result"`
}

type InfoRequest struct {
	ProductID []string `json:"product_id"`
}

type InfoResponseResultPicture struct {
	Is360     bool                           `json:"is_360"`
	IsColor   bool                           `json:"is_color"`
	IsPrimary bool                           `json:"is_primary"`
	ProductID int64                          `json:"product_id"`
	State     InfoResponseResultPictureState `json:"state"`
	URL       string                         `json:"url"`
}

type InfoResponseResult struct {
	Pictures []InfoResponseResultPicture `json:"pictures"`
}

type InfoResponse struct {
	Result InfoResponseResult `json:"result"`
}
