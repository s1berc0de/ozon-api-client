package v1

//go:generate go-enum -f=$GOFILE --marshal --names

// ImportBySKURequestItemCurrencyCode
// ENUM(RUB=0, BYN, KZT, EUR, USD, CNY)
type ImportBySKURequestItemCurrencyCode int32
