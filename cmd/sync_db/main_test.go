package main

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestShouldWriteFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "a.md")
	if err := os.WriteFile(path, []byte("x"), 0644); err != nil {
		t.Fatal(err)
	}
	oldTime := time.Now().Add(-2 * time.Hour)
	newTime := time.Now().Add(2 * time.Hour)
	if ok, _, _ := shouldWriteFile(path, oldTime); ok {
		t.Fatalf("should not write when file is newer")
	}
	if ok, _, _ := shouldWriteFile(path, newTime); !ok {
		t.Fatalf("should write when db is newer")
	}
}
