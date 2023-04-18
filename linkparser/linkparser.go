package main

import (
    "io"
    "golang.org/x/net/html"
    "golang.org/x/net/html/atom"
    "strings"
)

type Link struct {
    Href string
    Text string
}

func fetchHref(n *html.Node) string {
    for _, a := range n.Attr {
        if a.Key == "href" { return a.Val }
    }
    return ""
}

func fetchText(n *html.Node) string {
    if n == nil { return "" }

    if n.Type == html.TextNode {
        return strings.Join(strings.Fields(n.Data), " ")
    }

    var texts []string
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        text := fetchText(c)
        if text != "" {
            texts = append(texts, text)
        }
    }

    return strings.Join(texts, " ")
}

func fetchLinks(n *html.Node, resultCh chan Link) {
    if n.Type == html.ElementNode && n.DataAtom == atom.A {
        resultCh <- Link{fetchHref(n), fetchText(n)}
    } else {
        for c := n.FirstChild; c != nil; c = c.NextSibling {
            fetchLinks(c, resultCh)
        }
    }
}

func CollectLinks(reader io.Reader) ([]Link, error) {
    doc, err := html.Parse(reader)
    if err != nil { return nil, err }

    resultCh := make(chan Link)
    go func() {
        fetchLinks(doc, resultCh)
        close(resultCh)
    }()
    var links []Link
    for link := range resultCh {
        links = append(links, link)
    }
    return links, nil
}
