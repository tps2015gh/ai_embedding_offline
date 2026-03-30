package embedding

import (
	"ai_embedding_offline/internal/logger"
	"crypto/sha256"
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Vector represents a high-dimensional embedding
type Vector struct {
	Text      string    `json:"text"`
	Embedding []float64 `json:"embedding"`
	Source    string    `json:"source,omitempty"`
}

// CreateEmbeddings generates 40-dimensional embeddings for texts
// Uses a deterministic hash-based approach for offline use
func CreateEmbeddings(texts []string, dimensions int) ([]Vector, error) {
	rand.Seed(time.Now().UnixNano())

	vectors := make([]Vector, 0, len(texts))
	errorCount := 0

	for _, text := range texts {
		if len(text) < 3 {
			logger.Warning("embedding", "CreateEmbeddings", "Text too short, skipping", text)
			continue
		}

		embedding := generateEmbedding(text, dimensions)
		vectors = append(vectors, Vector{
			Text:      text,
			Embedding: embedding,
		})
	}

	if errorCount > 0 {
		logger.Warning("embedding", "CreateEmbeddings", fmt.Sprintf("Completed with %d errors", errorCount), "")
	}

	return vectors, nil
}

// generateEmbedding creates a deterministic embedding from text
func generateEmbedding(text string, dimensions int) []float64 {
	// Use SHA256 hash for deterministic embedding
	hash := sha256.Sum256([]byte(text))

	embedding := make([]float64, dimensions)

	for i := 0; i < dimensions; i++ {
		// Use hash bytes to generate pseudo-random but deterministic values
		idx := i % 32
		value := float64(int(hash[idx])) / 255.0

		// Normalize to [-1, 1] range
		value = value*2 - 1

		// Add some variation based on position
		value += float64(i-dimensions/2) / float64(dimensions) * 0.1

		embedding[i] = value
	}

	// Normalize the vector (L2 normalization)
	norm := 0.0
	for _, v := range embedding {
		norm += v * v
	}
	norm = math.Sqrt(norm)

	if norm > 0 {
		for i := range embedding {
			embedding[i] /= norm
		}
	}

	return embedding
}

// CosineSimilarity calculates cosine similarity between two vectors
func CosineSimilarity(a, b []float64) float64 {
	if len(a) != len(b) {
		return 0
	}

	dotProduct := 0.0
	normA := 0.0
	normB := 0.0

	for i := range a {
		dotProduct += a[i] * b[i]
		normA += a[i] * a[i]
		normB += b[i] * b[i]
	}

	normA = math.Sqrt(normA)
	normB = math.Sqrt(normB)

	if normA == 0 || normB == 0 {
		return 0
	}

	return dotProduct / (normA * normB)
}

// EuclideanDistance calculates Euclidean distance between two vectors
func EuclideanDistance(a, b []float64) float64 {
	if len(a) != len(b) {
		return math.MaxFloat64
	}

	sum := 0.0
	for i := range a {
		diff := a[i] - b[i]
		sum += diff * diff
	}

	return math.Sqrt(sum)
}

// min returns the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
