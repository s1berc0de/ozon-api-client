package attributes

type UpdateRequestItemAttributeValue struct {
	DictionaryValueID int64  `json:"dictionary_value_id"`
	Value             string `json:"value"`
}

type UpdateRequestItemAttribute struct {
	ComplexID int64                             `json:"complex_id"`
	ID        int64                             `json:"id"`
	Values    []UpdateRequestItemAttributeValue `json:"values"`
}

type UpdateRequestItem struct {
	Attributes []UpdateRequestItemAttribute `json:"attributes"`
	OfferID    string                       `json:"offer_id"`
}

type UpdateRequest struct {
	Items []UpdateRequestItem `json:"items"`
}

type UpdateResponse struct {
	TaskID int64 `json:"task_id"`
}
