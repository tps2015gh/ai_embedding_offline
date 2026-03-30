package ngram

import (
	"encoding/json"
	"os"
	"sort"
	"strings"
	"sync"
)

// NGramModel stores n-gram frequencies
type NGramModel struct {
	Unigrams  map[string]int            `json:"unigrams"`
	Bigrams   map[string]map[string]int `json:"bigrams"`
	Trigrams  map[string]map[string]int `json:"trigrams"`
	TotalWords int                      `json:"total_words"`
	mu        sync.RWMutex
}

// NewModel creates a new n-gram model
func NewModel() *NGramModel {
	return &NGramModel{
		Unigrams:  make(map[string]int),
		Bigrams:   make(map[string]map[string]int),
		Trigrams:  make(map[string]map[string]int),
		TotalWords: 0,
	}
}

// Train trains the model on text
func (m *NGramModel) Train(text string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Tokenize
	words := tokenize(text)

	if len(words) == 0 {
		return
	}

	// Unigrams
	for _, word := range words {
		m.Unigrams[word]++
		m.TotalWords++
	}

	// Bigrams
	for i := 0; i < len(words)-1; i++ {
		if m.Bigrams[words[i]] == nil {
			m.Bigrams[words[i]] = make(map[string]int)
		}
		m.Bigrams[words[i]][words[i+1]]++
	}

	// Trigrams
	for i := 0; i < len(words)-2; i++ {
		key := words[i] + " " + words[i+1]
		if m.Trigrams[key] == nil {
			m.Trigrams[key] = make(map[string]int)
		}
		m.Trigrams[key][words[i+2]]++
	}
}

// Predict returns next word suggestions
func (m *NGramModel) Predict(text string, limit int) []Prediction {
	m.mu.RLock()
	defer m.mu.RUnlock()

	words := tokenize(text)
	suggestions := make(map[string]float64)

	if len(words) == 0 {
		// Return most common unigrams
		return m.getTopUnigrams(limit)
	}

	lastWord := words[len(words)-1]

	// Try trigrams first (most specific) - predicts 1 word ahead
	if len(words) >= 2 {
		key := words[len(words)-2] + " " + lastWord
		if trigramNext, ok := m.Trigrams[key]; ok {
			for next, count := range trigramNext {
				suggestions[next] += float64(count) * 3.0 // Weight trigrams higher
			}
		}
	}

	// Then bigrams
	if bigramNext, ok := m.Bigrams[lastWord]; ok {
		for next, count := range bigramNext {
			suggestions[next] += float64(count) * 2.0
		}
	}

	// Fallback to unigrams
	if len(suggestions) == 0 {
		return m.getTopUnigrams(limit)
	}

	// Sort by score
	return sortSuggestions(suggestions, limit)
}

// PredictPhrase returns multi-word phrase predictions
func (m *NGramModel) PredictPhrase(text string, maxWords int, limit int) []PhrasePrediction {
	m.mu.RLock()
	defer m.mu.RUnlock()

	words := tokenize(text)
	
	if len(words) == 0 {
		return []PhrasePrediction{}
	}

	phrases := make(map[string]float64)

	// Generate phrases by chaining predictions
	for i := 0; i < maxWords; i++ {
		currentWords := words
		phrase := ""
		score := 0.0

		for j := 0; j <= i; j++ {
			var nextWord string
			var wordScore float64

			// Find next word using trigrams if possible
			if len(currentWords) >= 2 {
				key := currentWords[len(currentWords)-2] + " " + currentWords[len(currentWords)-1]
				if trigramNext, ok := m.Trigrams[key]; ok && len(trigramNext) > 0 {
					// Get best match
					for w, c := range trigramNext {
						if float64(c) > wordScore {
							nextWord = w
							wordScore = float64(c) * 3.0
						}
					}
				}
			}

			// Fallback to bigrams
			if nextWord == "" {
				lastWord := currentWords[len(currentWords)-1]
				if bigramNext, ok := m.Bigrams[lastWord]; ok && len(bigramNext) > 0 {
					for w, c := range bigramNext {
						if float64(c) > wordScore {
							nextWord = w
							wordScore = float64(c) * 2.0
						}
					}
				}
			}

			if nextWord == "" {
				break
			}

			if phrase != "" {
				phrase += " "
			}
			phrase += nextWord
			score += wordScore
			currentWords = append(currentWords, nextWord)
		}

		if phrase != "" {
			phrases[phrase] = score
		}
	}

	// Convert to sorted list
	var preds []PhrasePrediction
	for phrase, score := range phrases {
		preds = append(preds, PhrasePrediction{Phrase: phrase, Score: score})
	}

	sort.Slice(preds, func(i, j int) bool {
		return preds[i].Score > preds[j].Score
	})

	if len(preds) > limit {
		preds = preds[:limit]
	}

	return preds
}

