# AI Developer Guide / คู่มือสำหรับนักพัฒนา AI

**Created:** 2025-03-30 23:45 (Thai Time UTC+7)  
**Last Updated:** 2025-03-31 00:05 (Thai Time UTC+7)  
**Development Time:** ~20 minutes  

---

## For Next AI Developer / สำหรับนักพัฒนา AI คนต่อไป

### Project Overview / ภาพรวมโปรเจกต์

This is an **AI learning project** that demonstrates text prediction using multiple approaches:

1. **N-gram Model** (Working ✅) - Simple, fast, practical
2. **Vector Embeddings** (Experimental ⚠️) - Hash-based, not semantic
3. **IntelliSense UI** (Working ✅) - IDE-style prediction interface

**Important:** This is a TEST/LEARNING project, NOT for production use.

---

## Architecture / สถาปัตยกรรม

```
ai_embedding_offline/
├── cmd/
│   ├── main.go              # Main app (vector embeddings + server)
│   └── ngram/
│       └── main.go          # N-gram training CLI
├── internal/
│   ├── ngram/
│   │   └── ngram.go         # N-gram model (unigram, bigram, trigram)
│   ├── embedding/
│   │   └── embedding.go     # Hash-based embeddings (40D)
│   ├── vectorstore/
│   │   └── vectorstore.go   # SQLite storage + cosine similarity
│   ├── scanner/
│   │   └── scanner.go       # File scanner with progress
│   ├── server/
│   │   └── server.go        # HTTP server + API endpoints
│   └── logger/
│       └── logger.go        # Error logging system
├── web/
│   ├── intellisense.html    # IDE-style prediction UI ⭐
│   ├── predict.html         # Simple prediction UI
│   └── index.html           # Vector visualization
├── data/
│   ├── ngram_model.json     # Trained n-gram model (~100KB)
│   ├── vectors.db           # Vector database (optional, ~GB)
│   └── errors.log           # Error logs
└── docs/
    ├── AI_DEV_GUIDE.md      # This file ⭐
    ├── NGRAM_PREDICTOR.md   # N-gram documentation
    ├── INTELLISENSE.md      # IntelliSense documentation
    ├── BEGINNER_GUIDE.md    # Beginner explanations
    └── DATASET_SUMMARY.md   # Dataset statistics
```

---

## Key Concepts / แนวคิดหลัก

### 1. N-gram Model (RECOMMENDED)

**What it is:** Statistical language model that predicts next word based on previous N-1 words.

**Implementation:**
```go
// internal/ngram/ngram.go

type NGramModel struct {
    Unigrams  map[string]int            // Word frequencies
    Bigrams   map[string]map[string]int // Word → next word
    Trigrams  map[string]map[string]int // Word1 Word2 → next word
}
```

**Training:**
```go
model := ngram.NewModel()
model.Train("func main() { fmt.Println(\"Hello\") }")
model.Train("if err != nil { return err }")
// ... train on more text

model.Save("data/ngram_model.json")
```

**Prediction:**
```go
// Single word
predictions := model.Predict("func", 5)
// Result: [{Word: "main", Score: 10.0}, ...]

// Multi-word phrase
phrases := model.PredictPhrase("if err", 3, 5)
// Result: [{Phrase: "!= nil { return", Score: 45.0}, ...]
```

**Why N-gram is Good for Learning:**
- ✅ Simple to understand
- ✅ Fast training (seconds)
- ✅ Small model size (~100KB)
- ✅ Works offline
- ✅ No external dependencies
- ❌ Limited context (only looks at last 2-3 words)
- ❌ No semantic understanding

### 2. Vector Embeddings (EXPERIMENTAL)

**What it is:** Converts text to high-dimensional vectors for similarity search.

**Current Implementation:** Hash-based (NOT semantic)
```go
// internal/embedding/embedding.go

func generateEmbedding(text string, dimensions int) []float64 {
    hash := sha256.Sum256([]byte(text))
    // Convert hash to 40-dimensional vector
    // L2 normalize
    return embedding
}

func CosineSimilarity(a, b []float64) float64 {
    // Calculate cosine of angle between vectors
    return dotProduct / (normA * normB)
}
```

**Problem:** Hash-based embeddings don't capture meaning!
- "cat" and "feline" → NOT similar (different hashes)
- "cat" and "cat" → Similar (same hash)

