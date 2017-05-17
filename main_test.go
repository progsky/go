package main

import (
	"reflect"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKnightMoves(t *testing.T) {
	c := func(place string, movePlaces []string) bool {
		got := knightMoves(place)
		sort.Strings(got)
		sort.Strings(movePlaces)
		return reflect.DeepEqual(got, movePlaces)
	}
	assert.True(t, c("a8", []string{"c7", "b6"}))
	assert.True(t, c("h8", []string{"f7", "g6"}))
	assert.True(t, c("h1", []string{"f2", "g3"}))
	assert.True(t, c("a1", []string{"b3", "c2"}))
	assert.True(t, c("b7", []string{"d8", "d6", "c5", "a5"}))
	assert.True(t, c("g7", []string{"e6", "e8", "h5", "f5"}))
	assert.True(t, c("g2", []string{"e1", "e3", "f4", "h4"}))
	assert.True(t, c("b2", []string{"a4", "c4", "d3", "d1"}))
	assert.True(t, c("b2", []string{"a4", "c4", "d3", "d1"}))
	assert.True(t, c("d7", []string{"b8", "b6", "c5", "e5", "f6", "f8"}))
	assert.True(t, c("g4", []string{"h6", "f6", "e5", "e3", "f2", "h2"}))
	assert.True(t, c("d2", []string{"b1", "b3", "c4", "e4", "f3", "f1"}))
	assert.True(t, c("b5", []string{"a7", "c7", "d6", "d4", "a3", "c3"}))
	assert.True(t, c("d4", []string{"b5", "b3", "c6", "e6", "f5", "f3", "c2", "e2"}))
}
