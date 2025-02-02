package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func Test_Foo(t *testing.T) {
	var items = []*gildedrose.Item{
		{"foo", 0, 0, true},
	}

  for _, item := range items {
    item.UpdateQuality()
  }

	if items[0].Name != "fixme" {
		t.Errorf("Name: Expected %s but got %s ", "fixme", items[0].Name)
	}
}
