package adapter

import "abstract-factory-example/domain"

var (
	_ domain.Chair            = (*ModernChair)(nil)
	_ domain.Sofa             = (*ModernSofa)(nil)
	_ domain.FurnitureFactory = (*ModernFactory)(nil)

	_ domain.Chair            = (*VictorianChair)(nil)
	_ domain.Sofa             = (*VictorianSofa)(nil)
	_ domain.FurnitureFactory = (*VictorianFactory)(nil)
)

// Modern Implementation
type ModernChair struct{}

func (c *ModernChair) SitOn() string { return "Sitting on a Modern chair." }

type ModernSofa struct{}

func (s *ModernSofa) LieOn() string { return "Lying on a Modern sofa." }

type ModernFactory struct{}

func (f *ModernFactory) CreateChair() domain.Chair { return &ModernChair{} }
func (f *ModernFactory) CreateSofa() domain.Sofa   { return &ModernSofa{} }

// Victorian Implementation
type VictorianChair struct{}

func (c *VictorianChair) SitOn() string { return "Sitting on a Victorian chair." }

type VictorianSofa struct{}

func (s *VictorianSofa) LieOn() string { return "Lying on a Victorian sofa." }

type VictorianFactory struct{}

func (f *VictorianFactory) CreateChair() domain.Chair { return &VictorianChair{} }
func (f *VictorianFactory) CreateSofa() domain.Sofa   { return &VictorianSofa{} }
