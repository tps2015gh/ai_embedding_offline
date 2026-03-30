package server

import (
	"ai_embedding_offline/internal/logger"
	"ai_embedding_offline/internal/ngram"
	"ai_embedding_offline/internal/vectorstore"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

var (
	ngramModel *ngram.NGramModel
	modelOnce  sync.Once
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
	logger.Info("server", "StartServer", "Starting server", addr)

	if err := vectorstore.InitDB(); err != nil {
		logger.Error("server", "StartServer", "Failed to initialize database", err.Error())
		return err
	}

	// Load n-gram model
	modelOnce.Do(func() {
		var err error
		ngramModel, err = ngram.LoadModel("data/ngram_model.json")
		if err != nil {
			logger.Warning("server", "StartServer", "No n-gram model found, will create on demand", "")
		} else {
			logger.Info("server", "StartServer", fmt.Sprintf("N-gram model loaded: %d words", ngramModel.TotalWords), "")
		}
	})

	// Serve static files
	webDir := filepath.Join(".", "web")
	fs := http.FileServer(http.Dir(webDir))
	http.Handle("/", fs)

	// API endpoints
	http.HandleFunc("/api/search", handleSearch)
	http.HandleFunc("/api/vectors", handleGetVectors)
	http.HandleFunc("/api/stats", handleStats)
	http.HandleFunc("/api/suggest", handleSuggest)
	http.HandleFunc("/api/ngram/predict", handleNGramPredict)
	http.HandleFunc("/api/ngram/predict-phrase", handleNGramPredictPhrase)
	http.HandleFunc("/api/ngram/stats", handleNGramStats)

	logger.Info("server", "StartServer", "Server starting on "+addr, "")
	log.Printf("Server starting on %s", addr)
	return http.ListenAndServe(addr, nil)
}

// handleSearch handles similarity search
func handleSearch(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	if r.Method != "POST" {
		logger.Warning("server", "handleSearch", "Method not allowed", r.Method)
		sendError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Query string `json:"query"`
		Limit int    `json:"limit"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error("server", "handleSearch", "Invalid request", err.Error())
		sendError(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if req.Limit <= 0 {
		req.Limit = 20
	}

	logger.Info("server", "handleSearch", "Search query", req.Query)

	results, err := vectorstore.SearchSimilar(req.Query, req.Limit)
	if err != nil {
		logger.Error("server", "handleSearch", "Search failed", err.Error())
		sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Info("server", "handleSearch", fmt.Sprintf("Found %d results", len(results)), req.Query)
	sendResponse(w, results)
}

// handleGetVectors returns all vectors for visualization
func handleGetVectors(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	vectors, err := vectorstore.GetAllVectors()
	if err != nil {
		logger.Error("server", "handleGetVectors", "Failed to get vectors", err.Error())
		sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Info("server", "handleGetVectors", fmt.Sprintf("Returning %d vectors", len(vectors)), "")
	sendResponse(w, vectors)
}

// handleStats returns database statistics
func handleStats(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	stats, err := vectorstore.GetVectorStats()
	if err != nil {
		logger.Error("server", "handleStats", "Failed to get stats", err.Error())
		sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendResponse(w, stats)
}

// handleSuggest provides next-word suggestions
func handleSuggest(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	if r.Method != "POST" {
		logger.Warning("server", "handleSuggest", "Method not allowed", r.Method)
		sendError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Text string `json:"text"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error("server", "handleSuggest", "Invalid request", err.Error())
		sendError(w, "Invalid request", http.StatusBadRequest)
		return
	}

	logger.Info("server", "handleSuggest", "Suggestion request", req.Text)

	// Find similar vectors and extract potential next words
	suggestions, err := generateSuggestions(req.Text)
	if err != nil {
		logger.Error("server", "handleSuggest", "Failed to generate suggestions", err.Error())
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

// handleNGramPredict handles n-gram prediction requests
func handleNGramPredict(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	if r.Method != "POST" {
		sendError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Text  string `json:"text"`
		Limit int    `json:"limit"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error("server", "handleNGramPredict", "Invalid request", err.Error())
		sendError(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if req.Limit <= 0 {
		req.Limit = 5
	}

	// Load model if not loaded
	if ngramModel == nil {
		var err error
		ngramModel, err = ngram.LoadModel("data/ngram_model.json")
		if err != nil {
			logger.Error("server", "handleNGramPredict", "Failed to load model", err.Error())
			sendError(w, "Model not trained yet", http.StatusInternalServerError)
			return
		}
	}

	predictions := ngramModel.Predict(req.Text, req.Limit)
	logger.Info("server", "handleNGramPredict", fmt.Sprintf("Predicted %d suggestions", len(predictions)), req.Text)
	sendResponse(w, predictions)
}

// handleNGramStats returns n-gram model statistics
func handleNGramStats(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	if ngramModel == nil {
		var err error
		ngramModel, err = ngram.LoadModel("data/ngram_model.json")
		if err != nil {
			sendError(w, "Model not trained yet", http.StatusInternalServerError)
			return
		}
	}

	stats := map[string]interface{}{
		"total_words": ngramModel.TotalWords,
		"unigrams":    len(ngramModel.Unigrams),
		"bigrams":     len(ngramModel.Bigrams),
		"trigrams":    len(ngramModel.Trigrams),
	}

	sendResponse(w, stats)
}

// handleNGramPredictPhrase handles multi-word phrase prediction
func handleNGramPredictPhrase(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	if r.Method != "POST" {
		sendError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Text     string `json:"text"`
		MaxWords int    `json:"maxWords"`
		Limit    int    `json:"limit"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error("server", "handleNGramPredictPhrase", "Invalid request", err.Error())
		sendError(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if req.MaxWords <= 0 {
		req.MaxWords = 3
	}
	if req.Limit <= 0 {
		req.Limit = 5
	}

	// Load model if not loaded
	if ngramModel == nil {
		var err error
		ngramModel, err = ngram.LoadModel("data/ngram_model.json")
		if err != nil {
			logger.Error("server", "handleNGramPredictPhrase", "Failed to load model", err.Error())
			sendError(w, "Model not trained yet", http.StatusInternalServerError)
			return
		}
	}

	phrases := ngramModel.PredictPhrase(req.Text, req.MaxWords, req.Limit)
	logger.Info("server", "handleNGramPredictPhrase", fmt.Sprintf("Predicted %d phrases", len(phrases)), req.Text)
	sendResponse(w, phrases)
}
