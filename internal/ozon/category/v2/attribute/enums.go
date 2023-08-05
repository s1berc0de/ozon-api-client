package attribute

//go:generate go-enum -f=$GOFILE --marshal --names

// Language
// ENUM(DEFAULT=0, RU, EN)
type Language int32
