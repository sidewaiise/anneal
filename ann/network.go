package ann

import "errors"

var (
	ErrNoInputLayer   = errors.New("No Input Layer, please create one")
	ErrNoOutputLayer  = errors.New("No Input Layer, please create one")
	ErrNoHiddenLayers = errors.New("No hidden layers, please create some")
	ErrNoNodes        = errors.New("No Nodes in layer, please create some")
)

type Layers struct {
	Input  *Layer
	Hidden []*Layer
	Output *Layer
}

type Network struct {
	Output chan []StandardUnit
	Layers Layers
}

func (n *Network) GetInput() *Layer {
	return n.Layers.Input
}

func (n *Network) SetInput(layer *Layer) {
	n.Layers.Input = layer
}

func (n *Network) GetOutput() *Layer {
	return n.Layers.Output
}

func (n *Network) SetOutput(layer *Layer) {
	n.Layers.Output = layer
}

func (n *Network) AddHidden(layer *Layer) {
	n.Layers.Hidden = append(n.Layers.Hidden, layer)
}

func (n *Network) GetHidden() []*Layer {
	return n.Layers.Hidden
}

func (n *Network) GetLastHidden() *Layer {
	hidden := len(n.Layers.Hidden)
	if hidden == 0 {
		return nil
	}
	return n.Layers.Hidden[hidden-1]
}

func (n *Network) CountTotalLayers() int {
	var count = 0

	if n.Layers.Input != nil {
		count++
	}

	if n.Layers.Output != nil {
		count++
	}

	count += len(n.GetHidden())

	return count
}

// Feed is for feeding data into the network via input nodes
func (n *Network) Feed(values []StandardUnit) error {
	if n.Layers.Input == nil {
		return ErrNoInputLayer
	}

	if len(n.Layers.Hidden) == 0 {
		return ErrNoHiddenLayers
	}

	for i, v := range values {
		n.Layers.Input.nodes[i].Set(v)
	}

	return nil
}
