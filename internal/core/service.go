package core

import "encoding/json"

type StoragePort interface {
	Write(filePath string, data []byte) error
	Read(filePath string) ([]byte, error)
}
type MetaService struct {
	storage StoragePort
}

func NewMetaService(storage StoragePort) *MetaService {
	return &MetaService{storage: storage}
}

func (ms *MetaService) Write(filePath string, metadata *Metadata) error {
	if err := ValidateMetadata(metadata); err != nil {
		return err
	}

	data, err := json.Marshal(metadata)
	if err != nil {
		return err
	}

	return ms.storage.Write(filePath, data)
}

func (ms *MetaService) Read(filePath string) (*Metadata, error) {
	data, err := ms.storage.Read(filePath)
	if err != nil {
		return nil, err
	}

	var metadata Metadata
	if err := json.Unmarshal(data, &metadata); err != nil {
		return nil, err
	}

	return &metadata, nil
}
