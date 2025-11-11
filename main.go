// main.go
package main

import (
	"context"
	"fmt"
	"github.com/problem-01/problem/internal/service"
	"github.com/problem-01/problem/providers/aws"
	"os"
	"strings"
	"time"
)

/*
  This file represents a deliberately messy implementation.
  Your job is to refactor and extend it per the README:

  - Add S3Uploader interface + dependency injection
  - Add context.Context support
  - Add exponential backoff retries
  - Add tag support
  - Remove global/hardcoded configuration
  - Add at least one unit test (mock-based)
*/

// ----------------- Current Quick Implementation (Refactor Target) -----------------
func main() {
	if len(os.Args) < 3 {
		fmt.Println("usage: go run main.go upload <path> [key=value,key2=value2]")
		return
	}
	cmd := os.Args[1]

	serv := service.NewUpload(service.WithUploader(aws.New()))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	switch cmd {
	case "upload":
		path := os.Args[2]

		tags := parseTagsArg("")
		if len(os.Args) >= 4 {
			tags = parseTagsArg(os.Args[3])
		}
		_ = path
		_ = tags

		err := serv.Upload(ctx, path, tags)
		if err != nil {
			fmt.Printf("error %v", err)
		}

		// This is where DI should be introduced after refactor
		fmt.Println("upload succeeded")

	default:
		fmt.Println("unknown cmd:", cmd)
	}
}

func parseTagsArg(arg string) map[string]string {
	m := map[string]string{}
	if arg == "" {
		return m
	}
	pairs := strings.Split(arg, ",")
	for _, p := range pairs {
		if p == "" {
			continue
		}
		kv := strings.SplitN(p, "=", 2)
		if len(kv) == 2 {
			m[kv[0]] = kv[1]
		}
	}
	return m
}

// QuickUpload currently uses mock S3 client directly and performs no retry.
// You must refactor and improve this.
