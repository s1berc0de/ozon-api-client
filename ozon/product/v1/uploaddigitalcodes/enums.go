package uploaddigitalcodes

//go:generate go-enum -f=$GOFILE --marshal --names

// InfoResponseResultStatus
// ENUM(pending=1, imported, failed)
type InfoResponseResultStatus int32
