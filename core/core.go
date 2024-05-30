package core

import (
	"bytes"
	"fmt"
	"log"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"go.abhg.dev/goldmark/frontmatter"
)

func Run() {
	src := []byte(`---
title: My Document
tags: [example, demo]
---

# Heading

This is some *Markdown* content.
`)

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
