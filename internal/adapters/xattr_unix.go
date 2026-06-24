//go:build linux || darwin

package adapters

import "github.com/pkg/xattr"

const metaAttrName = "user.metadriver.protocol"

type XattrAdapter struct{}

func NewXattrAdapter() *XattrAdapter {
	return &XattrAdapter{}
}

func (xa *XattrAdapter) Write(filePath string, data []byte) error {
	xattr.Set(filePath, metaAttrName, metadata.Value())
}

func (xa *XattrAdapter) Read(filePath string) ([]byte, error) {
	return xattr.Get(filePath, metaAttrName)
}
