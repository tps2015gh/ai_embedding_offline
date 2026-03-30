@echo off
REM Train n-gram model with code examples

echo ========================================
echo Training N-gram Model with Code Examples
echo ========================================
echo.

.\ngram.exe train

echo.
echo Model trained! Testing predictions...
echo.

echo Test 1: "func"
.\ngram.exe predict "func"
echo.

echo Test 2: "if err"
.\ngram.exe predict "if err"
echo.

echo Test 3: "SELECT"
.\ngram.exe predict "SELECT"
echo.

echo ========================================
echo Done! Start server with: .\ai_embedding.exe serve
echo ========================================
pause
