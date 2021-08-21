## Probability

Frequency interpretation: the probability of the outcome will be observable as being the proportion of the experiments that result in the outcome

Bayesian interpretation: the probability of an outcome is considered a statement about the beliefs of the person who is quoting the probability

### Sample Space

- sample space: set of all possible outcomes of an experiment

- event: any subset *E* of the sample space
- set operation
  - Intersaction: $E \cap F$
  - Union: $E \cup F$
  - Complement: $E^c$
  - mutually exclusive: $E \cap F=\phi$
  - E is contained in F: $E \subset F$
  - communative law: $E \cup F=F\cup E$
  - associative law: $(E\cup F)\cup G=E\cup (F\cup G)$
  - distributive law: $(E\cup F)\cap G=(E\cap G) \cup (F\cap G)$

![image-20210725182157150](https://i.loli.net/2021/08/21/uGvIayV7fZ6zseS.png)

### Conditional Probability

- general formula

$$
P(E|F)=\frac{P(EF)}{P(F)}
$$

- Bayes formula

$$
P(E)=P(EF)+P(EF^c)=P(E|F)P(F)+P(E|F^c)(1-P(F))
$$

$$
P(E)=\sum_{i=1}^nP(EF_i)=\sum_{i=1}^nP(E|F_i)P(F_i)
$$

$$
P(F_j|E)=\frac{P(E|F_j)P(F_j)}{\sum_{i=1}^nP(E|F_i)P(F_i)}
$$

### Independence

- Two events

$$
P(EF)=P(E)P(F)
$$

- Three events

$$
P(EFG)=P(E)P(F)P(G)
$$

$$
P(EF)=P(E)P(F)
$$

$$
P(EG)=P(E)P(G)
$$

$$
P(FG)=P(F)P(G)
$$

## Random Variables

- random variables: quantities of interest that are determined by the result of the experiment

- discrete random variables: Random variables whose set of possible values can be written either as a finite sequence *x*1, . . . , *x**n*, or as an infinite sequence *x*1, . . . 

- continuous random variables: random variables that take on a continuum of possible values

- cumulative distribution function: $F(x)=P(X\le x)$

- probability mass function: $p(a)=P(X=a)$

  for discrete random variables: $F(a)=\sum_{x\le a}p(x)$

  for continuous random variables: $F(a)=\int_{-inf}^ap(x)dx$

### Jointly Distributed Random Variables

- joint cumulative probability distribution: $F(x,y)=P(X\le x,Y\le y)$
- marginal probability distribution: $F_X(x)=F(x, inf), F_Y(y)=F(inf, y)$
- joint probability mass function: $f(x_i, y_j)=P(X=x_i, Y=y_j)$
- marginal probability mass function: $f_X(x)=\int f(x, y)dy, f_Y(y)=\int f(x, y)dx$

### Independent Random Variables

- $F(a,b)=F_X(a)F_Y(b)$
- $p(x,y)=p_X(x)p_Y(y)$

### Conditional Distribution

- $p_{X|Y}(x|y)=\frac{p(x,y)}{p_Y(y)}$

### Expectation

- Definition
  - discrete variables: $E[X]=\sum_{i=1}^n x_ip(x_i)$
  - continuous variables: $E[X]=\int xf(x)$
- Properties
  - $E[g(x)]=\sum_xg(x)p(x)$, $E[g(x)]=\int g(x)f(x)dx$
  - $E[aX+b]=aE[X]+b$
  - $E[g(X,Y)]=\sum_y \sum_xg(x,y)p(x,y)$, $E[g(X,Y)]=\int \int g(x,y)f(x, y)dxdy$
  - $E[X+Y]=E[X]+E[Y]$

### Variance

- Definition: $Var(X)=E[(X-\mu)^2]=E[X^2]-E[X]^2$
- Properties: $Var(aX+b)=a^2Var(X)$

### Covariance

- Definition: $Cov(X, Y)=E[(X-\mu_x)(Y-\mu_y)]=E[XY]-E[X]E[Y]$
- Properties:
  - $Cov(X,Y)=Cov(Y,X)$
  - $Cov(X, X)=Var(X)$
  - $Cov(aX, Y)=aCov(X,Y)$
  - $Cov(X_1+X_2,Y)=Cov(X_1,Y)+Cov(X_2,Y)$
  - $Cov(\sum_i X_i, \sum_j Y_j)=\sum_i\sum_jCov(X_i, Y_j)$
  - $Var(X+Y)=Var(X)+Var(Y)+2Cov(X, Y)$
  - $Var(\sum_i X_i)=\sum_i Var(X_i)+\sum_i\sum_jCov(X_i, X_j)$
  - Independent random variables X Y => $Cov(X,Y)=0$
  - $Corr(X,Y)=\frac{Cov(X,Y)}{\sqrt{Var(X)Var(Y)}}$

### Moment Generating

- Definition: $\phi(t)=E[e^{tX}]$
- Properties:
  - $\phi'(t)=E[Xe^{tX}]$
  - $\phi^n(0)=E[X^n]$
  - sum of independent random variables is just the product of the individual moment generating functions: $\phi_{X+Y}(t)=\phi_X(t)\phi_Y(t)$
  - the moment generating function uniquely determines the distribution

### Law of large numbers

- Markov's inequality:  X is a random variable that takes only nonnegative values, for any value a > 0

$$
P(X\ge a)\le \frac{E[X]}{a}
$$

- Chebyshev's inequality: X is a random variable with mean $\mu$ and variance $\sigma^2$, then for any value k > 0

$$
P\{|X-\mu|\ge k \}\le\frac{\sigma^2}{k^2}
$$

- Weak law of large numbers: X1,X2,..., be a sequence of independent and identically distributed random variables, each having mean $E[X_i] = \mu$, then for any $\epsilon$ > 0, with n -> inf

$$
P\{|\frac{X_1+...+X_n}{n}-\mu|>\epsilon\}->0
$$

## Special Random Variables

### Bernoulli

An experiment whose outcome can be classified as either a “success” or as a “failure”
$$
P(x=0)=1-p
$$

$$
P(x=1)=p
$$

$$
E[X]=p
$$

$$
Var(X)=p(1-p)
$$

### Bionomial

*n* independent trials, each of which results in a “success” with probability *p* and in a “failure” with probability 1 − *p*, If *X* represents the number of successes that occur in the *n* trials, then *X* is said to be a *binomial* random variable with parameters (*n*, *p*)
$$
P(x=i)=C_n^ip^i(1-p)^{n-i}
$$

$$
E[X]=np
$$

$$
Var(X)=np(1-p)
$$

### Possion

A random variable *X*, taking on one of the values 0, 1, 2,..., is said to be a Poisson random variable with parameter$\lambda$, $\lambda$ > 0, if its probability mass function is given by
$$
P(X=i)=e^{-\lambda}\frac{\lambda^i}{i!}
$$

$$
\phi(t)=exp(\lambda(e^t-1))
$$

$$
E[X]=\phi'(0)=\lambda
$$

$$
Var(X)=\phi''(0)-(E[X])^2=\lambda
$$

$$
\frac{P(X=i+1)}{P(X=i)}=\frac{\lambda}{i+1}
$$

The Poisson random variable may be used as an approximation for a binomial random variable with parameters (*n*, *p*) when *n* is large and *p* is small, in that case $\lambda=np$

<img src="https://i.loli.net/2021/08/04/Xy32mSYOxFnrVwz.png" alt="image-20210726235942973" style="zoom:25%;" />

### Hypergeometric

A bin contains *N* + *M* batteries, of which *N* are of acceptable quality and the other *M* are defective. A sample of size *n* is to be randomly chosen (without replacements) in the sense that the set of sampled batteries is equally likely to be any of the 􏰍*N*+*M*􏰀 subsets of *n* size *n*. If we let *X* denote the number of acceptable batteries in the sample, then
$$
P(X=i)=\frac{C_N^iC_{n-i}^M}{C_{N+M}^n}
$$

$$
E[X]=\frac{nN}{N+M}
$$

$$
Var(X)=\sum_iVar(X_i)+2\sum\sum Cov(X_i,X_j)=\frac{nNM}{(N+M)^2}(1-\frac{n-1}{N+M-1})
$$

For fixed *p*, as *N* + *M* increases to ∞, Var(*X*) converges to *np*(1 − *p*), which is the variance of a binomial random variable with parameters (*n*, *p*)

### Uniform

$$
f(x)=\frac{1}{\beta-\alpha}, if \alpha\le x \le \beta
$$

$$
E[X]=\frac{\alpha+\beta}{2}
$$

$$
Var(X)=\frac{(\beta-\alpha)^2}{12}
$$

### Normal

$$
f(x)=\frac{1}{\sqrt{2\pi}\sigma}e^{-(x-\mu)^2/2\sigma ^2}
$$

### Exponential

$$
f(x)=\lambda e^{-\lambda x}, if x\ge0
$$

$$
E[X]=1/\lambda
$$

$$
Var(X)=1/\lambda^2
$$

$$
P(X>s+t|X>t)=P(X>s)
$$

Exponentially distributed random variables are memoryless

### Chi-square

$$
X=Z_1^2+Z_2^2+...+Z_n^2
$$

$$
E[X]=n
$$

$$
Var[X]=2n
$$

### t-Distribution

$$
T_n=\frac{Z}{\sqrt{X_n^2/n}}​
$$

### F-Distribution

$$
F_{n,m}=\frac{X_n^2/n}{X_m^2/m}
$$



