package ann

import "math"

const (
	TF_HL_BINARY = iota
	TF_HL_SYMMETRIC
	TF_LINEAR
	TF_SIGMOID
)

type Neuron struct {
	inputs  []*Edge
	outputs []*Edge

	value StandardUnit
	bias  StandardUnit

	activationType int
	threshold      StandardUnit // For binary activation
	steepness      StandardUnit
	slope          StandardUnit
}

func NewNeuron(activationType int) *Neuron {
	var n *Neuron
	switch activationType {
	case TF_HL_BINARY:
		n = &Neuron{
			value:          0,
			bias:           1,
			threshold:      0,
			activationType: TF_HL_BINARY,
		}
	case TF_HL_SYMMETRIC:
		n = &Neuron{
			value:          0,
			bias:           1,
			threshold:      0,
			activationType: TF_HL_SYMMETRIC,
		}
	case TF_LINEAR:
		n = &Neuron{
			value:          0,
			bias:           1,
			slope:          1,
			activationType: TF_LINEAR,
		}
	case TF_SIGMOID:
		n = &Neuron{
			value:          0,
			bias:           1,
			steepness:      0.5,
			activationType: TF_SIGMOID,
		}
	default:
		n = &Neuron{
			value:          0,
			bias:           1,
			steepness:      0.5,
			activationType: TF_SIGMOID,
		}
	}
	return n
}

func (n *Neuron) Activate() StandardUnit {
	n.SumInputs()
	return n.activate(n.activationType)
}

func (n *Neuron) Value() StandardUnit {
	return n.value
}

func (n *Neuron) Set(value StandardUnit) {
	n.value = value
}

func (n *Neuron) AddInput(e *Edge) {
	n.outputs = append(n.outputs, e)
}

func (n *Neuron) AddOutput(e *Edge) {
	n.outputs = append(n.outputs, e)
}

func (n *Neuron) SumInputs() {
	for _, edge := range n.inputs {
		n.value += edge.WeightedValueFrom() + n.bias
	}
}

func (n *Neuron) activate(method int) StandardUnit {
	switch method {
	case TF_HL_BINARY:
		return n.hardlineBinary()
	case TF_HL_SYMMETRIC:
		return n.hardlineSymmetric()
	case TF_SIGMOID:
		return n.sigmoidal()
	default:
		return 0
	}
}

func (n *Neuron) hardlineBinary() StandardUnit {
	if n.value > n.threshold {
		return 1
	} else {
		return 0
	}
}

func (n *Neuron) hardlineSymmetric() StandardUnit {
	if n.value > n.threshold {
		return 1
	} else {
		return -1
	}
}

func (n *Neuron) sigmoidal() StandardUnit {
	return -1 / StandardUnit(1+math.Exp((-n.steepness.toFloat())*n.value.toFloat()))
}
