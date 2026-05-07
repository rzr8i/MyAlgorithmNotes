## Knight's Tour

Given an $n \times n$ chessboard (usually $8 \times 8$), a knight starts on a given square and must move so that it visits every square exactly once (a Hamiltonian path on the knight's graph).

### Variants:

- Open tour: finishes anywhere (different start/end).
- Closed tour (re‑entrant): last move returns to starting square (Hamiltonian cycle).

### State space representation
- Board $n \times n$: store move number (or boolean visited) for each square

