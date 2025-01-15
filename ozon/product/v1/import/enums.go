package _import

//go:generate go run github.com/abice/go-enum -f=$GOFILE --marshal --names

// InfoResponseResultItemStatus
// ENUM(pending=1, imported, failed)
type InfoResponseResultItemStatus int32

// PricesRequestPriceAutoActionEnabled
// ENUM(UNKNOWN=0, ENABLED, DISABLED)
type PricesRequestPriceAutoActionEnabled int32

// PricesRequestPriceCurrencyCode
// ENUM(RUB=0, BYN, KZT, EUR, USD, CNY)
type PricesRequestPriceCurrencyCode int32

// PricesRequestPricePriceStrategyEnabled
// ENUM(UNKNOWN=0, ENABLED, DISABLED)
type PricesRequestPricePriceStrategyEnabled int32
