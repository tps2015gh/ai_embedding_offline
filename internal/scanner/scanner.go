package scanner

import (
	"ai_embedding_offline/internal/logger"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// File extensions to scan
var textExtensions = map[string]bool{
	".txt": true, ".go": true, ".py": true, ".js": true, ".ts": true,
	".java": true, ".c": true, ".cpp": true, ".h": true, ".hpp": true,
	".cs": true, ".rb": true, ".php": true, ".rs": true, ".swift": true,
	".kt": true, ".scala": true, ".r": true, ".md": true, ".json": true,
	".xml": true, ".yaml": true, ".yml": true, ".html": true, ".css": true,
	".sql": true, ".sh": true, ".bat": true, ".ps1": true,
}

// ScanDirectory scans a directory recursively for text files
func ScanDirectory(rootPath string, progressChan chan<- string) ([]string, error) {
	var texts []string

	if progressChan != nil {
		progressChan <- fmt.Sprintf("🔍 Scanning: %s", rootPath)
	}

	fileCount := 0
	chunkCount := 0
	skipCount := 0

	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil // Skip inaccessible files
		}

		if info.IsDir() {
			// Skip common non-essential directories
			skipDirs := map[string]bool{
				"node_modules": true, ".git": true, ".svn": true,
				"vendor": true, "bin": true, "obj": true,
				"__pycache__": true, ".venv": true, "venv": true,
				"dist": true, "build": true, "target": true,
			}
			if skipDirs[info.Name()] {
				skipCount++
				return filepath.SkipDir
			}
			return nil
		}

		ext := strings.ToLower(filepath.Ext(path))
		if !textExtensions[ext] {
			return nil
		}

		fileCount++

		// Show progress every 50 files
		if fileCount%50 == 0 && progressChan != nil {
			progressChan <- fmt.Sprintf("📄 Files: %d | Chunks: %d | Skipped dirs: %d", fileCount, chunkCount, skipCount)
		}

		// Read file content
		content, err := ioutil.ReadFile(path)
		if err != nil {
			logger.Error("scanner", "ReadFile", err.Error(), path)
			return nil
		}

		// Split into chunks (by lines or size)
		chunks := splitIntoChunks(string(content), path)
		texts = append(texts, chunks...)
		chunkCount += len(chunks)

		return nil
	})

	logger.Info("scanner", "ScanDirectory", "Scan complete", fmt.Sprintf("Found %d chunks from %d files, skipped %d dirs", chunkCount, fileCount, skipCount))
	if progressChan != nil {
		progressChan <- fmt.Sprintf("✅ Complete: %d files, %d chunks, %d dirs skipped", fileCount, chunkCount, skipCount)
	}
	return texts, err
}

// splitIntoChunks splits text into manageable chunks for embedding
func splitIntoChunks(content, source string) []string {
	var chunks []string

	// Split by newlines first
	lines := strings.Split(content, "\n")

	var currentChunk strings.Builder
	chunkSize := 0
	maxChunkSize := 500 // characters

	for _, line := range lines {
		if chunkSize+len(line) > maxChunkSize && currentChunk.Len() > 0 {
			chunk := strings.TrimSpace(currentChunk.String())
			if len(chunk) > 10 { // Skip very small chunks
				chunks = append(chunks, chunk)
			}
			currentChunk.Reset()
			chunkSize = 0
		}
		currentChunk.WriteString(line)
		currentChunk.WriteString(" ")
		chunkSize += len(line) + 1
	}

	// Add remaining chunk
	if currentChunk.Len() > 0 {
		chunk := strings.TrimSpace(currentChunk.String())
		if len(chunk) > 10 {
			chunks = append(chunks, chunk)
		}
	}

	return chunks
}
