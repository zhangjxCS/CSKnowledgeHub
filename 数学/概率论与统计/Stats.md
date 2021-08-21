## Outline

Week 1: introduction; data and presentations; descriptive statistics

Week 2: sample spaces, set operations, axioms probability; conditional probability 

Week 3: random variables (rv), pmfs/pdfs/cdfs; independence, expected value, joint rvs

Week 4: conditional rvs, variance/covariance/correlations; inequalities, weak law of large numbers (LLN)

Week 5: discrete named distributions (Bernoulli, binomial, hypergeometric, Poisson, etc)

Week 6: continuous named distributions (normal, exponential, uniform,Chi-square); Poisson process

Week 7: parameters, likelihood, maximum likelihood estimation (MLE), central limit theorem (CLT) continuity corrections, general confidence intervals

Week 8: z-scores, z-intervals, prediction intervals, sampling distribution of the mean estimator; t-intervals, binomial intervals

Week 9: hypothesis testing, z-tests under different situations 

Week 10: one sample t-tests, two samples, paired, binomial tests

Week 11: simple regression overview, MLE estimates and their distributions; multiple regression optimization

Week 12: weighted least squares, multiple linear regression, one way/two way analysis of variance (ANOV A)

Week 13: Chi-square goodness of fit; review

## Fundamental

### Data Collection

- Collect data from real word, sensor, etc.

- Use statistical theory to design an appropriate experiment to generate data (Random Experiment)

### Statistics

- Descriptive statistics: description and summarization of data

- Inferential statistics: statistics for drawing of conclusions

### Population and sample

- population: a total collection of elements

- sample: subgroup of a population

In practice, a given sample generally cannot be assumed to be representative of a population unless that sample has been chosen in a random manner. 

### Descriptive Statistics

- Statistics
  - Frequency: times of occurrences, can be drawed into tables and graphs
  - Relative Frequency: frequency/total, can be described as tables and graphs
  - Sample Mean: the arithmetic average of values; sample mean makes use of all the data values and is affected by extreme values that are much larger or smaller than the others
  - Sample Median: the value in position (*n* + 1)/2; if *n* is even, it is the average of the values in positions *n*/2 and *n*/2 + 1; the sample median makes use of only one or two of the middle values and is thus not affected by extreme values
  - Sample Mode: the value that occurs with the greatest frequency
  - Sample Variance: $s^2=\sum_{i=1}^n(x_i-\bar x)^2/(n-1)$
  - Sample Standard Deviation:   $s=\sqrt{\sum_{i=1}^n(x_i-\bar x)^2/(n-1)}$
  - Sample Percentile: at least 100*p* percent of the data are less than or equal to it and at least 100(1 − *p*) percent are greater than or equal to it

- Graph

  line graph*: plots the distinct data values on the horizontal axis and indicates their frequencies by the heights of vertical lines

<img src="https://i.loli.net/2021/08/03/r6f7OXvjzBysDmi.png" alt="image-20210725132021997" style="zoom:25%;" />

bar graph*: lines in a line graph are given added thickness

<img src="https://i.loli.net/2021/08/03/zu3dkCpZD2eiOxK.png" alt="image-20210725132047571" style="zoom: 25%;" />

*frequency polygon*: plots the frequencies of the different data values on the vertical axis, and then 	connects the plotted points with straight lines

<img src="https://i.loli.net/2021/08/03/PcKB3TWx4kVQCfn.png" alt="image-20210725132109777" style="zoom:25%;" />

*pie chart*: indicate relative frequencies when the data are not numerical

<img src="https://i.loli.net/2021/08/03/TY1qfDdj9sRIQgE.png" alt="image-20210725132313820" style="zoom:25%;" />

*frequency histogram*: a bar graph plot of class data representing frequency

<img src="https://i.loli.net/2021/08/03/eQMw1v9nO7oBbYC.png" alt="image-20210725133002966" style="zoom:25%;" />

*relative frequency histogram*: a bar graph plot of class data representing relative frequency

*stem and leaf plot*: dividing each data value into two parts — its stem and its leaf

<img src="https://i.loli.net/2021/08/03/Vf1wojqmRcvl5Fb.png" alt="image-20210725133055239" style="zoom:25%;" />

*box plot*: 

<img src="https://i.loli.net/2021/08/03/FykGmlfaPcWtOq5.png" alt="image-20210725154041422" style="zoom:25%;" />

