@echo off
REM Build script with CGO enabled for SQLite support

echo ========================================
echo Building AI Embedding Offline
echo ========================================
echo.

echo Setting CGO_ENABLED=1...
set CGO_ENABLED=1

echo Checking for C compiler...
where gcc >nul 2>nul
if %ERRORLEVEL% NEQ 0 (
    echo [ERROR] GCC (MinGW) not found!
    echo.
    echo Please install MinGW from: https://www.mingw-w64.org/
    echo Or install with chocolatey: choco install mingw
    echo.
    pause
    exit /b 1
)

echo [OK] GCC found
gcc --version | findstr /C:"gcc"
echo.

echo Building main application with SQLite support...
go build -o ai_embedding.exe ./cmd/
if %ERRORLEVEL% NEQ 0 (
    echo [ERROR] Build failed!
    echo.
    echo Make sure you have:
    echo   1. Go installed
    echo   2. GCC (MinGW) installed
    echo   3. Run: go mod tidy
    echo.
    pause
    exit /b 1
)
echo [OK] Built ai_embedding.exe
echo.

echo Building n-gram tool...
go build -o ngram.exe ./cmd/ngram/
if %ERRORLEVEL% NEQ 0 (
    echo [ERROR] N-gram build failed!
    pause
    exit /b 1
)
echo [OK] Built ngram.exe
echo.

echo ========================================
echo Build Complete! ✅
echo ========================================
echo.
echo To start server:
echo   .\ai_embedding.exe serve
echo.
echo To train model:
echo   .\ngram.exe train
echo.
pause
