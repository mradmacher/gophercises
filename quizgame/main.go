package main

import (
    "encoding/csv"
    "errors"
    "flag"
    "fmt"
    "io"
    "log"
    "os"
    "bytes"
    "time"
)

type Problem struct {
    Question string
    Answer string
}

type Quiz struct {
    problems []Problem
}

type Score struct {
    Total int
    Solved int
}

func (problem *Problem) isSolved(guess string) bool {
    return guess == problem.Answer
}

func (quiz *Quiz) loadProblems(source io.Reader) error {
    quiz.problems = make([]Problem, 0)
    csvReader := csv.NewReader(source)
    for {
        record, err := csvReader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            return err
        }
        if len(record) != 2 {
          return errors.New("wrong number of fields")
        }
        quiz.problems = append(quiz.problems, Problem{record[0], record[1]})
    }
    return nil
}

func (quiz *Quiz) run(questionCh, answerCh chan string, quitCh chan bool) Score {
    doneCh := make(chan bool)
    score := Score{Total: len(quiz.problems), Solved: 0}

    go func() {
        for _, problem := range quiz.problems {
            questionCh <- problem.Question
            if problem.isSolved(<-answerCh) {
                score.Solved++
            }
        }
        doneCh <- true
    }()
    select {
        case <-quitCh:
        case <-doneCh:
    }
    return score
}

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
