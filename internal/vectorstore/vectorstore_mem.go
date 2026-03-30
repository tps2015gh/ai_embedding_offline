package vectorstore

import (
	"ai_embedding_offline/internal/embedding"
	"ai_embedding_offline/internal/logger"
	"encoding/json"
	"math"
	"os"
	"sort"
	"sync"
)

var (
	memDB      *MemoryDB
	memDBOnce  sync.Once
	dbFilePath = "data/vectors.json"
)

// MemoryDB stores vectors in memory (no SQLite required)
type MemoryDB struct {
	Vectors []VectorRecord `json:"vectors"`
	mu      sync.RWMutex
}

// VectorRecord represents a stored vector
type VectorRecord struct {
	ID           int64   `json:"id"`
	Text         string  `json:"text"`
	EmbeddingStr string  `json:"embedding"` // JSON encoded
	PositionX    float64 `json:"position_x"`
	PositionY    float64 `json:"position_y"`
	// Runtime field (not serialized)
	Embedding []float64 `json:"-"`
}

// InitDB initializes the in-memory database
func InitDB() error {
	var err error
	memDBOnce.Do(func() {
		logger.Info("vectorstore", "InitDB", "Initializing in-memory database (no SQLite)", "")
		
		memDB = &MemoryDB{
			Vectors: make([]VectorRecord, 0),
		}
		
		// Try to load existing data
		if data, readErr := os.ReadFile(dbFilePath); readErr == nil {
			if err := json.Unmarshal(data, memDB); err != nil {
				logger.Warning("vectorstore", "InitDB", "Failed to load existing data", err.Error())
			} else {
				// Decode embeddings
				for i := range memDB.Vectors {
					var emb []float64
					if json.Unmarshal([]byte(memDB.Vectors[i].EmbeddingStr), &emb) == nil {
						memDB.Vectors[i].Embedding = emb
					}
				}
				logger.Info("vectorstore", "InitDB", "Loaded existing vectors", "")
			}
		}
	})
	
	return err
}

// StoreVectors stores vectors in memory
func StoreVectors(vectors []embedding.Vector) error {
	if memDB == nil {
		if err := InitDB(); err != nil {
			return err
		}
	}
	
	memDB.mu.Lock()
	defer memDB.mu.Unlock()
	
	logger.Info("vectorstore", "StoreVectors", "Storing vectors in memory", "")
	
	for _, v := range vectors {
		posX, posY := project2D(v.Embedding)
		embJSON, _ := json.Marshal(v.Embedding)
		
		memDB.Vectors = append(memDB.Vectors, VectorRecord{
			ID:           int64(len(memDB.Vectors) + 1),
			Text:         v.Text,
			Embedding:    v.Embedding,
			EmbeddingStr: string(embJSON),
			PositionX:    posX,
			PositionY:    posY,
		})
	}
	
	// Save to file
	return saveDB()
}

// saveDB saves the in-memory database to file
func saveDB() error {
	if err := os.MkdirAll("data", 0755); err != nil {
		return err
	}
	
	data, err := json.Marshal(memDB)
	if err != nil {
		return err
	}
	
	return os.WriteFile(dbFilePath, data, 0644)
}

// SearchSimilar finds similar vectors
func SearchSimilar(queryText string, limit int) ([]VectorRecord, error) {
	if memDB == nil {
		if err := InitDB(); err != nil {
			return nil, err
		}
	}
	
	memDB.mu.RLock()
	defer memDB.mu.RUnlock()
	
	// Generate embedding for query
	queryVectors, err := embedding.CreateEmbeddings([]string{queryText}, 40)
	if err != nil {
		return nil, err
	}
	queryEmbedding := queryVectors[0].Embedding
	
	type scoredVector struct {
		record VectorRecord
		score  float64
	}
	
	var scoredVecs []scoredVector
	
	for _, record := range memDB.Vectors {
		score := embedding.CosineSimilarity(queryEmbedding, record.Embedding)
		if score > 0.1 { // Filter low similarity
			scoredVecs = append(scoredVecs, scoredVector{
				record: record,
				score:  score,
			})
		}
	}
	
	// Sort by score
	sort.Slice(scoredVecs, func(i, j int) bool {
		return scoredVecs[i].score > scoredVecs[j].score
	})
	
	// Return top results
	result := make([]VectorRecord, 0, limit)
	for i := 0; i < len(scoredVecs) && i < limit; i++ {
		result = append(result, scoredVecs[i].record)
	}
	
	return result, nil
}

// GetAllVectors returns all vectors
func GetAllVectors() ([]VectorRecord, error) {
	if memDB == nil {
		if err := InitDB(); err != nil {
			return nil, err
		}
	}
	
	memDB.mu.RLock()
	defer memDB.mu.RUnlock()
	
	return append([]VectorRecord{}, memDB.Vectors...), nil
}

// GetVectorStats returns statistics
func GetVectorStats() (map[string]interface{}, error) {
	if memDB == nil {
		if err := InitDB(); err != nil {
			return nil, err
		}
	}
	
	memDB.mu.RLock()
	defer memDB.mu.RUnlock()
	
	stats := map[string]interface{}{
		"total_vectors": len(memDB.Vectors),
		"database_type": "memory (JSON)",
	}
	
	if len(memDB.Vectors) > 0 {
		avgX, avgY := 0.0, 0.0
		for _, v := range memDB.Vectors {
			avgX += v.PositionX
			avgY += v.PositionY
		}
		stats["center_x"] = math.Round((avgX / float64(len(memDB.Vectors))) * 100) / 100
		stats["center_y"] = math.Round((avgY / float64(len(memDB.Vectors))) * 100) / 100
	}
	
	return stats, nil
}

// project2D projects 40D vector to 2D
func project2D(vec []float64) (float64, float64) {
	if len(vec) < 2 {
		return 0, 0
	}
	
	x := vec[0] * 100
	y := vec[1] * 100
	
	for i := 2; i < len(vec) && i < 10; i++ {
		x += vec[i] * float64(10-i) * 10
		y += vec[i] * float64(i-1) * 10
	}
	
	return x, y
}

// CloseDB saves and closes
func CloseDB() error {
	if memDB != nil {
		return saveDB()
	}
	return nil
}
