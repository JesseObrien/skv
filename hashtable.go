package main

import (
	"errors"
	"fmt"
	"math"
)

const PRIMEONE int = 7559
const PRIMETWO int = 18043
const INITIAL_BASE_SIZE int = 10

type HashItem struct {
	Key   string
	Value string
}

type Hashtable struct {
	itemCount int
	items     []*HashItem
}

func NewHashtable() *Hashtable {
	size := nextPrime(INITIAL_BASE_SIZE)
	items := make([]*HashItem, size, size)

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

func getKeyHash(key string, numBuckets int, attempt int) int {
	hash_a := hashKey(key, PRIMEONE, numBuckets)
	hash_b := hashKey(key, PRIMETWO, numBuckets)
	hash := math.Mod(float64((hash_a + (attempt * (hash_b + 1)))), float64(numBuckets))
	return int(hash)
}

func (ht *Hashtable) up() {
	newCap := nextPrime(cap(ht.items) * 2)
	tempItems := make([]*HashItem, newCap, newCap)
	copy(tempItems, ht.items)
	ht.items = tempItems
}

func (ht *Hashtable) insert(key string, value string) {
	if (ht.count() + 2) >= cap(ht.items) {
		ht.up()
	}

	newItem := &HashItem{
		Key:   key,
		Value: value,
	}

	index := getKeyHash(newItem.Key, cap(ht.items), 0)

	curItem := ht.items[index]
	i := 1

	for curItem != nil {
		if curItem != (&HashItem{}) {
			if key == curItem.Key {
				ht.items[index] = newItem
				return
			}
		}

		index = getKeyHash(newItem.Key, cap(ht.items), i)
		curItem = ht.items[index]
		i++
	}

	ht.items[index] = newItem
	ht.itemCount++
}

func (ht *Hashtable) search(key string) (string, error) {
	index := getKeyHash(key, cap(ht.items), 0)
	item := ht.items[index]

	i := 1

	for item != nil {
		if item != (&HashItem{}) {
			if key == item.Key {
				return item.Value, nil
			}
		}

		if key == item.Key {
			return item.Value, nil
		}

		index := getKeyHash(key, cap(ht.items), i)
		item = ht.items[index]
		i++
	}

	return "", errors.New(fmt.Sprintf("Could not find value for key `%v`", key))
}

func (ht *Hashtable) remove(key string) {
	index := getKeyHash(key, cap(ht.items), 0)
	item := ht.items[index]

	i := 1

	for item != nil {
		if item != (&HashItem{}) {
			if key == item.Key {
				ht.items[index] = &HashItem{}
			}
		}
		index = getKeyHash(key, cap(ht.items), i)
		item = ht.items[index]
		i++
	}
	ht.itemCount--
}

func (ht *Hashtable) size() int {
	return len(ht.items)
}

func (ht *Hashtable) count() int {
	return ht.itemCount
}
