package v1

//go:generate go run github.com/abice/go-enum -f=$GOFILE --marshal --names

// ImportBySKURequestItemCurrencyCode
// ENUM(RUB=0, BYN, KZT, EUR, USD, CNY)
type ImportBySKURequestItemCurrencyCode int32
