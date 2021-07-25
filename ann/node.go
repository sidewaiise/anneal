package ann

type StandardUnit float64

func (s StandardUnit) toFloat() float64 {
	return float64(s)
}

type Node interface {
	Activate() StandardUnit
	Value() StandardUnit
	Set(StandardUnit)
	SumInputs()

	AddInput(*Edge)
	AddOutput(*Edge)
}
