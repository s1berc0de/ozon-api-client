package _import

type InfoRequest struct {
	TaskID string `json:"task_id"`
}

type InfoResponseResultItemErrorOptionalDescriptionElements struct {
	PropertyName string `json:"property name*"`
}

type InfoResponseResultItemError struct {
	Code                        string                                                 `json:"code"`
	Message                     string                                                 `json:"message"`
	State                       string                                                 `json:"state"`
	Level                       string                                                 `json:"level"`
	Description                 string                                                 `json:"description"`
	Field                       string                                                 `json:"field"`
	AttributeID                 int64                                                  `json:"attribute_id"`
	AttributeName               string                                                 `json:"attribute_name"`
	OptionalDescriptionElements InfoResponseResultItemErrorOptionalDescriptionElements `json:"optional_description_elements"`
}

type InfoResponseResultItem struct {
	OfferID   string                        `json:"offer_id"`
	ProductID int64                         `json:"product_id"`
	Status    InfoResponseResultItemStatus  `json:"status"`
	Errors    []InfoResponseResultItemError `json:"errors"`
}

type InfoResponseResult struct {
	Items []InfoResponseResultItem `json:"items"`
	Total int32                    `json:"total"`
}

type InfoResponse struct {
	Result InfoResponseResult `json:"result"`
}
