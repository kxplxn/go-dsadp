package main

import "fmt"

// Component serves as the Component Interface, describing operations that are
// common to both simple and complex elements of the tree.
type Component interface {
	Search(string)
}

// File is the Leaf — a basic element of the tree that doesn’t have
// sub-elements. Usually, leaf components end up doing most of the real work,
// since they don’t have anyone to delegate the work to.
type File struct {
	name string
}

func (f *File) Search(keyword string) {
	fmt.Printf("searching for keyword %s in file %s\n", keyword, f.name)
}

// Folder is the Container (AKA Composite). which is an element that has
// sub-elements: leaves or other containers. A container doesn’t know the
// concrete classes of its children. It works with all sub-elements only via the
// component interface.
//
// Upon receiving a request, a container delegates the work to its sub-elements,
// processes intermediate results and then returns the final result to the
// client.
type Folder struct {
	name       string
	components []Component
}

func (f *Folder) Search(keyword string) {
	fmt.Printf("searching recursively for keyword %s in folder %s\n", keyword, f.name)
	for _, v := range f.components {
		v.Search(keyword)
	}
}

func (f *Folder) Add(component Component) {
	f.components = append(f.components, component)
}

func main() {
	file1 := &File{name: "File 1"}
	file2 := &File{name: "File 2"}
	file3 := &File{name: "File 3"}

	folder1 := &Folder{name: "Folder 1"}
	folder1.Add(file1)

	folder2 := &Folder{name: "Folder 2"}
	folder2.Add(file2)
	folder2.Add(file3)
	folder2.Add(folder1)

	folder2.Search("baban")
}
