# Vector Database Usage Guide

**Created:** 2025-03-30 23:05 (Thai Time UTC+7)  
**Last Updated:** 2025-03-30 23:20 (Thai Time UTC+7)  
**Development Time:** ~15 minutes  

---

## English Version

### Overview
This system uses SQLite as a lightweight vector database. Vectors are stored as JSON-encoded arrays with metadata.

### Database Schema

```sql
CREATE TABLE vectors (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    text TEXT NOT NULL,           -- Original text content
    embedding TEXT NOT NULL,      -- JSON array of 40 floats
    position_x REAL DEFAULT 0,    -- 2D projection X coordinate
    position_y REAL DEFAULT 0     -- 2D projection Y coordinate
);

CREATE INDEX idx_vectors_text ON vectors(text);
```

### Operations

#### 1. Initialize Database
```bash
go run cmd/main.go init
```
Creates `data/vectors.db` with the schema above.

#### 2. Store Vectors
```go
import "ai_embedding_offline/internal/vectorstore"

vectors := []embedding.Vector{
    {Text: "Hello World", Embedding: []float64{...}},
}

err := vectorstore.StoreVectors(vectors)
```

#### 3. Search Similar Vectors
```go
results, err := vectorstore.SearchSimilar("query text", 20)
// Returns top 20 most similar vectors
```

#### 4. Get All Vectors (for visualization)
```go
vectors, err := vectorstore.GetAllVectors()
```

#### 5. Get Statistics
```go
stats, err := vectorstore.GetVectorStats()
// Returns: total_vectors, center_x, center_y
```

### API Usage

#### Search Endpoint
```bash
curl -X POST http://localhost:8080/api/search \
  -H "Content-Type: application/json" \
  -d '{"query": "your text", "limit": 20}'
```

Response:
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "text": "similar content",
      "position_x": 12.5,
      "position_y": -3.2
    }
  ]
}
```

#### Get All Vectors
```bash
curl http://localhost:8080/api/vectors
```

#### Get Statistics
```bash
curl http://localhost:8080/api/stats
```

### Optimization Details

**Distance Calculation Optimization:**

1. **Filter First**: Only process vectors with similarity > 0.1
   ```go
   if scoredVecs[i].score > 0.1 {
       result = append(result, scoredVecs[i].record)
   }
   ```

2. **Sort by Score**: Descending order for quick top-N selection
   ```go
   sort.Slice(scoredVecs, func(i, j int) bool {
       return scoredVecs[i].score > scoredVecs[j].score
   })
   ```

3. **Limit Results**: Return only top N results
   ```go
   for i := 0; i < len(scoredVecs) && i < limit; i++ {
       // ...
   }
   ```

### 2D Projection

High-dimensional vectors (40D) are projected to 2D for visualization:

```go
func project2D(vec []float64) (float64, float64) {
    x := vec[0] * 100
    y := vec[1] * 100
    
    // Add contribution from other dimensions
    for i := 2; i < len(vec) && i < 10; i++ {
        x += vec[i] * float64(10-i) * 10
        y += vec[i] * float64(i-1) * 10
    }
    
    return x, y
}
```

---

## เวอร์ชันภาษาไทย

### ภาพรวม
ระบบนี้ใช้ SQLite เป็นฐานข้อมูลเวกเตอร์แบบเบา เวกเตอร์ถูกเก็บเป็นอาเรย์ที่เข้ารหัส JSON พร้อมข้อมูลเมตา

### โครงสร้างฐานข้อมูล

```sql
CREATE TABLE vectors (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    text TEXT NOT NULL,           -- เนื้อหาข้อความเดิม
    embedding TEXT NOT NULL,      -- อาเรย์ JSON ของเลขทศนิยม 40 ตัว
    position_x REAL DEFAULT 0,    -- พิกัด X สำหรับแสดงผล 2D
    position_y REAL DEFAULT 0     -- พิกัด Y สำหรับแสดงผล 2D
);

CREATE INDEX idx_vectors_text ON vectors(text);
```

### การดำเนินการ

#### 1. เริ่มต้นฐานข้อมูล
```bash
go run cmd/main.go init
```
สร้าง `data/vectors.db` พร้อมโครงสร้างข้างต้น

#### 2. เก็บเวกเตอร์
```go
import "ai_embedding_offline/internal/vectorstore"

vectors := []embedding.Vector{
    {Text: "Hello World", Embedding: []float64{...}},
}

err := vectorstore.StoreVectors(vectors)
```

#### 3. ค้นหาเวกเตอร์ที่คล้ายกัน
```go
results, err := vectorstore.SearchSimilar("query text", 20)
// ส่งคืนเวกเตอร์ที่คล้ายกัน 20 อันดับแรก
```

#### 4. ดึงเวกเตอร์ทั้งหมด (สำหรับแสดงผล)
```go
vectors, err := vectorstore.GetAllVectors()
```

#### 5. ดึงสถิติ
```go
stats, err := vectorstore.GetVectorStats()
// ส่งคืน: total_vectors, center_x, center_y
```

### การใช้ API

#### Endpoint ค้นหา
```bash
curl -X POST http://localhost:8080/api/search \
  -H "Content-Type: application/json" \
  -d '{"query": "your text", "limit": 20}'
```

การตอบกลับ:
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "text": "similar content",
      "position_x": 12.5,
      "position_y": -3.2
    }
  ]
}
```

#### ดึงเวกเตอร์ทั้งหมด
```bash
curl http://localhost:8080/api/vectors
```

#### ดึงสถิติ
```bash
curl http://localhost:8080/api/stats
```

### รายละเอียดการเพิ่มประสิทธิภาพ

**การเพิ่มประสิทธิภาพการคำนวณระยะทาง:**

1. **กรองก่อน**: ประมวลผลเฉพาะเวกเตอร์ที่มีความคล้ายคลึง > 0.1
2. **เรียงตามคะแนน**: เรียงจากมากไปน้อยสำหรับการเลือก Top-N อย่างรวดเร็ว
3. **จำกัดผลลัพธ์**: ส่งกลับเฉพาะผลลัพธ์ N อันดับแรก

### การแปลง 2D

เวกเตอร์มิติสูง (40 มิติ) ถูกแปลงเป็น 2D สำหรับการแสดงผล:

```go
func project2D(vec []float64) (float64, float64) {
    x := vec[0] * 100
    y := vec[1] * 100
    
    // เพิ่มส่วนร่วมจากมิติอื่นๆ
    for i := 2; i < len(vec) && i < 10; i++ {
        x += vec[i] * float64(10-i) * 10
        y += vec[i] * float64(i-1) * 10
    }
    
    return x, y
}
```

### File Locations / ตำแหน่งไฟล์

| File | Purpose | Location |
|------|---------|----------|
| vectors.db | SQLite database | `data/vectors.db` |
| main.go | Application entry | `cmd/main.go` |
| index.html | Web interface | `web/index.html` |

### Troubleshooting

**Problem**: Database not found
```
Solution: Run `go run cmd/main.go init` first
```

**Problem**: No vectors in database
```
Solution: Run `go run cmd/main.go scan` to populate data
```

**Problem**: Slow search performance
```
Solution: Ensure index exists: `CREATE INDEX idx_vectors_text ON vectors(text)`
```
