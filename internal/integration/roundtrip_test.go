package integration_test

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"

	"github.com/PaulPAdm/meta-driver/internal/adapters"
	"github.com/PaulPAdm/meta-driver/internal/core"
)

// TestRoundTrip проверяет полный цикл сервиса с реальным адаптером:
// MetaService.Write сериализует метаданные в JSON и пишет их через
// XattrAdapter, а MetaService.Read читает и десериализует обратно.
// Прочитанные метаданные должны быть идентичны записанным.
func TestRoundTrip(t *testing.T) {
	// Реальный файл на диске, к которому адаптер привяжет метаданные
	// (NTFS ADS на Windows, xattr на Unix). Базовый файл должен существовать.
	filePath := filepath.Join(t.TempDir(), "sample.bin")
	if err := os.WriteFile(filePath, []byte("payload"), 0600); err != nil {
		t.Fatalf("setup: не удалось создать файл: %v", err)
	}

	// Реальный адаптер, выбранный для текущей ОС через build-теги.
	storage := adapters.NewXattrAdapter()
	service := core.NewMetaService(storage)

	ttl := 30
	// time.Date без монотонных часов и в UTC, чтобы round-trip через
	// RFC3339 давал побитово равное значение при сравнении.
	want := &core.Metadata{
		ProtocolVersion: "1.0",
		SourceURL:       "https://example.com/file",
		SourcePage:      "https://example.com/page",
		SiteHost:        "example.com",
		AccountName:     "alice",
		IsAuthenticated: true,
		DownloadedAt:    time.Date(2026, 6, 25, 10, 30, 0, 0, time.UTC),
		FileType:        core.FileTypeDocument,
		Category: []core.Category{
			{Name: "work", Weight: 0.8},
			{Name: "study", Weight: 0.2},
		},
		Importance: core.ImportanceCritical,
		TTLDays:    &ttl,
	}

	if err := service.Write(filePath, want); err != nil {
		t.Fatalf("Write: %v", err)
	}

	got, err := service.Read(filePath)
	if err != nil {
		t.Fatalf("Read: %v", err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("round-trip изменил метаданные:\n записано: %+v\n прочитано: %+v", want, got)
	}
}

// TestRoundTrip_Minimal проверяет цикл на минимально валидных метаданных,
// где все omitempty-поля опущены, а указатель TTLDays равен nil.
func TestRoundTrip_Minimal(t *testing.T) {
	filePath := filepath.Join(t.TempDir(), "minimal.bin")
	if err := os.WriteFile(filePath, []byte("x"), 0600); err != nil {
		t.Fatalf("setup: не удалось создать файл: %v", err)
	}

	service := core.NewMetaService(adapters.NewXattrAdapter())

	want := &core.Metadata{
		ProtocolVersion: "1.0",
		DownloadedAt:    time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
		FileType:        core.FileTypeOther,
		Importance:      core.ImportanceLow,
	}

	if err := service.Write(filePath, want); err != nil {
		t.Fatalf("Write: %v", err)
	}

	got, err := service.Read(filePath)
	if err != nil {
		t.Fatalf("Read: %v", err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("round-trip изменил метаданные:\n записано: %+v\n прочитано: %+v", want, got)
	}
}
