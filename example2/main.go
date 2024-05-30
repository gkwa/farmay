package example2

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"go.abhg.dev/goldmark/frontmatter"
	"gopkg.in/yaml.v3"
)

func Run() {
	markdown := "testdata/example2.md"

	file, err := os.Open(markdown)
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

	var meta map[string]interface{}
	if err := fm.Decode(&meta); err != nil {
		log.Fatal(err)
	}

	// Add a new key-value pair to the front matter
	meta["fart2"] = "new value"

	// Encode the updated front matter back to YAML
	updatedFrontMatter, err := yaml.Marshal(meta)
	if err != nil {
		log.Fatal(err)
	}

	// Write the updated front matter and content back to the file
	file, err = os.Create(markdown)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.Write([]byte("---\n"))
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.Write(updatedFrontMatter)
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.Write([]byte("---\n"))
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.Write(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Title: %s\n", meta["title"])
	fmt.Printf("Fart: %s\n", meta["fart"])
	fmt.Printf("Fart2: %s\n", meta["fart2"])
	fmt.Printf("Tags: %v\n", meta["tags"])

	fmt.Println("\nParsed Markdown:")
	fmt.Println(buf.String())
}
