package ann

func NewEdge(from, to Node, weight StandardUnit) *Edge {
	e := &Edge{
		from:   from,
		to:     to,
		weight: weight,
	}

	from.AddOutput(e)
	to.AddInput(e)

	return e
}

type Edge struct {
	from   Node
	to     Node
	weight StandardUnit
}

func (e *Edge) WeightedValueFrom() StandardUnit {
	return e.from.Value() * e.weight
}
