# meatadriver

A Go tool for attaching structured **provenance metadata** to files and storing it
directly in the filesystem's **extended attributes (xattr)** — so the metadata travels
with the file rather than living in a separate database.

Useful for tracking where a downloaded file came from (source URL, site, account,
download time, authentication state) and classifying it by type, importance, and a
time-to-live.

## Tech stack

- **Go 1.26**
- `github.com/pkg/xattr` — cross-platform extended-attribute access
- `github.com/spf13/cobra` — CLI framework
- **Hexagonal (ports & adapters)** architecture

## Architecture

```
internal/
├── core/
│   ├── domain.go      # Metadata model: FileType, Importance, Category, provenance fields
│   ├── service.go     # MetaService: validate → (de)serialize JSON → StoragePort
│   └── validator.go   # metadata validation
└── adapters/
    ├── xattr_unix.go      # StoragePort impl for Linux/macOS
    └── xattr_windows.go   # StoragePort impl for Windows
```

The domain is provider-agnostic: `MetaService` depends only on a `StoragePort` interface
(`Read` / `Write`), and the OS-specific xattr adapters implement it. Metadata is
serialized to JSON before being written to the attribute.

### Metadata model (excerpt)

- provenance: `source_url`, `source_page`, `site_host`, `account_name`,
  `is_authenticated`, `downloaded_at`
- classification: `file_type` (document / image / video / audio / archive / code / …),
  `category` (named, weighted), `importance` (critical / normal / low / trash)
- lifecycle: `ttl_days`

## Building

```bash
go build ./...
```

## Status

Early stage. The core domain, service, validation, and platform xattr adapters are
implemented; the cobra-based CLI entry point is not wired up yet (`main.go` is still a
placeholder).
