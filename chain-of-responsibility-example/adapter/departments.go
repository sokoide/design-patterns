package adapter

import (
	"chain-of-responsibility-example/domain"
	"fmt"
)

type Reception struct {
	next domain.Department
}

func (r *Reception) Execute(p *domain.Patient) {
	if p.RegistrationDone {
		fmt.Println("Patient registration already done")
		r.next.Execute(p)
		return
	}
	fmt.Println("Reception registering patient")
	p.RegistrationDone = true
	if r.next != nil {
		r.next.Execute(p)
	}
}

func (r *Reception) SetNext(next domain.Department) { r.next = next }

type Doctor struct {
	next domain.Department
}

func (d *Doctor) Execute(p *domain.Patient) {
	if p.DoctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		d.next.Execute(p)
		return
	}
	fmt.Println("Doctor checking patient")
	p.DoctorCheckUpDone = true
	if d.next != nil {
		d.next.Execute(p)
	}
}

func (d *Doctor) SetNext(next domain.Department) { d.next = next }

type Cashier struct {
	next domain.Department
}

func (c *Cashier) Execute(p *domain.Patient) {
	if p.PaymentDone {
		fmt.Println("Payment already done")
		return
	}
	fmt.Println("Cashier getting money")
	p.PaymentDone = true
}

func (c *Cashier) SetNext(next domain.Department) { c.next = next }
