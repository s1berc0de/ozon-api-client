package pictures

//go:generate go run github.com/abice/go-enum -f=$GOFILE --marshal --names

// ImportResponseResultPictureState
// ENUM(imported=1, uploaded, failed)
type ImportResponseResultPictureState int32

// InfoResponseResultPictureState
// ENUM(imported=1, uploaded, failed)
type InfoResponseResultPictureState int32
