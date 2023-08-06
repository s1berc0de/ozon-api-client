package info

import "time"

type LimitRequest struct{}

type LimitResponseDailyCreate struct {
	Limit   int64     `json:"limit"`
	ResetAt time.Time `json:"reset_at"`
	Usage   int64     `json:"usage"`
}

type LimitResponseDailyUpdate struct {
	Limit   int64     `json:"limit"`
	ResetAt time.Time `json:"reset_at"`
	Usage   int64     `json:"usage"`
}

type LimitResponseTotal struct {
	Limit int64 `json:"limit"`
	Usage int64 `json:"usage"`
}

type LimitResponse struct {
	DailyCreate LimitResponseDailyCreate `json:"daily_create"`
	DailyUpdate LimitResponseDailyUpdate `json:"daily_update"`
	Total       LimitResponseTotal       `json:"total"`
}
