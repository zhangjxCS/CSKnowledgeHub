## Reduction

<img src="/Users/shawnzhang/Library/Application Support/typora-user-images/image-20211206152700332.png" alt="image-20211206152700332" style="zoom: 50%;" />

- Input: give inputs to the two problems
- Transformation: show how to transform the original problem to the new problem
- Proof: prove the equivalence of the two problems

## P与NP

An optimization problem can be transformed in to a equivalent decision problem

- Supply a target value for the quantity to be optimized
- ask whether the value can be attained

Class P: the set of decisions problems that can be solved by polynomial-time algorithms

Class NP: the set of decisions problems that have an efficient certifier
$$
P\subseteq NP
$$
Class NP-Complete: $X\in NP$, for all $Y\in NP$, $Y\le_P X$

X is solvable in polynomial time if and only if P=NP

Prove that a problem Y is NP-Complete

- $Y\in NP$
- $X\le_P Y$

## NPC Examples

### IS

Independent set: a set of vertexes that there is no edge between any pair of nodes in the set

IS(D): given a graph G and an integer k, does G have an independent set of size k

### VC

Vertex cover: a set of vertexes that every edge in the graph has at least on endpoint in the set

VC(D): given a graph G and an integer k, does G have a vertex cover of size k

### SAT

Literal: a variable or its negation

CNF: a boolean formula consists of conjunctions of clauses each of wich is a disjunction of literals

truth assignment: an assignment of truth values from {0, 1} to each xi

SAT: given a formula in CNF with n variables and m clauses, is the formula satisfiable

3SAT: SAT problem and each clause has exactly 3 literals

Circuit SAT$\le_P$SAT, SAT$\le_P$3SAT, 3SAT$\le_P$IS(D)

### Hamiltonian Cycle

Given a graph G=(V,E), is there a simple cycle that visits every vertex exactly once?

### TSP

Tour: a simple cycle that visits every vertex exactly once

reduced from undirected Hamiltonian cycle; set the weight for edges 1 if $e\in E$, else 2

### Set Packing

given a set U of a elements, a collection $S_1, S_2, ..., S_b$, of subsets of U, and a number k, is there a collection of at least k subsets such that no two of them intersect?

reduced from independent set