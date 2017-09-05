package main

import (
	"fmt"
	"testing"

	"github.com/satori/go.uuid"
)

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

func TestRetrieve(t *testing.T) {
	expected := "value"
	ht := NewHashtable()
	ht.insert("key", expected)

	value, err := ht.search("key")
	if err != nil {
		t.Error(err)
	}

	if value != expected {
		t.Error(fmt.Sprintf("Expected %v, got", expected), ht.count())
	}

	noKey := "none"
	value, err = ht.search(noKey)

	if value != "" {
		t.Error(fmt.Sprintf("Expected no value for %v, got", noKey), value)
	}
}

func testRemove(t *testing.T) {
	aKey := "key"
	ht := NewHashtable()
	ht.insert(aKey, "value")

	ht.remove(aKey)

	value, _ := ht.search(aKey)

	if value != "" {
		t.Error(fmt.Sprintf("Expected no value for %v, got", aKey), value)
	}

	if ht.count() > 0 {
		t.Error("Expected count of 0 for %v, got", ht.count())
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

func TestLarge(t *testing.T) {
	ht := NewHashtable()

	for i := 0; i < 20; i++ {
		ht.insert(fmt.Sprintf("%v", uuid.NewV4()), "val")
	}

	if ht.count() != 20 {
		t.Error("Expected 19 size, got ", ht.count())
	}
}
