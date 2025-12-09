package adapter

import (
	"composite-example/domain"
	"fmt"
)

type Developer struct {
	Name string
}

func (d *Developer) ShowDetails(indent string) {
	fmt.Println(indent + d.Name + " (Developer)")
}

type Designer struct {
	Name string
}

func (d *Designer) ShowDetails(indent string) {
	fmt.Println(indent + d.Name + " (Designer)")
}

type Manager struct {
	Name         string
	Subordinates []domain.Employee
}

func (m *Manager) ShowDetails(indent string) {
	fmt.Println(indent + m.Name + " (Manager)")
	for _, e := range m.Subordinates {
		e.ShowDetails(indent + "  ")
	}
}

func (m *Manager) Add(e domain.Employee) {
	m.Subordinates = append(m.Subordinates, e)
}
