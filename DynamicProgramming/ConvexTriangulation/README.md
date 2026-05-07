## Convex Triangulation

The general formula for `dp[i][j]` (where `i < j`) in the convex polygon triangulation problem is:

$$
dp[i][j] = 
\begin{cases} 
0 & \text{if } j = i+1 \\
\displaystyle \min_{i < k < j} \big( dp[i][k] + dp[k][j] + w[i][k] + w[k][j] + w[i][j] \big) & \text{if } j \ge i+2
\end{cases}
$$

- `dp[i][j]` = minimum cost to triangulate the subpolygon formed by vertices `v_i, v_{i+1}, ..., v_j` in cyclic order.
- `w[a][b]` = given cost of edge between `v_a` and `v_b` (boundary or diagonal).
- The base case `j = i+1` is a single edge with no area, so cost = 0.
- For `j ≥ i+2`, we try every possible vertex `k` between `i` and `j` to form a triangle `(v_i, v_k, v_j)`, splitting the polygon into two smaller subpolygons `(i..k)` and `(k..j)`, plus the cost of the three edges of the triangle.

The final answer for the whole polygon (vertices 0 to n-1) is `dp[0][n-1]`.

To get the actual answer use an array `choice[i][j]` and store the `k` that lead to minimum.

