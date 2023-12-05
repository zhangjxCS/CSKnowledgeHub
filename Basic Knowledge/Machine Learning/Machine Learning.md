[toc]

# Machine Learning

## Concept

## Evaluation

### Cross-Validation

Cross-validation is a technique for model evaluation that involves splitting the dataset into multiple subsets (folds). The model is trained and tested on different combinations of these folds to ensure robust evaluation.

**Example:** In k-fold cross-validation, the dataset is divided into k subsets. The model is trained on k-1 subsets and tested on the remaining subset. This process is repeated k times, and the results are averaged to obtain a reliable estimate of model performance.

#### Stratified K-Fold Cross-Validation

Stratified K-fold cross-validation is similar to K-fold cross-validation, but it **ensures that each fold contains roughly the same proportion of samples from each class** in the dataset. This is particularly useful for imbalanced datasets, where some classes have very few samples.



#### Advantages of Cross-Validation

Cross-validation has several advantages, including:

- It provides a more **accurate estimate** of the model’s performance compared to using a single dataset for training and testing.
- It helps to **identify overfitting**, which can lead to better generalization of the model.
- It can be used to **compare different models** and select the best one for a given task.

#### Disadvantages of Cross-Validation

Cross-validation also has some disadvantages, including:

- It can be **computationally expensive**, especially for large datasets and complex models.
- It can be **sensitive to the choice of k**, and different values can lead to different results.
- It **assumes that the data is independent and identically distributed**, which may not be true in some cases.

### Generative Model and Descrimitive Model

They differ in their primary objectives and how they model the underlying data distribution.

- Generative models model the joint probability distribution of the input and output, i.e., they can generate new data points.
  - **Gaussian Mixture Models (GMM)**: These models use a mixture of Gaussian distributions to model the underlying data distribution. They are commonly used for clustering and density estimation.
  - **Hidden Markov Models (HMM)**: HMMs model sequential data and are used in applications like speech recognition and part-of-speech tagging.
  - **Variational Autoencoders (VAE)**: VAEs are used to learn a probabilistic mapping between high-dimensional data and a lower-dimensional latent space. They are often used in image generation tasks.
  - **Generative Adversarial Networks (GAN)**: GANs consist of a generator and a discriminator, and they are used for generating realistic images and data.
- Discriminative models model the conditional probability of the output given the input, typically used for classification tasks.
  - **Logistic Regression**: A simple and widely used model for binary classification.
  - **Support Vector Machines (SVM):** SVMs find the optimal hyperplane that best separates different classes in the data.
  - **Decision Trees**: Decision trees recursively split the data based on the most discriminative features, making them excellent for classification.
  - **Neural Networks (Deep Learning)**: Deep learning models, like feedforward neural networks and convolutional neural networks (CNNs), are often used for a wide range of discriminative tasks, including image classification and natural language processing.

### Model Metrics

(e.g., accuracy, precision, recall, F1-score, ROC AUC) ![Screen Shot 2023-10-28 at 1.35.53 AM](/Users/moya_shoreline/Library/Application Support/typora-user-images/Screen Shot 2023-10-28 at 1.35.53 AM.png)

1. **Precision**: This metric assesses the accuracy of the positive predictions made by the model. It's calculated as the ratio of true positive predictions to the total number of predicted positives (true positives + false positives). High precision indicates that among the predicted positive instances, a significant proportion is actually positive.

   Formula: Precision = True Positives / (True Positives + False Positives)

2. **Recall (Sensitivity)**: Recall evaluates the model's ability to identify all positive instances. It's calculated as the ratio of true positive predictions to the total number of actual positives (true positives + false negatives). High recall means the model captures most positive instances.

   Formula: Recall = True Positives / (True Positives + False Negatives)

Tradeoff between precision and recall:

- **Balancing Precision and Recall**: There's often an inverse relationship between precision and recall. If you increase one, the other may decrease. For instance, a model can achieve perfect precision by making fewer positive predictions, but this might lead to missing some actual positives, resulting in lower recall. Conversely, maximizing recall can sometimes decrease precision because of higher numbers of positive predictions.

Examples of when to focus which:

**High Recall:**

- **Disease Detection in Medical Tests:** In medical diagnostics, especially in cancer screenings, high recall is crucial. It's better to have false positives (predicting someone has cancer when they don't) but avoid false negatives (missing a cancer diagnosis) because missing a true case could be life-threatening.
- **Fraud Detection in Banking:** When dealing with fraudulent transactions, high recall is essential. Missing even a single fraudulent transaction could be highly detrimental. Therefore, the detection model needs to have a high recall to catch as much fraudulent activity as possible.

**High Precision:**

- **Search Engine Results:** In search engines, users want the results to be highly relevant to their query. High precision ensures that most of the results are what the user is looking for, even if it might miss some potentially relevant content. False positives are less tolerable in this case.
- **Spam Detection in Emails:** While high recall is crucial to avoid missing important emails, high precision is also important in this scenario. Marking a non-spam email as spam is an inconvenience, but it might be more tolerable than missing an important email. Therefore, email spam filters often focus on achieving high precision.

Type I and Type II errors:

- **Type I Error**: Also known as a "false positive," this occurs when the model incorrectly predicts a positive (True) outcome when the actual result is negative (False).

- **Type II Error**: This is a "false negative" where the model incorrectly predicts a negative (False) outcome when the actual result is positive (True).

  

### Overfitting and Underfitting 

**bias–variance tradeoff** describes the relationship between a model's complexity, the accuracy of its predictions, and how well it can make predictions on previously unseen data that were not used to train the model.

- Overfitting: High variance, model too complex, fit training data too closely, resulting in poor generalization to unseen data
- Underfitting: high bias, model too simple, fail to capture pattern of training data, low accuracy

Preventing Overfitting: Common techniques include:

- Regularization (L1, L2, or both).

- Cross-validation for hyperparameter tuning.

- Increasing the size of the training dataset.

- Feature selection or reduction.