// Prediction represents a word suggestion
type Prediction struct {
	Word  string  `json:"word"`
	Score float64 `json:"score"`
}

// PhrasePrediction represents a multi-word phrase suggestion
type PhrasePrediction struct {
	Phrase string  `json:"phrase"`
	Score  float64 `json:"score"`
}

func (m *NGramModel) getTopUnigrams(limit int) []Prediction {
	type wordCount struct {
		word  string
		count int
	}

	var counts []wordCount
	for word, count := range m.Unigrams {
		// Skip very common words
		if isStopWord(word) {
			continue
		}
		counts = append(counts, wordCount{word, count})
	}

	sort.Slice(counts, func(i, j int) bool {
		return counts[i].count > counts[j].count
	})

	result := make([]Prediction, 0, limit)
	for i := 0; i < len(counts) && i < limit; i++ {
		result = append(result, Prediction{
			Word:  counts[i].word,
			Score: float64(counts[i].count),
		})
	}

	return result
}

func sortSuggestions(suggestions map[string]float64, limit int) []Prediction {
	var preds []Prediction
	for word, score := range suggestions {
		preds = append(preds, Prediction{Word: word, Score: score})
	}

	sort.Slice(preds, func(i, j int) bool {
		return preds[i].Score > preds[j].Score
	})

	if len(preds) > limit {
		preds = preds[:limit]
	}

	return preds
}

// Save saves the model to a file
func (m *NGramModel) Save(path string) error {
	m.mu.RLock()
	defer m.mu.RUnlock()

	data, err := json.Marshal(m)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

// Load loads a model from a file
func LoadModel(path string) (*NGramModel, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var model NGramModel
	err = json.Unmarshal(data, &model)
	return &model, err
}

// Tokenize splits text into words
func tokenize(text string) []string {
	text = strings.ToLower(text)
	// Simple tokenization - split on non-alphanumeric
	var words []string
	current := ""
	for _, r := range text {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '_' {
			current += string(r)
		} else {
			if current != "" {
				words = append(words, current)
				current = ""
			}
		}
	}
	if current != "" {
		words = append(words, current)
	}
	return words
}

// isStopWord checks if a word is a common stop word
func isStopWord(word string) bool {
	stopWords := map[string]bool{
		"a": true, "an": true, "the": true, "and": true, "or": true,
		"but": true, "in": true, "on": true, "at": true, "to": true,
		"for": true, "of": true, "with": true, "by": true, "from": true,
		"is": true, "are": true, "was": true, "were": true, "be": true,
		"been": true, "being": true, "have": true, "has": true, "had": true,
		"do": true, "does": true, "did": true, "will": true, "would": true,
		"could": true, "should": true, "may": true, "might": true, "must": true,
		"that": true, "this": true, "these": true, "those": true, "it": true,
		"its": true, "as": true, "if": true, "then": true, "than": true,
		"so": true, "just": true, "not": true, "no": true, "yes": true,
		"i": true, "you": true, "he": true, "she": true, "we": true, "they": true,
		"what": true, "which": true, "who": true, "when": true, "where": true,
		"why": true, "how": true, "all": true, "each": true, "every": true,
		"both": true, "few": true, "more": true, "most": true, "other": true,
		"some": true, "such": true, "only": true, "own": true, "same": true,
		"into": true, "over": true, "after": true, "before": true, "between": true,
	}
	return stopWords[word]
}
