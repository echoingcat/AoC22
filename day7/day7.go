// INCOMPLETE
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	i, _ := os.ReadFile("input.txt")
	input := strings.Split(string(i), "\n")[1:]

	root := &dir{name: "/"}
	root.parent = root
	c := root

	for _, s := range input {
		a := strings.Fields(s)

		if a[0] == "$" {
			if a[1] == "cd" {
				if a[2] == ".." {
					c = c.parent
				} else {
					c = c.subs[get(c.subs, a[2])]
				}
			}
		} else if a[0] == "dir" {
			temp := &dir{name: a[1], parent: c}
			c.subs = append(c.subs, temp)
		} else {
			c.files = append(c.files, a[0])
		}
	}
}

func dirSize(d *dir) int {
	res := 0
	for i := range d.files {
		temp, _ := strconv.Atoi((string(d.files[i])))
		res += temp
	}
	return res
}

func print(d *dir) {
	fmt.Println("Name:", d.name)
	fmt.Println("Parent:", d.parent.name)
	fmt.Print("Subs:")
	for i := range d.subs {
		fmt.Print(" ", d.subs[i].name)
	}
	fmt.Print("\nFiles:")
	for i := range d.files {
		fmt.Print(" ", d.files[i])
	}
	fmt.Println("")
}

func get(d []*dir, s string) int {
	for i, x := range d {
		if x.name == s {
			return i
		}
	}
	return -1
}

type dir struct {
	parent *dir
	name   string
	subs   []*dir
	files  []string
}