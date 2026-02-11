package shared

import (
	"errors"
)

// ===========================ERRORS============================

var (
	ErrIncompatibleUnits = errors.New("incompatible units")
	ErrInvalidUnit       = errors.New("invalid unit")
)

// ===========================MODELS============================

type Dimension int

const (
	Mass Dimension = iota
	Volume
	Count
)

type Unit struct {
	Name      string
	Dimension Dimension
	factor    float64
}

var Units = map[string]Unit{
	"mg":  {"mg", Mass, 0.001},
	"g":   {"g", Mass, 1},
	"kg":  {"kg", Mass, 1000},
	"ml":  {"ml", Volume, 0.001},
	"l":   {"l", Volume, 1},
	"pcs": {"pcs", Count, 1},
}

// ===========================METHODS===========================

func NewUnit(name string) (Unit, error) {
	unit, ok := Units[name]
	if !ok {
		return Unit{}, ErrInvalidUnit
	}
	return unit, nil
}

func Convert(value float64, from, to Unit) (float64, error) {
	if from.Dimension != to.Dimension {
		return 0, ErrIncompatibleUnits
	}
	base := value * from.factor
	return base / to.factor, nil
}
