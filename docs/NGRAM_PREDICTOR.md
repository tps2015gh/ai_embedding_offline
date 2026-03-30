# N-gram Text Predictor / การทำนายข้อความด้วย N-gram

## English Version

### What is N-gram?

**N-gram** is a simple but effective text prediction technique that looks at sequences of N words to predict what comes next.

**How it works:**
1. **Unigram (1-gram)**: Counts individual words
   - "func" appears 10 times
   - "if" appears 8 times

2. **Bigram (2-gram)**: Counts word pairs
   - "func main" appears 5 times
   - "if err" appears 3 times

3. **Trigram (3-gram)**: Counts word triplets
   - "if err != nil" appears 2 times

### Prediction Example

```
Input: "func"

Looking at training data:
- "func main()" → suggests "main"
- "func New()" → suggests "new"
- "func Test()" → suggests "test"

Output: ["main", "new", "test"]
```

### Why N-gram is Better for This Use Case

| Feature | Hash Embedding | N-gram |
|---------|---------------|--------|
| Semantic meaning | ❌ No | ✅ Yes (contextual) |
| Next word prediction | ❌ Poor | ✅ Excellent |
| Model size | Large (GB) | Small (MB) |
| Training time | Hours | Seconds |
| Query speed | Slow | Fast |
| Easy to understand | ❌ Complex | ✅ Simple |

### How to Use

**1. Train the model:**
```bash
# Using the demo tool
ngram.exe train

# Or via Go
go run cmd/ngram/main.go train
```

**2. Start the server:**
```bash
ai_embedding.exe serve
```

**3. Open the web UI:**
```
http://localhost:8080/predict.html
```

**4. Or use the API:**
```bash
curl -X POST http://localhost:8080/api/ngram/predict \
  -H "Content-Type: application/json" \
  -d '{"text": "func", "limit": 5}'
```

### Model Storage

- **File**: `data/ngram_model.json`
- **Size**: ~100KB - 10MB (depending on training data)
- **Format**: JSON

### Training Data

The model learns from:
1. Sample code snippets (built-in)
2. Files in `c:\dev\ai_embedding_offline`
3. Any text you add

### API Endpoints

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/api/ngram/predict` | POST | Get next word suggestions |
| `/api/ngram/stats` | GET | Get model statistics |

**Request:**
```json
{
  "text": "if err",
  "limit": 5
}
```

**Response:**
```json
{
  "success": true,
  "data": [
    {"word": "nil", "score": 15.0},
    {"word": "!= nil", "score": 10.0},
    {"word": "return", "score": 5.0}
  ]
}
```

---

## เวอร์ชันภาษาไทย

### N-gram คืออะไร?

**N-gram** เป็นเทคนิคการทำนายข้อความที่ง่ายแต่มีประสิทธิภาพ โดยดูจากลำดับของ N คำเพื่อทำนายคำถัดไป

**การทำงาน:**
1. **Unigram (1-gram)**: นับคำแต่ละคำ
2. **Bigram (2-gram)**: นับคู่คำ
3. **Trigram (3-gram)**: นับกลุ่มสามคำ

### ตัวอย่างการทำนาย

```
อินพุต: "func"

ดูจากข้อมูลฝึก:
- "func main()" → แนะนำ "main"
- "func New()" → แนะนำ "new"
- "func Test()" → แนะนำ "test"

เอาต์พุต: ["main", "new", "test"]
```

### ทำไม N-gram ดีกว่าสำหรับกรณีนี้

| คุณสมบัติ | Hash Embedding | N-gram |
|-----------|---------------|--------|
| ความหมายเชิงบริบท | ❌ ไม่มี | ✅ มี |
| ทำนายคำถัดไป | ❌ แย่ | ✅ ยอดเยี่ยม |
| ขนาดโมเดล | ใหญ่ (GB) | เล็ก (MB) |
| เวลาฝึก | ชั่วโมง | วินาที |
| ความเร็ว query | ช้า | เร็ว |
| เข้าใจง่าย | ❌ ยาก | ✅ ง่าย |

### วิธีใช้

**1. ฝึกโมเดล:**
```bash
ngram.exe train
```

**2. เริ่มเซิร์ฟเวอร์:**
```bash
ai_embedding.exe serve
```

**3. เปิดหน้าเว็บ:**
```
http://localhost:8080/predict.html
```

### การเก็บโมเดล

- **ไฟล์**: `data/ngram_model.json`
- **ขนาด**: ~100KB - 10MB
- **รูปแบบ**: JSON

### ตัวอย่างการใช้งานจริง

**สำหรับโปรแกรมเมอร์:**
```
พิมพ์: "func "
แนะนำ: "main", "New", "Test"

พิมพ์: "if err "
แนะนำ: "nil", "!= nil", "return"

พิมพ์: "SELECT * FROM "
แนะนำ: "users", "vectors", "WHERE"
```

**สำหรับเขียนเอกสาร:**
```
พิมพ์: "The quick brown "
แนะนำ: "fox", "dog", "cat"

พิมพ์: "In conclusion, "
แนะนำ: "we", "this", "the"
```

### ข้อจำกัด

1. **ต้องการข้อมูลฝึก**: ยิ่งมีข้อมูลมาก ยิ่งทำนายแม่นยำ
2. **ไม่เข้าใจความหมายลึก**: ดูแค่รูปแบบ ไม่ใช่ความหมายจริง
3. **คำใหม่ไม่ได้**: ถ้าคำไม่เคยเห็น จะทำนายไม่ได้

### การปรับปรุง

1. เพิ่มข้อมูลฝึกมากขึ้น
2. ใช้ 4-gram, 5-gram สำหรับบริบทที่ยาวขึ้น
3. รวมกับเทคนิคอื่น (เช่น Word2Vec, BERT)

---

## Quick Start / เริ่มต้นอย่างรวดเร็ว

```bash
# 1. Train
ngram.exe train

# 2. Serve
ai_embedding.exe serve

# 3. Open browser
# http://localhost:8080/predict.html
```

**That's it!** Start typing and see predictions.
