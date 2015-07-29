package bricker

import (
	"fmt"
	"sort"
)

type Pair struct {
	Key   string
	Value int64
}
type PairList []Pair

func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }

// A function to turn a map into a PairList, then sort and return it.
func SortMapByValue(m map[string]int64) PairList {
	p := make(PairList, len(m))
	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i += 1
	}
	sort.Sort(p)
	return p
}

type Part struct {
	Color    string
	Item     string
	Name     string
	Quantity int
}

type PartSet map[Part]bool

func (self *Part) SearchUrl() string {
	return fmt.Sprintf("search.asp?colorID=%s&itemID=%s&sz=500", self.Color, self.Item)
}

func (self Part) String() string {
	return fmt.Sprintf("%s:%s", self.Color, self.Item)
}

type Supply struct {
	Name     string
	Quantity int
	BuyUrl   string
	Price    float64
}

type CartItem struct {
	Part     Part
	Price    float64
	Quantity int
	BuyUrl   string
	Name     string
}

type SupplierScores (map[string]int64)
type PartSupply map[string][]Supply
type Carts map[string][]CartItem
