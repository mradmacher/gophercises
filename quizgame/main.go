package main

import (
    "flag"
    "fmt"
    "log"
    "os"
    "bytes"
    "time"
)

func main() {
    fileName := flag.String("file", "problems.csv", "a CSV file with the quiz in format of 'question,answer'")
    timeLimit := flag.Int("limit", 10, "time limit in seconds to solve the quiz")
    flag.Parse()

    buffer, err := os.ReadFile(*fileName)
    if err != nil {
        log.Fatal(err)
    }

    quiz := Quiz{}

    err = quiz.loadProblems(bytes.NewReader(buffer))
    if err != nil {
        log.Fatal(err)
    }

    questionCh := make(chan string)
    answerCh := make(chan string)
    quitCh := make(chan bool)

    go func() {
        for question := range questionCh {
            var guess string
            fmt.Printf("%v: ", question)
            fmt.Scanln(&guess)
            answerCh <- guess
        }
    }()
    go func() {
        <- time.After(time.Duration(*timeLimit) * time.Second)
        quitCh <- true
    }()

    score := quiz.run(questionCh, answerCh, quitCh)

    fmt.Printf("\nYour score: %v/%v\n", score.Solved, score.Total)
}
