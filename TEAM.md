# AI Embedding Offline - Team Registry

## Agent Registration Log

| Agent | Model | Role | Work Done | DateTime |
|-------|-------|------|-----------|----------|
| Qwen-Code | Qwen-2.5-Coder | Lead Programmer + Go Developer + Web Developer | Project setup, folder structure, core implementation, build verification | 2025-03-30 |
| Qwen-Code | Qwen-2.5-Coder | Debugger | Fixed unused import in vectorstore.go | 2025-03-30 |
| Qwen-Code | Qwen-2.5-Coder | UX/UI Designer | Created interactive HTML visualization with zoom/pan/12px labels | 2025-03-30 |
| Qwen-Code | Qwen-2.5-Coder | Error Monitor | Added error logging system to all packages | 2025-03-30 |
| Qwen-Code | Qwen-2.5-Coder | Documentation Writer | Added Beginner's Guide (EN/TH) with visual explanations | 2025-03-30 |
| Qwen-Code | Qwen-2.5-Coder | ML Engineer | Added N-gram text predictor (practical alternative) | 2025-03-30 |

## Agent Roles

1. **Lead Programmer** - Overall architecture, code review, coordination
2. **Web Developer** - HTML/CSS/JS frontend, visualization
3. **Debugger** - Testing, error fixing
4. **Server Developer** - Backend API, server logic
5. **Go Lang Developer** - Go implementation, optimization
6. **UX/UI Designer** - User interface, visualization design
7. **Error Monitor** - Monitors errors, logs to text file, tracks resolutions
8. **Documentation Writer** - Creates user guides and technical documentation
9. **ML Engineer** - Implements prediction models, optimizes for practical use

## Instructions for Agents

- Register your model when starting work
- Log completed work in this file
- Use minimal tokens in communication
- Append to log files instead of re-reading

## Project Status: ✅ COMPLETE

All requirements from prompt01.txt have been implemented.

**Pivot Decision:** Original hash-based embedding approach was NOT suitable for meaningful text prediction. Switched to N-gram model which:
- Actually predicts next words based on context
- Small model size (~100KB vs GB)
- Fast training (seconds vs hours)
- Easy to understand and debug

## Error Monitoring System

- Errors are logged to `data/errors.log`
- Resolved errors are tracked in `docs/SOLVED_ERRORS.md`
- All packages now have error logging integrated

## Documentation

- `README.md` - Main documentation (EN/TH)
- `docs/BEGINNER_GUIDE.md` - Simple explanations for beginners
- `docs/ABSTRACT.md` - Technical abstract
- `docs/VECTOR_DB_USAGE.md` - Database usage guide
- `docs/NGRAM_PREDICTOR.md` - N-gram prediction guide
- `TODO.md` - Future improvements and BCP
