package adapter

import (
	"fmt"
	"prototype-example/domain"
)

type File struct {
	Name string
}

func (f *File) Print(indent string) {
	fmt.Println(indent + f.Name)
}

func (f *File) Clone() domain.Inode {
	return &File{Name: f.Name + "_clone"}
}

type Folder struct {
	Children []domain.Inode
	Name     string
}

func (f *Folder) Print(indent string) {
	fmt.Println(indent + f.Name)
	for _, i := range f.Children {
		i.Print(indent + indent)
	}
}

func (f *Folder) Clone() domain.Inode {
	cloneFolder := &Folder{Name: f.Name + "_clone"}
	var tempChildren []domain.Inode
	for _, i := range f.Children {
		copy := i.Clone()
		tempChildren = append(tempChildren, copy)
	}
	cloneFolder.Children = tempChildren
	return cloneFolder
}
