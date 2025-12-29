package domain

import (
	"fmt"
	"strings"
)

type Logger interface {
	Log(message string)
}

type Component interface {
	Search(keyword string)
	Display(indent string)
}

type File struct {
	Name   string
	Logger Logger
}

func (f *File) Search(keyword string) {
	if strings.Contains(f.Name, keyword) {
		f.Logger.Log(fmt.Sprintf("Found '%s' in File: %s", keyword, f.Name))
	}
}

func (f *File) Display(indent string) {
	f.Logger.Log(fmt.Sprintf("%s- %s", indent, f.Name))
}

type Directory struct {
	Name       string
	Components []Component
	Logger     Logger
}

func (d *Directory) Search(keyword string) {
	if strings.Contains(d.Name, keyword) {
		d.Logger.Log(fmt.Sprintf("Found '%s' in Directory: %s", keyword, d.Name))
	}
	for _, c := range d.Components {
		c.Search(keyword)
	}
}

func (d *Directory) Display(indent string) {
	d.Logger.Log(fmt.Sprintf("%s+ %s/", indent, d.Name))
	for _, c := range d.Components {
		c.Display(indent + "  ")
	}
}

func (d *Directory) Add(c Component) {
	d.Components = append(d.Components, c)
}
