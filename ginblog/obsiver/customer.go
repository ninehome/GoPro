package obsiver

import "fmt"

// 实现了 observer 接口
type Customer struct {
	Id   string
	Name string
}

func (c *Customer) Update(itemName string) {
	fmt.Printf("Sending email to %s customer %s for item %s\n", c.Name, c.Id, itemName)
}

func (c *Customer) GetID() string {
	return c.Id
}
