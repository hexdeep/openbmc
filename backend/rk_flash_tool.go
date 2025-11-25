package main

import (
	"context"
	"fmt"
	"os/exec"
)

type RKFlashTool struct {
	ToolPath string
}

func (t *RKFlashTool) Run(ctx context.Context, arg ...string) (string, error) {
	cmd := exec.CommandContext(ctx, t.ToolPath, arg...)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func (t *RKFlashTool) DetectDevice(ctx context.Context) error {

	if _, err := t.Run(ctx, "id"); err != nil {
		return fmt.Errorf("failed to detect device: %w", err)
	}

	return nil
}
