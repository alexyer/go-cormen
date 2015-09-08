package hashes

import "testing"

const MAXSIZE = 42

var (
	empty   SCHashTable
	one     SCHashTable
	several SCHashTable
	many    SCHashTable
)

func init() {
	empty = NewSCHashTable()
	one = NewSCHashTable()
	several = NewSCHashTable()
	many = NewSCHashTable()

	one.Set("One", 1)

	several.Set("One", 1)
	several.Set("Two", 2)

	for i := 0; i < MAXSIZE; i++ {
		many.Set(string(i), i)
	}
}

func TestSCHashGet(t *testing.T) {
	if i, _ := one.Get("One"); i != 1 {
		t.Errorf("One hash get failed")
	}

	first, _ := several.Get("One")
	second, _ := several.Get("Two")
	if first != 1 || second != 2 {
		t.Errorf("Several hash get failed")
	}

	for i := 0; i < MAXSIZE; i++ {
		if i, _ := many.Get(string(i)); i != i {
			t.Errorf("Many hash get failed")
		}
	}
}

func TestSCHashGetNotInTable(t *testing.T) {
	if _, err := empty.Get("One"); err == nil {
		t.Errorf("Wrong behavior")
	}

	if _, err := one.Get("Two"); err == nil {
		t.Errorf("Wrong behavior")
	}

	if _, err := several.Get("three"); err == nil {
		t.Errorf("Wrong behavior")
	}
}

func TestSCHashSetChangeValue(t *testing.T) {
	one.Set("One", 10)

	if val, _ := one.Get("One"); val != 10 {
		t.Errorf("Wrong update behavior")
	}

	several.Set("Two", 10)

	if val, _ := several.Get("Two"); val != 10 {
		t.Errorf("Wrong update behavior")
	}
}

func TestSCHashSetValue(t *testing.T) {
	one.Set("Two", 10)

	if val, _ := one.Get("Two"); val != 10 {
		t.Errorf("Wrong update behavior")
	}

	several.Set("Ten", 10)

	if val, _ := several.Get("Ten"); val != 10 {
		t.Errorf("Wrong update behavior")
	}
}

func TestSCHashDelete(t *testing.T) {
	one.Del("One")

	if _, err := one.Get("One"); err == nil {
		t.Errorf("Wrong delete behavior")
	}

	several.Del("Two")

	if _, err := several.Get("Two"); err == nil {
		t.Errorf("Wrong delete behavior")
	}
}
