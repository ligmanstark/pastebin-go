package services

func GetExtensionMime(mimeType string) string {
	extensions := map[string]string{
		"image/jpeg": ".jpg",
		"image/png":  ".png",
		"image/webp": ".webp",
		"image/gif":  ".gif",
	}
	if ext, ok := extensions[mimeType]; ok {
		return ext
	}
	return ".bin"
}
