package v3

//go:generate go run github.com/abice/go-enum -f=$GOFILE --marshal --names

// AttributeType
// ENUM(ALL=0, REQUIRED, OPTIONAL)
type AttributeType int32

// Language
// ENUM(DEFAULT=0, RU, EN, TR)
type Language int32