**Better Approach for Production:**
```python
# Use sentence-transformers
from sentence_transformers import SentenceTransformer
model = SentenceTransformer('all-MiniLM-L6-v2')
embeddings = model.encode(["cat", "feline"])
# Now "cat" and "feline" are similar!
```

### 3. IntelliSense UI

**How it works:**
```
User types → Debounce (150ms) → API call → N-gram lookup → Display popup
```

**Key features:**
- Real-time prediction
- Keyboard navigation
- Multi-word phrases
- VS Code-like styling

---

## API Reference / API อ้างอิง

### N-gram Endpoints

#### POST /api/ngram/predict
Predict next word(s).

**Request:**
```json
{
  "text": "func",
  "limit": 5
}
```

**Response:**
```json
{
  "success": true,
  "data": [
    {"word": "main", "score": 10.0},
    {"word": "New", "score": 4.0}
  ]
}
```

#### POST /api/ngram/predict-phrase
Predict multi-word phrases.

**Request:**
```json
{
  "text": "if err",
  "maxWords": 3,
  "limit": 5
}
```

**Response:**
```json
{
  "success": true,
  "data": [
    {"phrase": "!= nil { return", "score": 45.0},
    {"phrase": "!= nil { log", "score": 30.0}
  ]
}
```

#### GET /api/ngram/stats
Get model statistics.

**Response:**
```json
{
  "success": true,
  "data": {
    "total_words": 3622,
    "unigrams": 850,
    "bigrams": 1240,
    "trigrams": 980
  }
}
```

---

## How to Extend / วิธีขยายระบบ

### 1. Add Better Embeddings

Replace hash-based with ML-based:

```go
// internal/embedding/embedding.go

// Option 1: Use Python bridge
import "github.com/go-python/gopy"

func CreateEmbeddings(texts []string) ([]Vector, error) {
    // Call Python sentence-transformers
    // Return real semantic embeddings
}

// Option 2: Use ONNX runtime
import "github.com/owulveryck/onnx-go"

func CreateEmbeddings(text text) ([]float64, error) {
    // Load pre-trained model (MiniLM, BERT)
    // Run inference
    // Return embeddings
}
```

### 2. Add 4-gram, 5-gram Support

```go
// internal/ngram/ngram.go

type NGramModel struct {
    // ... existing fields ...
    Fourgrams map[string]map[string]int  // Word1 Word2 Word3 → next
    Fivegrams map[string]map[string]int  // Word1 Word2 Word3 Word4 → next
}

func (m *NGramModel) Train(text string) {
    // ... existing code ...
    
    // Add 4-grams
    for i := 0; i < len(words)-3; i++ {
        key := words[i] + " " + words[i+1] + " " + words[i+2]
        if m.Fourgrams[key] == nil {
            m.Fourgrams[key] = make(map[string]int)
        }
        m.Fourgrams[key][words[i+3]]++
    }
}
```

### 3. Add Smart Context Detection

Detect if user is typing code, comments, or strings:

```go
// web/intellisense.html

function detectContext(text) {
    const lastLine = text.split('\n').pop();
    
    if (lastLine.trim().startsWith('//')) {
        return 'comment';
    }
    if (lastLine.includes('"') || lastLine.includes("'")) {
        return 'string';
    }
    return 'code';
}

// Load different models per context
```

### 4. Add User Feedback Loop

Track which suggestions users accept:

```go
// internal/ngram/ngram.go

type NGramModel struct {
    // ... existing fields ...
    AcceptedCounts map[string]int  // Track accepted suggestions
}

func (m *NGramModel) RecordAcceptance(word string) {
    m.AcceptedCounts[word]++
    // Boost score for frequently accepted words
}

func (m *NGramModel) Predict(text string) []Prediction {
    // ... existing prediction ...
    
    // Boost by acceptance count
    for _, pred := range predictions {
        if count, ok := m.AcceptedCounts[pred.Word]; ok {
            pred.Score += float64(count) * 0.5
        }
    }
}
```

### 5. Add Model Compression

For large models:

```go
// internal/ngram/ngram.go

func (m *NGramModel) Compress() {
    // Remove low-frequency n-grams
    for word, count := range m.Unigrams {
        if count < 3 {
            delete(m.Unigrams, word)
        }
    }
    // Quantize scores (float64 → float32)
    // Use more efficient data structures
}
```

---

## Performance Optimization / การปรับปรุงประสิทธิภาพ

### Current Performance

