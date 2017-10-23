package fruit

import (
	"log"
)

type Product struct {
	Name	string
}

type Data struct {
	Products	[]*Product
}

func init() {
	log.Println("include fruit.")
}

func GetList() Data {
	p1 := Product{Name:"りんご"}
	p2 := Product{Name:"なし"}
	p3 := Product{Name:"ばなな"}
	return Data{Products: []*Product{&p1, &p2, &p3}}
}
