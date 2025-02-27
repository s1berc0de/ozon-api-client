package info

//go:generate go run github.com/abice/go-enum -f=$GOFILE --marshal --names

// AttributesRequestFilterVisibility
// ENUM(ALL=0, VISIBLE, INVISIBLE, EMPTY_STOCK, NOT_MODERATED, MODERATED, DISABLED, STATE_FAILED, READY_TO_SUPPLY, VALIDATION_STATE_PENDING, VALIDATION_STATE_FAIL, VALIDATION_STATE_SUCCESS, TO_SUPPLY, IN_SALE, REMOVED_FROM_SALE, BANNED, OVERPRICED, CRITICALLY_OVERPRICED, EMPTY_BARCODE, BARCODE_EXISTS, QUARANTINE, ARCHIVED, OVERPRICED_WITH_STOCK, PARTIAL_APPROVED, IMAGE_ABSENT, MODERATION_BLOCK)
type AttributesRequestFilterVisibility int32

// AttributesResponseResultDimensionUnit
// ENUM(mm=1, cm, in)
type AttributesResponseResultDimensionUnit int32