### Chebyshev's inequality

- Two-Sided chebyshev's inequality

$$
S_k=\{i, 1\le i\le n:|x_i-\bar x|<ks\}
$$

$$
\frac{|S_k|}{n}\ge 1-\frac{n-1}{nk^2}\ge 1-\frac{1}{k^2}
$$

- One-Sided chebyshev's inequality

$$
N_k=\{i, 1\le i\le n:x_i-\bar x\ge ks\}
$$

$$
\frac{N(k)}{n}\le \frac{1}{1+k^2}
$$

### Normal Data

- normal histograms

<img src="https://i.loli.net/2021/08/03/HUOQtzA3ey5kXaj.png" alt="image-20210725160122767" style="zoom:25%;" />

- approximatedly normal

<img src="https://i.loli.net/2021/08/03/cIMksd6wjrZHWbK.png" alt="image-20210725160204936" style="zoom:25%;" />

- skewed to the left

<img src="https://i.loli.net/2021/08/04/imH1VU87vINjOr2.png" alt="image-20210725160227931" style="zoom:25%;" />

- skewed to the right

<img src="https://i.loli.net/2021/08/04/gTyaECiBHuhAlX3.png" alt="image-20210725160253390" style="zoom:25%;" />

- bimodal

<img src="https://i.loli.net/2021/08/04/gx6K3oLamYtqJQV.png" alt="image-20210725160342698" style="zoom:25%;" />

### Paired Data

- sample correlation coefficient: 

$$
r=\frac{\sum_{i=1}^n(x_i-\bar x)(y_i-\bar y)}{\sqrt{\sum_{i=1}^n(x_i-\bar x)^2\sum_{i=1}^n(y_i-\bar y)^2}}
$$

When *r* > 0 we say that the sample data pairs are *positively correlated*, and when *r* < 0 we say that they are *negatively correlated*

Correlation measures association, not causation, we can not conclude that high correlation means one factor contributes to another

## Parameter Estimation

Week 8: z-scores, z-intervals, prediction intervals, sampling distribution of the mean estimator; t-intervals, binomial intervals

- Parametric Inference: Problems in which the form of the population distribution is specified up to a set of unknown parameters are called *parametric* inference
- Nonparametric Inference: Problems that nothing is assumed about the form of population distribution are called *nonparametric* inference problems

### Sample Statistics

If *X*1,...,*Xn* are independent random variables having a common distribution *F*, then we say that they constitute a *sample* (sometimes called a *random sample*) from the distribution *F*

- sample mean: $\bar X=\frac{X_1+...+X_n}{n}$
- expectation: $E[X]=\mu$
- variance: $Var(\bar X)=\frac{\sigma^2}{n}$

### Central Limit Theorem

$$
\frac{\bar X-\mu}{\sigma/\sqrt n}\sim N(0,1)
$$

No matter how nonnormal the underlying population distribution is, the sample mean of a sample of size at least 30 will be approximately normal

### Maximum Likelihood Estimation

The maximum likelihood estimate $\theta$ is defined to be that value of $\theta$ maximizing $f (x_1,...,x_n|\theta)$ where $x_1,...,x_n$ are the observed values. The function  is $f (x_1,...,x_n|\theta)$ often referred to as the *likelihood* function of $\theta$

### Interval Estimation

