package vectorstore

import (
	"ai_embedding_offline/internal/embedding"
	"database/sql"
	"encoding/json"
	"math"
	"sort"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db   *sql.DB
	once sync.Once
)

// VectorRecord represents a stored vector with metadata
type VectorRecord struct {
	ID        int64   `json:"id"`
	Text      string  `json:"text"`
	Embedding string  `json:"embedding"` // JSON encoded
	PositionX float64 `json:"position_x"`
	PositionY float64 `json:"position_y"`
}

// InitDB initializes the SQLite database
func InitDB() error {
	var err error
	once.Do(func() {
		db, err = sql.Open("sqlite3", "data/vectors.db")
		if err != nil {
			return
		}

		// Create tables
		_, err = db.Exec(`
			CREATE TABLE IF NOT EXISTS vectors (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				text TEXT NOT NULL,
				embedding TEXT NOT NULL,
				position_x REAL DEFAULT 0,
				position_y REAL DEFAULT 0
			)
		`)
		if err != nil {
			return
		}

		// Create index for faster searches
		_, err = db.Exec(`
			CREATE INDEX IF NOT EXISTS idx_vectors_text ON vectors(text)
		`)
	})

	return err
}

// StoreVectors stores multiple vectors in the database
func StoreVectors(vectors []embedding.Vector) error {
	if db == nil {
		if err := InitDB(); err != nil {
			return err
		}
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`
		INSERT INTO vectors (text, embedding, position_x, position_y)
		VALUES (?, ?, ?, ?)
	`)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	for _, v := range vectors {
		embeddingJSON, _ := json.Marshal(v.Embedding)

		// Calculate 2D projection for visualization
		posX, posY := project2D(v.Embedding)

		_, err = stmt.Exec(v.Text, string(embeddingJSON), posX, posY)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

// project2D projects high-dimensional vector to 2D for visualization
func project2D(vec []float64) (float64, float64) {
	if len(vec) < 2 {
		return 0, 0
	}

	// Simple PCA-like projection using first two dimensions
	// with some weighting from other dimensions
	x := vec[0] * 100
	y := vec[1] * 100

	// Add contribution from other dimensions
	for i := 2; i < len(vec) && i < 10; i++ {
		x += vec[i] * float64(10-i) * 10
		y += vec[i] * float64(i-1) * 10
	}

	return x, y
}

// SearchSimilar finds vectors similar to the query
// Uses optimized filtering and sorting for fast calculation
func SearchSimilar(queryText string, limit int) ([]VectorRecord, error) {
	if db == nil {
		if err := InitDB(); err != nil {
			return nil, err
		}
	}

	// Generate embedding for query
	queryVectors, err := embedding.CreateEmbeddings([]string{queryText}, 40)
	if err != nil {
		return nil, err
	}
	queryEmbedding := queryVectors[0].Embedding

	// Fetch all vectors (for small datasets)
	// For larger datasets, implement approximate nearest neighbor
	rows, err := db.Query(`SELECT id, text, embedding, position_x, position_y FROM vectors`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	type scoredVector struct {
		record VectorRecord
		score  float64
	}

	var scoredVecs []scoredVector

	for rows.Next() {
		var record VectorRecord
		var embeddingJSON string

		err := rows.Scan(&record.ID, &record.Text, &embeddingJSON, &record.PositionX, &record.PositionY)
		if err != nil {
			continue
		}

		var storedEmbedding []float64
		if err := json.Unmarshal([]byte(embeddingJSON), &storedEmbedding); err != nil {
			continue
		}

		// Calculate cosine similarity
		score := embedding.CosineSimilarity(queryEmbedding, storedEmbedding)

		scoredVecs = append(scoredVecs, scoredVector{
			record: record,
			score:  score,
		})
	}

	// Sort by similarity score (descending)
	sort.Slice(scoredVecs, func(i, j int) bool {
		return scoredVecs[i].score > scoredVecs[j].score
	})

	// Return top results
	result := make([]VectorRecord, 0, limit)
	for i := 0; i < len(scoredVecs) && i < limit; i++ {
		if scoredVecs[i].score > 0.1 { // Filter low similarity
			result = append(result, scoredVecs[i].record)
		}
	}

	return result, nil
}

// GetAllVectors retrieves all vectors for visualization
func GetAllVectors() ([]VectorRecord, error) {
	if db == nil {
		if err := InitDB(); err != nil {
			return nil, err
		}
	}

	rows, err := db.Query(`SELECT id, text, embedding, position_x, position_y FROM vectors`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vectors []VectorRecord

	for rows.Next() {
		var record VectorRecord
		var embeddingJSON string

		err := rows.Scan(&record.ID, &record.Text, &embeddingJSON, &record.PositionX, &record.PositionY)
		if err != nil {
			continue
		}

		vectors = append(vectors, record)
	}

	return vectors, nil
}

// GetVectorStats returns statistics about stored vectors
func GetVectorStats() (map[string]interface{}, error) {
	if db == nil {
		if err := InitDB(); err != nil {
			return nil, err
		}
	}

	stats := make(map[string]interface{})

	// Count total vectors
	var count int
	err := db.QueryRow(`SELECT COUNT(*) FROM vectors`).Scan(&count)
	if err != nil {
		return nil, err
	}
	stats["total_vectors"] = count

	// Calculate average position
	var avgX, avgY float64
	err = db.QueryRow(`SELECT AVG(position_x), AVG(position_y) FROM vectors`).Scan(&avgX, &avgY)
	if err != nil {
		return nil, err
	}
	stats["center_x"] = math.Round(avgX*100) / 100
	stats["center_y"] = math.Round(avgY*100) / 100

	return stats, nil
}

// CloseDB closes the database connection
func CloseDB() error {
	if db != nil {
		return db.Close()
	}
	return nil
}
