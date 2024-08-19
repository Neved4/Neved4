package main

import (
	"io"
	"net/http"
	"os"
	"regexp"
)

func main() {
	raw := "https://raw.githubusercontent.com/"
	src := raw + "Neved4/homebrew-tap/main/README.md"
	out := "README.md"

	syncFile(src, out)
}

func syncFile(src, dst string) {
	resp, err := http.Get(src)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	re := regexp.MustCompile(`(?s)<!-- START SYNC -->(.*?)<!-- END SYNC -->`)
	match := re.FindSubmatch(body)
	output, _ := os.ReadFile(dst)
	os.WriteFile(dst, re.ReplaceAll(output, match[0]), 0644)
}