| Operation | Time | Memory |
|-----------|------|--------|
| Train (100 texts) | ~2s | ~10MB |
| Predict (1 word) | ~5ms | ~1MB |
| Predict (phrase) | ~20ms | ~2MB |
| Model load | ~50ms | ~5MB |

### Optimization Ideas

1. **Use Trie for faster lookups**
```go
type TrieNode struct {
    children map[rune]*TrieNode
    count    int
}
```

2. **Cache frequent predictions**
```go
var cache = make(map[string][]Prediction)

func Predict(text string) []Prediction {
    if cached, ok := cache[text]; ok {
        return cached
    }
    // ... compute ...
    cache[text] = result
    return result
}
```

3. **Parallel training**
```go
func TrainParallel(texts []string) {
    var wg sync.WaitGroup
    for _, text := range texts {
        wg.Add(1)
        go func(t string) {
            defer wg.Done()
            model.Train(t)
        }(text)
    }
    wg.Wait()
}
```

---

## Testing / การทดสอบ

### Unit Tests

```go
// internal/ngram/ngram_test.go

package ngram

import "testing"

func TestTrain(t *testing.T) {
    model := NewModel()
    model.Train("func main() {}")
    
    if model.TotalWords != 3 {
        t.Errorf("Expected 3 words, got %d", model.TotalWords)
    }
}

func TestPredict(t *testing.T) {
    model := NewModel()
    model.Train("func main() {}")
    model.Train("func New() {}")
    
    preds := model.Predict("func", 5)
    
    if len(preds) == 0 {
        t.Error("Expected predictions, got none")
    }
}
```

### Integration Tests

```bash
# Test training
go run cmd/ngram/main.go train

# Test prediction
curl -X POST http://localhost:8080/api/ngram/predict \
  -H "Content-Type: application/json" \
  -d '{"text": "func", "limit": 5}'

# Expected: {"success": true, "data": [...]}
```

---

## Troubleshooting / แก้ปัญหา

### Problem: Predictions are slow

**Solution:**
1. Check model size: `ls -lh data/ngram_model.json`
2. If > 10MB, consider compression
3. Add caching layer

### Problem: Predictions are not relevant

**Solution:**
1. Train on more relevant data
2. Increase n-gram order (4-gram, 5-gram)
3. Add context detection

### Problem: UI doesn't show predictions

**Solution:**
1. Check browser console for errors
2. Verify server is running: `curl http://localhost:8080/api/ngram/stats`
3. Check CORS headers

### Problem: Model won't save/load

**Solution:**
1. Check file permissions: `ls -la data/`
2. Verify JSON is valid: `cat data/ngram_model.json | python -m json.tool`
3. Check disk space: `df -h`

---

## Resources / ทรัพยากร

### Learn More About N-grams

- [Wikipedia: N-gram](https://en.wikipedia.org/wiki/N-gram)
- [Stanford NLP: N-grams](https://web.stanford.edu/~jurafsky/slp3/3.pdf)
- [NLTK Book: N-grams](https://www.nltk.org/book/ch02.html)

### Learn More About Text Prediction

- [The Unreasonable Effectiveness of Recurrent Neural Networks](http://karpathy.github.io/2015/05/21/rnn-effectiveness/)
- [Attention Is All You Need (Transformer paper)](https://arxiv.org/abs/1706.03762)

### Go Resources

- [Effective Go](https://golang.org/doc/effective_go)
- [Go by Example](https://gobyexample.com/)

---

## Quick Start for Next Dev / เริ่มต้นอย่างรวดเร็ว

```bash
# 1. Clone and build
cd C:\dev\ai_embedding_offline
go build -o ai_embedding.exe ./cmd/

# 2. Train model
ngram.exe train

# 3. Start server
ai_embedding.exe serve

# 4. Test API
curl http://localhost:8080/api/ngram/stats
curl -X POST http://localhost:8080/api/ngram/predict \
  -H "Content-Type: application/json" \
  -d '{"text": "func", "limit": 5}'

# 5. Open UI
# http://localhost:8080/intellisense.html
```

---

## Contact / ติดต่อ

This project was created by Qwen-2.5-Coder on 2025-03-30.

For questions or improvements, check:
- `TODO.md` - Future improvements
- `TEAM.md` - Agent registration
- `docs/` - All documentation

---

**Good luck with your AI journey! 🚀**
**โชคดีกับการเดินทาง AI ของคุณ! 🚀**
