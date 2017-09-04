package main

import (
	"math"
)

const PRIMEONE int = 7559
const PRIMETWO int = 18043

type HashItem struct {
	Key   string
	Value string
}

type Hashtable struct {
	baseSize  int
	itemCount int
	items     []*HashItem
}

func NewHashtable() *Hashtable {
	baseSize := 53

	var items []*HashItem
	items = make([]*HashItem, baseSize, baseSize)

	return &Hashtable{
		items: items,
	}
}

func hashKey(key string, prime int, numBucket int) int {
	str := []rune(key)
	hash := 0.0
	len_s := len(key)
	for i := 0; i < len_s; i++ {
		hash += math.Pow(float64(prime), float64(len_s-(i+1))) * float64(str[i])
		hash = math.Mod(hash, float64(numBucket))
	}

	return int(hash)
}

func doubleHash(key string, numBuckets int, attempt int) int {
	hash_a := hashKey(key, PRIMEONE, numBuckets)
	hash_b := hashKey(key, PRIMETWO, numBuckets)
	hash := math.Mod(float64((hash_a + (attempt * (hash_b + 1)))), float64(numBuckets))
	return int(hash)
}

func (ht *Hashtable) insert(key string, value string) {
	newItem := &HashItem{
		Key:   key,
		Value: value,
	}

	index := doubleHash(newItem.Key, ht.size(), 0)

	curItem := ht.items[index]
	i := 1

	for curItem != nil {
		index = doubleHash(newItem.Key, ht.size(), i)
		curItem = ht.items[index]
		i++
	}

	ht.items[index] = newItem
	ht.itemCount++
}

func (ht *Hashtable) size() int {
	return len(ht.items)
}

func (ht *Hashtable) count() int {
	return ht.itemCount
}
