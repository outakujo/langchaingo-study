package main

import (
	"context"
	"fmt"
	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/llms/local"
	"github.com/tmc/langchaingo/tools"
)

type MyTool struct {
}

func (m MyTool) Name() string {
	return "mytool"
}

func (m MyTool) Description() string {
	return "Description"
}

func (m MyTool) Call(ctx context.Context, s string) (string, error) {
	return fmt.Sprintf("%s 哈哈", s), nil
}

func main() {
	bin := local.WithBin("./mockllm")
	llm, err := local.New(bin)
	if err != nil {
		panic(err)
	}
	prompt := "哈哈哈"
	completion, err := llm.Call(context.Background(), prompt)
	if err != nil {
		panic(err)
	}
	fmt.Println(completion)
	exe, err := agents.Initialize(llm, []tools.Tool{MyTool{}},
		agents.ZeroShotReactDescription)
	if err != nil {
		panic(err)
	}
	res, err := exe.Call(context.Background(), map[string]any{
		"input": "sdsd",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(res["output"])
}
