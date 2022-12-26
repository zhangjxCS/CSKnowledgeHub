## Coding

A variation on binary search

graph traversal problem (bfs/dfs) (shortest path, etc.) (Dijkstra)

matrix addition and multiplication for sparse matrices

Backtracking

String manipulation (recursion)

Maximum increasing subsequence in a binary tree

## ML

### Fundamental

nonconvex or convex, difficult to optimize

vanishing/exploding gradient (ReLU or Sigmoid) (closer to the beginning or end)

### System Design

#### Recommender System

Clarify questions

- Engagement: new user retention, session length, number of sessions 
- Revenue: ctr, ctc, etc

Product requirements

- real-time or batch?
- apply to all user segments?

- how many items? How many users?
- peak number of requests

High Level Design

- candidate generation stage to quickly produce a medium-sized list of items to consider
- ranking model to select the top x items to return

Data sources

- the user-side features: age, gender, nationality, language, pre-selected favorites
- Ads/content side features: length, tag, transcript, image
- interactions between the user and content: view, like, comment
- the search query: text query, audio query, image query

feature representation

- numeric values: scaling, outliers, binning, quantization
- categorical values: one hot encoding, embedding
- text: bert and word2vec
- Complex:representing each post as its own embedding, combine the embedding

Feature selection:

- forward regression?
- regularization

candidate sources

heuristic to generate an initial list of candidates

- items close to the user in an embedded space (collaborative filtering)
- Items similar to items the user has interacted with (item based filtering)
- global popular items
- items that friends of friends like
- items that match some type of search criteria

ranking model

- a binary classifier that predicts whether the user will interact with the given post
- a multi classifier with predictions for comment, like, hide
- a regressor predicts the number of interactions a user will make with the post

offline training and evaluation

- Data for training: data imbalances
- data for evaluation: k-fold, hold out, metrics to compare models

models selection

- logistic regression/linear regression
- gradient boosted trees
- deep neural nets

online evaluation

- A/B testing: metrics you'd measure and statistical tests
- Business KPI: revenue, retention, engagement

#### Spam Filtering System

clarify question:

- how much training data? labeled or unlabeled?
- offline or online?
- peak number of requests?

Data sources:

- training data with labels, spam non-spam (text data)
- sender ip address, sender-related information

Data Preprocessing:

- lower case
- remove punctuation, stop words, white spaces, hyperlink
- Word stemming, lemmatization

Feature Extraction

- word counter
- tf idf vectorizer
- word embedding

Modeling:

- logistic regression
- svm
- Xgboost
- neural nets

Metrics:

- accuracy
- precision, recall
- Auc score

## Company Related

questions

- what's the monetization org doing
- what's the future growth of the org? what's the priority
- who's the stakeholder

- what's the normal work for a machine learning engineer on Pinterest?
- what's the tech stack for a machine learning engineer
- what's the career path for a machine learning engineer
- how to cooperate with others? PM DS?

### BQ

- Why Pinterest

Also a pinner, use Pinterest to find ideas for photography. I want to join the company to make it better

Good working environment, got a chance to meet with a few pinterest employees, want to work with pinterest tanlents

- Why Ads

Previously, I have some experience with ads org, I know some background knowledge about ads, want to continue on this area

Ads is the core team of the company, has many technical challenges such as large scale machine learning and big data, I think I could learn a lot and grow fast in the ads team

- Why you?

experience with ads and recommender system, so could ramp up quickly and help others

I have strong interest in the ads org, although a new grad, but learn fast

- Short term and long term goals

Contribute to the team and company

Grow fast, learn about Pinterest's business and technologies

Grow technical and communication skills

### Ads Marketplace

- What's the ads marketplace? Is it a pricing team for ads?

- How many time spent on infra, how many time on coding, how many time on meeting?
- Who's the decision maker, who decide to do the project?
- What's the future work for the team?

### Ads Targeting

- Interest Targeting: match relevant users for advertisers to target; keyword targeting product;
- Audience Targeting: ads targeting, ads retrieval, and ads ranking; real-time data; demographic/age/geo-targeting

find relevant users for advertisers as soon as ads are created

import act-alike model performance

reduce data processing and serving latency

Question:

- What's the data like? What model are you using? NLP technology for what? How to use the sequence model?
- What's the architecture for ads targeting? What tech stacks are you using? Cloud? Big data?
- How many time spent on infra, how many time on coding, how many time on meeting?
- Who's the decision maker, who decide to do the project?
- What's the future work for the team?

### Senior

- Share about the current project
- Share about the tech stack you used
- Share about the org and team, how many people, how to cooperate with others
- Share about the daily work, agile, oncall
- What's the future project, and future plan for the team
- How about the career path, how long to get promoted
- Tech layoff does it affect Pinterest and the team

## 面试

### 一面

coding面试

简单介绍亚麻实习项目

做题 二维矩阵 0 1，1代表障碍，给定起点终点，bfs计算最短距离

follow up: 如果有一把钥匙 一个门 拿到钥匙可以开门 怎么计算最短距离？

在visited数组和queue中维护一个变量state 代表是否拿到钥匙

follow up: 怎样返回路径 用一个res数组记录路径

### 二面

coding面试

设计一个calendar 类 记录 start end description attendee

add 给calendar加入记录，并且按照start排序，用插入排序

print 打印所有记录

计算两个事件是否conflict 如果attendee有重复并且start end time又重复则conflict

list有序，每次遍历之后的记录即可

### 三面

ml 面

设计一个spam filtering system

data：

用户report data 能否用？不能直接用

需要找labeler标记数据

在所有数据中采样作为负样本，labeler标记的数据作为正样本

feature

content相关：text做embedding处理（bert提取特征），image做embedding处理（resnet提取特征）还有一些tags

user相关：用户是否标记unsafe 用户行为 点赞 评论等

publisher相关：publisher历史记录等

model

logistic regression：简单二分类问题，有什么优缺点

tree based model：缺点是feature维度高，embedding feature无法用tree base模型有效处理

neural nets模型：dnn 激活函数前面的用relu 最后一层用sigmoid

evaluation

precision recall的区别，本题用recall

accuracy auc score

线上实验：用户report 率作为指标

### 四面

ml面试

设计一个404 页面分类器

data

数据找labeler进行标记：标记的比较准

或者通过用户反馈获得数据：准确性低

feature

image数据：ocr embedding

text数据：embedding，如果embedding维度过高怎么办，用max pooling或者self attention

tabular数据：用户停留时长等

evaluation

roc auc的含义

precision recall

accuracy

如果数据集分布不同 90% 10% 99% 1%这些指标会变吗

根据指标的公式进行分析
