package main

import (
	"testing"
	"testing/fstest"

	"github.com/stretchr/testify/assert"
)

// https://apitest.dev/

func TestReadFileFromVolume(t *testing.T) {
	// Mock a filesystem file (in this case, messages.txt) https://stackoverflow.com/questions/16742331/how-to-mock-abstract-filesystem-in-go
	fs := fstest.MapFS{
		"messages.txt": {
			Data: []byte("2022-11-22T11:01:32.149Z 1 MSG_{1} to compse140.o\n2022-11-22T11:01:35.155Z 2 Got MSG_{2} to compse140.i\n2022-11-22T11:01:38.156Z 3 Got MSG_{3} to compse140.i"),
		},
	}
	expectedData, err := fs.ReadFile("messages.txt")
	if err != nil {
		panic(err)
	}

	assert.Equal(t, expectedData, ReadFileFromVolume("messages.txt"), "should be equal")
}
