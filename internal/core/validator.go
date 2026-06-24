package core

import (
	"errors"
	"fmt"
	"strings"
)

var validFileTypes = map[FileType]bool{
	FileTypeDocument:     true,
	FileTypeImage:        true,
	FileTypeVideo:        true,
	FileTypeVideoLoop:    true,
	FileTypeAudio:        true,
	FileTypeArchive:      true,
	FileTypeCodeSource:   true,
	FileTypeCodeCompiled: true,
	FileTypeInstaller:    true,
	FileTypeApplication:  true,
	FileTypeOther:        true,
}

var importance = map[Importance]bool{
	ImportanceCritical: true,
	ImportanceNormal:   true,
	ImportanceLow:      true,
	ImportanceTrash:    true,
}

var validCategories = map[string]bool{
	"work":     true,
	"study":    true,
	"personal": true,
	"other":    true,
}

func ValidateMetadata(data *Metadata) error {
	var errs []string

	if data.ProtocolVersion == "" {
		errs = append(errs, "ProtocolVersion is required")
	}

	if data.FileType == "" || !validFileTypes[data.FileType] {
		errs = append(errs, "invalid file type")
	}

	if data.Importance == "" || !importance[data.Importance] {
		errs = append(errs, "invalid importance")
	}

	for i, category := range data.Category {
		if !validCategories[category.Name] {
			errs = append(errs, fmt.Sprintf("category[%d]: invalid name %q", i, category.Name))
		}

		if category.Weight < 0 || category.Weight > 1 {
			errs = append(errs, fmt.Sprintf("category[%d]: invalid weight %q", i, category.Name))
		}
	}

	if len(errs) > 0 {
		return errors.New(strings.Join(errs, "; "))
	}

	return nil
}
