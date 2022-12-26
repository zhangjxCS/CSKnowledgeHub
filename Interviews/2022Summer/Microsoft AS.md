## 一面

中国小姐姐

1.自我介绍 实习项目

2.顺着实习项目接着问决策树模型

决策树怎么找最优切分变量和最优切分点的

回归树、分类树的区别，loss function的形式

GBDT的原理，能否拟合残差的2次项 3次项？为什么

3.类似lc542的题，找connected component数量

## 二面

印度大哥

1.自我介绍 实习项目

2.问了实习项目中的Granger Causality Test，为什么用这个test，这个test的前提条件是什么

3.问了实习项目中的Time series model，ARIMA的p d q怎么确定

4.问了实习项目中怎么做feature engineering的，最终为什么用xgboost

5.recommender system design 你怎么设计一个text ads的推荐系统，怎么构建特征，怎么选模型；问了推荐系统召回排序的细节，召回过程数据量太多怎么办；排序过程中目前的state of the art模型是什么，特征应该怎么构建

6.两道统计问题

扔骰子一次 数字乘1000 扔1000次骰子每次一块钱 你怎么选择

先后扔两次骰子，最先扔到6的人赢，哪个人占优势，概率是多少

## 三面

美国大哥

1.自我介绍 实习项目

2.实习项目中的时间序列部分，为什么选这个模型，模型结果是啥，模型特征怎么构建

3.文本特征的挖掘：embedding的细节，word2vec，cbow，你的embedding vector自己fine tune的吗

4.给一个视频推荐系统的三张表，user video action，怎么做推荐系统

特征构建：离散变量做one hot encoding, target encoding，详细问了target encoding的细节，为什么构建这个特征；问了如果时间序列很长，用户兴趣有变化怎么办；用滑动窗口构建特征，怎么选择滑动窗口大小，怎么tune hyperparameters

任务：是回归任务，还是分类任务，分类任务的话，negative sampling怎么操作

## 四面

中国人hiring manager

上来没有自我介绍，直接给了一个case

给一些用户review的表，你怎么构建模型预测他们对产品的rating 从1到5

什么任务？回归还是分类

建立rnn模型，输入输出是什么，loss function是什么

用gru lstm相比rnn的优势是什么，lstm几个gate的细节

如果不能用rnn deep learning 你怎么做这个case？

tf-idf的特征怎么构建，特征数量多于数据数量怎么办

特征数量太多 加regularization怎么操作，如果sparse用哪种regularization