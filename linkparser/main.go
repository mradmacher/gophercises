package main

import (
    "fmt"
    "log"
    "strings"
    "golang.org/x/net/html"
    "golang.org/x/net/html/atom"
)

type Link struct {
    Href string
    Text string
}

func collectHref(n *html.Node) string {
    for _, a := range n.Attr {
        if a.Key == "href" { return a.Val }
    }
    return ""
}

func collectText(n *html.Node) string {
    if n == nil { return "" }

    var texts []string
    if n.Type == html.TextNode {
        texts = append(texts, n.Data)
    } else {
        for c := n.FirstChild; c != nil; c = c.NextSibling {
            texts = append(texts, collectText(c))
        }

    }
    return strings.Join(texts, " ")
}

func collectLinks(n *html.Node, resultCh chan Link) {
    if n.Type == html.ElementNode && n.DataAtom == atom.A {
        resultCh <- Link{collectHref(n), collectText(n)}
    } else {
        for c := n.FirstChild; c != nil; c = c.NextSibling {
            collectLinks(c, resultCh)
        }
    }
}

func reportLinks(doc *html.Node) chan Link {
    resultCh := make(chan Link)
    go func() {
        collectLinks(doc, resultCh)
        close(resultCh)
    }()
    return resultCh
}

func main() {
    s := `<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">xyz</a></ul><a href="text"><span>Test<i>cur</i></span><span>Dupa</span></a>`
    doc, err := html.Parse(strings.NewReader(s))
    if err != nil {
        log.Fatal(err)
    }
    resultCh := reportLinks(doc)
    for link := range resultCh {
        fmt.Printf("%v => %v\n", link.Href, link.Text)
    }
}
