# Solved Errors Log

This file tracks errors that have been identified and fixed.

## Error Resolution Process

1. Errors are logged to `data/errors.log` automatically
2. When an error is fixed, move it here with the resolution
3. Format: Date - Error - Resolution

---

## Resolved Errors

### [Template for future errors]

**Date:** YYYY-MM-DD HH:MM:SS
**Component:** [scanner|embedding|vectorstore|server|main]
**Error:** [Error message]
**Context:** [Additional context]
**Resolution:** [How it was fixed]
**Fixed By:** [Agent name]

---

## Error Summary

| Component | Total Errors | Resolved | Pending |
|-----------|-------------|----------|---------|
| scanner   | 0           | 0        | 0       |
| embedding | 0           | 0        | 0       |
| vectorstore | 0         | 0        | 0       |
| server    | 0           | 0        | 0       |
| main      | 0           | 0        | 0       |

---

## Monitoring Instructions

### Check for new errors:
```bash
# View recent errors
type data\errors.log

# Or use the API (when server is running)
curl http://localhost:8080/api/errors
```

### After fixing an error:
1. Copy the error entry from `data/errors.log` to this file
2. Add the resolution details
3. Remove or clear the resolved error from the log file
4. Update the summary table above

### Clear error log (after all resolved):
```bash
# Manual clear
echo. > data\errors.log

# Or use the application
go run cmd/main.go clear-errors
```

---

## Common Error Patterns

### Scanner Errors
- **Access denied**: File/folder permissions issue
- **Read failed**: File locked or corrupted

### Embedding Errors
- **Text too short**: Skipped automatically
- **Dimension mismatch**: Check vector dimension parameter

### Vectorstore Errors
- **Database locked**: Another process is using the DB
- **Insert failed**: Data format issue

### Server Errors
- **Port in use**: Change port in main.go
- **Connection failed**: Database not initialized

---

*Last updated: 2025-03-30*
