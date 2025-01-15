package attribute

//go:generate go run github.com/abice/go-enum -f=$GOFILE --marshal --names

// Language
// ENUM(DEFAULT=0, RU, EN)
type Language int32
