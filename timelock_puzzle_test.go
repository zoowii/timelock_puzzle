package timelock_puzzle

import (
	"testing"
	"log"
	"github.com/stretchr/testify/assert"
	"encoding/json"
)

func TestGeneratePrime(t *testing.T) {
	x, err := GeneratePrime(DefaultBits)
	if err != nil {
		log.Println(err.Error())
	}
	assert.True(t, err == nil)
	log.Println("x =", x.Text(10))
}

func TestGenerateRandomBigInt(t *testing.T) {
	x, err := GenerateRandomBigInt(DefaultBits)
	if err != nil {
		log.Println(err.Error())
	}
	assert.True(t, err == nil)
	log.Println("x =", x.Text(10))
}

func TestGenerateTimeLockPuzzle(t *testing.T) {
	puzzleT, err := GenerateRandomBigInt(100)
	if err != nil {
		log.Println(err.Error())
	}
	assert.True(t, err == nil)
	log.Println("T =", puzzleT.Text(10))
	x, err := GenerateRandomBigInt(100)
	if err != nil {
		log.Println(err.Error())
	}
	assert.True(t, err == nil)
	log.Println("x =", x.Text(10))
	puzzle, err := GenerateTimeLockPuzzle(DefaultBits, puzzleT, x)
	if err != nil {
		log.Println(err.Error())
	}
	assert.True(t, err == nil)
	puzzleBytes, err := json.Marshal(puzzle)
	assert.True(t, err == nil)
	log.Println("puzzle =", string(puzzleBytes))
}

func TestSolveByPrivate(t *testing.T) {
	puzzleT, err := GenerateRandomBigInt(10)
	if err != nil {
		log.Println(err.Error())
	}
	assert.True(t, err == nil)
	log.Println("T =", puzzleT.Text(10))
	x, err := GenerateRandomBigInt(10)
	if err != nil {
		log.Println(err.Error())
	}
	assert.True(t, err == nil)
	log.Println("x =", x.Text(10))
	puzzle, err := GenerateTimeLockPuzzle(10, puzzleT, x)
	if err != nil {
		log.Println(err.Error())
	}
	assert.True(t, err == nil)
	puzzleBytes, err := json.Marshal(puzzle)
	assert.True(t, err == nil)
	log.Println("puzzle =", string(puzzleBytes))
	y, err := SolveByPrivate(puzzle)
	if err != nil {
		log.Println(err.Error())
	}
	assert.True(t, err == nil)
	println("y =", y.Text(10))
}

func TestSolveByPublic(t *testing.T) {
	puzzleT, err := GenerateRandomBigInt(3)
	if err != nil {
		log.Println(err.Error())
	}
	assert.True(t, err == nil)
	log.Println("T =", puzzleT.Text(10))
	x, err := GenerateRandomBigInt(6)
	if err != nil {
		log.Println(err.Error())
	}
	assert.True(t, err == nil)
	log.Println("x =", x.Text(10))
	puzzle, err := GenerateTimeLockPuzzle(6, puzzleT, x)
	if err != nil {
		log.Println(err.Error())
	}
	assert.True(t, err == nil)
	puzzleBytes, err := json.Marshal(puzzle)
	assert.True(t, err == nil)
	log.Println("puzzle =", string(puzzleBytes))
	y, err := SolveByPrivate(puzzle)
	if err != nil {
		log.Println(err.Error())
	}
	assert.True(t, err == nil)
	println("y by private =", y.Text(10))
	yByPublic, err := SolveByPublic(puzzle)
	if err != nil {
		log.Println(err.Error())
	}
	assert.True(t, err == nil)
	println("y by public =", yByPublic.Text(10))
}
