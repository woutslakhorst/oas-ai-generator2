package server

import "testing"

func TestNew(t *testing.T) {
	r, db, err := New(":memory:")
	if err != nil {
		t.Fatalf("failed to create server: %v", err)
	}
	if r == nil || db == nil {
		t.Fatalf("expected router and db")
	}
}
