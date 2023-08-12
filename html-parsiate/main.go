package main

import (
	"errors"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"os"
	"strings"
)

type Link struct {
	href string
	text string
}

func main() {
	if len(os.Args) < 2 {
		exitGracefully(errors.New("A filepath argument is required"))
	}

	htmlFile := os.Args[1]

	htmlContent, err := parseHtmlFile(htmlFile)
	if err != nil {
		exitGracefully(err)
	}

	doc := strings.NewReader(htmlContent)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	_, err = Parse(doc)
	if err != nil {
		exitGracefully(err)
	}
}

func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		exitGracefully(err)
	}

	nodes := linkNodes(doc)
	var links []Link
	for _, node := range nodes {
		links = append(links, buildLink(node))
	}
	fmt.Println(links)
	return nil, nil
}

func buildLink(n *html.Node) Link {
	var ret Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.href = attr.Val
			break
		}
	}
	ret.text = "TODO"
	return ret
}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}

	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}
	return ret
}

func dfs(n *html.Node, padding string) {
	msg := n.Data

	if n.Type == html.ElementNode {
		msg = "<" + msg + ">"
	}
	fmt.Println(msg, padding)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, "  ")
	}
}

func parseHtmlFile(fileName string) (string, error) {
	b, err := os.ReadFile(fileName)

	if err != nil {
		exitGracefully(err)
	}

	return string(b), nil
}

func exitGracefully(err error) {
	fmt.Println("Error:", err)
	os.Exit(1)
}
