package main

import (
	"ai_embedding_offline/internal/embedding"
	"ai_embedding_offline/internal/logger"
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
		fmt.Println("  init     - Initialize vector database")
		fmt.Println("  scan     - Scan directories and create embeddings")
		fmt.Println("  serve    - Start the web server")
		return
	}

	// Initialize logger first
	if err := logger.InitLogger("data"); err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.CloseLogger()

	command := os.Args[1]

	switch command {
	case "init":
		log.Println("Initializing vector database...")
		logger.Info("main", "init", "Starting database initialization", "")
		if err := vectorstore.InitDB(); err != nil {
			logger.Error("main", "init", "Database initialization failed", err.Error())
			log.Fatalf("Failed to initialize DB: %v", err)
		}
		logger.Info("main", "init", "Database initialized successfully", "")
		log.Println("Database initialized successfully")

	case "scan":
		log.Println("Scanning directories...")
		logger.Info("main", "scan", "Starting directory scan", "")

		// Default scan directories - customize these for your system
		// You can also pass paths as arguments: ai_embedding.exe scan C:\my\code D:\projects
		dirs := []string{
			"c:\\dev\\",
			"C:\\Users\\admin\\Documents",
			"C:\\Users\\admin\\Downloads",
		}

		// Override with command line arguments if provided
		if len(os.Args) > 2 {
			dirs = os.Args[2:]
			log.Printf("Scanning custom paths: %v", dirs)
		}

		// Create progress channel
		progressChan := make(chan string, 100)
		doneChan := make(chan bool)

		// Progress display goroutine
		go func() {
			for {
				select {
				case msg := <-progressChan:
					fmt.Println("  ", msg)
				case <-doneChan:
					return
				}
			}
		}()

		// Scan all directories
		var allTexts []string
		for _, dir := range dirs {
			texts, err := scanner.ScanDirectory(dir, progressChan)
			if err != nil {
				logger.Warning("main", "scan", fmt.Sprintf("Error scanning %s", dir), err.Error())
				log.Printf("Warning: Error scanning %s: %v", dir, err)
				continue
			}
			allTexts = append(allTexts, texts...)
		}

		close(doneChan)

		logger.Info("main", "scan", fmt.Sprintf("Found %d text chunks", len(allTexts)), "")
		log.Printf("Found %d text chunks", len(allTexts))

		// Create embeddings
		log.Println("Creating embeddings...")
		logger.Info("main", "scan", "Creating embeddings", "")
		vectors, err := embedding.CreateEmbeddings(allTexts, 40)
		if err != nil {
			logger.Error("main", "scan", "Embedding creation failed", err.Error())
			log.Fatalf("Failed to create embeddings: %v", err)
		}

		// Store in database
		log.Println("Storing vectors in database...")
		logger.Info("main", "scan", "Storing vectors in database", "")
		if err := vectorstore.StoreVectors(vectors); err != nil {
			logger.Error("main", "scan", "Vector storage failed", err.Error())
			log.Fatalf("Failed to store vectors: %v", err)
		}

		logger.Info("main", "scan", "Scan and embedding complete", "")
		log.Println("✅ Scan and embedding complete!")

	case "serve":
		log.Println("Starting web server on :8080...")
		logger.Info("main", "serve", "Starting web server on :8080", "")
		if err := server.StartServer(":8080"); err != nil {
			logger.Error("main", "serve", "Server failed", err.Error())
			log.Fatalf("Server failed: %v", err)
		}

	default:
		logger.Warning("main", "main", fmt.Sprintf("Unknown command: %s", command), "")
		log.Printf("Unknown command: %s", command)
	}
}
