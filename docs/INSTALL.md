# Installation Guide / คู่มือการติดตั้ง

## Quick Start (Windows) / เริ่มต้นอย่างรวดเร็ว

### Option 1: Automated Setup (Recommended) ⭐

```bash
# Run the setup script
setup.bat
```

This will:
1. Check if Go is installed
2. Download dependencies
3. Build executables
4. Train the n-gram model
5. Create data directories

### Option 2: Manual Setup

```bash
# 1. Install Go (if not installed)
# Download from: https://go.dev/dl/

# 2. Navigate to project
cd C:\dev\ai_embedding_offline

# 3. Download dependencies
go mod tidy

# 4. Build main application
go build -o ai_embedding.exe ./cmd/

# 5. Build n-gram tool
go build -o ngram.exe ./cmd/ngram/

# 6. Train model
ngram.exe train

# 7. Start server
ai_embedding.exe serve
```

---

## System Requirements / ความต้องการระบบ

### Minimum / ขั้นต่ำ

- **OS:** Windows 10 / macOS / Linux
- **Go:** Version 1.21 or higher
- **RAM:** 2 GB
- **Disk:** 500 MB free space

### Recommended / แนะนำ

- **OS:** Windows 11 / macOS 12+ / Ubuntu 22.04+
- **Go:** Version 1.21 or higher
- **RAM:** 4 GB+
- **Disk:** 2 GB free space

---

## Step-by-Step Installation / การติดตั้งแบบละเอียด

### Step 1: Install Go

**Windows:**
1. Download from https://go.dev/dl/
2. Run installer
3. Restart terminal

**Verify installation:**
```bash
go version
# Should show: go version go1.21.x ...
```

### Step 2: Clone/Download Project

```bash
# If using git
git clone <repository-url>
cd ai_embedding_offline

# Or extract downloaded zip
cd C:\dev\ai_embedding_offline
```

### Step 3: Install Dependencies

```bash
go mod tidy
```

This downloads:
- `github.com/mattn/go-sqlite3` - SQLite driver
- `gonum.org/v1/gonum` - Math library

### Step 4: Build Executables

```bash
# Main application (server + vector tools)
go build -o ai_embedding.exe ./cmd/

# N-gram training tool
go build -o ngram.exe ./cmd/ngram/
```

**Output files:**
- `ai_embedding.exe` (~7.6 MB)
- `ngram.exe` (~2.8 MB)

### Step 5: Train N-gram Model

```bash
ngram.exe train
```

**Output:**
- `data/ngram_model.json` (~100 KB)

### Step 6: Start Server

```bash
ai_embedding.exe serve
```

**Server starts on:** http://localhost:8080

---

## Verify Installation / ตรวจสอบการติดตั้ง

### Test 1: Check Executables

```bash
# Should show help
ai_embedding.exe
ngram.exe
```

### Test 2: Check API

```bash
# Test n-gram API
curl http://localhost:8080/api/ngram/stats

# Expected response:
{"success":true,"data":{"total_words":3622,...}}
```

### Test 3: Open Web UI

Open in browser:
- http://localhost:8080/intellisense.html
- http://localhost:8080/predict.html

---

## Troubleshooting / แก้ปัญหา

### Problem: "Go is not installed"

**Solution:**
```bash
# Download and install Go
# Windows: https://go.dev/dl/
# Then restart terminal
```

### Problem: "module not found"

**Solution:**
```bash
go mod tidy
```

### Problem: "build failed"

**Solution:**
```bash
# Clean build cache
go clean -cache

# Rebuild
go build -o ai_embedding.exe ./cmd/
```

### Problem: "port 8080 already in use"

**Solution:**
1. Find process using port 8080
2. Kill it or change port in `internal/server/server.go`

### Problem: Model not trained

**Solution:**
```bash
# Train manually
ngram.exe train

# Verify file exists
dir data\ngram_model.json
```

---

## Uninstall / ลบการติดตั้ง

```bash
# Remove executables
del ai_embedding.exe
del ngram.exe

# Remove trained model
del data\ngram_model.json

# Remove database (if created)
del data\vectors.db

# Remove logs
del data\errors.log
```

**Keep:**
- Source code (*.go files)
- Documentation (*.md files)
- Web UI (web/*.html)

---

## Update / อัพเดท

```bash
# Pull latest changes
git pull

# Rebuild
go build -o ai_embedding.exe ./cmd/
go build -o ngram.exe ./cmd/ngram/

# Retrain model (if needed)
ngram.exe train
```

---

## File Structure After Install / โครงสร้างไฟล์หลังติดตั้ง

```
ai_embedding_offline/
├── ai_embedding.exe      ⭐ Main server
├── ngram.exe             ⭐ Training tool
├── setup.bat             ⭐ Setup script
├── start.bat             Quick start script
├── data/
│   ├── ngram_model.json  ⭐ Trained model
│   └── errors.log        Error logs
├── cmd/
├── internal/
├── web/
└── docs/
```

---

## Next Steps / ขั้นตอนถัดไป

After installation:

1. **Start Server:**
   ```bash
   ai_embedding.exe serve
   ```

2. **Open Browser:**
   ```
   http://localhost:8080/intellisense.html
   ```

3. **Start Typing!**
   - Type code or text
   - See AI predictions
   - Press Tab to accept

4. **Read Documentation:**
   - [docs/INDEX.md](docs/INDEX.md) - Documentation map
   - [docs/INTELLISENSE.md](docs/INTELLISENSE.md) - IntelliSense guide
   - [docs/AI_DEV_GUIDE.md](docs/AI_DEV_GUIDE.md) - Developer guide

---

## Support / การสนับสนุน

**Issues:** Check [docs/SOLVED_ERRORS.md](docs/SOLVED_ERRORS.md)

**Documentation:** [docs/INDEX.md](docs/INDEX.md)

**Learning:** [docs/BEGINNER_GUIDE.md](docs/BEGINNER_GUIDE.md)

---

**Installation Time:** ~2-5 minutes  
**Difficulty:** Easy ⭐⭐⭐⭐⭐
