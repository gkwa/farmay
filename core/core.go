package core

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"go.abhg.dev/goldmark/frontmatter"
)

func Run() {
	file, err := os.Open("testdata/example.md")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var buf bytes.Buffer
	md := goldmark.New(
		goldmark.WithExtensions(&frontmatter.Extender{}),
	)

	ctx := parser.NewContext()
	if err := md.Convert(file, &buf, parser.WithContext(ctx)); err != nil {
		log.Fatal(err)
	}

	fm := frontmatter.Get(ctx)

	var meta struct {
		Title string   `yaml:"title"`
		Tags  []string `yaml:"tags"`
	}
	if err := fm.Decode(&meta); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Title: %s\n", meta.Title)
	fmt.Printf("Tags: %v\n", meta.Tags)

	fmt.Println("\nParsed Markdown:")
	fmt.Println(buf.String())
}
