package sort

import (
	"fmt"
	"sort"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int16
}

func (p Person) String() string {
	return p.FirstName + " " + p.LastName + " is " + fmt.Sprint(p.Age)
}

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func PrintSortPeoples() {
	people := []Person{
		{"Akash", "sen", 52},
		{FirstName: "Raj", LastName: "verma", Age: 30},
		{FirstName: "Disha", LastName: "verma", Age: 39},
		{FirstName: "Aman", LastName: "Lal", Age: 51},
	}

	fmt.Println("before sort:", people)
	sort.Sort(ByAge(people))
	fmt.Println("after sort:", people)
}
