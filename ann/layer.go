package ann

import "errors"

var (
	ErrMismatchingNodesToEdgeWeights = errors.New("The number of nodes and numbers of edge weights do not match.")
)

func NewLayer(inputWeights [][]StandardUnit, fromLayer *Layer) (*Layer, error) {
	l := Layer{
		nodes: make([]Node, len(inputWeights)),
		edges: make([]*Edge, len(inputWeights)*len(inputWeights)),
	}

	if fromLayer != nil {

	}
	for _, weights := range inputWeights {
		if fromLayer != nil && len(weights) != len(fromLayer.Nodes()) {
			return nil, ErrMismatchingNodesToEdgeWeights
		}

		toNode := NewNeuron(TF_SIGMOID)
		l.AddNode(toNode)

		if fromLayer != nil {
			for i, fromNode := range fromLayer.Nodes() {
				e := NewEdge(fromNode, toNode, weights[i])
				l.AddEdge(e)
			}
		}
	}

	return &l, nil
}

// Layer is a collection of nodes (neurons) with edges TO nodes in another layer
type Layer struct {
	nodes []Node
	edges []*Edge
}

// AddNode adds a node to a layer
func (l *Layer) AddNode(node Node) {
	l.nodes = append(l.nodes, node)
}

// AddEdge adds an edge to a layer
func (l *Layer) AddEdge(edge *Edge) {
	l.edges = append(l.edges, edge)
}

// Nodes gets all the nodes for a layer
func (l *Layer) Nodes() []Node {
	return l.nodes
}

// Edges gets all the edges for a layer
func (l *Layer) Edges() []*Edge {
	return l.edges
}
