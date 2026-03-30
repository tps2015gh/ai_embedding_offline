@echo off
echo ========================================
echo AI Embedding Offline - Quick Start
echo ========================================
echo.

echo [1/3] Initializing database...
ai_embedding.exe init
if errorlevel 1 (
    echo Error: Failed to initialize database
    pause
    exit /b 1
)

echo.
echo [2/3] Scanning directories (this may take a while)...
ai_embedding.exe scan
if errorlevel 1 (
    echo Error: Failed to scan directories
    pause
    exit /b 1
)

echo.
echo [3/3] Starting web server...
echo.
echo Open http://localhost:8080 in your browser
echo Press Ctrl+C to stop the server
echo.
ai_embedding.exe serve
