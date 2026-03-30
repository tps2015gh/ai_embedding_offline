# AI Embedding Offline

**Created:** 2025-03-30 23:00 (Thai Time UTC+7)  
**Last Updated:** 2025-03-31 00:18 (Thai Time UTC+7)  
**Total Development Time:** ~2 hours  
**Git Commits:** 20

---

## English Version

### Overview
AI Embedding Offline is a vector-based text similarity system that creates 40-dimensional embeddings from your code and text files. It uses cosine similarity to find related content and provides an interactive 2D visualization.

### What is Vector Embedding? (สำหรับผู้ที่เริ่มศึกษา)

**Vector Embedding** คือการแปลงข้อความให้เป็นตัวเลขชุดหนึ่ง (เวกเตอร์) ที่คอมพิวเตอร์เข้าใจ

**Cosine Similarity** คือวิธีการวัดว่าเวกเตอร์สองอันคล้ายกันแค่ไหน โดยดูจากมุมระหว่างเวกเตอร์

**40 Dimensions** หมายถึง เวกเตอร์แต่ละอันมีตัวเลข 40 ตัว ซึ่งแต่ละตัวแทน "ลักษณะ" ของข้อความ

📖 **New to AI embeddings?** Read our [Beginner's Guide](docs/BEGINNER_GUIDE.md) for simple explanations in English and Thai!

### Features

#### 🔮 AI IntelliSense (NEW!)
- **IDE-style prediction** like VS Code
- Real-time suggestions as you type
- Multi-word phrase predictions
- Keyboard navigation (↑↓ Tab Enter Esc)
- Dark theme UI
- **Try it**: http://localhost:8080/intellisense.html

#### 📊 Text Prediction
- N-gram based next word prediction
- Small model size (~100KB - 10MB)
- Fast predictions (milliseconds)
- Works offline
- See [docs/NGRAM_PREDICTOR.md](docs/NGRAM_PREDICTOR.md)

#### 📊 Vector Embeddings
- 40-dimensional vector embeddings
- Cosine similarity search
- Interactive 2D visualization

### Installation

**Quick Setup (Windows):**
```bash
# Run automated setup
setup.bat
```

**Manual Setup:**
```bash
# 1. Install Go from https://go.dev/dl/

# 2. Download dependencies
go mod tidy

# 3. Build executables
go build -o ai_embedding.exe ./cmd/
go build -o ngram.exe ./cmd/ngram/

# 4. Train model
ngram.exe train

# 5. Start server
ai_embedding.exe serve
```

📖 **Full installation guide:** [docs/INSTALL.md](docs/INSTALL.md)

### Usage

**Option 1: Using the executable (Windows)**

```bash
# PowerShell (use .\ prefix)
.\ngram.exe train
.\ai_embedding.exe serve

# CMD (no prefix needed)
ngram.exe train
ai_embedding.exe serve
```

**Option 2: Using Go commands**

```bash
# 1. Train the n-gram model
go run cmd/ngram/main.go train

# 2. Initialize the vector database (optional, for vector search)
go run cmd/main.go init

# 3. Scan directories (optional, takes time)
go run cmd/main.go scan

# 4. Start the web server
go run cmd/main.go serve
```

Then open:
- **AI IntelliSense**: http://localhost:8080/intellisense.html (Recommended!)
- **Text Predictor**: http://localhost:8080/predict.html
- **Vector Visualizer**: http://localhost:8080/

### Scan Progress

The scan command shows real-time progress:
- 🔍 Scanning: [directory]
- 📄 Files: X | Chunks: Y | Skipped dirs: Z
- ✅ Complete: X files, Y chunks, Z dirs skipped

### Data Storage

- **Database**: `data/vectors.db` (SQLite)
- **Error Log**: `data/errors.log`
- **Scanned Paths**: `c:\dev\`, `C:\Users\admin\Documents`, `C:\Users\admin\Downloads`

**Note**: Large codebases (10K+ files) may take 5-10 minutes to scan.

### Project Structure (Company Organization)

```
ai_embedding_offline/
├── cmd/                    # Main application entry point
│   └── main.go            # Lead Programmer: Application orchestration
├── internal/
│   ├── scanner/           # Data Scanner Agent
│   │   └── scanner.go     # Scans directories for text files
│   ├── embedding/         # Embedding Engine
│   │   └── embedding.go   # Creates 40D vectors with cosine similarity
│   ├── vectorstore/       # Database Agent
│   │   └── vectorstore.go # SQLite storage with optimized queries
│   └── server/            # Server Agent
│       └── server.go      # HTTP API server
├── web/                   # Web Developer
│   └── index.html         # Interactive visualization (zoom, pan, 12px font)
├── data/                  # Database storage
├── docs/                  # Documentation
├── scripts/               # Utility scripts
├── TEAM.md               # Agent registration log
├── README.md             # This file
└── go.mod                # Go module definition
```

### How It Works

1. **Scanning**: The scanner reads all text files from `c:\dev\`, `C:\Users\admin\Documents`, and `C:\Users\admin\Downloads`
2. **Chunking**: Text is split into manageable chunks (500 chars each)
3. **Embedding**: Each chunk is converted to a 40-dimensional vector using hash-based embedding
4. **Storage**: Vectors are stored in SQLite with 2D projection for visualization
5. **Search**: Cosine similarity finds the most similar vectors to your query
6. **Visualization**: Interactive canvas shows all vectors as nodes with 12px labels

### Optimization

The distance calculation is optimized by:
- **Filtering first**: Only vectors with similarity > 0.1 are considered
- **Sorting**: Results are sorted by similarity score (descending)
- **Limiting**: Only top N results are returned

### API Endpoints

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/api/search` | POST | Search similar vectors |
| `/api/vectors` | GET | Get all vectors for visualization |
| `/api/stats` | GET | Get database statistics |
| `/api/suggest` | POST | Get word suggestions |
| `/api/ngram/predict` | POST | Get next word prediction |
| `/api/ngram/predict-phrase` | POST | Get multi-word phrase prediction |
| `/api/ngram/stats` | GET | Get n-gram model statistics |

### Documentation

**📑 Start Here:** [docs/INDEX.md](docs/INDEX.md) - Complete documentation map with reading order

**For AI Developers:**
- **📘 [AI_DEV_GUIDE.md](docs/AI_DEV_GUIDE.md)** - Start here! Complete guide for next developers
- **📗 [INTELLISENSE.md](docs/INTELLISENSE.md)** - IntelliSense feature documentation
- **📙 [NGRAM_PREDICTOR.md](docs/NGRAM_PREDICTOR.md)** - N-gram model guide

**For Users:**
- **📕 [BEGINNER_GUIDE.md](docs/BEGINNER_GUIDE.md)** - Simple explanations (EN/TH)
- **📔 [DATASET_SUMMARY.md](docs/DATASET_SUMMARY.md)** - Dataset statistics

**Technical:**
- **📓 [VECTOR_DB_USAGE.md](docs/VECTOR_DB_USAGE.md)** - Vector database guide
- **📜 [ABSTRACT.md](docs/ABSTRACT.md)** - Technical abstract
- **📋 [SOLVED_ERRORS.md](docs/SOLVED_ERRORS.md)** - Error tracking

**Project:**
- **📝 [TODO.md](TODO.md)** - Future improvements
- **👥 [TEAM.md](TEAM.md)** - Agent registry

### Security

Sensitive data is protected via `.gitignore`:
- Database files (`*.db`)
- Environment files (`*.env`)
- Credentials and keys
- Log files

---

## เวอร์ชันภาษาไทย

### ภาพรวม
AI Embedding Offline คือระบบค้นหาข้อความที่คล้ายกันโดยใช้เวกเตอร์ 40 มิติ สร้างจากไฟล์โค้ดและข้อความของคุณ ใช้ Cosine Similarity ในการหาเนื้อหาที่เกี่ยวข้อง และมีภาพแสดงผล 2D แบบโต้ตอบได้

### เวกเตอร์เอ็มเบดดิ้งคืออะไร?

**เวกเตอร์เอ็มเบดดิ้ง (Vector Embedding)** คือการแปลงข้อความเป็นตัวเลขชุดหนึ่ง (เวกเตอร์) ที่คอมพิวเตอร์เข้าใจ

**Cosine Similarity** คือวิธีการวัดว่าเวกเตอร์สองอันคล้ายกันแค่ไหน โดยดูจากมุมระหว่างเวกเตอร์

**40 มิติ (Dimensions)** หมายถึง เวกเตอร์แต่ละอันมีตัวเลข 40 ตัว ซึ่งแต่ละตัวแทน "ลักษณะ" ของข้อความ

### คุณสมบัติ
- 📁 สแกนไฟล์โค้ดและข้อความจากหลายไดเรกทอรี
- 🔢 สร้างเวกเตอร์ 40 มิติ
- 🔍 ค้นหาเนื้อหาที่คล้ายกันด้วย Cosine Similarity
- 🎨 ภาพแสดงผล 2D แบบโต้ตอบได้ (ซูมและแพน)
- 💡 แสดงคำแนะนำคำถัดไปจากความใกล้เคียงของเวกเตอร์
- ⚡ คำนวณระยะทางอย่างมีประสิทธิภาพด้วยการกรองข้อมูล

### การติดตั้ง

```bash
# ติดตั้ง Go dependencies
go mod tidy

# ติดตั้ง SQLite driver
go get github.com/mattn/go-sqlite3
```

### วิธีใช้

```bash
# 1. เริ่มต้นฐานข้อมูลเวกเตอร์
go run cmd/main.go init

# 2. สแกนไดเรกทอรีและสร้างเอ็มเบดดิ้ง
go run cmd/main.go scan

# 3. เริ่มต้นเว็บเซิร์ฟเวอร์
go run cmd/main.go serve
```

จากนั้นเปิด http://localhost:8080 ในเบราว์เซอร์

### โครงสร้างโปรเจกต์

```
ai_embedding_offline/
├── cmd/                    # จุดเข้าใช้งานหลัก
├── internal/
│   ├── scanner/           # ตัวสแกนข้อมูล
│   ├── embedding/         # ตัวสร้างเวกเตอร์
│   ├── vectorstore/       # ฐานข้อมูล
│   └── server/            # เซิร์ฟเวอร์
├── web/                   # หน้าเว็บ
├── data/                  # ข้อมูลฐานข้อมูล
├── docs/                  # เอกสาร
└── scripts/               # สคริปต์เสริม
```

### วิธีการทำงาน

1. **การสแกน**: อ่านไฟล์ข้อความจาก `c:\dev\`, `C:\Users\admin\Documents`, และ `C:\Users\admin\Downloads`
2. **การแบ่ง**: แบ่งข้อความเป็นชิ้นเล็กๆ (500 ตัวอักษร)
3. **การสร้างเอ็มเบดดิ้ง**: แปลงแต่ละชิ้นเป็นเวกเตอร์ 40 มิติ
4. **การเก็บ**: เก็บเวกเตอร์ใน SQLite พร้อมพิกัด 2D สำหรับแสดงผล
5. **การค้นหา**: ใช้ Cosine Similarity หาเวกเตอร์ที่คล้ายกับคำค้นหา
6. **การแสดงผล**: แสดงเวกเตอร์ทั้งหมดเป็นโหนดพร้อมป้ายชื่อ 12px

### การเพิ่มประสิทธิภาพ

การคำนวณระยะทางได้รับการปรับปรุงโดย:
- **กรองก่อน**: พิจารณาเฉพาะเวกเตอร์ที่มีความคล้ายคลึง > 0.1
- **เรียงลำดับ**: เรียงผลลัพธ์ตามคะแนนความคล้ายคลึง (จากมากไปน้อย)
- **จำกัด**: ส่งกลับเฉพาะผลลัพธ์ Top N

### คำอธิบายสำหรับผู้เริ่มต้น

**AI Embedding** คืออะไร?
- คือการทำให้คอมพิวเตอร์ "เข้าใจ" ข้อความโดยแปลงเป็นตัวเลข

**Vector Similarity** คืออะไร?
- คือการวัดว่าข้อความสองข้อความคล้ายกันแค่ไหน

**ทำไมต้องใช้ 40 มิติ?**
- มิติ越多 หมายถึงสามารถบอกลักษณะของข้อความได้ละเอียด越多
- 40 มิติเป็นจุดที่สมดุลระหว่างความละเอียดและความเร็ว

### การใช้งานหน้าเว็บ

1. พิมพ์ข้อความในช่องค้นหา
2. กด "Search Similar" หรือกด Enter
3. ระบบจะไฮไลท์โหนดที่คล้ายกัน
4. ดูคำแนะนำคำถัดไปที่ด้านล่าง
5. คลิกที่โหนดเพื่อดูรายละเอียด
6. ใช้เมาส์เพื่อแพน (ลาก) และซูม (ล้อเมาส์)

### ทีมพัฒนา (AI Agents)

| Agent | บทบาท | หน้าที่ |
|-------|-------|--------|
| Lead Programmer | สถาปนิกหลัก | ออกแบบระบบและประสานงาน |
| Web Developer | นักพัฒนาเว็บ | สร้างหน้า HTML และ visualization |
| Server Developer | นักพัฒนาเซิร์ฟเวอร์ | สร้าง API และ backend |
| Go Developer | นักพัฒนา Go | เขียนโค้ด Go และปรับปรุงประสิทธิภาพ |
| Data Scanner | ตัวสแกนข้อมูล | อ่านและประมวลผลไฟล์ |
| UX/UI Designer | นักออกแบบ | ออกแบบประสบการณ์ผู้ใช้ |

### License

MIT License - Free to use and modify
