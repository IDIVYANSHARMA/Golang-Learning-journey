package main

import "fmt"

type class struct {
	name string
	section string 
}

type student struct {
	name string
	age int
	marks float64
	class
}

// func newStudent(name string, age int, marks float64) *student{
// 	s1 := student {
// 		name: name,
// 		age: age,
// 		marks: marks,
// 	}
// 	return &s1
// }

// func (s *student) changeStatus(age int) {
// 	s.age = age 
// }

func main() {
    
	newClass := class {
		name: "b.Tech",
		section: "CSE",
	}

	s1 := student {
		name: "jhon",
		age: 22,
		marks: 88.5,
		class: newClass,
	}
    // s1 := newStudent("jhonny", 22, 88.25)
	// fmt.Println(s1)
    s1.class.name = "MBA"
    // language := struct {
	// 	name string
	// 	isGood bool
	// } {
	// 	"go", true,
	// }
    // fmt.Println(language)
	// fmt.Println(s1.name)
	// fmt.Println(s1.age)
	// fmt.Println(s1.marks)
	// s1.changeStatus(24)
	// fmt.Println(s1.age)
	fmt.Println(s1)
}