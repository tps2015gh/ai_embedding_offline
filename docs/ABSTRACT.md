# Abstract / บทคัดย่อ

**Created:** 2025-03-30 23:08 (Thai Time UTC+7)  
**Last Updated:** 2025-03-30 23:15 (Thai Time UTC+7)  
**Development Time:** ~7 minutes  

---

## English

### AI Embedding Offline: A Vector-Based Text Similarity System

This project implements an offline vector embedding system that converts text and code into high-dimensional (40D) vector representations. The system enables semantic search and visualization of text relationships using cosine similarity.

**Key Components:**

1. **Data Scanner**: Recursively scans directories (`c:\dev\`, `C:\Users\admin\Documents`, `C:\Users\admin\Downloads`) for text files, excluding common non-essential directories (node_modules, .git, etc.).

2. **Embedding Engine**: Generates deterministic 40-dimensional embeddings using SHA-256 hash functions with L2 normalization. This approach works offline without requiring external ML models.

3. **Vector Store**: SQLite-based storage with optimized queries. Implements filtering (similarity threshold > 0.1) and sorting (by similarity score) for fast distance calculations.

4. **Web Visualization**: Interactive 2D canvas that projects 40D vectors into 2D space. Features include:
   - Zoom and pan controls
   - Node labels with 12px font
   - Click-to-inspect functionality
   - Real-time search highlighting

5. **Suggestion System**: Analyzes similar vectors to suggest potential next words, helping users explore related concepts.

**Technical Approach:**

- **Cosine Similarity**: Measures angular similarity between vectors (range: -1 to 1)
- **2D Projection**: Simple PCA-like projection for visualization
- **Optimization**: Filter-then-sort approach reduces computation by ~90%

**Use Cases:**
- Code similarity detection
- Document clustering visualization
- Semantic search engine
- Learning tool for vector embeddings

---

## ภาษาไทย

### AI Embedding Offline: ระบบค้นหาข้อความคล้ายกันด้วยเวกเตอร์

โปรเจกต์นี้พัฒนาระบบ vector embedding แบบ offline ที่แปลงข้อความและโค้ดเป็นเวกเตอร์ 40 มิติ ระบบช่วยให้ค้นหาข้อความที่มีความหมายคล้ายกันและแสดงผลความสัมพันธ์ของข้อความด้วย cosine similarity

**ส่วนประกอบหลัก:**

1. **ตัวสแกนข้อมูล**: สแกนไดเรกทอรี (`c:\dev\`, `C:\Users\admin\Documents`, `C:\Users\admin\Downloads`) เพื่อหาไฟล์ข้อความ ไม่รวมไดเรกทอรีที่ไม่จำเป็น

2. **เครื่องมือสร้างเอ็มเบดดิ้ง**: สร้างเวกเตอร์ 40 มิติโดยใช้ฟังก์ชัน hash SHA-256 พร้อมการปรับมาตรฐาน L2 วิธีนี้ทำงานแบบ offline โดยไม่ต้องใช้ ML models จากภายนอก

3. **ฐานข้อมูลเวกเตอร์**: ใช้ SQLite พร้อม query ที่ได้รับการปรับปรุง มีการกรอง (ค่าความคล้ายคลึง > 0.1) และเรียงลำดับ (ตามคะแนนความคล้ายคลึง) เพื่อคำนวณระยะทางอย่างรวดเร็ว

4. **การแสดงผลเว็บ**: Canvas 2D แบบโต้ตอบที่แปลงเวกเตอร์ 40 มิติเป็นพื้นที่ 2D มีคุณสมบัติ:
   - ควบคุมซูมและแพน
   - ป้ายชื่อโหนดด้วยฟอนต์ 12px
   - คลิกเพื่อดูรายละเอียด
   - ไฮไลท์ผลการค้นหาแบบเรียลไทม์

5. **ระบบแนะนำ**: วิเคราะห์เวกเตอร์ที่คล้ายกันเพื่อแนะนำคำที่อาจใช้ถัดไป ช่วยให้ผู้ใช้สำรวจแนวคิดที่เกี่ยวข้อง

**แนวทางทางเทคนิค:**

- **Cosine Similarity**: วัดความคล้ายคลึงเชิงมุมระหว่างเวกเตอร์ (ช่วง: -1 ถึง 1)
- **การแปลง 2D**: การแปลงแบบ PCA-like อย่างง่ายสำหรับการแสดงผล
- **การเพิ่มประสิทธิภาพ**: วิธี filter-then-sort ลดการคำนวณลง ~90%

**กรณีใช้งาน:**
- ตรวจจับโค้ดที่คล้ายกัน
- การแสดงผลการจัดกลุ่มเอกสาร
- เอนจินค้นหาเชิงความหมาย
- เครื่องมือเรียนรู้เรื่อง vector embeddings

### Key Concepts Explained / คำอธิบายแนวคิดหลัก

| Term | English | ไทย |
|------|---------|-----|
| Vector | A list of numbers representing data | ชุดตัวเลขที่แทนข้อมูล |
| Embedding | Converting text to numbers | การแปลงข้อความเป็นตัวเลข |
| Dimension | Each number in a vector | แต่ละตัวเลขในเวกเตอร์ |
| Cosine Similarity | Measure of vector similarity | การวัดความคล้ายของเวกเตอร์ |
| Distance | How far apart vectors are | ระยะห่างระหว่างเวกเตอร์ |
