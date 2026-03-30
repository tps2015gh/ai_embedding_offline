@echo off
setlocal EnableDelayedExpansion

echo ========================================
echo AI Embedding Offline - Setup Installer
echo ========================================
echo.

REM Check if Go is installed
where go >nul 2>nul
if %ERRORLEVEL% NEQ 0 (
    echo [ERROR] Go is not installed!
    echo.
    echo Please install Go from: https://go.dev/dl/
    echo.
    pause
    exit /b 1
)

echo [OK] Go is installed
go version
echo.

REM Step 1: Initialize Go modules
echo [1/4] Initializing Go modules...
go mod tidy
if %ERRORLEVEL% NEQ 0 (
    echo [ERROR] Failed to initialize Go modules
    pause
    exit /b 1
)
echo [OK] Go modules initialized
echo.

REM Step 2: Build main application
echo [2/4] Building main application...
go build -o ai_embedding.exe ./cmd/
if %ERRORLEVEL% NEQ 0 (
    echo [ERROR] Failed to build main application
    pause
    exit /b 1
)
echo [OK] Built ai_embedding.exe
echo.

REM Step 3: Build n-gram tool
echo [3/4] Building n-gram tool...
go build -o ngram.exe ./cmd/ngram/
if %ERRORLEVEL% NEQ 0 (
    echo [ERROR] Failed to build n-gram tool
    pause
    exit /b 1
)
echo [OK] Built ngram.exe
echo.

REM Step 4: Train n-gram model
echo [4/4] Training n-gram model...
ngram.exe train
if %ERRORLEVEL% NEQ 0 (
    echo [WARNING] Model training failed, but you can train later
    echo.
)
echo.

REM Create data directory
if not exist data mkdir data
echo [OK] Data directory ready
echo.

echo ========================================
echo Setup Complete! ✅
echo ========================================
echo.
echo Files created:
echo   - ai_embedding.exe (main server)
echo   - ngram.exe (training tool)
echo   - data/ngram_model.json (trained model)
echo.
echo To start the server:
echo   ai_embedding.exe serve
echo.
echo Then open in browser:
echo   http://localhost:8080/intellisense.html
echo.
echo ========================================
pause
