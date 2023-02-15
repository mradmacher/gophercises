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
)

type Problem struct {
    Question string
    Answer string
}

type Score struct {
    Correct int
    Wrong int
}

type Quiz struct {
    problems []Problem
    score Score
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

func (quiz *Quiz) resetScore() {
    quiz.score = Score{}
}

func (quiz *Quiz) updateScore(solved bool) {
    if solved {
        quiz.score.Correct++
    } else {
        quiz.score.Wrong++
    }
}

func (quiz *Quiz) result() string {
    return fmt.Sprintf("%v/%v", quiz.score.Correct, quiz.score.Correct + quiz.score.Wrong)
}

func getGuess(question string) string {
    fmt.Printf("%v: ", question)
    var guess string
    fmt.Scanln(&guess)
    return guess
}

func printScore(score string) {
    fmt.Printf("Your score: %s\n", score)
}

func (problem *Problem) isSolved(guess string) bool {
    return guess == problem.Answer
}

func main() {
    fileName := flag.String("file", "problems.csv", "a CSV file with the quiz in format of 'question,answer'")
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

    quiz.resetScore()
    for _, problem := range quiz.problems {
        quiz.updateScore(problem.isSolved(getGuess(problem.Question)))
    }
    printScore(quiz.result())
}
