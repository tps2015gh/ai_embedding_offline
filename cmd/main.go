package main

import (
	"ai_embedding_offline/internal/embedding"
	"ai_embedding_offline/internal/scanner"
	"ai_embedding_offline/internal/server"
	"ai_embedding_offline/internal/vectorstore"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ai_embedding_offline <command> [args]")
		fmt.Println("Commands:")
		fmt.Println("  scan     - Scan directories and create embeddings")
		fmt.Println("  serve    - Start the web server")
		fmt.Println("  init     - Initialize vector database")
		return
	}

	command := os.Args[1]

	switch command {
	case "init":
		log.Println("Initializing vector database...")
		if err := vectorstore.InitDB(); err != nil {
			log.Fatalf("Failed to initialize DB: %v", err)
		}
		log.Println("Database initialized successfully")

	case "scan":
		log.Println("Scanning directories...")
		dirs := []string{
			"c:\\dev\\",
			"C:\\Users\\admin\\Documents",
			"C:\\Users\\admin\\Downloads",
		}

		// Scan all directories
		var allTexts []string
		for _, dir := range dirs {
			texts, err := scanner.ScanDirectory(dir)
			if err != nil {
				log.Printf("Warning: Error scanning %s: %v", dir, err)
				continue
			}
			allTexts = append(allTexts, texts...)
		}

		log.Printf("Found %d text chunks", len(allTexts))

		// Create embeddings
		log.Println("Creating embeddings...")
		vectors, err := embedding.CreateEmbeddings(allTexts, 40)
		if err != nil {
			log.Fatalf("Failed to create embeddings: %v", err)
		}

		// Store in database
		log.Println("Storing vectors in database...")
		if err := vectorstore.StoreVectors(vectors); err != nil {
			log.Fatalf("Failed to store vectors: %v", err)
		}

		log.Println("Scan and embedding complete!")

	case "serve":
		log.Println("Starting web server on :8080...")
		if err := server.StartServer(":8080"); err != nil {
			log.Fatalf("Server failed: %v", err)
		}

	default:
		log.Printf("Unknown command: %s", command)
	}
}
