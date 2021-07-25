package anneal

import "github.com/sidewaiise/anneal/ann"

// NewNetwork creates a new network
// @weights - layer / node / input edge weights
func NewNetwork(weights [][][]ann.StandardUnit) (*ann.Network, error) {
	network := ann.Network{
		Output: make(chan []ann.StandardUnit),
		Layers: ann.Layers{
			Input:  nil,
			Hidden: make([]*ann.Layer, 0),
			Output: nil,
		},
	}

	for i, layer := range weights {
		var fromLayer *ann.Layer
		if i == 0 {
			l, err := ann.NewLayer(layer, fromLayer)
			if err != nil {
				return nil, err
			}

			network.SetInput(l)
			continue
		}

		// If there's no hidden layers, we use input layer as from layer
		if lastHidden := network.GetLastHidden(); lastHidden == nil {
			fromLayer = network.GetInput()
		} else {
			fromLayer = lastHidden
		}

		l, err := ann.NewLayer(layer, fromLayer)
		if err != nil {
			return nil, err
		}

		if i > 0 && i < len(weights) {
			network.AddHidden(l)
		} else {
			network.SetOutput(l)
		}
	}

	return &network, nil
}
