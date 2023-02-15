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

func TestScoreIsResetAtTheBeginning(t *testing.T) {
    quiz := Quiz{}

    assert.Equal(t, 0, quiz.score.Correct)
    assert.Equal(t, 0, quiz.score.Wrong)
}

func TestUpdatesScoreWithCorrectGuess(t *testing.T) {
    quiz := Quiz{}
    quiz.updateScore(true)

    assert.Equal(t, 1, quiz.score.Correct)
    assert.Equal(t, 0, quiz.score.Wrong)

    quiz.updateScore(true)

    assert.Equal(t, 2, quiz.score.Correct)
    assert.Equal(t, 0, quiz.score.Wrong)
}

func TestUpdatesScoreWithWrongGuess(t *testing.T) {
    quiz := Quiz{}
    quiz.updateScore(false)

    assert.Equal(t, 0, quiz.score.Correct)
    assert.Equal(t, 1, quiz.score.Wrong)

    quiz.updateScore(false)

    assert.Equal(t, 0, quiz.score.Correct)
    assert.Equal(t, 2, quiz.score.Wrong)
}

func TestUpdatesScoreWithManyGuesses(t *testing.T) {
    quiz := Quiz{}
    quiz.updateScore(false)

    assert.Equal(t, 0, quiz.score.Correct)
    assert.Equal(t, 1, quiz.score.Wrong)

    quiz.updateScore(true)

    assert.Equal(t, 1, quiz.score.Correct)
    assert.Equal(t, 1, quiz.score.Wrong)

    quiz.updateScore(true)

    assert.Equal(t, 2, quiz.score.Correct)
    assert.Equal(t, 1, quiz.score.Wrong)
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
