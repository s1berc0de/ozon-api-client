package _import

//go:generate go-enum -f=$GOFILE --marshal --names

// InfoResponseResultItemStatus
// ENUM(pending=1, imported, failed)
type InfoResponseResultItemStatus int32
