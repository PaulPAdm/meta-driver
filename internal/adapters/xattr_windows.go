//go:build windows

package adapters

import "os"

const metaAttrName = "metadriver.protocol"
const permissionLevel os.FileMode = 0600

type XattrAdapter struct{}

func NewXattrAdapter() *XattrAdapter {
	return &XattrAdapter{}
}

func (xa *XattrAdapter) Write(filePath string, data []byte) error {
	return os.WriteFile(streamPath(filePath), data, permissionLevel)
}

func (xa *XattrAdapter) Read(filePath string) ([]byte, error) {
	return os.ReadFile(streamPath(filePath))
}

func streamPath(fileName string) string {
	return fileName + ":" + metaAttrName
}