- Early stopping during training.

  - **Loss**: This metric measures how well a model is performing. It represents the error in predictions made by the model during training. In essence, the goal is to minimize this metric. Therefore, a decrease in loss indicates an improvement in the model's performance. Loss can increase when the model starts to generalize less effectively, leading to overfitting on the training data.
  - **Accuracy**: Accuracy measures the proportion of correctly classified instances out of the total instances. As accuracy increases, it indicates a better ability to make correct predictions. However, sometimes accuracy might stagnate or drop, especially in cases of overfitting, where the model performs well on training data but poorly on unseen (validation or test) data.

  During training, it's common for loss to decrease and accuracy to increase simultaneously as the model learns from the data. However, when the model begins to overfit the training data, the loss might start to increase, while accuracy can decrease, or its growth might plateau. 

### L1 and L2 regularization

- L1 regularization encourages sparsity by adding absolute value of the coefficients as a penalty term

- Why L1 Results in Sparsity:

  here sparsity refer to  many of the coefficients in the model are exactly zero

  The L1 regularization penalty has the effect of "shrinking" some coefficients all the way to zero, effectively removing them from the model

  L1 regularization encourages some coefficients to be exactly zero because the L1 norm has sharp corners at zero, making it more likely to eliminate irrelevant features.

  **Advantages of Sparsity:**

  - **Feature Selection:** Sparse models have the property of selecting a subset of features, effectively performing feature selection. This can be valuable in situations where not all features are equally informative or relevant.
  - **Interpretability:** A model with fewer non-zero coefficients is often simpler and more interpretable.
  - **Reduced Overfitting:** Sparsity can also help in reducing overfitting, as the model is less likely to memorize noise in the training data.

