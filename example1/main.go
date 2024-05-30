package example1

import (
	"bytes"
	"fmt"
	"io"
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

	src, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	var buf bytes.Buffer
	md := goldmark.New(
		goldmark.WithExtensions(&frontmatter.Extender{}),
	)

	ctx := parser.NewContext()
	if err := md.Convert(src, &buf, parser.WithContext(ctx)); err != nil {
		log.Fatal(err)
	}

	fm := frontmatter.Get(ctx)

	var meta struct {
		Title string   `yaml:"title"`
		Fart  string   `yaml:"fart"`
		Fart2 string   `yaml:"fart2"`
		Tags  []string `yaml:"tags"`
	}
	if err := fm.Decode(&meta); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Title: %s\n", meta.Title)
	fmt.Printf("Fart: %s\n", meta.Fart)
	fmt.Printf("Fart2: %s\n", meta.Fart2)
	fmt.Printf("Tags: %v\n", meta.Tags)

	fmt.Println("\nParsed Markdown:")
	fmt.Println(buf.String())
}
