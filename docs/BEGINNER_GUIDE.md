# Beginner's Guide / คู่มือสำหรับผู้เริ่มต้น

**Created:** 2025-03-30 23:15 (Thai Time UTC+7)  
**Last Updated:** 2025-03-30 23:25 (Thai Time UTC+7)  
**Development Time:** ~10 minutes  

---

## English Version

### What is AI Embedding? (Explained Simply)

Imagine you have a library with thousands of books. Instead of organizing them by title or author, you want to organize them by **meaning** and **topic**.

**AI Embedding** is like giving each book a "coordinate" on a map:
- Books about cooking end up near each other
- Books about programming cluster together
- A book about "Python programming" would be between "programming" and "snakes" (because Python has two meanings!)

### What is a Vector?

A **vector** is just a list of numbers. In our system, each vector has **40 numbers** (40 dimensions).

Think of it like describing a person:
- Height: 175 cm
- Weight: 70 kg
- Age: 25 years
- ... and 37 more characteristics

Each text gets 40 "characteristics" that describe its meaning.

### What is Cosine Similarity?

**Cosine Similarity** measures how similar two vectors are by looking at the angle between them:

- **1.0** = Exactly the same direction (very similar)
- **0.0** = Perpendicular (no similarity)
- **-1.0** = Opposite direction (opposite meaning)

```
        Text A
       /
      /   ← Small angle = Similar!
     /
    ●────── Text B
```

### What is Vector Distance?

**Distance** tells us how far apart two texts are in meaning:
- **Small distance** = Similar meaning
- **Large distance** = Different meaning

### How This System Works

```
1. SCAN: Read all your text files
   ↓
2. CONVERT: Turn text into 40-number vectors
   ↓
3. STORE: Save vectors in database
   ↓
4. SEARCH: Find vectors close to your query
   ↓
5. VISUALIZE: Show as dots on a 2D map
   ↓
6. SUGGEST: Recommend related words
```

### Real Example

If you search for: **"function"**

The system might find these nearby:
- "method" (similar concept in programming)
- "procedure" (related term)
- "variable" (often mentioned together)
- "return" (associated with functions)

---

## เวอร์ชันภาษาไทย

### AI Embedding คืออะไร? (อธิบายแบบง่าย)

ลองนึกภาพคุณมีห้องสมุดที่มีหนังสือเป็นพันเล่ม แทนที่จะจัดเรียงตามชื่อหรือผู้แต่ง คุณต้องการจัดตาม **ความหมาย** และ **หัวข้อ**

**AI Embedding** ก็เหมือนการให้ "พิกัด" กับหนังสือแต่ละเล่ม:
- หนังสือทำอาหารจะอยู่ใกล้กัน
- หนังสือโปรแกรมมิ่งจะ扎堆กัน
- หนังสือ "Python programming" จะอยู่ระหว่าง "programming" และ "snakes" (เพราะ Python มีสองความหมาย!)

### Vector คืออะไร?

**Vector** ก็แค่ชุดของตัวเลข ในระบบของเรา เวกเตอร์แต่ละอันมี **40 ตัวเลข** (40 มิติ)

คิดเหมือนการอธิบายบุคคล:
- ส่วนสูง: 175 ซม.
- น้ำหนัก: 70 กก.
- อายุ: 25 ปี
- ... และอีก 37 ลักษณะ

ข้อความแต่ละอันได้ 40 "ลักษณะ" ที่อธิบายความหมาย

### Cosine Similarity คืออะไร?

**Cosine Similarity** วัดว่าเวกเตอร์สองอันคล้ายกันแค่ไหนโดยดูจากมุมระหว่างพวกมัน:

- **1.0** = ทิศทางเดียวกันเป๊ะ (คล้ายกันมาก)
- **0.0** = ตั้งฉาก (ไม่คล้ายกันเลย)
- **-1.0** = ทิศตรงข้าม (ความหมายตรงข้าม)

### ระยะทางเวกเตอร์คืออะไร?

**ระยะทาง** บอกเราว่าข้อความสองข้อความห่างกันในแง่ความหมายแค่ไหน:
- **ระยะใกล้** = ความหมายคล้ายกัน
- **ระยะไกล** = ความหมายต่างกัน

### ระบบนี้ทำงานอย่างไร

```
1. สแกน: อ่านไฟล์ข้อความทั้งหมด
   ↓
2. แปลง: เปลี่ยนข้อความเป็นเวกเตอร์ 40 ตัวเลข
   ↓
3. เก็บ: บันทึกเวกเตอร์ในฐานข้อมูล
   ↓
4. ค้นหา: หาเวกเตอร์ที่ใกล้กับคำค้นหา
   ↓
5. แสดงผล: แสดงเป็นจุดบนแผนที่ 2D
   ↓
6. แนะนำ: แนะนำคำที่เกี่ยวข้อง
```

### ตัวอย่างจริง

ถ้าคุณค้นหา: **"function"**

ระบบอาจพบสิ่งเหล่านี้ใกล้เคียง:
- "method" (แนวคิดคล้ายกันในโปรแกรมมิ่ง)
- "procedure" (คำที่เกี่ยวข้อง)
- "variable" (มักพูดถึงด้วยกัน)
- "return" (เกี่ยวข้องกับ functions)

---

## Visual Explanation / คำอธิบายด้วยภาพ

### 40 Dimensions Visualization

```
Text: "Hello World"

Vector (simplified to 10 dimensions for display):
[0.82, -0.15, 0.43, 0.91, -0.67, 0.23, -0.88, 0.55, -0.32, 0.76, ...]
 ↑     ↑     ↑     ↑     ↑     ↑     ↑     ↑     ↑     ↑
 1st   2nd   3rd   4th   5th   6th   7th   8th   9th   10th
```

Each number represents a different "aspect" of the text's meaning.

### 2D Projection

Our system projects 40D vectors to 2D for visualization:

```
40D Vector → [PCA-like projection] → 2D Point (x, y)

Example:
Vector A (40D) → (125.5, -43.2) on screen
Vector B (40D) → (130.1, -40.8) on screen  ← Close to A!
Vector C (40D) → (-50.0, 200.3) on screen  ← Far from A!
```

---

## Common Questions / คำถามที่พบบ่อย

**Q: Why 40 dimensions?**
- More dimensions = more precise meaning representation
- 40 is a good balance between accuracy and speed
- Too few (e.g., 5) = loses meaning
- Too many (e.g., 1000) = slow computation

**Q: Can I use this for Thai language?**
- Yes! The system works with any text
- Hash-based embedding is language-agnostic

**Q: How accurate is the similarity?**
- Hash-based (current): Fast but not truly semantic
- For better accuracy: Consider ML-based embeddings (see TODO.md)

**Q: What if I have millions of documents?**
- Current system: Good for thousands
- For millions: Need approximate nearest neighbor (ANN) search

---

## Glossary / ศัพท์บัญญัติ

| English | Thai | Meaning |
|---------|------|---------|
| Embedding | การฝัง | Converting text to numbers |
| Vector | เวกเตอร์ | List of numbers |
| Dimension | มิติ | Each number in vector |
| Similarity | ความคล้ายคลึง | How alike two things are |
| Cosine | โคไซน์ | Mathematical function for angles |
| Distance | ระยะทาง | How far apart |
| Projection | การฉาย | Converting high-D to low-D |
| Query | คำค้นหา | Text you search with |

---

*For technical details, see README.md and docs/VECTOR_DB_USAGE.md*
