package gildedrose

import "errors"

type Updatable interface{
  UpdateQuality()
}

type Item struct {
	Name            string
	SellIn, Quality int
  BetterAged bool
}

type Backstage struct {
  Item
}

type Conjured struct {
  Item
} 


func isQualityNegative(items *Item) (bool, error){
    if items.Quality <= 0 {
        return false, errors.New("quality cannot be negative")
    }
    return true, nil;
}

func isSellInNegative(items *Item) (bool){
    if items.SellIn < 0 {
        return true
    }
    return false;
}

func verifyQualityLimit(items *Item) bool{
  if items.Quality > 50{
    return false
  }
  return true

}


func isBetterAged(items *Item) bool{
  if items.BetterAged == true {
    return true
  }
  return false

}

func reduceSellIn(items *Item){
  if items.Name != "Sulfuras, Hand of Ragnaros"{
   items.SellIn--
  }
}

func reduceQuality(items *Item, times int) error{
  result, err:= isQualityNegative(items)
  if result && items.Name != "Sulfuras, Hand of Ragnaros"{
    items.Quality = items.Quality - (1 * times)
  }
  return err
}

func increaseQuality(items *Item){
  result := verifyQualityLimit(items)

  if (result) {
    items.Quality++
  }
}

func limitBackstageQuality(backstage *Backstage){
  if backstage.Quality > 50{
    backstage.Quality = 50
  }
}

func NewItem (name string, sellIn, quality int, betterAged bool) *Item{
  return &Item{name, sellIn, quality, betterAged}
}

func NewBackstage(name string, sellIn, quality int, betterAged bool) *Backstage {
  return &Backstage{Item{name, sellIn, quality, betterAged}}
}

func NewConjured(name string, sellIn, quality int, betterAged bool) *Conjured  {
  return &Conjured{Item{name, sellIn, quality, betterAged}}
}

func (items *Item) UpdateQuality() {
	if !isBetterAged(items){
    reduceQuality(items, 1)
	} else {
    increaseQuality(items) 
	}

  reduceSellIn(items)
  
  if isSellInNegative(items) {
    if isBetterAged(items) {
      increaseQuality(items)
    }else{
     reduceQuality(items, 1) 
    }
  }
}

func (conjuredItem *Conjured) UpdateQuality(){
  if !isBetterAged(&conjuredItem.Item){
    reduceQuality(&conjuredItem.Item, 2)
	} else {
    increaseQuality(&conjuredItem.Item) 
	}

  reduceSellIn(&conjuredItem.Item)
  
  if isSellInNegative(&conjuredItem.Item) {
    if isBetterAged(&conjuredItem.Item) {
      increaseQuality(&conjuredItem.Item)
    }else{
     reduceQuality(&conjuredItem.Item, 2)
    }
  }
}

func (backstageItem *Backstage) UpdateQuality(){
  switch  {

  case backstageItem.SellIn > 10:
    backstageItem.Quality++

  case backstageItem.SellIn > 5:
    backstageItem.Quality += 2

  case backstageItem.SellIn > 0:
    backstageItem.Quality += 3

  default:
    backstageItem.Quality = 0

  }
  
  limitBackstageQuality(backstageItem)

  reduceSellIn(&backstageItem.Item)

}
