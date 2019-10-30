package collection

import "fmt"

type Item struct {
	field    int
	NextItem *Item
	PrevItem *Item
}

type Collection struct {
	FirstItem *Item
	LastItem  *Item
	Length    int
}

func (c *Collection) Add(Value int) {
	newItem := &Item{}
	newItem.field = Value
	if c.Length == 0 {
		c.FirstItem = newItem
		c.LastItem = newItem
	} else {

		newItem.PrevItem = c.LastItem
		c.LastItem.NextItem = newItem
		c.LastItem = newItem
	}
	c.Length++
}

func (c *Collection) Print() {
	x := c.FirstItem
	for i := 0; i < c.Length; i++ {
		fmt.Println(x.field)
		x = x.NextItem
	}
}

func (c *Collection) Get(index int) int {
	result := c.FirstItem
	for i := 0; i < index; i++ {
		result = result.NextItem
	}
	return result.field
}

func (c *Collection) Remove(index int) {
	result := c.FirstItem
	for i := 0; i < index; i++ {
		result = result.NextItem
	}
	result.PrevItem.NextItem = result.NextItem
	result.NextItem.PrevItem = result.PrevItem
	c.Length--

}

func (c *Collection) First() int {
	return c.FirstItem.field
}
func (c *Collection) Last() int {
	return c.LastItem.field
}
func (c *Collection) LengthCheck() int {
	return c.Length
}
