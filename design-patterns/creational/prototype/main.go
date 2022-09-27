package main

import "fmt"

// Node is the Prototype Interface. A Prototype Interface defines the cloning
// methods, which, in most cases, is a single clone method.
type Node interface {
	print(string)
	clone() Node
}

// File is a Contrete Prototype. The Concrete Prototypes implement the cloning
// method. In addition to copying the original object’s data to the clone, this
// method may also handle some edge cases of the cloning process related to
// cloning linked objects, untangling recursive dependencies, etc.
type File struct {
	name string
}

func (f *File) print(indentation string) {
	fmt.Println(indentation + f.name)
}

func (f *File) clone() Node {
	return &File{name: f.name + "_clone"}
}

// Folder is a Contrete Prototype. The Concrete Prototypes implement the cloning
// method. In addition to copying the original object’s data to the clone, this
// method may also handle some edge cases of the cloning process related to
// cloning linked objects, untangling recursive dependencies, etc.
type Folder struct {
	name     string
	children []Node
}

func (f *Folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, v := range f.children {
		v.print(indentation + indentation)
	}
}

func (f *Folder) clone() Node {
	cloneFolder := &Folder{
		name:     f.name + "_clone",
		children: make([]Node, len(f.children)),
	}
	for i, v := range f.children {
		cloneFolder.children[i] = v.clone()
	}
	return cloneFolder
}

func main() {
	folder := &Folder{
		name: "Folder #1",
		children: []Node{
			&File{name: "File #1"},
			&File{name: "File #2"},
			&File{name: "File #3"},
		},
	}

	fmt.Println("Printing hierarchy for the original folder.")
	folder.print("  ")

	cloneFolder := folder.clone()

	fmt.Println("Printing hierarchy for the clone folder.")
	cloneFolder.print("  ")
}
