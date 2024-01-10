package structural

import "fmt"

type FileSystemComponent interface {
	Display(indentation string)
}

type File struct {
	name string
}

func (f *File) Display(indentation string) {
	fmt.Println(indentation + f.name)
}

type Directory struct {
	name       string
	components []FileSystemComponent
}

func (d *Directory) Display(indentation string) {
	fmt.Println(indentation + "Directory: " + d.name + "/")
	for _, component := range d.components {
		component.Display(indentation + "    ")
	}
}

func (d *Directory) Add(component FileSystemComponent) {
	d.components = append(d.components, component)
}

func (d *Directory) Remove(component FileSystemComponent) {
	for i, c := range d.components {
		if c == component {
			d.components = append(d.components[:i], d.components[i+1:]...)
			break
		}
	}
}

func Composite() {
	file1 := &File{name: "file1.txt"}
	file2 := &File{name: "file2.txt"}
	file3 := &File{name: "file3.txt"}

	dir1 := &Directory{name: "Dir1"}
	dir2 := &Directory{name: "Dir2"}

	dir1.Add(file1)
	dir1.Add(file2)

	dir2.Add(file3)

	rootDir := &Directory{name: "Root"}
	rootDir.Add(dir1)
	rootDir.Add(dir2)

	rootDir.Display("")
}
