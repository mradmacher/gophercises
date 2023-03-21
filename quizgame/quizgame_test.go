package main

import (
    "testing"
    "gotest.tools/assert"
    "strings"
)

func TestProblemIsSolvedWithCorrectGuess(t *testing.T) {
    problem := Problem{"question", "answer"}

    assert.Assert(t, problem.isSolved("answer"))
}

func TestProblemIsNotSolvedWithWrongGuess(t *testing.T) {
    problem := Problem{"question", "answer"}

    assert.Assert(t, !problem.isSolved("wrong"))
}

func TestLoadsProblemsFromReader(t *testing.T) {
    source := "q1,a1\nq2,a2\nq3,a3"
    quiz := Quiz{}
    assert.NilError(t, quiz.loadProblems(strings.NewReader(source)))

    assert.Equal(t, 3, len(quiz.problems))
    assert.Equal(t, "q1", quiz.problems[0].Question)
    assert.Equal(t, "a1", quiz.problems[0].Answer)

    assert.Equal(t, "q2", quiz.problems[1].Question)
    assert.Equal(t, "a2", quiz.problems[1].Answer)

    assert.Equal(t, "q3", quiz.problems[2].Question)
    assert.Equal(t, "a3", quiz.problems[2].Answer)
}

func TestFailsReadingProblemsFromIncorrectSource(t *testing.T) {
    quiz := Quiz{}

    source := "q1\nq2,a2\nq3,a3"
    err := quiz.loadProblems(strings.NewReader(source))
    assert.ErrorContains(t, err, "wrong number of fields")

    source = "q1,a1\nq2\nq3,a3"
    err = quiz.loadProblems(strings.NewReader(source))
    assert.ErrorContains(t, err, "wrong number of fields")
}

func TestRunsQuizProvidingQuestionsAndReceivingAnswers(t *testing.T) {
    quiz := Quiz{}
    source := "q1,a1\nq2,a2\nq3,a3"
    quiz.loadProblems(strings.NewReader(source))
    questionCh := make(chan string)
    answerCh := make(chan string)
    quitCh := make(chan bool)
    scoreCh := make(chan Score)

    var question string
    go func() {
        scoreCh <- quiz.run(questionCh, answerCh, quitCh)
    }()

    question = <- questionCh
    assert.Equal(t, "q1", question)
    answerCh <- "a1"

    question = <- questionCh
    assert.Equal(t, "q2", question)
    answerCh <- "wrong"

    question = <- questionCh
    assert.Equal(t, "q3", question)
    answerCh <- "a3"

    score := <- scoreCh

    assert.Equal(t, 2, score.Solved)
    assert.Equal(t, 3, score.Total)
}

func TestQuizCanBeTimeouted(t *testing.T) {
    quiz := Quiz{}
    source := "q1,a1\nq2,a2"
    quiz.loadProblems(strings.NewReader(source))
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

    quitCh <- true
    score := <- scoreCh

    assert.Equal(t, 1, score.Solved)
    assert.Equal(t, 2, score.Total)
}
