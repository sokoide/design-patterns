package adapter

import "abstract-factory-example/domain"

// Modern Implementation
type ModernChair struct{}

func (c *ModernChair) SitOn() string { return "Sitting on a modern chair." }

type ModernSofa struct{}

func (s *ModernSofa) LieOn() string { return "Lying on a modern sofa." }

type ModernFactory struct{}

func (f *ModernFactory) CreateChair() domain.Chair { return &ModernChair{} }
func (f *ModernFactory) CreateSofa() domain.Sofa   { return &ModernSofa{} }

// Victorian Implementation
type VictorianChair struct{}

func (c *VictorianChair) SitOn() string { return "Sitting on a victorian chair." }

type VictorianSofa struct{}

func (s *VictorianSofa) LieOn() string { return "Lying on a victorian sofa." }

type VictorianFactory struct{}

func (f *VictorianFactory) CreateChair() domain.Chair { return &VictorianChair{} }
func (f *VictorianFactory) CreateSofa() domain.Sofa   { return &VictorianSofa{} }
