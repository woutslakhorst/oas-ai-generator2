package main

import (
	"math/rand"
	"sort"
	"testing"

	"example.com/blobapi/internal/models"
)

func sortedRunes(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
	return string(r)
}

func TestScrambleNames(t *testing.T) {
	orig := []models.Blob{{Name: "abc"}, {Name: "xyz"}}
	blobs := []models.Blob{{Name: "abc"}, {Name: "xyz"}}
	rand.Seed(1)
	scrambleNames(blobs)
	for i := range blobs {
		if sortedRunes(blobs[i].Name) != sortedRunes(orig[i].Name) {
			t.Fatalf("letters changed for blob %d", i)
		}
		if blobs[i].Name == orig[i].Name {
			t.Fatalf("blob %d was not scrambled", i)
		}
	}
}
