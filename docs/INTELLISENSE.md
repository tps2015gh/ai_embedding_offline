# AI IntelliSense - IDE-style Code Prediction

**Created:** 2025-03-30 23:55 (Thai Time UTC+7)  
**Last Updated:** 2025-03-31 00:10 (Thai Time UTC+7)  
**Development Time:** ~15 minutes  

---

## English Version

### What is AI IntelliSense?

AI IntelliSense is an **IDE-style code prediction feature** that shows suggestions as you type, similar to VS Code's IntelliSense or ReSharper.

### Features

✅ **Real-time Prediction**: Shows suggestions after every keystroke (150ms debounce)
✅ **Multi-word Phrases**: Predicts complete phrases, not just single words
✅ **Keyboard Navigation**: Use ↑↓ arrows, Tab/Enter to accept, Esc to close
✅ **VS Code-like UI**: Dark theme, popup styling, type icons
✅ **150ms Response**: Fast predictions without lag

### How to Use

**1. Start the server:**
```bash
ai_embedding.exe serve
```

**2. Open IntelliSense:**
```
http://localhost:8080/intellisense.html
```

**3. Start typing!**
- Type `func` → See suggestions like `main() {`, `NewModel()`, etc.
- Type `if` → See `err != nil`, `true`, etc.
- Type `SELECT` → See `* FROM`, `COUNT`, etc.

### Keyboard Shortcuts

| Key | Action |
|-----|--------|
| `↑` `↓` | Navigate suggestions |
| `Tab` or `Enter` | Accept selected suggestion |
| `Esc` | Close suggestions |

### API Endpoints

**Single Word Prediction:**
```bash
POST /api/ngram/predict
{
  "text": "func",
  "limit": 5
}

Response:
{
  "success": true,
  "data": [
    {"word": "main", "score": 10.0},
    {"word": "New", "score": 4.0}
  ]
}
```

**Multi-word Phrase Prediction:**
```bash
POST /api/ngram/predict-phrase
{
  "text": "if err",
  "maxWords": 3,
  "limit": 5
}

Response:
{
  "success": true,
  "data": [
    {"phrase": "!= nil { return", "score": 45.0},
    {"phrase": "!= nil { log", "score": 30.0}
  ]
}
```

### Example Predictions

| Input | Predictions |
|-------|-------------|
| `func` | `main() {`, `NewModel() *NGramModel`, `Test()` |
| `if err` | `!= nil { return err }`, `!= nil { log.Fatal }` |
| `SELECT` | `* FROM users`, `COUNT(*) FROM` |
| `docker` | `run -p`, `build -t`, `compose up` |
| `git` | `commit -m`, `add .`, `push origin` |

### How It Works

```
User types: "f"
     ↓
Debounce (150ms wait)
     ↓
API Call: /api/ngram/predict-phrase
     ↓
N-gram Model chains predictions:
  - "f" → "func" (bigram)
  - "func" → "main" (trigram)
  - "func main" → "()" (trigram)
     ↓
Result: "func main() { ... }"
     ↓
Display in popup with score
     ↓
User presses Tab → Insert prediction
```

---

## เวอร์ชันภาษาไทย

### AI IntelliSense คืออะไร?

AI IntelliSense เป็น **ฟีเจอร์ทำนายโค้ดสไตล์ IDE** ที่แสดงคำแนะนำขณะพิมพ์ คล้ายกับ VS Code IntelliSense หรือ ReSharper

### คุณสมบัติ

✅ **การทำนายแบบเรียลไทม์**: แสดงคำแนะนำหลังทุกการกดปุ่ม (รอ 150ms)
✅ **วลีหลายคำ**: ทำนายทั้งวลี ไม่ใช่แค่คำเดียว
✅ **นำทางด้วยคีย์บอร์ด**: ใช้ ↑↓, Tab/Enter ยอมรับ, Esc ปิด
✅ **UI แบบ VS Code**: ธีมมืด, popup, ไอคอนประเภท
✅ **ตอบสนอง 150ms**: เร็วไม่กระตุก

### วิธีใช้

**1. เริ่มเซิร์ฟเวอร์:**
```bash
ai_embedding.exe serve
```

**2. เปิด IntelliSense:**
```
http://localhost:8080/intellisense.html
```

**3. เริ่มพิมพ์!**
- พิมพ์ `func` → ดูคำแนะนำเช่น `main() {`, `NewModel()`
- พิมพ์ `if` → ดู `err != nil`, `true`
- พิมพ์ `SELECT` → ดู `* FROM`, `COUNT`

### คีย์ลัด

| คีย์ | การกระทำ |
|------|---------|
| `↑` `↓` | เลื่อนคำแนะนำ |
| `Tab` หรือ `Enter` | ยอมรับคำแนะนำที่เลือก |
| `Esc` | ปิดคำแนะนำ |

### ตัวอย่างการทำนาย

| อินพุต | คำแนะนำ |
|--------|---------|
| `func` | `main() {`, `NewModel() *NGramModel`, `Test()` |
| `if err` | `!= nil { return err }`, `!= nil { log.Fatal }` |
| `SELECT` | `* FROM users`, `COUNT(*) FROM` |
| `docker` | `run -p`, `build -t`, `compose up` |
| `git` | `commit -m`, `add .`, `push origin` |

### วิธีการทำงาน

```
ผู้ใช้พิมพ์: "f"
     ↓
รอ 150ms (debounce)
     ↓
เรียก API: /api/ngram/predict-phrase
     ↓
โมเดล N-gram ทำนายต่อเนื่อง:
  - "f" → "func" (bigram)
  - "func" → "main" (trigram)
  - "func main" → "()" (trigram)
     ↓
ผลลัพธ์: "func main() { ... }"
     ↓
แสดงใน popup พร้อมคะแนน
     ↓
ผู้ใช้กด Tab → แทรกคำแนะนำ
```

---

## Technical Details

### Debounce Timing

- **150ms**: Fast enough for real-time, slow enough to avoid spam
- Configurable in `web/intellisense.html`

### Phrase Generation

```go
// Chains predictions to form phrases
for i := 0; i < maxWords; i++ {
    // Find next word using trigrams
    // Fallback to bigrams
    // Append to phrase
}
```

### Response Time

| Operation | Time |
|-----------|------|
| API call | ~10-50ms |
| N-gram lookup | ~1-5ms |
| UI render | ~5-10ms |
| **Total** | **~20-70ms** |

---

## Screenshots

### IntelliSense Popup
```
┌─────────────────────────────────┐
│ 📝 main() {              45.0   │
│ 🔧 NewModel() *NGram     30.0   │
│ ⚡ Test()                25.0   │
│ 📄 init()                20.0   │
└─────────────────────────────────┘
```

### Status Bar
```
📊 3,622 words trained  ⌨️ 45 chars  🟢 Ready
```

---

## Comparison with Other Tools

| Feature | AI IntelliSense | VS Code | Vim Complete |
|---------|----------------|---------|--------------|
| Offline | ✅ Yes | ❌ Needs LSP | ✅ Yes |
| Custom training | ✅ Yes | ⚠️ Complex | ⚠️ Manual |
| Multi-word | ✅ Yes | ❌ Single | ❌ Single |
| Language agnostic | ✅ Yes | ⚠️ Per-lang | ⚠️ Per-lang |
| Model size | ~100 KB | ~500 MB | ~50 MB |

---

## Future Improvements

- [ ] Ghost text (inline prediction like GitHub Copilot)
- [ ] Multiple model support (different models per file type)
- [ ] Smart context detection (code vs comments vs strings)
- [ ] User feedback loop (accept/reject tracking)
- [ ] Model auto-save on training
