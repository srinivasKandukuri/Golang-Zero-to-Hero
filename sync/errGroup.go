package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	urls := []string{
		"https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf",
		"https://www.learningcontainer.com/wp-content/uploads/2020/05/sample-zip-file.zip",
	}
	targetDir := "./downloads"
	os.MkdirAll(targetDir, 0755)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()
	g, ctx := errgroup.WithContext(ctx)

	for _, url := range urls {
		u := url
		g.Go(func() error {
			return downloadFile(ctx, u, targetDir)
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Done")
	}
}

func downloadFile(ctx context.Context, url string, targetDir string) error {
	start := time.Now()
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)

	if err != nil {
		fmt.Println(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status for %s: %s", url, resp.Status)
	}

	filePath := filepath.Join(targetDir, getFileName(url))
	out, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("creating file: %w", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("saving file: %w", err)
	}

	fmt.Println("✅ Downloaded:", filePath)
	fmt.Println("<UNK> Took:", time.Since(start))
	return nil
}

// getFileName extracts filename from URL
func getFileName(url string) string {
	tokens := strings.Split(url, "/")
	return tokens[len(tokens)-1]
}
