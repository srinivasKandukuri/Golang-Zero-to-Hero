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

// call http
// download the file
// create local dir
// save the file in local dir
// return
func fetchURL(ctx context.Context, url string, id int, outdir string) error {

	res, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}

	client := http.Client{Timeout: 5 * time.Second}

	resp, err := client.Do(res)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status %s : %d", url, resp.StatusCode)
	}

	tokens := strings.Split(url, "/")
	filePath := tokens[len(tokens)-1]

	file := filepath.Join(outdir, filePath)
	filedir, err := os.Create(fmt.Sprintf("%s_%d", file, id))
	if err != nil {
		return err
	}

	defer func(filedir *os.File) {
		err := filedir.Close()
		if err != nil {

		}
	}(filedir)

	_, err = io.Copy(filedir, resp.Body)

	if err != nil {
		return err
	}
	fmt.Printf("Sucessfully downloaded %s\n", filedir.Name())
	return nil
}

func main() {
	urls := []string{
		"https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf",
		"https://www.learningcontainer.com/wp-content/uploads/2020/05/sample-zip-file.zip",
		"https://nonexistent.example.com",
	}
	outdir := "download"
	err := os.MkdirAll(outdir, 0755)

	if err != nil {
		fmt.Errorf("failed to create folder")
	}

	// using err group download above urls parallel using go routines
	// use context timeout
	// download the files from given url

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	g, gtx := errgroup.WithContext(ctx)

	for id, url := range urls {

		id := id
		url := url

		g.Go(func() error {
			return fetchURL(gtx, url, id, outdir)
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Printf("Error occured %v\n", err)
	} else {
		fmt.Println("all urls fetched sucessfully")
	}
}
