package postings

//go:generate go-enum -f=$GOFILE --marshal --names

// Language
// ENUM(DEFAULT=1, RU, EN)
type Language int32
