# Recipe 01: Go-Brain with Local Ollama (Gemma)

This example demonstrates how to orchestrate a deterministic FSM agent running completely locally using [Ollama](https://ollama.com/) and Google's [Gemma](https://blog.google/technology/developers/gemma-open-models/) model. This setup guarantees that no proprietary code or PII leaves your enterprise network.

## 🛠️ Prerequisites & Installation

To run this recipe, you need three basic tools: **Go 1.22+**, the **Ollama Engine**, and the **Gemma Model**.

### 1. Install Go
If you don't have Go installed, use your package manager:
- **macOS:** `brew install go`
- **Linux (Debian/Ubuntu):** `sudo apt update && sudo apt install golang-go`
- **Windows:** `winget install GoLang.Go`

### 2. Install Ollama
Ollama is a lightweight engine to run LLMs locally on CPU or GPU.
- **macOS:** Download the app from [Ollama's website](https://ollama.com/download) or run `brew install ollama`.
- **Linux:** `curl -fsSL https://ollama.com/install.sh | sh`
- **Windows:** Download the executable installer from the [Ollama Download Page](https://ollama.com/download/windows).

### 3. Download the Target Model (Gemma)
Once Ollama is installed, verify the service is running, then pull the Gemma model (often referred to casually in prompt instructions as Gemma 4b/7b/2b). In your terminal, run:

```bash
ollama pull gemma
```
*Note: This will download a few gigabytes (depending on the precision), so it might take a few minutes on the first run.*

### 4. Authenticate Enterprise SDK
Because `go-brain` is a proprietary enterprise framework, you must authorize your local Go environment to securely fetch the private dependencies:
```bash
go env -w GOPRIVATE="github.com/gherk-lib/*"
```

## 🚀 Running the Project

Once the prerequisites are ready, booting your AI Node is instantaneous.

1. **Start the Ollama daemon:** (On macOS/Windows, just open the Ollama app. On Linux, ensure `systemctl status ollama` is active).
2. **Execute the Go script:**

```bash
# Navigate to the recipe folder
cd 01-ollama-gemma/

# Run the orchestration node natively
go run main.go
```

### Expected Output
The framework will compile and boot immediately. It binds to the local Ollama socket (`http://127.0.0.1:11434`), issues the prompt defined in the FSM state, and streams the answer back synchronously without relying on Python or network latency.

```bash
>> [Node: Initializing FSM State]

🚀 Booting Go-Brain Framework...
   Adapter: Ollama
   Model: gemma


🤖 Gemma Response:
Statically typed FSM orchestration provides strict compile-time error catching and absolute routing determinism. This eliminates runtime hallucinations common in dynamic DAGs and guarantees massive enterprise-grade scalability.
```
