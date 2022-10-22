package helpers

import "os"

func GetBaseStoragePath() string {
	baseStoragePath := os.Getenv("OBJECT_STORAGE_BASE_PATH")
	if baseStoragePath == "" {
		baseStoragePath = "/tmp"
	}
	return baseStoragePath
}
