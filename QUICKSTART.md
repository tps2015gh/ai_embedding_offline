# Quick Start Guide / คู่มือเริ่มต้นอย่างรวดเร็ว

**Updated:** 2025-03-31 00:30 (Thai Time UTC+7)

---

## Step 1: Train the Model (if not done)

```powershell
.\ngram.exe train
```

**Expected output:**
```
📚 Training n-gram model...
✅ Model trained! Total words: 68,266
   Model saved to: data/ngram_model.json
```

---

## Step 2: Start the Server

```powershell
.\ai_embedding.exe serve
```

**Expected output:**
```
2026/03/31 00:30:00 Starting web server on :8080...
2026/03/31 00:30:00 Server starting on :8080
```

**Server is now running!** ✅

---

## Step 3: Open in Browser

Choose one:

### 🎯 AI IntelliSense (Recommended!)
```
http://localhost:8080/intellisense.html
```
- IDE-style prediction
- Type code, get suggestions
- Press Tab to accept

### 📝 Simple Text Predictor
```
http://localhost:8080/predict.html
```
- Simple interface
- Click suggestions

### 📊 Vector Visualizer (if you have vector data)
```
http://localhost:8080/
```
- 2D visualization
- Zoom and pan

---

## Test the API

```powershell
# Get model stats
curl http://localhost:8080/api/ngram/stats

# Get predictions
curl -X POST http://localhost:8080/api/ngram/predict `
  -H "Content-Type: application/json" `
  -d '{"text": "func", "limit": 5}'
```

---

## Stop the Server

Press `Ctrl+C` in the terminal.

---

## Troubleshooting

### "ngram.exe: The term is not recognized"

Use `.\` prefix in PowerShell:
```powershell
.\ngram.exe train
```

### "Port 8080 already in use"

Find and kill the process, or change port in code.

### "No predictions showing"

1. Check if model is trained: `dir data\ngram_model.json`
2. Retrain if needed: `.\ngram.exe train`
3. Restart server

---

## What You Can Do

### Try typing in IntelliSense:

| Type | Get suggestions like |
|------|---------------------|
| `func` | `main() {`, `New()`, `Test()` |
| `if err` | `!= nil { return`, `!= nil { log` |
| `SELECT` | `* FROM`, `COUNT(*) FROM` |
| `docker` | `run -p`, `build -t` |
| `git` | `commit -m`, `add .`, `push` |

---

**Have fun! 🚀**
