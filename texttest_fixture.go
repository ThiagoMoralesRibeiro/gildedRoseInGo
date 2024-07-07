package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func main() {
	fmt.Println("OMGHAI!")

	var items = []gildedrose.Updatable{
		gildedrose.NewItem("+5 Dexterity Vest", 10, 20, false),
		gildedrose.NewItem("Aged Brie", 2, 0, true),
		gildedrose.NewItem("Elixir of the Mongoose", 5, 7, false),
		gildedrose.NewItem("Sulfuras, Hand of Ragnaros", 0, 80, false),
		gildedrose.NewItem("Sulfuras, Hand of Ragnaros", -1, 80, false),
		gildedrose.NewBackstage("Backstage passes to a TAFKAL80ETC concert", 15, 20, true),
		gildedrose.NewBackstage("Backstage passes to a TAFKAL80ETC concert", 10, 49, true),
		gildedrose.NewBackstage("Backstage passes to a TAFKAL80ETC concert", 5, 49, true),
		gildedrose.NewConjured("Conjured Mana Cake", 3, 6, false), // <-- :O
	}

	days := 2
	var err error
	if len(os.Args) > 1 {
		days, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		days++
	}

	for day := 0; day < days; day++ {
		fmt.Printf("-------- day %d --------\n", day)
		fmt.Println("Name, SellIn, Quality")
		for i := 0; i < len(items); i++ {
			fmt.Println(items[i])
		}
		fmt.Println("")
    
    for _, item := range items {
      item.UpdateQuality()
    }

	}
}
