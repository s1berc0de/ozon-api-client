package v1

//go:generate go run github.com/abice/go-enum -f=$GOFILE --marshal --names

// SummaryResponseGroupItemChangeDirection
// ENUM(DIRECTION_UNKNOWN=1, DIRECTION_NONE, DIRECTION_RISE, DIRECTION_FALL)
type SummaryResponseGroupItemChangeDirection int32

// SummaryResponseGroupItemChangeMeaning
// ENUM(MEANING_UNKNOWN=1, MEANING_NONE, MEANING_GOOD, MEANING_BAD)
type SummaryResponseGroupItemChangeMeaning int32

// SummaryResponseGroupItemRatingDirection
// ENUM(UNKNOWN_DIRECTION=1, NEUTRAL, HIGHER_IS_BETTER, LOWER_IS_BETTER)
type SummaryResponseGroupItemRatingDirection int32

// SummaryResponseGroupItemStatus
// ENUM(UNKNOWN_STATUS=1, OK, WARNING, CRITICAL)
type SummaryResponseGroupItemStatus int32

// SummaryResponseGroupItemValueType
// ENUM(UNKNOWN_VALUE=1, INDEX, PERCENT, TIME, RATIO, REVIEW_SCORE, COUNT)
type SummaryResponseGroupItemValueType int32
