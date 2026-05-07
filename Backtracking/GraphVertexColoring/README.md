## Graph Vertex Coloring

Given an undirected graph `G=(V,E)` and a positive integer `k` (number of colors), assign a color (from `1..k`) to each vertex such that no two adjacent vertices share the same color.

### State space representation

- Can be represented using an array, `color[v]` ($v \in V(G)$). (`0` for uncolored or `1..k`)
- Total state space: $k^{|V|}$
- We can sort the vertices by their degrees to reduce the state space and more efficient pruning.

