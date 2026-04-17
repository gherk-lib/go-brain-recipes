package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gherk-lib/go-brain"
	"github.com/gherk-lib/go-brain/adapters/ollama"
	"github.com/gherk-lib/go-brain/router"
)

func main() {
	// 1. Adapter Configuration
	adapter := ollama.New("", "gemma")

	// 2. Framework Instantiation
	agent := brain.NewAgent("LocalGemmaNode", adapter)

	// 3. NeuroSwitch FSM Routing
	// Initialize router starting at state START
	agent = agent.WithRouter("START")
	
	agent.Router().AddState("START", func(ctx context.Context, rctx *router.Context) (string, error) {
		fmt.Println(">> [Node: Initializing FSM State]")
		
		llmVal, ok := rctx.GetVar("go-brain:llm")
		if !ok {
			return "", fmt.Errorf("llm engine disconnected")
		}
		llm := llmVal.(brain.LLM)
		
		prompt := "You are an elite Enterprise Software Architect. " +
			"Summarize the technical benefits of moving away from Python Langchain " +
			"towards statically-typed compiled Go orchestration. Keep it under 50 words."
			
		resp, err := llm.Generate(ctx, prompt)
		if err != nil {
			return "", err
		}
		
		rctx.SetResult(resp)
		
		// Conclude the execution gracefully (empty string halts the FSM)
		return "", nil
	})

	// 4. Execution & Observability
	fmt.Printf("\n🚀 Booting Go-Brain Framework...\n")
	fmt.Printf("   Adapter: Ollama\n")
	fmt.Printf("   Model: gemma\n\n")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	// Run the FSM synchronously
	outCtx, err := agent.Run(ctx, nil)
	if err != nil {
		log.Fatalf("Fatal Engine Exception: %v", err)
		os.Exit(1)
	}

	fmt.Println("\n🤖 Gemma Response:")
	fmt.Printf("%s\n", outCtx.GetResult().(string))
}
