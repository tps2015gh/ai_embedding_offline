package server

import (
	"ai_embedding_offline/internal/vectorstore"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// Response represents API response
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// StartServer starts the HTTP server
func StartServer(addr string) error {
	// Initialize database
	if err := vectorstore.InitDB(); err != nil {
		return err
	}

	// Serve static files
	webDir := filepath.Join(".", "web")
	fs := http.FileServer(http.Dir(webDir))
	http.Handle("/", fs)

	// API endpoints
	http.HandleFunc("/api/search", handleSearch)
	http.HandleFunc("/api/vectors", handleGetVectors)
	http.HandleFunc("/api/stats", handleStats)
	http.HandleFunc("/api/suggest", handleSuggest)

	log.Printf("Server starting on %s", addr)
	return http.ListenAndServe(addr, nil)
}

// handleSearch handles similarity search
func handleSearch(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	if r.Method != "POST" {
		sendError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Query string `json:"query"`
		Limit int    `json:"limit"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if req.Limit <= 0 {
		req.Limit = 20
	}

	results, err := vectorstore.SearchSimilar(req.Query, req.Limit)
	if err != nil {
		sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendResponse(w, results)
}

// handleGetVectors returns all vectors for visualization
func handleGetVectors(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	vectors, err := vectorstore.GetAllVectors()
	if err != nil {
		sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendResponse(w, vectors)
}

// handleStats returns database statistics
func handleStats(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	stats, err := vectorstore.GetVectorStats()
	if err != nil {
		sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendResponse(w, stats)
}

// handleSuggest provides next-word suggestions
func handleSuggest(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	if r.Method != "POST" {
		sendError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Text string `json:"text"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Find similar vectors and extract potential next words
	suggestions, err := generateSuggestions(req.Text)
	if err != nil {
		sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendResponse(w, suggestions)
}

// generateSuggestions generates word suggestions based on vector proximity
func generateSuggestions(text string) ([]map[string]interface{}, error) {
	results, err := vectorstore.SearchSimilar(text, 10)
	if err != nil {
		return nil, err
	}

	suggestions := make([]map[string]interface{}, 0)

	for _, r := range results {
		// Extract potential next words from similar texts
		words := extractWords(r.Text)
		for _, word := range words {
			suggestions = append(suggestions, map[string]interface{}{
				"word":   word,
				"score":  r.PositionX + r.PositionY, // Use position as proxy for relevance
			})
		}
	}

	return suggestions, nil
}

// extractWords extracts meaningful words from text
func extractWords(text string) []string {
	// Simple word extraction
	words := make([]string, 0)
	// Implementation would tokenize and filter stopwords
	// For now, return empty - can be enhanced
	return words
}

// Helper functions

func sendResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{
		Success: true,
		Data:    data,
	})
}

func sendError(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(Response{
		Success: false,
		Error:   message,
	})
}

func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r := os.Getenv("REQUEST_METHOD"); r == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
}
