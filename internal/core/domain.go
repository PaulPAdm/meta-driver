package core

import "time"

type FileType string
type Importance string

const (
	FileTypeDocument     FileType = "document"
	FileTypeImage        FileType = "image"
	FileTypeVideo        FileType = "video"
	FileTypeVideoLoop    FileType = "video_loop"
	FileTypeAudio        FileType = "audio"
	FileTypeArchive      FileType = "archive"
	FileTypeCodeSource   FileType = "code_source"
	FileTypeCodeCompiled FileType = "code_compiled"
	FileTypeInstaller    FileType = "installer"
	FileTypeApplication  FileType = "application"
	FileTypeOther        FileType = "other"
)

const (
	ImportanceCritical Importance = "critical"
	ImportanceNormal   Importance = "normal"
	ImportanceLow      Importance = "low"
	ImportanceTrash    Importance = "trash"
)

type Category struct {
	Name   string  `json:"name"`
	Weight float64 `json:"weight"`
}

type Metadata struct {
	ProtocolVersion string     `json:"protocol_version"`
	SourceURL       string     `json:"source_url,omitempty"`
	SourcePage      string     `json:"source_page,omitempty"`
	SiteHost        string     `json:"site_host,omitempty"`
	AccountName     string     `json:"account_name,omitempty"`
	IsAuthenticated bool       `json:"is_authenticated"`
	DownloadedAt    time.Time  `json:"downloaded_at"`
	FileType        FileType   `json:"file_type,omitempty"`
	Category        []Category `json:"category,omitempty"`
	Importance      Importance `json:"importance,omitempty"`
	TTLDays         *int       `json:"ttl_days,omitempty"`
}
