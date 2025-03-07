package filex

import (
	"github.com/suyuan32/simple-admin-file/api/internal/enum/filetype"
)

// ConvertFileTypeToUint8 converts file type string to uint8.
func ConvertFileTypeToUint8(fileType string) uint8 {
	switch fileType {
	case "other":
		return filetype.Other
	case "image":
		return filetype.Image
	case "video":
		return filetype.Video
	case "audio":
		return filetype.Audio
	default:
		return filetype.Other
	}
}
