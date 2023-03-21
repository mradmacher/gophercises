package main

import (
    "encoding/csv"
    "errors"
    "io"
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