![img](https://lh7-us.googleusercontent.com/l_APqRwD1P0hFB5U1ixrRR6G2ijb66SBYlNm3Vn9SfDZozMiK9qWs0mM2BL1ftvezcoitLMTwVcRj5BGVfWzugXk3qBH5f_ZYsXnWEnV7wHZ6NT0Gqk6hHLQh-MNDolk70hhJDomk0Jz53uQK1UxN0Q)

- L2 regularization(Ridge) adds the square of coefficients as a penalty term to prevent overfitting
- ![img](https://lh7-us.googleusercontent.com/M0NDnJYRKzPz6hrOegUUx_6UCpKKK_2YpELYO0XkH4iH2nGDLMH0zChzZ1yZvudlrEAAfxu2cQoJfDfQpxC4AgPXE22qClkt3GMqLXqU3Ofr1-ipKZOrwL09AIzMatE-qPMH5D3IQSNU-EyffP81_4Y)

#### Model evaluation metrics selection

1. **Accuracy:**
   - **Use Case:** When the class distribution is roughly equal and there is no significant class imbalance.
   - **Example:** In a binary classification problem where you want to predict whether an email is spam or not, and both spam and non-spam emails are equally important.
2. **Precision and Recall:**
   - **Use Case:** When there is an imbalance in class distribution, and false positives or false negatives have different costs.
   - **Example:** In a medical diagnosis scenario, where identifying positive cases (e.g., disease presence) is crucial, you might use recall. However, if false positives (misdiagnosing a healthy person) have significant consequences, you may also consider precision.
3. **F1 Score:**
   - **Use Case:** When there is an uneven class distribution, and you want a balance between precision and recall.
   - **Example:** In a fraud detection scenario, where fraud cases are rare compared to legitimate transactions, and both false positives and false negatives are undesirable.
4. **Area Under the Receiver Operating Characteristic (AUROC) Curve and Area Under the Precision-Recall (AUC-PR) Curve:**
   - **Use Case:** When you want to evaluate the model's ability to discriminate between positive and negative instances across different probability thresholds.
   - **Example:** In a marketing campaign where you predict whether a customer will purchase a product, and you want to assess the model's ability to rank customers by their likelihood to buy.
5. **Log Loss (Cross-Entropy Loss):**
   - **Use Case:** When you want a probabilistic interpretation of the model's predictions.
   - **Example:** In a binary classification problem where you want to assess how well the predicted probabilities align with the actual outcomes.
6. **Matthews Correlation Coefficient (MCC):**
   - **Use Case:** When dealing with imbalanced datasets and you want a balanced metric that considers all quadrants of the confusion matrix.
   - **Example:** In a rare event prediction, such as detecting a particular type of rare disease where true negatives greatly outnumber the other classes.
7. **Specificity and Sensitivity:**
   - **Use Case:** In certain applications where minimizing false positives or false negatives is critical.
   - **Example:** In a diagnostic test where false positives lead to unnecessary treatments and false negatives result in missed opportunities for early intervention.

### Hyperparameter Tuning

- **Grid Search:** Understanding the process of defining a grid of hyperparameters and exhaustively searching through the grid to find the best parameter combination.
- **Random Search:** Using a random search over the hyperparameter space, which might be more efficient than grid search for large search spaces.
- **Bayesian Optimization:** Exploring an approach that uses probabilistic models to predict the model's performance under different hyperparameter settings and selects the most promising hyperparameters to evaluate.
- **Hyperparameter Impact:** Assessing how various hyperparameters influence the model's performance and efficiency.

### Model Selection

- **Model Evaluation Metrics:** Understanding different evaluation metrics based on the nature of the problem (accuracy, precision, recall, F1-score, AUC-ROC, etc.).
- **Model Comparison:** Learning methods to compare models based on performance, interpretability, computational complexity, and scalability.
- **Trade-offs:** Evaluating the trade-offs between various models, such as the simplicity of linear models versus the complexity of neural networks or the interpretability of decision trees versus the robustness of ensemble methods.

## Feature Engineering

- **Creating New Features:** Employing domain knowledge to create new features or engineer existing features for better model performance.
- **Normalization and Scaling:** Using techniques like Min-Max scaling or Z-score normalization to scale numerical features to a similar range.
- **Handling Missing Data:** Dealing with missing values through imputation techniques or encoding missingness as a separate feature.
- **Dimensionality Reduction:** Exploring methods like PCA, t-SNE, or feature selection techniques to reduce the number of features and focus on the most important ones.

### Encoding

- **Label Encoding:** Converting categorical features to ordinal values, useful when dealing with ordinal data.
- **One-Hot Encoding:** Transforming categorical variables into a binary matrix, suitable for nominal data.
- **Target Encoding:** Encoding categorical variables based on target variable statistics, which can be effective in certain situations.
- **Handling High-Cardinality Features:** Techniques for managing features with a large number of distinct categories to prevent the curse of dimensionality.

### Feature Importance

- **Understanding Feature Importance:** Realizing the role of different features in predicting the target variable.
- **Techniques to Calculate Feature Importance:** Employing methods like mean decrease in impurity (for decision trees), coefficient weights (for linear models), SHAP values, or permutation importance.
- **Feature Selection Based on Importance:** Selecting a subset of features based on their importance to improve model performance and interpretability.v

## Regression

### Linear Regression

**Linear regression** is a supervised learning algorithm used for predicting and forecasting values that fall within a continuous range, such as sales numbers or housing prices. It is a technique derived from statistics and is commonly used to establish a relationship between an input variable (X) and an output variable (Y) that can be represented by a straight line.

##### MLE

The linear regression model is given by:

*y*=*β*0+*β*1*x*+*ϵ*

The likelihood function for a single data point  (*x**i*,*y**i*) is given by the probability density function (PDF) of the normal distribution:

*L*(*β*0,*β*1,*σ*2∣*x**i*,*y**i*)=2*π**σ*21exp(−2*σ*2(*y**i*−(*β*0+*β*1*x**i*))2)

The joint likelihood function for the entire dataset {(*x**i*,*y**i*)} is the product of the individual likelihoods:

*L*(*β*0,*β*1,*σ*2∣{(*x**i*,*y**i*)})=∏*i*=1*n*2*π**σ*21exp(−2*σ*2(*y**i*−(*β*0+*β*1*x**i*))2)

The log-likelihood function is often used:

log*L*=−2*n*log(2*π**σ*2)−∑*i*=1*n*2*σ*2(*y**i*−(*β*0+*β*1*x**i*))2

The MLE estimates for simple linear regression are:

*β*^1=∑*i*=1*n*(*x**i*−*x*ˉ)2∑*i*=1*n*(*x**i*−*x*ˉ)(*y**i*−*y*ˉ)

*β*^0=*y*ˉ−*β*^1*x*ˉ

*σ*^2=*n*∑*i*=1*n*(*y**i*−(*β*^0+*β*^1*x**i*))2

#### Assumptions

Ref: https://www.statisticssolutions.com/free-resources/directory-of-statistical-analyses/assumptions-of-linear-regression/

- Linear relationship

  linear regression needs the relationship between the independent and dependent variables to be linear. It is also important to check for outliers since linear regression is sensitive to outlier effects.

- Multivariate normality

  the linear regression analysis requires all variables to be multivariate normal. This assumption can best be checked with a histogram or a Q-Q-Plot. Normality can be checked with a goodness of fit test, e.g., the Kolmogorov-Smirnov test. When the data is not normally distributed a non-linear transformation (e.g., log-transformation) might fix this issue.

- No or little [multicollinearity](https://www.statisticssolutions.com/multicollinearity/)
  linear regression assumes that there is little or no multicollinearity in the data. Multicollinearity occurs when the independent variables are too highly correlated with each other. It can be problematic because it makes it difficult to isolate the individual effect of each variable on the dependent variable.

  Detecting multicollinearity can be done using correlation matrices or variance inflation factors (VIF). Remedies include removing one of the correlated variables or using techniques like ridge regression or principal component analysis (PCA).

  **Consequences of Multicollinearity:**

  Multicollinearity can have the following consequences:

  1. **Unstable Coefficients:** The coefficients of the correlated variables can become unstable. Small changes in the data can lead to substantial changes in the coefficients.
  2. **Reduced Interpretability:** It becomes challenging to interpret the effect of each variable on the dependent variable because their relationships are intertwined.
  3. **Difficulty in Feature Selection:** Multicollinearity can make it hard to decide which variables to include or exclude from the model, as the importance of each variable is unclear.

  **How to Detect and Deal with Multicollinearity:**

  To detect multicollinearity, you can use the following methods:

  1. **Correlation Matrix:** Calculate the correlation matrix between independent variables. Correlation coefficients close to +1 or -1 indicate strong multicollinearity.
  2. **Variance Inflation Factor (VIF):** VIF quantifies how much the variance of the estimated coefficients is increased due to multicollinearity. A high VIF (typically above 5 or 10) indicates multicollinearity.

  To address multicollinearity, you can consider the following strategies:

  1. **Remove Redundant Variables:** If two or more variables are highly correlated, you can choose to remove one of them, retaining the most informative variable.
  2. **Regularization:** Techniques like ridge regression or LASSO regression can help mitigate multicollinearity by adding a penalty term to the loss function.

- No auto-correlation

  Autocorrelation occurs when the residuals are not independent from each other. In other words when the value of y(x+1) is not independent from the value of y(x).

  While a scatterplot allows you to check for autocorrelations, you can test the linear regression model for autocorrelation with the Durbin-Watson test. Durbin-Watson’s d tests the null hypothesis that the residuals are not linearly auto-correlated. While d can assume values between 0 and 4, values around 2 indicate no autocorrelation. As a rule of thumb values of 1.5 < d < 2.5 show that there is no auto-correlation in the data. However, the Durbin-Watson test only analyses linear autocorrelation and only between direct neighbors, which are first order effects.

- [Homoscedasticity](https://www.statisticssolutions.com/free-resources/directory-of-statistical-analyses/homoscedasticity/)

  The scatter plot is good way to check whether the data are homoscedastic (meaning the residuals are equal across the regression line).  

  ##### Correlation vs Collinearity vs Multicollinearity: https://quantifyinghealth.com/correlation-collinearity-multicollinearity/

  ### Linear Regression

  #### Case Study:

  #### Exploratory Analysis:

  1. **Data Collection:** Gather data on house prices, including features like square footage, number of bedrooms, and location.
  2. **Data Preprocessing:** Clean the data, handle missing values, and encode categorical variables.
  3. **Exploratory Data Analysis (EDA):** Explore the distribution of house prices, identify correlations between features and the target variable.

  #### Feature Extraction:

  1. **Feature Engineering:** Create additional features such as a 'price per square foot' and 'proximity to amenities.'
  2. **Feature Selection:** Select relevant features using techniques like correlation analysis and feature importance.

  #### Model Selection:

  1. **Model Choice:** Choose linear regression as the modeling technique.
  2. **Train-Test Split:** Split the data into a training set and a testing set.

  #### Model Evaluation:

  1. **Model Training:** Train the linear regression model on the training data.
  2. **Model Testing:** Evaluate the model's performance using metrics like Mean Absolute Error (MAE), Mean Squared Error (MSE), and R-squared (R²).
  3. **Residual Analysis:** Analyze the residuals to check for patterns or heteroscedasticity.

  #### Tuning to Improve Accuracy:

  1. **Hyperparameter Tuning:** Experiment with different regression techniques (e.g., Lasso or Ridge) to regularize the model.
  2. **Feature Transformation:** Apply logarithmic transformations to the target variable to address non-normality.
  3. **Feature Scaling:** Standardize or normalize features to improve model convergence.

### Polynomial Regression

### Ridge and Lasso Regression

- L1 regularization (Lasso) encourages sparsity by adding the absolute values of the coefficients as a penalty term.![img](https://lh7-us.googleusercontent.com/1Dga_Me2ASc80lC158LOpPTWS38u_6mrFR_1osFqdS3p-iHqkIDNVEa7Whd0U4iRmO8KJM8jky9UN-Xu6YwOHncqyMquKJreAoLWSdNz3x0CvGnL9mkaUCaHm_RMDCRJBhtkrFeKuFpaBS6lJiE2JzE)
- L2 regularization (Ridge) adds the square of the coefficients as a penalty term to prevent overfitting.![img](https://lh7-us.googleusercontent.com/SFVD9ohGHwXY-un10xCHdmDXdsVVbG0wma6MzKJXpkclw3fvLEjWZiAHX0_C2wFufLKnfTP9_3pxo8E8wQN8TaO9pFVxL6npOmkBbqO8cfVjs4Gb5yh-pggPuX7TVkU3egwn6CeUaguthUBQRijUxhI)

Why L1 Results in Sparsity:

- L1 regularization encourages some coefficients to be exactly zero because the L1 norm has sharp corners at zero, making it more likely to eliminate irrelevant features.

Why Regularization Works:

- Regularization prevents overfitting by adding a penalty term to the loss function, which discourages the model from fitting noise in the data. It encourages simpler models.

## Classification

### Logistic Regression

Logistic Regression models the probability that a binary outcome belongs to a particular category. The logistic function (sigmoid) is commonly used for this purpose:

##### MLE 

*P*(*Y*=1∣*X*)=1+*e*−(*β*0+*β*1*X*)1

The likelihood function for a single data point (*x**i*,*y**i*) is the Bernoulli distribution:

*L*(*β*0,*β*1∣*x**i*,*y**i*)=*P*(*Y*=*y**i*∣*X*=*x**i*)=(1+*e*−(*β*0+*β*1*x**i*)1)*y**i*(1+*e*−(*β*0+*β*1*x**i*)*e*−(*β*0+*β*1*x**i*))1−*y**i*

The joint likelihood function for the entire dataset {(*x**i*,*y**i*)} is the product of the individual likelihoods:

*L*(*β*0,*β*1∣{(*x**i*,*y**i*)})=∏*i*=1*n*(1+*e*−(*β*0+*β*1*x**i*)1)*y**i*(1+*e*−(*β*0+*β*1*x**i*)*e*−(*β*0+*β*1*x**i*))1−*y**i*

The log-likelihood function:

log*L*=∑*i*=1*n**y**i*log(1+*e*−(*β*0+*β*1*x**i*)1)+(1−*y**i*)log(1+*e*−(*β*0+*β*1*x**i*)*e*−(*β*0+*β*1*x**i*))

#### Case Study:

##### Exploratory Analysis:

1. **Data Collection:** Collect a dataset of emails with labeled categories (spam or not spam) and features based on email content.
2. **Data Preprocessing:** Clean and preprocess the text data, including text tokenization and vectorization.
3. **Exploratory Data Analysis (EDA):** Explore the distribution of spam and non-spam emails, analyze word frequency, and investigate patterns in the data.

##### Feature Extraction:

1. **Feature Engineering:** Create features like the number of capitalized words, the presence of specific keywords, and email length.
2. **Feature Selection:** Select relevant features based on statistical tests and domain knowledge.

##### Model Selection:

1. **Model Choice:** Choose logistic regression as the classification technique.
2. **Train-Test Split:** Split the data into a training set and a testing set.

##### Model Evaluation:

1. **Model Training:** Train the logistic regression model on the training data.
2. **Model Testing:** Evaluate the model's performance using metrics like accuracy, precision, recall, and F1-score.
3. **ROC Curve:** Analyze the Receiver Operating Characteristic (ROC) curve and calculate the Area Under the Curve (AUC).

##### Tuning to Improve Accuracy:

1. **Hyperparameter Tuning:** Experiment with different regularization strengths (C parameter) for logistic regression.
2. **Feature Engineering:** Create additional text features and explore techniques like TF-IDF.
3. **Cross-Validation:** Use k-fold cross-validation to assess model stability and generalization.

### Decision Trees

1. **Gini Impurity:**
   - **Definition:** Gini impurity measures the frequency at which a randomly chosen element would be incorrectly labeled.
   - **Split Decision:** At each node, the algorithm chooses the feature that minimizes the weighted sum of Gini impurities in the child nodes after the split.
   - **Use Case:** Gini impurity is often used in classification problems. It is a measure of how often a randomly chosen element would be incorrectly classified.
2. **Entropy:**
   - **Definition:** Entropy is a measure of disorder or impurity in a set. In the context of decision trees, it measures the information gain—the reduction in entropy—achieved by a particular split.
   - **Split Decision:** At each node, the algorithm chooses the feature that maximizes the information gain, which is the difference between the entropy before and after the split.
   - **Use Case:** Entropy is commonly used in decision trees for classification. It is particularly popular in the context of the ID3 algorithm.

Both Gini impurity and entropy are used to evaluate the homogeneity of a node. The goal is to create splits that result in nodes with predominantly one class, making the decision tree more predictive. The decision tree algorithm compares the impurity or entropy of different splits and selects the one that improves homogeneity the most.

#### Key Points:

- Decision trees are versatile models used for both classification and regression.
- They create a tree-like structure of decisions based on features.

#### Example:

- Predicting customer churn (whether a telecom customer is likely to leave) based on factors like contract length, usage, and customer service interactions.

#### Case Study:

##### Exploratory Analysis:

1. **Data Collection:** Collect customer data, including contract details, call records, and customer service interactions.
2. **Data Preprocessing:** Clean and preprocess the data, handling missing values and encoding categorical variables.
3. **Exploratory Data Analysis (EDA):** Analyze customer churn rates, identify correlations, and understand customer behavior.

##### Feature Extraction:

1. **Feature Engineering:** Create new features like customer tenure, calculate total call duration, and derive customer usage patterns.
2. **Feature Selection:** Select relevant features that impact customer churn.

##### Model Selection:

1. **Model Choice:** Choose a decision tree algorithm and consider ensembles like Random Forest.
2. **Train-Test Split:** Split the data into training and testing sets.

##### Model Evaluation:

1. **Model Training:** Train the decision tree model on the training data.
2. **Model Testing:** Evaluate the model's performance using accuracy, precision, recall, and the Gini impurity for decision trees.
3. **Visualize the Tree:** Visualize the decision tree to understand the decision-making process.

##### Tuning to Improve Accuracy:

1. **Hyperparameter Tuning:** Experiment with tree depth, minimum samples per leaf, and other hyperparameters.
2. **Pruning:** Prune the tree to prevent overfitting.
3. **Feature Importance:** Analyze feature importance scores to understand which features drive customer churn.
4. **Ensemble Methods:** Consider using a Random Forest ensemble to improve model accuracy.

### Random Forest

Random Forest is an ensemble learning method used for both classification and regression tasks. It constructs multiple decision trees during the training phase. The model randomly selects samples from the dataset with replacement (bootstrapping) to build each tree. Also, at each node of a decision tree, it selects a random subset of features to make a split. The final prediction in a Random Forest is the mode (for classification) or the mean (for regression) of the individual tree predictions.

Random Forest is considered "random" due to its two primary sources of randomness in the model:

1. **Bootstrapping:** Each tree in the forest is trained on a random sample of the original dataset, known as bootstrapping or bagging. Random sampling involves selecting subsets of the training data, allowing for the construction of different trees.
2. **Feature Randomness:** During the construction of each tree in the forest, only a random subset of features is considered at each node. This feature selection method aims to reduce the likelihood of trees relying heavily on one specific feature, adding randomness to the decision-making process.

By introducing these random elements into the model's construction, Random Forest improves robustness and decreases overfitting. It achieves this by aggregating the outputs of individual decision trees, leading to more accurate and stable predictions on new, unseen data.

1. **How does Random Forest differ from a single decision tree?**

   - RF comprises multiple decision trees, providing predictions based on the ensemble, while a single decision tree may be prone to overfitting.

2. **What are the advantages of using Random Forest?**

   - Robustness to overfitting, handling missing values, handling high-dimensional datasets, and providing feature importance measures.

3. **Why is RF considered a "black-box" model?**

   - Due to its ensemble nature and multiple trees, interpreting every decision might be complex compared to a single decision tree.

4. **How does Random Forest prevent overfitting?**

   1. By introducing randomness in tree construction through bootstrapping and random feature selection.

5. **How does RF handle missing data?**

   1. By utilizing the mean or mode of available samples for the missing feature.

6. **What is Out-of-Bag (OOB) error in Random Forest?**

   - The error rate is calculated on the predictions from the data not used while training each individual tree

7. **how random forest different from gradient boosting** 

   Random Forest and Gradient Boosting are both ensemble learning techniques, but they differ in several key aspects:

   ##### 1. **Training Process:**

   - Random Forest:
     - **Training Method:** Random Forest builds multiple decision trees independently and merges them together.
     - **Tree Independence:** Each tree in a Random Forest is trained independently of the others.
   - Gradient Boosting:
     - **Training Method:** Gradient Boosting builds trees sequentially, and each new tree corrects the errors of the previous one.
     - **Sequential Learning:** Trees in a Gradient Boosting model are trained sequentially, with each tree trying to correct the errors of the combined ensemble.

   ##### 2. **Tree Construction:**

   - Random Forest:
     - **Tree Diversity:** Each tree in a Random Forest is built from a random subset of features at each split, providing diversity among the trees.
     - **Parallel Construction:** Trees are constructed independently of each other.
   - Gradient Boosting:
     - **Sequential Learning:** Trees in a Gradient Boosting model are built sequentially, with each tree focusing on the errors made by the ensemble so far.
     - **Shallow Trees:** Trees in a Gradient Boosting model are often shallow.

   ##### 3. **Handling of Errors:**

   - Random Forest:
     - **Error Handling:** Errors are controlled by averaging the predictions of multiple trees.
     - **No Adaptation:** Each tree is built independently, without consideration of the errors made by others.
   - Gradient Boosting:
     - **Error Handling:** Each new tree in a Gradient Boosting model is trained to correct the errors of the existing ensemble.
     - **Adaptive Learning:** Trees are built sequentially to adapt to the mistakes of the ensemble.

   ##### 4. **Predictions:**

   - Random Forest:
     - **Prediction:** Predictions are made by averaging (regression) or voting (classification) of individual tree predictions.
   - Gradient Boosting:
     - **Prediction:** Predictions are made by summing the predictions of all the trees, with each tree weighted by a factor.

   ##### 5. **Handling Overfitting:**

   - Random Forest:
     - **Robust to Overfitting:** Random Forests are generally robust to overfitting due to the diversity among the trees.
   - Gradient Boosting:
     - **Potential for Overfitting:** Gradient Boosting can be prone to overfitting, especially if the number of trees is high.

   ##### 6. **Hyperparameters:**

   - Random Forest:
     - **Key Hyperparameters:** Number of trees, maximum depth of trees, and the number of features considered at each split.
   - Gradient Boosting:
     - **Key Hyperparameters:** Learning rate, number of trees, and tree-specific parameters.

   **Case Study: Predicting Customer Churn in a Telecommunications Company**

   **Problem Statement:** A telecommunications company aims to reduce customer churn by predicting which customers are likely to terminate their contracts. They possess a vast dataset with various customer attributes.

   **Application of RF:**

   1. **Feature Extraction/Selection:**
      - The dataset contains numerous features like customer demographics, service usage, and billing information. RF is used for feature selection to determine which factors are most influential in predicting churn. Through analyzing feature importance, it's found that service usage patterns and customer support interaction are significant predictors.
   2. **Model Selection - Why RF is Better:**
      - Compared to other machine learning models like Logistic Regression or Decision Trees, RF is chosen because of its capability to handle large, high-dimensional datasets and effectively capture non-linear relationships between variables. This is crucial in a telecommunications dataset with diverse features.
   3. **Evaluation and Tuning:**
      - Initially, the dataset is split into training and testing sets. RF models are trained with different hyperparameters such as the number of trees, tree depth, and minimum samples per split. The models are evaluated using various metrics such as accuracy, precision, recall, and F1-score. Tuning is conducted by using cross-validation to identify the best-performing set of hyperparameters.
   4. **Improving Model Performance:**
      - Further model improvement is achieved by applying techniques like feature engineering based on RF feature importance and removing less informative attributes. By using the gained insights, the company develops targeted customer retention strategies based on predicted churn probability.

   **Outcome:** The RF model accurately predicts potential churners by considering various features, thereby enabling the telecommunications company to take proactive measures to retain at-risk customers. By implementing these strategies, the company observes a significant reduction in customer churn, leading to increased customer satisfaction and enhanced profitability.

   This case study demonstrates how RF can be applied for feature extraction, selection, model selection, evaluation, and tuning, specifically to address the critical issue of customer churn prediction in a telecommunications company.

   

### Support Vector Machines

The objective of the support vector machine algorithm is to find a hyperplane in an N-dimensional space(N — the number of features) that distinctly classifies the data points.![Screen Shot 2023-11-20 at 1.49.44 PM](/Users/moya_shoreline/Library/Application Support/typora-user-images/Screen Shot 2023-11-20 at 1.49.44 PM.png)

Ref: https://towardsdatascience.com/support-vector-machine-introduction-to-machine-learning-algorithms-934a444fca47

### Basic Concept:

1. **Hyperplane:**
   - In a two-dimensional space, a hyperplane is a simple line. In a three-dimensional space, it's a plane. In higher dimensions, it's a hyperplane. For a binary classification problem (two classes), SVM aims to find the hyperplane that best separates the data points of one class from another.
2. **Support Vectors:**
   - Support vectors are the data points that are closest to the hyperplane and have the most influence on its position. These are the critical points for determining the decision boundary.
3. **Margin:**
   - The margin is the distance between the hyperplane and the nearest data point from either class. SVM aims to maximize this margin, leading to a more robust and generalized model.

### Steps in SVM:

1. **Data Preparation:**
   - SVM works well for both linear and non-linear data. The data needs to be preprocessed and feature-scaled, as SVM is sensitive to the scale of features.
2. **Choosing the Kernel Function:**
   - SVM can use different kernel functions to map the input data into higher-dimensional spaces. Common kernels include linear, polynomial, radial basis function (RBF), and sigmoid. The choice of the kernel depends on the nature of the data.
3. **Training the Model:**
   - SVM finds the optimal hyperplane by adjusting the weights and biases in the kernel function. The training process involves finding the support vectors and optimizing the hyperplane to maximize the margin.
4. **Optimizing the Hyperplane:**
   - The optimization process involves minimizing a cost function while satisfying the constraint that the data points are correctly classified and the margin is maximized. The cost function includes a regularization term to prevent overfitting.
5. **Prediction:**
   - Once the SVM model is trained, it can be used to make predictions on new, unseen data. The decision function of the SVM classifies the data point based on its position relative to the hyperplane.

### k-Nearest Neighbors

### Naive Bayes

#### Key Points:

- Naive Bayes is a probabilistic classification algorithm based on Bayes' theorem.
- It's widely used for text classification and spam detection.

#### Example:

- Classifying news articles into categories such as sports, politics, and entertainment based on the words used in the articles.

#### Case Study:

##### Exploratory Analysis:

1. **Data Collection:** Gather a dataset of news articles with labeled categories.
2. **Data Preprocessing:** Clean and preprocess the text data by removing stopwords, punctuation, and stemming/lemmatizing words.
3. **Exploratory Data Analysis (EDA):** Perform EDA to understand the distribution of categories, word frequency, and any patterns in the text data.

##### Feature Extraction:

1. **Feature Engineering:** Transform text data into numerical features using techniques like TF-IDF (Term Frequency-Inverse Document Frequency) or Bag of Words.
2. **Feature Selection:** Select relevant features that contribute to classification accuracy.

##### Model Selection:

1. **Model Choice:** Choose the Naive Bayes algorithm, specifically Multinomial Naive Bayes or Bernoulli Naive Bayes, depending on the nature of the text data.
2. **Train-Test Split:** Split the data into training and testing sets for model evaluation.

##### Model Evaluation:

1. **Model Training:** Train the Naive Bayes model on the training data.
2. **Model Testing:** Evaluate the model's performance on the testing data using metrics like accuracy, precision, recall, and F1-score.
3. **Confusion Matrix:** Analyze the confusion matrix to understand where the model makes errors.

##### Tuning to Improve Accuracy:

1. **Hyperparameter Tuning:** Experiment with different hyperparameters like Laplace smoothing (additive smoothing) for Naive Bayes.
2. **Cross-Validation:** Perform k-fold cross-validation to assess the model's generalization performance.
3. **Grid Search:** Use grid search to systematically explore hyperparameter combinations.
4. **Ensemble Methods:** Explore ensemble methods like AdaBoost to improve model accuracy further.

## Clustering

### K-Means

K-means clustering is a popular unsupervised machine learning algorithm used for partitioning a dataset into K distinct, non-overlapping subsets (clusters). Each data point belongs to the cluster with the nearest mean, serving as a prototype of the cluster.

**Steps:**

1. **Initialization:**
   - Choose the number of clusters (K) you want to create.
   - Randomly initialize the K cluster centroids (points representing the center of each cluster).
2. **Assignment:**
   - Assign each data point to the nearest cluster centroid. Use a distance metric like Euclidean distance.
3. **Update Centroids:**
   - Recalculate the centroids of the clusters based on the data points assigned to them.
   - The new centroid is the mean of all data points in the cluster.
4. **Repeat:**
   - Repeat steps 2 and 3 until convergence is reached. Convergence occurs when the assignment of data points to clusters and the centroids no longer change significantly.

**Algorithm Iteration:**

Repeat the following until convergence:

- **Assignment Step:** Assign each data point to the nearest centroid.

  *c*(*i*)=argmin*k*∥*x*(*i*)−*μ**k*∥2

- **Update Step:** Recalculate the centroids based on the assigned data points. 

  *μ**k*=∣*C**k*∣1∑*i*∈*C**k**x*(*i*) where C**k* is the set of data points assigned to cluster *k*.

**Termination:**

The algorithm terminates when one of the following conditions is met:

- Centroids no longer change significantly.
- A maximum number of iterations is reached.
- Other convergence criteria are satisfied.

**Choosing K:**

- The choice of K is crucial and may require domain knowledge or exploration through methods like the elbow method or silhouette analysis.

**Pros:**

- Simple and computationally efficient.
- Works well with large datasets.

**Cons:**

- Sensitive to initial centroid placement.
- Assumes spherical clusters with equal variance.

It's important to note that K-means clustering is not deterministic, and different initializations may result in different final cluster assignments. Repeated runs with different initializations can help mitigate this.

### Hierarchical Clustering

Hierarchical clustering is another unsupervised machine learning algorithm used for grouping similar data points into clusters. Unlike K-means, hierarchical clustering doesn't require the pre-specification of the number of clusters (K). Instead, it creates a hierarchy of clusters, often represented as a tree-like structure (dendrogram).

**Steps:**

1. **Initialization:**
   - Treat each data point as a singleton cluster.
   - Compute the distance (similarity) matrix between all pairs of clusters.
2. **Merge:**
   - Find the closest (most similar) pair of clusters based on the distance matrix.
   - Merge these two clusters into a new cluster.
   - Update the distance matrix.
3. **Repeat:**
   - Repeat steps 2 until only a single cluster remains, or a desired number of clusters is obtained.

**Algorithm Iteration:**

Repeat the following until the desired number of clusters is reached:

- **Compute Similarities:** Compute the similarity (or dissimilarity) matrix between all clusters.
- **Merge Closest Clusters:** Identify the two clusters with the smallest dissimilarity and merge them into a new cluster.

**Termination:**

The algorithm terminates when:

- A specified number of clusters is obtained.
- A linkage criterion (e.g., maximum or average linkage) is met.
- The dendrogram is cut at a specific height.

**Dendrogram:**

- The output of hierarchical clustering is often visualized as a dendrogram, which is a tree-like diagram showing the arrangement of clusters and their relationships.

**Linkage Methods:**

- Different linkage methods determine how the distance between clusters is calculated. Common methods include:
  - **Single Linkage:** Minimum distance between points in different clusters.
  - **Complete Linkage:** Maximum distance between points in different clusters.
  - **Average Linkage:** Average distance between points in different clusters.

**Pros:**

- No need to specify the number of clusters in advance.
- Provides a hierarchical structure.

**Cons:**

- Computationally more expensive, especially for large datasets.

**Choosing the Number of Clusters:**

- Decide the number of clusters by visually inspecting the dendrogram or using other criteria.

Hierarchical clustering is versatile and applicable to various domains, but its scalability can be a limitation for very large datasets.

## Dimensionality Reduction

### PCA

Principal Component Analysis (PCA) is a dimensionality reduction technique commonly used in data analysis and machine learning. Its primary goal is to transform a dataset into a new coordinate system, called the principal component space, where the data's variability is maximized along the new axes (principal components). Here's a step-by-step explanation of how PCA works:

1. **Standardization:**
   - PCA works best when the features are on a similar scale. Therefore, the first step is often to standardize the data by subtracting the mean and dividing by the standard deviation for each feature.
2. **Compute Covariance Matrix:**
   - The next step is to compute the covariance matrix of the standardized data. The covariance matrix provides information about the relationships between different features.
3. **Compute Eigenvectors and Eigenvalues:**
   - The eigenvectors and eigenvalues of the covariance matrix are then calculated. Eigenvectors represent the directions of maximum variance (principal components), and eigenvalues indicate the magnitude of the variance in those directions.
4. **Sort Eigenvectors by Eigenvalues:**
   - The eigenvectors are arranged in descending order based on their corresponding eigenvalues. The higher the eigenvalue, the more variance is captured along that eigenvector.
5. **Select Principal Components:**
   - Decide on the number of principal components to retain. This decision is often based on the explained variance. If, for example, you want to retain 95% of the variance, you select the first k eigenvectors that cumulatively explain at least 95% of the total variance.
6. **Create Projection Matrix:**
   - Form a projection matrix by stacking the selected eigenvectors as columns. This matrix is used to transform the original data into the new principal component space.
7. **Transform Data:**
   - Multiply the standardized data by the projection matrix to obtain the data in the principal component space. The transformed data has a reduced number of dimensions.

**Key Concepts:**

- **Explained Variance:**
  - Each principal component explains a certain amount of variance in the original data. The cumulative sum of the eigenvalues provides a measure of how much variance is retained by the selected principal components.
- **Reduced Dimensionality:**
  - PCA reduces the dimensionality of the data by retaining the most important features (principal components) while discarding less important ones. This is useful for visualization, noise reduction, and improving the performance of machine learning models.
- **Orthogonality:**
  - The principal components are orthogonal, meaning they are uncorrelated. This simplifies the interpretation of the transformed data.

### T-SNE

## Ensemble Learning

##### Strong and week learners

1. **Weak Learners:**
   - **Definition:** A weak learner is a model or algorithm that performs slightly better than random chance. It may have limited predictive accuracy on its own, but it is still able to learn something from the data.
   - **Characteristics:** Weak learners are often simple and have limitations, but they contribute valuable information when combined in an ensemble. Examples of weak learners include shallow decision trees, linear models, or models that perform just above random chance.
2. **Strong Learners:**
   - **Definition:** A strong learner is a model or algorithm that achieves high predictive accuracy and performs well on its own. It is capable of capturing complex patterns and relationships in the data.
   - **Characteristics:** Strong learners are typically more complex and have a higher capacity to model the underlying structure of the data. Examples of strong learners include deep neural networks, complex ensemble models (e.g., Random Forest, Gradient Boosting), or sophisticated machine learning algorithms.

The concepts of weak learners and strong learners are often associated with ensemble learning, a technique where multiple models are combined to improve overall performance. The idea is that by combining weak learners, their individual weaknesses or errors can be mitigated, leading to a stronger, more robust model.

### Boosting

Boosting is another ensemble learning technique, but unlike bagging, it focuses on building a sequence of weak learners to create a strong learner. In boosting, each model in the sequence corrects the errors of its predecessor. The key concept in boosting is assigning weights to data points, emphasizing the misclassified points more in each iteration.

The key steps of boosting (e.g., AdaBoost algorithm) are:

1. **Assign Initial Weights:** Assign equal weights to all data points in the training set.
2. **Model Training:** Train a base model on the training set, giving more weight to the misclassified points.
3. **Compute Error:** Compute the error of the model on the training set.
4. **Update Weights:** Increase the weights of misclassified points, making them more influential in the next iteration.
5. **Repeat:** Repeat steps 2-4 for a specified number of iterations or until a stopping criterion is met.
6. **Aggregation:** Combine the predictions of all models with weighted voting.

Boosting aims to improve the performance of the model by focusing on the weaknesses of the current ensemble. It is particularly effective when used with weak learners.

##### why boosting more prone to overfit

1. **Complexity of Weak Learners:**
   - Boosting often uses weak learners (e.g., shallow decision trees) as base models. While these weak learners are individually simple, when combined sequentially, they can create a complex model capable of fitting the training data very closely.
   - As boosting continues, it focuses on reducing errors on the training set, potentially leading to a model that fits the noise in the data rather than the underlying patterns. This can result in overfitting.
2. **Adaptation to Residual Errors:**
   - Boosting builds models sequentially, each one focusing on the errors (residuals) made by the previous models. While this adaptability is a strength in reducing bias and improving accuracy, it also makes the ensemble more prone to fitting noise in the training data.
   - The boosting process may become too tailored to the peculiarities of the training set, leading to poor generalization on new, unseen data.
3. **Sensitivity to Noisy Data:**
   - Boosting is sensitive to noisy data or outliers. If there are outliers or noise in the training data, boosting algorithms may give these instances more attention in an attempt to correct their errors, even if they are not representative of the underlying patterns.
4. **Overemphasis on Difficult Examples:**
   - Boosting assigns higher weights to misclassified examples at each iteration. While this helps the algorithm focus on challenging instances, it can lead to overfitting if the difficult examples are outliers or noise.
5. **Overemphasizing Rare Classes:**
   - In imbalanced datasets, boosting may overemphasize the minority class, leading to overfitting on rare instances and potentially neglecting the majority class.

### Mitigating Overfitting in Boosting:

1. **Regularization Parameters:**
   - Many boosting algorithms have hyperparameters that control model complexity. These include parameters like the learning rate, tree depth, and the number of iterations. Tuning these parameters can help control overfitting.
2. **Early Stopping:**
   - Monitoring the performance on a validation set and stopping the boosting process when performance ceases to improve can prevent overfitting.
3. **Limiting Tree Depth:**
   - Constraining the depth of the weak learners (e.g., decision trees) can prevent them from becoming overly complex.
4. **Subsampling:**
   - Using subsampling techniques (e.g., Stochastic Gradient Boosting) can introduce randomness and reduce overfitting.

### Bagging

Bagging is an ensemble learning technique that aims to improve the stability and accuracy of a machine learning model by training multiple instances of the model on different subsets of the training data. The subsets are created using bootstrap sampling, which involves randomly drawing samples with replacement from the original training dataset. Each model in the ensemble is trained on a different bootstrap sample.

The key steps of bagging are:

1. **Bootstrap Sampling:** Randomly draw samples with replacement from the original training dataset to create multiple bootstrap samples.
2. **Model Training:** Train a base model (e.g., a decision tree) on each bootstrap sample.
3. **Aggregation:** Combine the predictions of each model to make a final prediction. For regression, this might involve averaging the predictions, and for classification, it might involve a majority vote.

Bagging helps reduce overfitting and variance, particularly for models that are sensitive to the specific training data. The diversity introduced by training on different subsets helps improve the overall performance of the ensemble.

### Stacking

Stacking, also known as stacked generalization, is an ensemble learning technique that combines multiple models to improve predictive performance. The basic idea is to train several base models and then use a meta-model to combine their predictions. The key steps in the stacking process are as follows:

### Steps in Stacking:

1. **Data Splitting:**
   - Divide the dataset into training and testing sets. The training set is used to train the base models, and the testing set is used to evaluate their performance.
2. **Base Model Training:**
   - Train multiple diverse base models on the training set. Diversity is important to ensure that the models capture different aspects of the data.
3. **Base Model Prediction:**
   - Use the trained base models to make predictions on the testing set. These predictions become the input features for the meta-model.
4. **Meta-Model Training:**
   - Train a meta-model (also called a blender or a second-level model) using the predictions from the base models as features and the true labels as the target variable. This meta-model learns how to combine the predictions of the base models.
5. **Stacking:**
   - For a new, unseen dataset, make predictions using the base models and then use those predictions as input features for the meta-model. The meta-model provides the final prediction.
