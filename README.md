# Gopilot

## Architecture

### Document Processing Pipeline

* Markdown Parser: Read and parse .md files from a local directory
* Text Chunker: Split documents into smaller chunks (e.g., 500-1000 tokens)
* Metadata Extractor: Extract file names, headers, creation dates

### Vector Database

* Use Qdrant
* Store document chunks with their embeddings
* Enable similarity search

### Embedding Service

* Use OpenAI's embedding API (text-embedding-3-small)
* Create embeddings for each document chunk
* Store in vector DB

### RAG Pipeline

* Query Processing: Convert user questions to embeddings
* Retrieval: Find top-k similar chunks from vector DB
* Context Building: Combine retrieved chunks
* Generation: Send context + query to OpenAI Chat API

### API/Interface

* Simple REST API or gRPC service
* Web UI (could use HTMX + Go templates for simplicity)


## Project Structure

```
gopilot/
├── build/                # Build artifacts
├── cmd/                  # Main applications
│   └── server/           # HTTP server entry point
├── deployments/          # Docker, K8s configs
├── docs/                 # Documentation & samples
├── internal/             # Private application code
│   ├── indexer/          # Document indexing logic
│   ├── embeddings/       # OpenAI embedding wrapper
│   ├── vectordb/         # Vector DB abstraction
│   ├── rag/              # RAG pipeline orchestration
│   ├── api/              # HTTP handlers & middleware
│   └── config/           # Configuration management
├── pkg/                  # Public libraries (if needed)
│   ├── chunker/          # Text chunking algorithms
│   └── markdown/         # Markdown parsing utilities
├── scripts/              # Build/deployment scripts
└── tests/                # Test folder
└── web/                  # Web assets
```

## Tools & Dependencies

Gin Tonic
Qdrant