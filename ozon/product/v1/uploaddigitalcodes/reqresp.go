package uploaddigitalcodes

type InfoRequest struct {
	TaskID int64 `json:"task_id"`
}

type InfoResponseResult struct {
	Status InfoResponseResultStatus `json:"status"`
}

type InfoResponse struct {
	Result InfoResponseResult `json:"result"`
}
