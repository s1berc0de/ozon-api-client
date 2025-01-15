package v2

//go:generate go run github.com/abice/go-enum -f=$GOFILE --marshal --names

// ImportRequestItemCurrencyCode
// ENUM(RUB=0, BYN, KZT, EUR, USD, CNY)
type ImportRequestItemCurrencyCode int32

// ImportRequestItemDimensionUnit
// ENUM(mm=1, cm, in)
type ImportRequestItemDimensionUnit int32

// ImportRequestItemServiceType
// ENUM(IS_CODE_SERVICE=0, IS_NO_CODE_SERVICE)
type ImportRequestItemServiceType int32

// ImportRequestItemWeightUnit
// ENUM(g=1, kg, lb)
type ImportRequestItemWeightUnit int32

// ListRequestFilterVisibility
// ENUM(ALL=0, VISIBLE, INVISIBLE, EMPTY_STOCK, NOT_MODERATED, MODERATED, DISABLED, STATE_FAILED, READY_TO_SUPPLY, VALIDATION_STATE_PENDING, VALIDATION_STATE_FAIL, VALIDATION_STATE_SUCCESS, TO_SUPPLY, IN_SALE, REMOVED_FROM_SALE, BANNED, OVERPRICED, CRITICALLY_OVERPRICED, EMPTY_BARCODE, BARCODE_EXISTS, QUARANTINE, ARCHIVED, OVERPRICED_WITH_STOCK, PARTIAL_APPROVED, IMAGE_ABSENT, MODERATION_BLOCK)
type ListRequestFilterVisibility int32

// InfoResponseResultCurrencyCode
// ENUM(RUB=1, BYN, KZT, EUR, USD, CNY)
type InfoResponseResultCurrencyCode int32

// InfoResponseResultPriceIndexesPriceIndex
// ENUM(WITHOUT_INDEX=1, PROFIT, AVG_PROFIT, NON_PROFIT)
type InfoResponseResultPriceIndexesPriceIndex int32
