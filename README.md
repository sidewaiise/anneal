# Anneal

A simple Artificial neurla network library written in go

## Setting up a neural net

```
weights := [][][]ann.StandardUnit{
    // Layers
    {
        // Nodes
        {
            // Input Weights
            1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        },
    },
}

network, err := anneal.NewNetwork(weights)
```

TODO:
- package to prepare default weights a lot easier
- training package
