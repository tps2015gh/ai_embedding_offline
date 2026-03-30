# Build Instructions / คำแนะนำการคอมไพล์

**Updated:** 2025-03-31 00:28 (Thai Time UTC+7)

---

## No C Compiler Required! ✅

This project now uses **in-memory vector storage** (JSON-based) instead of SQLite.

**Benefits:**
- ✅ No GCC/MinGW required
- ✅ No CGO dependencies
- ✅ Works on any Windows system
- ✅ Faster to build
- ✅ Simpler deployment

---

## Quick Build

```bash
# Just run (no special setup needed!)
go build -o ai_embedding.exe ./cmd/
go build -o ngram.exe ./cmd/ngram/
```

Or use the build script:
```bash
.\build.bat
```

---

## Old SQLite Version (Not Recommended)

If you want to use SQLite (requires GCC):

1. Restore the SQLite version:
```bash
move internal\vectorstore\vectorstore_sqlite.go.bak internal\vectorstore\vectorstore.go
del internal\vectorstore\vectorstore_mem.go
```

2. Install MinGW: https://www.mingw-w64.org/

3. Build with CGO:
```bash
set CGO_ENABLED=1
go build -o ai_embedding.exe ./cmd/
```

---

## Current Setup (In-Memory)

**Files:**
- `internal/vectorstore/vectorstore_mem.go` - Main implementation
- `data/vectors.json` - JSON database (created on first use)

**Storage:**
- Vectors stored in memory (RAM)
- Persisted to `data/vectors.json` on save
- Fast lookups, no SQL overhead

---

**Status:** ✅ Working without SQLite!
