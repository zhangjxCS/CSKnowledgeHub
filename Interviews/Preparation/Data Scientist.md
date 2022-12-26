## Job Description

### Responsibilities

As a Data Scientist intern, you will formulate approaches to solve problems using well-defined algorithms and data sources. You will incorporate an understanding of product functionality and customer perspective to provide context for those problems. You will use data exploration techniques to discover new questions or opportunities within your problem area and propose applicability and limitations of the data. Successful Data Scientists will interpret the results of their analysis, validate their approach, and learn to monitor, analyze, and iterate to continuously improve.  

You will engage with peer stakeholders to produce clear, compelling, actionable insights that influence product and service improvements that will impact millions of customers. As a Data Scientist, you will also engage in the peer review process and act on feedback while learning innovative methods, algorithms, and tools to increase the impact and applicability of your results. 

### Qualifications

- Currently pursuing a bachelor's or master's in Computer Science, Mathematics, Economics, Statistics, Applied Sciences like Biostatistics, Physics, Chemistry, Computational Neurology and or other quant-focused field with at least one semester/quarter remaining after internship. 
- Some Engineering experience and or project course work using large data systems on SQL, Hadoop, etc. 
- Proficiency using one or more programming or scripting language to work with data such as: Python, Perl, or C#. 
- Some experience and or project course work performing data analysis and applying statistics working with tools such as: Excel, R, MATLAB, AMPL, or SAS. 
- Some experience and or project course work with product and service telemetry systems. 
- Some A/B Testing or experimentation (this can be from conducting real life science experiments, hypothesis testing in stats etc.) Not required but ideal. 
- Some experience or course work applying basic ML to a type of data and or used algorithms to conduct experiments on data. Machine Learning strongly encouraged. 
- Passion to learn from your peers, manager, and other stakeholders in the Data Science domain.  
- Ability to interact with peers and stakeholders to drive product and business impact. 
- Strong interpersonal and communications skills. 

## Tech Question

### Probs and Stats

- Three friends in Seattle told you it's rainy. Each has a probability of 1/3 of lying. What's the probability of Seattle is rainy.

$$
P(r|yyy)=\frac{P(yyy,r)}{P(yyy)}=\frac{P(yyy|r)P(r)}{P(yyy|r)P(r)+P(yyy|\bar r)P(\bar r)}=\frac{P(y|r)^3P(r)}{P(y|r)^3P(r)+P(y|\bar r)^3P(\bar r)}=\frac{P(r)}{P(r)+8P(\bar r)}
$$

- Difference between box plot and histogram

They are both for continuous variables visualization. Histograms are better in determining the distribution of the data, box plots shows the statistics of the data and is easy to compare multiple data sets better than histograms as they are less detailed and take up less space.

- Power

Statistical power is used in a binary hypothesis test, the definition for that is the probability correctly rejects the null hypothesis when the alternative hypothesis is true. The higher the statistical power, the better the test is. It is commonly used in experiment design to calculate the minimum sample size.

- Type 1 error

It is use to categorize error in a binary hypothesis test. The definition for that is mistakenly reject true null hypothesis. The larger the value, the less reliable a test is. It is commonly used in A/B testing when we observe differences but there is no difference.

- Type 2 error

It is used to categorize error in a binary hypothesis test. The definition for that is fail to reject false null hypothesis. The larger the value, the less reliable a test is. It is commonly used in A/B testing when we don't observe differences but there is a difference.

- Confidence interval

When we want to know how variable a sample result is. The definition for that is that it's an interval of numbers. It tells us how likely it covers the true value. The wider the interval, the more uncertainty about the sample. 

- P-Value

It is commonly used in hypothesis testing when connect observation and conclusion. It is a conditional probability that under the assumption that null hypothesis is true, and the probability of results as extreme as observed results. The lower the p value, the less the support of null hypothesis. It is commonly used in A/B testing when we test a metric in treatment and control group.

- Bernoulli test: test proportion

- Z-test: large sample or small sample from normal distribution with known population variance
- t-test: large sample or small sample from normal distribution with unknown population vairance

### ML

- Overfitting and Underfitting

Overfitting: The model is too complex. Good fitting performance on training set but bad predicting performance on test set

Underfitting: The model is too simple. Bad fitting performance on training set and bad predicting performance on test set

- What is bias-variance trade off ? 

If our model is too simple and has very few parameters then it may have high bias and low variance. On the other hand if our model has large number of parameters then it’s going to have high variance and low bias.

