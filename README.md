# Go-Brain Enterprise SDK - Swarm Recipe

This repository demonstrates how to orchestrate a high-performance deterministic FSM agent cluster running completely locally using [Ollama](https://ollama.com/) (Gemma). This setup guarantees that no proprietary code or PII leaves your enterprise network.

> [!WARNING]
> **Compilation Notice:** The `go-brain` core SDK is a closed-source B2B component. Attempting to download or compile this example without proper authentication to the `github.com/gherk-lib` registry will fail. You must configure your local Git environment using an Enterprise License Token first.
> ```bash
> go env -w GOPRIVATE="github.com/gherk-lib/*"
> ```

## 🛠️ Prerequisites & Installation

To run this cluster, you need three basic tools: **Go 1.22+**, the **Ollama Engine**, and the **Gemma Model**.

### 1. Install Go
- **macOS:** `brew install go`
- **Linux:** `sudo apt update && sudo apt install golang-go`

### 2. Install Ollama
Ollama is a lightweight engine to run LLMs locally on CPU or GPU.
Download it from [Ollama's website](https://ollama.com/download) or run `brew install ollama` on Mac.

### 3. Download the Target Model (Gemma)
Once Ollama is installed, verify the service is running, then pull the Gemma model in your terminal:
```bash
ollama pull gemma
```
*Note: This will download a few gigabytes so it might take a few minutes.*

### 4. Authenticate Enterprise SDK
Authorize your local Go environment to securely fetch the private framework logic:
```bash
go env -w GOPRIVATE="github.com/gherk-lib/*"
```

## 🚀 Running the Terminal Swarm

Once the prerequisites are ready, booting your AI orchestrator is instantaneous.

1. **Start the Ollama daemon:** Ensure the background service is running.
2. **Execute the Go script:**

```bash
# Run the interactive Swarm Terminal directly
go run main.go
```

The terminal will initialize the **Researcher** agent connected to a completely isolated Git Sandbox partition of your hard drive, alongside the **Analyst** agent. You will see a `User >` prompt to start the multi-agent investigation.
