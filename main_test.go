package anneal_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/sidewaiise/anneal"
	"github.com/sidewaiise/anneal/ann"
)

func TestNewNetwork(t *testing.T) {
	type networktest struct {
		weights        [][][]ann.StandardUnit
		expectedLayers int
		expectedError  error
	}

	tcs := map[string]networktest{
		"Network creates 1 layer": {
			weights: [][][]ann.StandardUnit{
				// Layers
				{
					// Nodes
					{
						// Input Weights
						1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
					},
				},
			},
			expectedLayers: 1,
		},
	}

	for name, test := range tcs {
		t.Run(name, func(t *testing.T) {
			n, err := anneal.NewNetwork(test.weights)
			if err != nil {
				if diff := cmp.Diff(err, test.expectedError, cmpopts.EquateErrors()); diff != "" {
					t.Fatalf("Not the error we were expecting")
				}
			}

			got := n.CountTotalLayers()
			if got != test.expectedLayers {
				t.Fatal("Didn't get expected number of layers")
			}
		})
	}
}
