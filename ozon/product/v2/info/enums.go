package info

//go:generate go run github.com/abice/go-enum -f=$GOFILE --marshal --names

// ListResponseResultItemCurrencyCode
// ENUM(RUB=1, BYN, KZT, EUR, USD, CNY)
type ListResponseResultItemCurrencyCode int32

// ListResponseResultItemPriceIndexesPriceIndex
// ENUM(WITHOUT_INDEX=1, PROFIT, AVG_PROFIT, NON_PROFIT)
type ListResponseResultItemPriceIndexesPriceIndex int32
