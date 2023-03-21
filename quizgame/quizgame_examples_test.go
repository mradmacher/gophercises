package main

import (
  "fmt"
  "strings"
)

func ExampleQuiz() {
    quiz := Quiz{}
    quiz.loadProblems(strings.NewReader("q1,a1\nq2,a2"))

    questionCh := make(chan string)
    answerCh := make(chan string)
    quitCh := make(chan bool)
    scoreCh := make(chan Score)

    go func() {
        scoreCh <- quiz.run(questionCh, answerCh, quitCh)
    }()

    <- questionCh
    answerCh <- "a1"
    <- questionCh
    answerCh <- "a2"

    score := <- scoreCh

    fmt.Printf("%v/%v", score.Solved, score.Total)

    // Output: 2/2
}
