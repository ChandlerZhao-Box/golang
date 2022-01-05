package main

import "fmt"

//别名
func main() {
	n := fakeNews{
		Name: "aaa",
	}
	n.Report()
}

type News struct {
	Name string
}

func (d News) Report()  {
	fmt.Println("I am news: " + d.Name)
}

type fakeNews = News
