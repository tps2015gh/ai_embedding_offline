package main

import (
	"ai_embedding_offline/internal/ngram"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ngram_demo <command> [args]")
		fmt.Println("Commands:")
		fmt.Println("  train    - Train model on text files")
		fmt.Println("  predict  - Predict next words from input")
		fmt.Println("  demo     - Run interactive demo")
		return
	}

	command := os.Args[1]

	switch command {
	case "train":
		fmt.Println("📚 Training n-gram model...")

		model := ngram.NewModel()

		// Sample training data (code-related text)
		trainingData := []string{
			"func main() { fmt.Println(\"Hello World\") }",
			"func NewModel() *NGramModel { return &NGramModel{} }",
			"if err != nil { return err }",
			"if err != nil { log.Fatal(err) }",
			"if err != nil { log.Fatalf(\"Failed: %v\", err) }",
			"for i := 0; i < len(words); i++ { words[i] = strings.ToLower(words[i]) }",
			"type Prediction struct { Word string; Score float64 }",
			"type Vector struct { Text string; Embedding []float64 }",
			"type NGramModel struct { Unigrams map[string]int; Bigrams map[string]map[string]int }",
			"func (m *Model) Train(text string) { words := tokenize(text) }",
			"func CosineSimilarity(a, b []float64) float64 { dot := 0.0; return dot / (normA * normB) }",
			"SELECT * FROM vectors WHERE id = ?",
			"SELECT COUNT(*) FROM vectors",
			"SELECT text, embedding FROM vectors ORDER BY similarity DESC LIMIT ?",
			"INSERT INTO vectors (text, embedding) VALUES (?, ?)",
			"CREATE TABLE vectors (id INTEGER PRIMARY KEY, text TEXT, embedding TEXT)",
			"npm install express react axios",
			"const App = () => { return <div>Hello</div> }",
			"docker build -t myapp . && docker run -p 8080:8080 myapp",
			"git add . && git commit -m \"update\" && git push origin main",
			"python3 -m venv venv && source venv/bin/activate",
			"go mod init && go mod tidy && go build -o app.exe",
			"go run cmd/main.go serve",
			"func TestMain(t *testing.T) { assert.NotNil(t, model) }",
			"func TestSearch(t *testing.T) { results, err := Search(\"query\", 10) }",
			"response, err := http.Get(url) if err != nil { return nil, err }",
			"defer response.Body.Close()",
			"log.Printf(\"Server starting on %s\", addr)",
			"log.Fatalf(\"Failed to create embeddings: %v\", err)",
		}

		for _, text := range trainingData {
			model.Train(text)
		}

		// Also train on code samples file
		if codeSamples, err := os.ReadFile("data/code_samples.txt"); err == nil {
			model.Train(string(codeSamples))
			fmt.Println("📁 Trained on code_samples.txt")
		}

		// Save model
		if err := model.Save("data/ngram_model.json"); err != nil {
			log.Fatalf("Failed to save model: %v", err)
		}

		fmt.Printf("✅ Model trained! Total words: %d\n", model.TotalWords)
		fmt.Println("   Model saved to: data/ngram_model.json")

	case "predict":
		model, err := ngram.LoadModel("data/ngram_model.json")
		if err != nil {
			log.Fatalf("Failed to load model: %v", err)
		}

		input := strings.Join(os.Args[2:], " ")
		if input == "" {
			input = "func"
		}

		predictions := model.Predict(input, 5)

		fmt.Printf("\nInput: \"%s\"\n", input)
		fmt.Println("Next word suggestions:")
		for i, p := range predictions {
			fmt.Printf("  %d. %s (score: %.2f)\n", i+1, p.Word, p.Score)
		}

	case "demo":
		model := ngram.NewModel()

		// Train on sample data
		samples := []string{
			"func main() { fmt.Println(\"Hello\") }",
			"if err != nil { log.Fatal(err) }",
			"for i := 0; i < 10; i++ { fmt.Println(i) }",
			"type User struct { Name string; Age int }",
			"SELECT * FROM users WHERE active = true",
			"const PORT = 8080 var db *sql.DB",
			"npm install react react-dom",
			"docker run -p 8080:80 app",
			"git commit -m \"fix bug\"",
			"python3 main.py --debug --verbose",
		}

		for _, s := range samples {
			model.Train(s)
		}

		fmt.Println("🎯 N-gram Text Predictor Demo")
		fmt.Println("Type text and press Enter to see predictions")
		fmt.Println("Type 'quit' to exit")
		fmt.Println()

		for {
			fmt.Print("> ")
			var input string
			fmt.Scanln(&input)

			if input == "quit" || input == "exit" {
				break
			}

			predictions := model.Predict(input, 5)
			fmt.Println("Suggestions:")
			for _, p := range predictions {
				fmt.Printf("  • %s\n", p.Word)
			}
			fmt.Println()
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
