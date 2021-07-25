package anneal_test

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"

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
		"Network fails to create with incorrect node/weight numbers": {
			weights: [][][]ann.StandardUnit{
				{
					{
						1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
					},
				},
				{
					{
						1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
					},
				},
				{
					{
						1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
					},
				},
				{
					{
						1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
					},
				},
				{
					{
						1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
					},
				},
			},
			expectedError: errors.New("The number of nodes and numbers of edge weights do not match."),
		},
		"Network Creates with 5 layers": {
			weights: [][][]ann.StandardUnit{
				{
					{
						1, 1, 1,
					},
				},
				{
					{
						1, 1, 1,
					},
					{
						1, 1, 1,
					},
					{
						1, 1, 1,
					},
				},
				{
					{
						1, 1, 1,
					},
					{
						1, 1, 1,
					},
					{
						1, 1, 1,
					},
				},
				{
					{
						1, 1, 1,
					},
					{
						1, 1, 1,
					},
					{
						1, 1, 1,
					},
				},
				{
					{
						1, 1, 1,
					},
					{
						1, 1, 1,
					},
					{
						1, 1, 1,
					},
				},
			},
			expectedError: errors.New("The number of nodes and numbers of edge weights do not match."),
		},
	}

	for name, test := range tcs {
		t.Run(name, func(t *testing.T) {
			n, err := anneal.NewNetwork(test.weights)

			if err != nil {
				if diff := cmp.Diff(err.Error(), test.expectedError.Error()); diff != "" {
					t.Log(err.Error())
					t.Fatalf("Not the error we were expecting: %s", test.expectedError.Error())
				}
			} else {
				got := n.CountTotalLayers()
				if got != test.expectedLayers {
					t.Fatal("Didn't get expected number of layers")
				}
			}
		})
	}
}
