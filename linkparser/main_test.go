package main

import (
  "testing"
  "bytes"
  "strings"
)

func TestMain_example1(t *testing.T) {
    args := []string{"linkparser", "examples/example1.html"}
    var stdout bytes.Buffer

    err := run(args, &stdout)
    if err != nil {
        t.Fatalf("Not expected error: %v", err)
    }

    output := stdout.String()
    want := "/other-page => A link to another page"
    if !strings.Contains(output, want) {
        t.Errorf("Got %s; want %s", output, want)
    }
}

func TestMain_example2(t *testing.T) {
    args := []string{"linkparser", "examples/example2.html"}
    var stdout bytes.Buffer

    err := run(args, &stdout)
    if err != nil {
        t.Fatalf("Not expected error: %v", err)
    }

    output := stdout.String()
    want := "https://www.twitter.com/joncalhoun => Check me out on twitter"
    if !strings.Contains(output, want) {
        t.Errorf("Got %s; want %s", output, want)
    }

    want = "https://github.com/gophercises => Gophercises is on Github !"
    if !strings.Contains(output, want) {
        t.Errorf("Got %s; want %s", output, want)
    }
}

func TestMain_example3(t *testing.T) {
    args := []string{"linkparser", "examples/example3.html"}
    var stdout bytes.Buffer

    err := run(args, &stdout)
    if err != nil {
        t.Fatalf("Not expected error: %v", err)
    }

    output := stdout.String()
    want := "# => Login"
    if !strings.Contains(output, want) {
        t.Errorf("Got %s; want %s", output, want)
    }

    want = "/lost => Lost? Need help?"
    if !strings.Contains(output, want) {
        t.Errorf("Got %s; want %s", output, want)
    }

    want = "https://twitter.com/marcusolsson => @marcusolsson"
    if !strings.Contains(output, want) {
        t.Errorf("Got %s; want %s", output, want)
    }
}

func TestMain_example4(t *testing.T) {
    args := []string{"linkparser", "examples/example4.html"}
    var stdout bytes.Buffer

    err := run(args, &stdout)
    if err != nil {
        t.Fatalf("Not expected error: %v", err)
    }

    output := stdout.String()
    want := "/dog-cat => dog cat"
    if !strings.Contains(output, want) {
        t.Errorf("Got %s; want %s", output, want)
    }
}
