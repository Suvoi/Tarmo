package shared

import "errors"

// ===========================ERRORS============================

var (
	ErrQuantityInvalid = errors.New("quantity must be greater than 0")
)

// ===========================MODELS============================

type Quantity struct {
	value float64
	unit  Unit
}

// ===========================GETTERS===========================

func (q *Quantity) Value() float64 { return q.value }
func (q *Quantity) Unit() Unit     { return q.unit }

// ===========================CONSTRUCTORS======================

func NewQuantity(value float64, unit Unit) (Quantity, error) {
	if value <= 0 {
		return Quantity{}, ErrQuantityInvalid
	}
	return Quantity{value: value, unit: unit}, nil
}

func ReconstructQuantity(db_value float64, db_unit string) (Quantity, error) {
	unit, err := NewUnit(db_unit)
	if err != nil {
		return Quantity{}, err
	}
	return Quantity{value: db_value, unit: unit}, nil
}

// ===========================METHODS===========================

func (q Quantity) ToBase() Quantity {
	return Quantity{
		value: q.value * q.unit.factor,
		unit: Unit{
			Name:      DimensionToString(q.unit.Dimension),
			Dimension: q.unit.Dimension,
			factor:    1,
		},
	}
}

// ===========================HELPER METHODS====================

func DimensionToString(dim Dimension) string {
	switch dim {
	case Mass:
		return "g"
	case Volume:
		return "l"
	case Count:
		return "pcs"
	default:
		return ""
	}
}
