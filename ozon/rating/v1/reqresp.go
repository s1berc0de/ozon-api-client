package v1

import "time"

type SummaryRequest struct {
}

type SummaryResponseGroupItemChange struct {
	Direction SummaryResponseGroupItemChangeDirection `json:"direction"`
	Meaning   SummaryResponseGroupItemChangeMeaning   `json:"meaning"`
}

type SummaryResponseGroupItem struct {
	Change          SummaryResponseGroupItemChange          `json:"change"`
	CurrentValue    float64                                 `json:"current_value"`
	Name            string                                  `json:"name"`
	PastValue       float64                                 `json:"past_value"`
	Rating          string                                  `json:"rating"`
	RatingDirection SummaryResponseGroupItemRatingDirection `json:"rating_direction"`
	Status          SummaryResponseGroupItemStatus          `json:"status"`
	ValueType       SummaryResponseGroupItemValueType       `json:"value_type"`
}

type SummaryResponseGroup struct {
	GroupName string                     `json:"group_name"`
	Items     []SummaryResponseGroupItem `json:"items"`
}

type SummaryResponse struct {
	Groups               []SummaryResponseGroup `json:"groups"`
	PenaltyScoreExceeded bool                   `json:"penalty_score_exceeded"`
	Premium              bool                   `json:"premium"`
}

type HistoryRequest struct {
	DateFrom          time.Time `json:"date_from"`
	DateTo            time.Time `json:"date_to"`
	Ratings           []string  `json:"ratings"`
	WithPremiumScores bool      `json:"with_premium_scores"`
}

type HistoryResponsePremiumScoreScore struct {
	Date        time.Time `json:"date"`
	RatingValue float64   `json:"rating_value"`
	Value       int32     `json:"value"`
}

type HistoryResponsePremiumScore struct {
	Rating string                             `json:"rating"`
	Scores []HistoryResponsePremiumScoreScore `json:"scores"`
}

type HistoryResponseRatingValueStatus struct {
	Danger  bool `json:"danger"`
	Premium bool `json:"premium"`
	Warning bool `json:"warning"`
}

type HistoryResponseRatingValue struct {
	DateFrom time.Time                        `json:"date_from"`
	DateTo   time.Time                        `json:"date_to"`
	Status   HistoryResponseRatingValueStatus `json:"status"`
	Value    float64                          `json:"value"`
}

type HistoryResponseRating struct {
	DangerThreshold  float64                      `json:"danger_threshold"`
	PremiumThreshold float64                      `json:"premium_threshold"`
	Rating           string                       `json:"rating"`
	Values           []HistoryResponseRatingValue `json:"values"`
	WarningThreshold float64                      `json:"warning_threshold"`
}

type HistoryResponse struct {
	PremiumScores []HistoryResponsePremiumScore `json:"premium_scores"`
	Ratings       []HistoryResponseRating       `json:"ratings"`
}
