# Go-Brain Enterprise SDK - Public Recipes

This repository contains public reference examples and architectural patterns for the proprietary **Go-Brain Enterprise SDK**.

> [!WARNING]
> **Compilation Notice:** The `go-brain` core SDK is a closed-source B2B component. Attempting to download or compile these examples without proper authentication to the `github.com/gherk-lib` registry will fail. You must configure your local Git environment using an Enterprise License Token first.
> ```bash
> go env -w GOPRIVATE="github.com/gherk-lib/*"
> ```

## 📚 Available Recipes

- **[01-ollama-gemma](01-ollama-gemma/)**: Synchronous execution using the Ollama adapter and the Gemma model. Demonstrates zero-network-dependency architecture.
- **[02-swarm-sandbox](02-swarm-sandbox/)**: Advanced multi-agent cluster (`Broker`), isolated local code execution (`GitWorktree`), and real-time terminal user input processing.