- Generative model and discriminative models

Generative model: Learn the joint probability distribution $p(x,y)$, including naive bayes, gaussian mixture model, hidden markov model, etc.

Discriminative model: Learn the conditional probability distribution $p(y|x)$, including linear regression, logistic regression, decision tree, etc.

- Expected risk, Empirical risk and structural risk

Expected risk: fitting performance on test set

Empirical risk: fitting performance on training set

Structural risk: model complexity

- Evaluation metric

TP：True Positive
FP：False Positive
TN：True Negative
FN：False Negative

Accuracy：TP+TN/TP+FP+TN+FN

Precision：TP/TP+FP

Recall：TP/TP+FN

F1 score：1/F1 = 1/P + 1/R

$MSE=\frac{1}{n}\sum_{i=1}^n||y_i-\hat y_i||_2^2$

$MAE=\frac{1}{n}\sum_{i=1}^n|y_i-\hat y_i|$

R squared

- How do you calculate AUC?

An ROC curve plots TPR vs. FPR at different classification thresholds. Lowering the classification threshold classifies more items as positive, thus increasing both False Positives and True Positives. The following figure shows a typical ROC curve.

 AUC stands for "Area under the ROC Curve." That is, AUC measures the entire two-dimensional area underneath the entire ROC curve (think integral calculus) from (0,0) to (1,1).

<img src="https://developers.google.com/machine-learning/crash-course/images/AUC.svg" alt="AUC (Area under the ROC Curve)." style="zoom:50%;" />

- Can you explain what regularization is. What's the difference between L1/L2 regularization

Regularization is to avoid learning a complex or flexible model, and to avoid the risk of overfitting.

The lasso penalty will force some of the coefficients quickly to zero. This means that variables are removed from the model, hence the  sparsity. 

Ridge regression will more or less compress the coefficients to  become smaller. This does not necessarily result in 0 coefficients and  removal of variables.

- Bagging, boosting, stacking

Bagging: considers homogeneous weak learners, learns them independently from each other in parallel and combines them following some kind of  deterministic averaging process

Boosting: considers homogeneous weak learners, learns them sequentially in a very  adaptative way (a base model depends on the previous ones) and combines  them following a deterministic strategy

Stacking: considers heterogeneous weak learners, learns them in parallel and combines them by training a meta-model to output a prediction based on  the different weak models predictions

- Describe how gradient boost works

Gradient boosting is a  type of machine learning boosting. It relies on the intuition that the  best possible next model, when combined with previous models, minimizes  the overall prediction error. The key idea is to **set the target outcomes for this next model in** order to minimize the error.

- introduce some dimension reduction technique

PCA: maximize the variance of projected data

LDA: supervised learning; maximize interclass variance / intraclass variance

NMF: give the dimensionality given non-negativity constraints

t-SNE: statistical method for visualizing high-dimensional data

- When do you choose to use neural network versus SVM?

high dimensional data, large scale data

- How do you detect if a new observation is outlier? 

Visualize the data distribution using boxplot, histogram, etc.

- How to deal with unbalanced binary classification

undersampling, oversampling

### Algorithm

Given pairs of characters, judge if all the characters could be split into two sets that the graph is a barpartite

Question about sorting an array

find the 5th largest element in an array

How to compute an inverse matrix faster by playing around with some computational tricks?

### Product

- Clarify the scenario
- Time sudden or stable; internal external
- Other product/feature by the same company
- Segment by user demographic and behavioral features: regions languages platform
- Decompose the metric
- Summarize overall approach

How to measure success

- Clarify question: what it does; how is it used; who is it for
- Define metrics: <3 metrics, 2 success metrics day number of bookings and conversion rate 1 guardrail metric cancelation rate

Launch or not

- Clarify the goal and define metrics
- Experimentation: how to design it; how long to run test
- Recommendation based on experiment results

### A/B test

make data driven decision

change to product - change to metrics

change to colors - change to use engagement

- prerequisite: define key metrics; change are easy to make; enough samples
- Experiment design: specific population vs all users; samples size; how long to run seasonality day of week primacy and novelty effect
- Running experiment: instrument loggings;
- Result to decision: sanity checks; tradeoffs between different metrics; cost of launching a feature; benefit should outweigh the cost
- Post launch monitoring: long-term effect

long term metric: company's mission

short term metric: driver metric
