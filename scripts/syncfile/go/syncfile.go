package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
)

func main() {
	raw := "https://raw.githubusercontent.com/"
	src := raw + "Neved4/homebrew-tap/main/README.md"
	out := "README.md"

	err := syncFile(src, out)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Files synchronized.")
}

func syncFile(src, dst string) error {
	resp, err := http.Get(src)
	if err != nil {
		return fmt.Errorf("fetching source: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("reading response body: %w", err)
	}

	re := regexp.MustCompile(`(?s)<!-- START SYNC -->(.*?)<!-- END SYNC -->`)

	match := re.FindSubmatch(body)
	if len(match) < 2 {
		return fmt.Errorf("reading output file: %w", err)
	}

	output, err := os.ReadFile(dst)
	if err != nil {
		return err
	}

	return os.WriteFile(dst, re.ReplaceAll(output, match[0]), 0644)
}
