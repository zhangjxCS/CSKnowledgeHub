## Dynamic Programming

The structure of dynamic programming problem

- Overlapping subproblems: the problem can be broken down into smaller versions of the original problem that are re-used multiple times
- Optimal substructures: an optimal solution can be formed from optimal solutions to the overlapping subproblems of the original problem.

Two way to solve dynamic programming

- Top-down: also known as memoization. A top-down implementation is usually much easier to write. This is because with recursion, the ordering of subproblems does not matter, whereas with tabulation, we need to go through a logical ordering of solving subproblems.
- Bottom-up: also known as tabulation. A bottom-up implementation's runtime is usually faster, as iteration does not have the overhead that recursion does.

Characteristic of dynamic programming

- Optimum value: DP problems will ask for the optimum value (maximum or minimum) of something, or the number of ways there are to do something.

- Dependency: future "decisions" depend on earlier decisions. Deciding to do something at one step may affect the ability to do something in a later step.

Top-down to Bottom-up

- State the subproblem
- Give the recurrence
- Give the boundary conditions
- Give the order to fill in the dp table
- Analyze the time and space complexity

## Linear Programming

### Fundamentals

Input:

- a set of variables
- a set of linear constraints on the variables
- a linear objective function to maximize

Output:

- an assignment of real values to the variables such that the constraints are satisfied and the objective function is maximized of minimized

- The optimum is achieved at a vertex of the feasible region, except for the infeasible linear program and unbounded optimum value

### Duality

- Weak duality: $V_P\le V_D$
- Strong duality: $V_P=V_D$

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h5hed9g4kaj20p205et8y.jpg" alt="image-20211208151122285" style="zoom:67%;" />

Linear Programming problem can be rewrite as follows

- a maximization or a minimization problem
- all constraints are inequalities in the same direction
- all variables are non-negative

$$
max_{\bold x\ge 0}\ \ \bold c^T\bold x
$$

$$
s.t.\ \ \bold A\bold x\le \bold b
$$

The dual problem can be given as follows
$$
min_{\bold y\ge 0}\ \ \bold b^T\bold y
$$

$$
s.t.\ \ \bold A^T\bold y\ge \bold c
$$

7-step to formulate a LP to its dual LP problem

- Work with a minimization problem
- Rewrite the LP in $\le$ and $=$ form
- Introduce the dual variables, non-negative for inequality constraint and unrestricted for equality constraint
- Maximize the sum of objective function and multiplication of dual variables and constraints
- Group terms by primal variables
- Eliminate the primal variables to ge the dual LP
- Formulate the dual problem

## Integer Programming

An NP problem

given a system of linear inequalityies in n variables and m constraints with integer coefficients and a integer target value k, does it have an integer solution of value k?
$$
max\ \ \bold c^T\bold x
$$

$$
s.t.\ \ \bold A\bold x\le \bold b
$$

$$
\bold x \in \bold Z^n
$$

## 网络流

### 网络流条件

- Capacity constraints: for all $e\in E$, $0\le f(e)\le c_e$
- Flow conservation: for all $v\in V-{s, t}$, $\sum f(u,v)=\sum f(v, w)$

### 网络流求解

通过残差图和增广路径求解，Ford-Fulkerson algorithm

forward edge: edges with leftover capacity

backward edge: edges that are already carrying flow so as to divert it to a different direction

Augment path: P is a simple s-t path in $G_f$, push at most $min c_f(e)$ capacity, which is the minimum residual capacity in the s-t path

- For a forward edge, $f'(e)=f(e)+c(P)$
- for a backward edge, $f'(e)=f(e)-c(P)$

最大流与最小割强对偶：in every flow network, the maximum value of an s-t flow equals the minimum capacity of an s-t cut

## Reduction

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h5hee1p8v1j20oi0bojs6.jpg" alt="image-20211206152700332" style="zoom: 50%;" />

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
