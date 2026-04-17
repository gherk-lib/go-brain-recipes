package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gherk-lib/go-brain"
	"github.com/gherk-lib/go-brain/adapters/ollama"
)

func main() {
	// =========================================================================
	// 1. Adapter Configuration
	// =========================================================================
	// We bind to a local Ollama instance. This ensures that IP constraints,
	// PII, and internal proprietary code never leaves the Enterprise subnet.
	// 
	// Pre-requisite:
	// Make sure Ollama is running and you have downloaded the target model:
	// $ ollama run gemma
	// =========================================================================
	adapter := ollama.NewAdapter(ollama.Options{
		// You can specify tags like "gemma:2b" or "gemma:7b".
		// We use "gemma" to map to your default locally pulled gemma variant.
		Model: "gemma", 
		// URL defaults to http://127.0.0.1:11434 if omitted
	})

	// =========================================================================
	// 2. Framework Instantiation
	// =========================================================================
	bot := brain.NewBot(brain.Options{
		Name:    "LocalGemmaNode",
		Adapter: adapter,
	})

	// =========================================================================
	// 3. NeuroSwitch FSM Routing
	// =========================================================================
	// Instead of black-box Python DAGs, we explicitly compile our instruction set.
	bot.WithRouter(func(ctx context.Context, state *brain.State) error {
		fmt.Println(">> [Node: Initializing FSM State]")
		
		state.Prompt(
			"You are an elite Enterprise Software Architect. " +
			"Summarize the technical benefits of moving away from Python Langchain " +
			"towards statically-typed compiled Go orchestration. Keep it under 50 words.",
		)
		
		// Conclude the execution gracefully
		return state.End()
	})

	// =========================================================================
	// 4. Execution & Observability
	// =========================================================================
	fmt.Printf("\n🚀 Booting Go-Brain Framework...\n")
	fmt.Printf("   Adapter: Ollama\n")
	fmt.Printf("   Model: gemma\n\n")

	// Inject a strict timeout context to guarantee deterministic bounds
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	// Run the FSM synchronously
	resp, err := bot.Run(ctx)
	if err != nil {
		log.Fatalf("Fatal Engine Exception: %v", err)
		os.Exit(1)
	}

	// Dump final deterministic memory projection
	fmt.Println("\n🤖 Gemma Response:")
	fmt.Printf("%s\n", resp.LastMessage())
}