![image-20210820114602135](https://i.loli.net/2021/08/20/ZFMNOIx6e3l7n1D.png)

![image-20210820115756232](https://i.loli.net/2021/08/20/5j1VPzKWsmy3Qrw.png)

<img src="https://i.loli.net/2021/08/20/BQWbNFfoG8M1q4n.png" alt="image-20210820115928895" style="zoom: 67%;" />

## Hypothesis Testing

Week 9: hypothesis testing, z-tests under different situations 

Week 10: one sample t-tests, two samples, paired, binomial tests

![image-20210820121250808](https://i.loli.net/2021/08/20/AoT9Dms5nWzwqBt.png)

![image-20210820121349544](https://i.loli.net/2021/08/20/KOhxwSGlo7rkBpJ.png)

![image-20210820121434365](https://i.loli.net/2021/08/20/Br7F6vSQAXzyMcu.png)

## Regression

### Overview

In many situations, there is a single *response* variable $Y$ , also called the *dependent* variable, which depends on the value of a set of *input*, also called *independent*, variables $*x_1,...,x_r$
$$
Y=\beta_0+\beta_1x_1+...+\beta_rx_r+e
$$

$$
E[Y|\bold x]=\beta_0+\beta_1x_1+...+\beta_rx_r
$$

### Least Squares Estimate

The method of least squares chooses as estimators of $\alpha$ and $\beta$ the values of *A* and *B* that minimize *SS*
$$
SS_R=\sum_{i=1}^n(Y_i-A-Bx_i)^2
$$

$$
\frac{\partial SS}{\partial A}=-2\sum(Y_i-A-Bx_i)=0
$$

$$
\frac{\partial SS}{\partial B}=-2\sum x_i(Y_i-A-Bx_i)=0
$$

$$
S_{xY}=\sum(x_i-\bar x)(Y_i-\bar Y)=\sum x_iY_i - n\bar X\bar Y
$$

$$
S_{xx}=\sum (x_i-\bar x)^2=\sum x_i^2-n\bar x^2
$$

$$
S_{YY}=\sum(Y_i-\bar Y)^2=\sum Y_i^2-n\bar Y^2
$$

$$
B=\frac{\sum x_iY_i-n\bar x \bar Y}{\sum x_i^2-n\bar x^2}=\frac{S_{xY}}{S_{xx}}
$$

$$
A=\bar Y-B\bar x
$$

$$
\frac{SS_R}{\sigma^2}\sim X_{n-2}^2
$$

### Parameter Distribution

<img src="../../../Library/Application Support/typora-user-images/image-20210821201317001.png" alt="image-20210821201317001" style="zoom:50%;" />

<img src="../../../Library/Application Support/typora-user-images/image-20210821201340782.png" alt="image-20210821201340782" style="zoom:50%;" />

### Determination

$$
R^2=\frac{S_{YY}-SS_R}{S_{YY}}=1-\frac{SS_R}{S_{YY}}
$$

$$
r^2=\frac{S_{xY}^2}{S_{xx}S_{YY}}=R^2
$$

### Analysis of Residual

$$
Standardized\ Residuals: \frac{Y_i-(A+Bx_i)}{\sqrt{SS_R/(n-2)}}
$$

When the simple linear regression model is correct, the standardized residuals are approximately independent standard normal random variables

### Weighted Least Squares

The variance of a response is not constant but rather depends on its input level. If these variances are known — at least up to a proportionality constant — then the regression parameters α and β should be estimated by minimizing a weighted sum of squares.
$$
Var(Y_i)=\frac{\sigma^2}{w_i}
$$

$$
\sum_i \frac{[Y_i-(A+Bx_i)]^2}{Var(Y_i)}=\frac{1}{\sigma ^2}\sum_i w_i(Y_i-A-Bx_i)^2
$$

### Multiple Linear Regression

$$
\bold Y=\bold X \beta+\bold e
$$

$$
\bold X'\bold X \bold B=\bold X'\bold Y
$$

$$
\bold B=(\bold X'\bold X)^{-1}\bold X'\bold Y
$$

## ANOVA

### Overview

The mean of a random variable depends only on a single factor is called a *one-way analysis of variance*

Models that assume that there are two factors that determine the mean value of a variable is called a *two-way analysis of variance*

In all of the models considered in this chapter, we assume that the data are normally distributed with the same (although unknown) variance $\sigma^2$

### ONE-WAY ANOVA

$$
H_0: \mu_1=\mu_2=...=\mu_m
$$

$$
H_1: not\ all\ means\ are\ equal
$$

![image-20210821205459560](https://i.loli.net/2021/08/21/LS9akngPIBhZi32.png)

### TWO-WAY ANOVA

$$
\mu_{ij}=E[X_{ij}]=\mu+\alpha_i+\beta_j
$$

$$
X_{ij}\sim N(\mu+\alpha_i+\beta_j, \sigma^2)
$$

![image-20210821210523786](https://i.loli.net/2021/08/21/FOQRldKtxqVo1Jv.png)

## Goodness of Fit

Statistical tests that determine whether a given probabilistic mechanism is appropriate are called *goodness of fit* tests
$$
H_0: P(Y=i)=p_i, \ i=1,...,k
$$

$$
H_1:P(Y=i)\ne p_i, \ i=1,...,k
$$

