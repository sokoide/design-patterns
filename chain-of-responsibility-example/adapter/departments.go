package adapter

import (
	"chain-of-responsibility-example/domain"
)

// --- Reception ---

type Reception struct {
	next   domain.Department
	logger domain.Logger
}

func NewReception(logger domain.Logger) *Reception {
	return &Reception{logger: logger}
}

func (r *Reception) Execute(p *domain.Patient) {
	if p.RegistrationDone {
		r.logger.Log("Patient registration already done")
		if r.next != nil {
			r.next.Execute(p)
		}
		return
	}
	r.logger.Log("Reception registering patient")
	p.RegistrationDone = true
	if r.next != nil {
		r.next.Execute(p)
	}
}

func (r *Reception) SetNext(next domain.Department) {
	r.next = next
}

// --- Doctor ---

type Doctor struct {
	next   domain.Department
	logger domain.Logger
}

func NewDoctor(logger domain.Logger) *Doctor {
	return &Doctor{logger: logger}
}

func (d *Doctor) Execute(p *domain.Patient) {
	if p.DoctorCheckUpDone {
		d.logger.Log("Doctor checkup already done")
		if d.next != nil {
			d.next.Execute(p)
		}
		return
	}
	d.logger.Log("Doctor checking patient")
	p.DoctorCheckUpDone = true
	if d.next != nil {
		d.next.Execute(p)
	}
}

func (d *Doctor) SetNext(next domain.Department) {
	d.next = next
}

// --- Cashier ---

type Cashier struct {
	next   domain.Department
	logger domain.Logger
}

func NewCashier(logger domain.Logger) *Cashier {
	return &Cashier{logger: logger}
}

func (c *Cashier) Execute(p *domain.Patient) {
	if p.PaymentDone {
		c.logger.Log("Payment already done")
		if c.next != nil {
			c.next.Execute(p)
		}
		return
	}
	c.logger.Log("Cashier getting money")
	p.PaymentDone = true
	if c.next != nil {
		c.next.Execute(p)
	}
}

func (c *Cashier) SetNext(next domain.Department) {
	c.next = next
}
