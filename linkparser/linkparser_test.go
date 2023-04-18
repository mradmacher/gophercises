package main

import (
    "testing"
    "strings"
)

func TestCollectLinks_fetchesAllLinks(t *testing.T) {
    htmlString := `
    <html>
      <body>
        <p>Links:</p>
        <ul>
          <li>
            <a href="foo">Foo</a>
            <a href="/bar/baz">Bar Baz</a>
          </li>
        </ul>
        <div><div><a href="/go/and/check/me">Here I am</a></div></div>
      </body>
    </html>
    `
    reader := strings.NewReader(htmlString)
    links, err := CollectLinks(reader)
    if err != nil {
        t.Fatalf("Error while collecting links: %v; no error expected", err)
    }

    if len(links) != 3 {
        t.Errorf("Collected %d links; expected 3", len(links))
    }

    var want Link
    want = Link{"foo", "Foo"}
    if got := links[0]; got != want {
        t.Errorf("Got %v; expected %v", got, want)
    }
    want = Link{"/bar/baz", "Bar Baz"}
    if got := links[1]; got != want {
        t.Errorf("Got %v; expected %v", got, want)
    }
    want = Link{"/go/and/check/me", "Here I am"}
    if got := links[2]; got != want {
        t.Errorf("Got %v; expected %v", got, want)
    }
}

func TestCollectLinks_skipsTagsInsideA(t *testing.T) {
    htmlString := `
      <div><a href="/go/and/check/me">Here<span>I<span>am</span></span></a></div>
    `
    reader := strings.NewReader(htmlString)
    links, err := CollectLinks(reader)
    if err != nil {
        t.Fatalf("Error while collecting links: %v; no error expected", err)
    }

    var want Link
    want = Link{"/go/and/check/me", "Here I am"}
    if links[0] != want {
        t.Errorf("Got %v; expected %v", links[0], want)
    }
}

func TestCollectLinks_removesUnnecessarySpaces(t *testing.T) {
    htmlString := `
      <div>
        <a href="/go/and/check/me">Look
          <span>at me.
            <strong>
             Here
            </strong>
            I am!
          </span>
        </a>
      </div>
    `
    reader := strings.NewReader(htmlString)
    links, err := CollectLinks(reader)
    if err != nil {
        t.Fatalf("Error while collecting links: %v; no error expected", err)
    }

    var want Link
    want = Link{"/go/and/check/me", "Look at me. Here I am!"}
    if links[0] != want {
        t.Errorf("Got %v; expected %v", links[0], want)
    }
}

func TestCollectLinks_handlesEmptyHrefAndText(t *testing.T) {
    htmlString := `
      <div><a></a></div>
    `
    reader := strings.NewReader(htmlString)
    links, err := CollectLinks(reader)
    if err != nil {
        t.Fatalf("Error while collecting links: %v; no error expected", err)
    }

    var want Link
    want = Link{"", ""}
    if links[0] != want {
        t.Errorf("Got %v; expected %v", links[0], want)
    }
}
