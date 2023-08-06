package request

//go:generate go-enum -f=$GOFILE --marshal --names

// ContentType
// ENUM(application/json=0)
type ContentType string
