package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/mock"
	// "github.com/stretchr/testify/http"
)

func TestAdd(t *testing.T) {
	assert.Equal(t, add(1, 2), 3)
	// assert.Equal(t, add(-1, 2), 3)
}
