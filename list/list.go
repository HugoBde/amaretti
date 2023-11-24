package list

import ()

type List struct {
	Title string
	Desc  string
	Items []Item
}

type Item struct {
	Desc string
}

func NewList() List {
	return List{
		Title: "my title",
		Desc:  "my list desc",
		Items: []Item{
			{Desc: "1"},
			{Desc: "2"},
			{Desc: "3"},
		},
	}
}
