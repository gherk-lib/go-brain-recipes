package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gherk-lib/go-brain"
	"github.com/gherk-lib/go-brain/adapters/ollama"
	"github.com/gherk-lib/go-brain/router"
	"github.com/gherk-lib/go-brain/sandbox"
	"github.com/gherk-lib/go-brain/swarm"
)

func main() {
	// =========================================================================
	// 0. Environment Setup
	// =========================================================================
	fmt.Println("🚀 Booting High-Performance Swarm Terminal...")

	// Default to Ollama with Gemma
	adapter := ollama.New("", "gemma")

	// The Broker orchestrates messages asynchronously across agents
	broker := swarm.NewBroker()

	// The Worktree allows agents to run parallel git commands natively 
	// against our local workspace without index.lock collisions.
	currentDir, _ := os.Getwd()
	fmt.Printf("📦 Mounting GitSandbox in: %s\n", currentDir)
	
	branchName := fmt.Sprintf("swarm-research-%d", time.Now().UnixNano())
	worktree, err := sandbox.NewWorktree(currentDir, branchName)
	if err != nil {
		log.Fatalf("Failed to mount sandbox. Are you running inside a git repo? %v", err)
	}
	defer worktree.Teardown() // Deletes the ghost branch before exiting

	// =========================================================================
	// 1. Agent Cluster Initialization
	// =========================================================================
	researcher := brain.NewAgent("Researcher", adapter).WithBroker(broker).WithRouter("WAIT")
	analyst := brain.NewAgent("Analyst", adapter).WithBroker(broker).WithRouter("WAIT")

	var wg sync.WaitGroup
	wg.Add(2)

	// ==========================================
	// Agent A: The Code Explorer (Researcher)
	// ==========================================
	researcher.Router().AddTransition("WAIT", "LOOP", "WAIT")
	researcher.Router().AddState("WAIT", func(c context.Context, rctx *router.Context) (string, error) {
		msgVal, ok := rctx.GetVar("go-brain:last_message")
		if !ok || msgVal == nil {
			// If no message is available, the agent blocks until Swarm preempts it
			<-c.Done()
			return "LOOP", nil
		}

		msg := msgVal.(swarm.Message)
		if msg.To == "Researcher" && msg.Type == "USER_PROMPT" {
			fmt.Printf("\n[Researcher 🔍]: Analyzing the user's objective: '%s'\n", msg.Content)
			fmt.Printf("[Researcher 🛠️ ]: Executing commands via GitSandbox enclosure...\n")
			
			// 1. Safe Native Execution
			outStatus, _ := worktree.Execute("git", "status")
			outLog, _ := worktree.Execute("git", "log", "-n", "3", "--oneline")

			// 2. Extracted LLM Reasoning
			llmVal, _ := rctx.GetVar("go-brain:llm")
			llm := llmVal.(brain.LLM)
			
			prompt := fmt.Sprintf(
				"You are an isolated Code Researcher. The user asked: '%s'. "+
				"You ran 'git status' and 'git log -n 3' and found:\nStatus:\n%s\nLog:\n%s\n"+
				"Summarize this repository's current state in exactly 2 concise sentences.",
				msg.Content, outStatus, outLog,
			)
			fmt.Printf("[Researcher 🧠]: Generating structural summary using local GPU...\n")
			resp, _ := llm.Generate(c, prompt)
			
			// 3. Publish findings downstream
			fmt.Printf("[Researcher 📤]: Handing off findings to Analyst Component via Broker...\n")
			broker.Publish("Analyst", swarm.Message{
				From:    "Researcher",
				To:      "Analyst",
				Content: resp,
				Type:    "RAW_FINDINGS",
			})
			
			rctx.SetVar("go-brain:last_message", nil)
		}
		
		return "LOOP", nil // Wait for next message
	})

	go func() {
		defer wg.Done()
		researcher.Run(context.Background(), nil)
	}()

	// ==========================================
	// Agent B: The Reviewer (Analyst)
	// ==========================================
	analyst.Router().AddTransition("WAIT", "LOOP", "WAIT")
	analyst.Router().AddState("WAIT", func(c context.Context, rctx *router.Context) (string, error) {
		msgVal, ok := rctx.GetVar("go-brain:last_message")
		if !ok || msgVal == nil {
			<-c.Done()
			return "LOOP", nil
		}

		msg := msgVal.(swarm.Message)
		if msg.To == "Analyst" && msg.Type == "RAW_FINDINGS" {
			fmt.Printf("\n[Analyst 🧐]: Received structured findings. Formulating strict response...\n")

			llmVal, _ := rctx.GetVar("go-brain:llm")
			llm := llmVal.(brain.LLM)
			
			prompt := fmt.Sprintf(
				"You are the Lead Analyst connecting natively to the user terminal. "+
				"The internal researcher sent you this code summary:\n%s\n"+
				"Re-package this information as a polite, highly corporate technical answer directly to the user.",
				msg.Content,
			)
			
			fmt.Printf("[Analyst 🧠]: Generating response using local GPU...\n")
			resp, _ := llm.Generate(c, prompt)
			
			fmt.Printf("\n==================== [FINAL OUTPUT] ====================\n")
			fmt.Printf("🤖 Analyst:\n%s\n", resp)
			fmt.Printf("========================================================\n")
			
			rctx.SetVar("go-brain:last_message", nil)
			
			// Mission accomplished, tear down application wrapper
			os.Exit(0)
		}
		
		return "LOOP", nil
	})

	go func() {
		defer wg.Done()
		analyst.Run(context.Background(), nil)
	}()

	// =========================================================================
	// User Input Terminal (Main Routine)
	// =========================================================================
	time.Sleep(2 * time.Second) // Lock step stabilization
	fmt.Println("\n==================================================")
	fmt.Println("💻 User Terminal Active (Waiting for standard input)")
	fmt.Println("==================================================")
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter requirement > ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Swarm Ignition Sequence
	broker.Publish("Researcher", swarm.Message{
		From:    "User",
		To:      "Researcher",
		Content: input,
		Type:    "USER_PROMPT",
	})

	wg.Wait()
}
