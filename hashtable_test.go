package main

import "testing"

func TestInsert(t *testing.T) {
	ht := NewHashtable()
	ht.insert("key", "value")

	if ht.count() != 1 {
		t.Error("Expected 1 insertion, got ", ht.count())
	}

	ht.insert("newkey", "value")

	if ht.count() != 2 {
		t.Error("Expected 1 insertion, got ", ht.count())
	}
}

func TestCollision(t *testing.T) {
	ht := NewHashtable()
	ht.insert("key", "value")
	ht.insert("key", "value")

	if ht.count() != 1 {
		t.Error("Expected 1 overwrite, got ", ht.count())
	}
}
