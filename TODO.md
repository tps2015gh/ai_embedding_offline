# TODO.md - Business Continuity Plan (BCP)

## For Other AI Agents

If you are continuing this project, here are the pending tasks and improvements:

### High Priority

1. **Improve Embedding Quality**
   - Current implementation uses hash-based embeddings (deterministic but not semantic)
   - Consider integrating offline ML models (e.g., sentence-transformers via Python bridge)
   - Or implement TF-IDF + SVD for better semantic representation

2. **Add Real Next-Word Prediction**
   - Current suggestion system extracts words from similar texts
   - Implement n-gram analysis for better predictions
   - Consider integrating a small language model (e.g., TinyLLaMA)

3. **Performance Optimization**
   - Implement approximate nearest neighbor (ANN) search for large datasets
   - Add HNSW (Hierarchical Navigable Small World) index
   - Batch database operations

### Medium Priority

4. **Better 2D Projection**
   - Implement proper PCA (Principal Component Analysis)
   - Or use t-SNE / UMAP for better clustering visualization
   - Current simple projection loses high-dimensional relationships

5. **Text Preprocessing**
   - Add stopword filtering
   - Implement stemming/lemmatization
   - Add language detection for multi-language support

6. **Enhanced Visualization**
   - Color-code nodes by source directory
   - Add clustering boundaries
   - Implement force-directed graph layout
   - Show connection lines between similar nodes

### Low Priority

7. **User Interface Improvements**
   - Add export functionality (PNG, SVG)
   - Implement search history
   - Add bookmarks for interesting nodes
   - Dark/light theme toggle

8. **API Enhancements**
   - Add authentication
   - Implement rate limiting
   - Add vector update/delete endpoints
   - Support batch operations

9. **Documentation**
   - Add API Swagger/OpenAPI specification
   - Create video tutorial
   - Add more examples in README

### Known Limitations

1. **Hash-based Embeddings**: Not truly semantic, similar texts may not have similar vectors
2. **2D Projection Loss**: 40D → 2D projection loses information
3. **Scalability**: Current approach loads all vectors into memory for search
4. **No True NLP**: Word suggestions are frequency-based, not context-aware

### Technical Debt

- [ ] Replace hash-based embeddings with real ML model
- [ ] Implement proper dimensionality reduction (PCA/t-SNE)
- [ ] Add caching layer for frequently searched queries
- [ ] Implement streaming for large file processing
- [ ] Add unit tests for all packages
- [ ] Add integration tests
- [ ] Set up CI/CD pipeline

### File Structure for Future Expansion

```
ai_embedding_offline/
├── cmd/
├── internal/
│   ├── scanner/
│   ├── embedding/
│   │   ├── embedding.go      # Current: hash-based
│   │   └── ml_embedding.go   # TODO: Add ML-based
│   ├── vectorstore/
│   ├── server/
│   └── nlp/                  # TODO: Add NLP utilities
├── web/
├── data/
├── docs/
├── scripts/
│   ├── train_model.py        # TODO: Model training script
│   └── convert_vectors.sh    # TODO: Migration script
└── tests/                    # TODO: Add tests
```

### Contact / Handover Notes

This project was created by Qwen-Code (Qwen-2.5-Coder) on 2025-03-30.

Key design decisions:
- Chose SQLite for simplicity and portability
- Hash-based embeddings for offline operation (no external dependencies)
- Simple 2D projection for visualization (trade-off: accuracy vs. speed)
- Filter-then-sort optimization for distance calculation

---

## สำหรับ AI Agent อื่นๆ

หากคุณทำงานต่อกับโปรเจกต์นี้ นี่คืองานที่เหลือและการปรับปรุงที่ควรทำ:

### ความสำคัญสูง

1. **ปรับปรุงคุณภาพเอ็มเบดดิ้ง**
   - ปัจจุบันใช้ hash-based embeddings (กำหนดได้แต่ไม่มีความหมาย)
   - พิจารณา integrate offline ML models
   - หรือใช้ TF-IDF + SVD สำหรับการแสดงที่ดีขึ้น

2. **เพิ่มการทำนายคำถัดไปจริง**
   - ระบบแนะนำปัจจุบันดึงคำจากข้อความที่คล้ายกัน
   - ใช้ n-gram analysis สำหรับคำแนะนำที่ดีขึ้น

3. **การเพิ่มประสิทธิภาพ**
   - ใช้ approximate nearest neighbor (ANN) search
   - เพิ่ม HNSW index

### ความสำคัญปานกลาง

4. **การแปลง 2D ที่ดีขึ้น**
   - ใช้ PCA จริง
   - หรือใช้ t-SNE / UMAP

5. **การประมวลผลข้อความ**
   - เพิ่ม stopword filtering
   - ใช้ stemming/lemmatization

### ความสำคัญต่ำ

6. **การปรับปรุงภาพแสดงผล**
   - โหนดสีตามไดเรกทอรีแหล่งที่มา
   - เพิ่มขอบเขต clustering
   - แสดงเส้นเชื่อมระหว่างโหนดที่คล้ายกัน

---

## Agent Work Log Template

| DateTime | Agent | Task | Status | Notes |
|----------|-------|------|--------|-------|
| 2025-03-30 | Qwen-2.5-Coder | Initial implementation | ✅ Complete | Core system functional |
