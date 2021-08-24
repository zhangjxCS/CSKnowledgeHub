# 数据科学学习笔记

## 关于本仓库

21世纪是信息的时代，随着数据量的不断增大，数据类型的多样化，如何处理数据，从数据中获取价值成为了一个科学问题。如下图所示，数据科学是计算机科学、数学统计和领域知识的结合。计算机科学研究如何提取、转化、加载数据，搭建高效的数据Pipelines；数学统计研究如何分析数据，找出数据之间的相关性，构建合理的机器学习模型；领域知识为模型构建提供指导。如果想成为一个优秀的数据科学家，这三类知识缺一不可。

<img src="https://www.datocms-assets.com/14946/1596797558-da-vs-ds-diagram.png" alt="Data science vs. data analytics" style="zoom:67%;" />

本仓库是自己在学习数据科学过程中的笔记汇总，主要分为数学、计算机、机器学习、专题研究四个部分，正是数据科学所对应的三个部分。数学主要是微积分、线性代数、概率论与统计的复习笔记，为后面机器学习打基础；计算机主要是数据结构算法、操作系统、计算机网络和数据库四门基础课程的学习笔记；机器学习分统计学习、深度学习和强化学习的学习笔记；专题研究则是针对不同的业务场景常用算法的学习，比如说文本分类、推荐系统、计算机视觉等等。

本仓库正确的阅读姿势是

1.本地阅读

```shell
#For Mac and Linux user, open terminal
cd ~
git clone git@github.com:zhangjx831/Data-Science-Notes.git
#For Windows user, download XSHELL
cd ~
git clone git@github.com:zhangjx831/Data-Science-Notes.git
```

下载免费的markdown文档阅读器Typora，通过Typora打开本仓库，阅读体验非常高

2.网上阅读

由于Github不支持latex公式渲染，所以在Github上看到的latex公式全部显示为原始格式，目前chrome和火狐浏览器都有针对latex公式的渲染插件，下载后阅读本仓库体验更佳

## 目录

### 计算机

- 编程语言：[C/C++](https://github.com/zhangjx831/Data-Science-Notes/blob/master/%E8%AE%A1%E7%AE%97%E6%9C%BA/%E7%BC%96%E7%A8%8B%E8%AF%AD%E8%A8%80/C%E4%B8%8EC%2B%2B.md), [Python](https://github.com/zhangjx831/Data-Science-Notes/blob/master/%E8%AE%A1%E7%AE%97%E6%9C%BA/%E7%BC%96%E7%A8%8B%E8%AF%AD%E8%A8%80/Python.md)

- 操作系统：[基础概念](https://github.com/zhangjx831/Data-Science-Notes/blob/master/%E8%AE%A1%E7%AE%97%E6%9C%BA/%E6%93%8D%E4%BD%9C%E7%B3%BB%E7%BB%9F/intro.md)
- 计算机网络：[基础概念](https://github.com/zhangjx831/Data-Science-Notes/blob/master/%E8%AE%A1%E7%AE%97%E6%9C%BA/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/%E5%9F%BA%E7%A1%80.md)

- 数据库与大数据：[基础概念](https://github.com/zhangjx831/Notes/blob/master/%E8%AE%A1%E7%AE%97%E6%9C%BA/%E6%95%B0%E6%8D%AE%E5%BA%93/%E5%9F%BA%E7%A1%80%E6%A6%82%E5%BF%B5.md), [数据库设计](https://github.com/zhangjx831/Notes/blob/master/%E8%AE%A1%E7%AE%97%E6%9C%BA/%E6%95%B0%E6%8D%AE%E5%BA%93/%E6%95%B0%E6%8D%AE%E5%BA%93%E8%AE%BE%E8%AE%A1.md), [SQL语言](https://github.com/zhangjx831/Notes/blob/master/%E8%AE%A1%E7%AE%97%E6%9C%BA/%E6%95%B0%E6%8D%AE%E5%BA%93/SQL%E8%AF%AD%E8%A8%80.md), [Spark框架](https://github.com/zhangjx831/Data-Science-Notes/blob/master/%E8%AE%A1%E7%AE%97%E6%9C%BA/%E6%95%B0%E6%8D%AE%E5%BA%93%E4%B8%8E%E5%A4%A7%E6%95%B0%E6%8D%AE/Spark.md)

- 效率工具：[Shell脚本](https://github.com/zhangjx831/Notes/blob/master/%E8%AE%A1%E7%AE%97%E6%9C%BA/%E6%95%88%E7%8E%87%E5%B7%A5%E5%85%B7/Shell%E8%84%9A%E6%9C%AC.md), [Vim编辑器](https://github.com/zhangjx831/Notes/blob/master/%E8%AE%A1%E7%AE%97%E6%9C%BA/%E6%95%88%E7%8E%87%E5%B7%A5%E5%85%B7/Vim%E7%BC%96%E8%BE%91%E5%99%A8.md), [正则表达式](https://github.com/zhangjx831/Data-Science-Notes/blob/master/%E8%AE%A1%E7%AE%97%E6%9C%BA/%E6%95%88%E7%8E%87%E5%B7%A5%E5%85%B7/%E6%AD%A3%E5%88%99%E8%A1%A8%E8%BE%BE%E5%BC%8F.md)

### 数学

- 概率论与统计：[概率论](https://github.com/zhangjx831/Data-Science-Notes/blob/master/%E6%95%B0%E5%AD%A6/%E6%A6%82%E7%8E%87%E8%AE%BA%E4%B8%8E%E7%BB%9F%E8%AE%A1/Probs.md), [统计](https://github.com/zhangjx831/Data-Science-Notes/blob/master/%E6%95%B0%E5%AD%A6/%E6%A6%82%E7%8E%87%E8%AE%BA%E4%B8%8E%E7%BB%9F%E8%AE%A1/Stats.md)

### 机器学习

- 统计学习：[基础概念](https://github.com/zhangjx831/Notes/blob/master/%E6%9C%BA%E5%99%A8%E5%AD%A6%E4%B9%A0/%E7%BB%9F%E8%AE%A1%E5%AD%A6%E4%B9%A0/%E5%9F%BA%E7%A1%80%E6%A6%82%E5%BF%B5.md), [监督学习](https://github.com/zhangjx831/Data-Science-Notes/blob/master/%E6%9C%BA%E5%99%A8%E5%AD%A6%E4%B9%A0/%E7%BB%9F%E8%AE%A1%E5%AD%A6%E4%B9%A0/%E7%9B%91%E7%9D%A3%E5%AD%A6%E4%B9%A0.md), [无监督学习](https://github.com/zhangjx831/Data-Science-Notes/blob/master/%E6%9C%BA%E5%99%A8%E5%AD%A6%E4%B9%A0/%E7%BB%9F%E8%AE%A1%E5%AD%A6%E4%B9%A0/%E6%97%A0%E7%9B%91%E7%9D%A3%E5%AD%A6%E4%B9%A0.md), [集成学习](https://github.com/zhangjx831/Data-Science-Notes/blob/master/%E6%9C%BA%E5%99%A8%E5%AD%A6%E4%B9%A0/%E7%BB%9F%E8%AE%A1%E5%AD%A6%E4%B9%A0/%E9%9B%86%E6%88%90%E5%AD%A6%E4%B9%A0.md), [手推公式](https://github.com/zhangjx831/Data-Science-Notes/blob/master/%E6%9C%BA%E5%99%A8%E5%AD%A6%E4%B9%A0/%E7%BB%9F%E8%AE%A1%E5%AD%A6%E4%B9%A0/%E6%89%8B%E6%8E%A8%E5%85%AC%E5%BC%8F.md), [时间序列](https://github.com/zhangjx831/Data-Science-Notes/blob/master/%E6%9C%BA%E5%99%A8%E5%AD%A6%E4%B9%A0/%E7%BB%9F%E8%AE%A1%E5%AD%A6%E4%B9%A0/%E6%97%B6%E9%97%B4%E5%BA%8F%E5%88%97.md)

- 深度学习：[神经网络基础](https://github.com/zhangjx831/Data-Science-Notes/blob/master/%E6%9C%BA%E5%99%A8%E5%AD%A6%E4%B9%A0/%E6%B7%B1%E5%BA%A6%E5%AD%A6%E4%B9%A0/DNN.md), [卷积神经网络](https://github.com/zhangjx831/Data-Science-Notes/blob/master/%E6%9C%BA%E5%99%A8%E5%AD%A6%E4%B9%A0/%E6%B7%B1%E5%BA%A6%E5%AD%A6%E4%B9%A0/CNN.md), [循环神经网络](https://github.com/zhangjx831/Data-Science-Notes/blob/master/%E6%9C%BA%E5%99%A8%E5%AD%A6%E4%B9%A0/%E6%B7%B1%E5%BA%A6%E5%AD%A6%E4%B9%A0/RNN.md)

- 强化学习：[蒙特卡洛方法](https://github.com/zhangjx831/Notes/blob/master/%E6%9C%BA%E5%99%A8%E5%AD%A6%E4%B9%A0/%E5%BC%BA%E5%8C%96%E5%AD%A6%E4%B9%A0/%E8%92%99%E7%89%B9%E5%8D%A1%E6%B4%9B%E6%96%B9%E6%B3%95.md)

- 专题研究：[推荐系统](https://github.com/zhangjx831/Data-Science-Notes/blob/master/%E4%B8%93%E9%A2%98%E5%AD%A6%E4%B9%A0/%E6%8E%A8%E8%8D%90%E7%B3%BB%E7%BB%9F.md), [因果推断](https://github.com/zhangjx831/Data-Science-Notes/blob/master/%E4%B8%93%E9%A2%98%E5%AD%A6%E4%B9%A0/%E5%9B%A0%E6%9E%9C%E6%A3%80%E9%AA%8C.md), [数据处理](https://github.com/zhangjx831/Data-Science-Notes/blob/master/%E4%B8%93%E9%A2%98%E5%AD%A6%E4%B9%A0/%E6%95%B0%E6%8D%AE%E5%A4%84%E7%90%86.md), [自然语言处理](https://github.com/zhangjx831/Data-Science-Notes/blob/master/%E4%B8%93%E9%A2%98%E5%AD%A6%E4%B9%A0/%E8%87%AA%E7%84%B6%E8%AF%AD%E8%A8%80%E5%A4%84%E7%90%86.md), [文本分类](https://github.com/zhangjx831/Data-Science-Notes/blob/master/%E4%B8%93%E9%A2%98%E5%AD%A6%E4%B9%A0/%E6%96%87%E6%9C%AC%E5%88%86%E7%B1%BB.md), [路径规划](https://github.com/zhangjx831/Data-Science-Notes/blob/master/%E4%B8%93%E9%A2%98%E5%AD%A6%E4%B9%A0/%E8%B7%AF%E5%BE%84%E8%A7%84%E5%88%92.md)

## 持续更新

本仓库将终生更新，永久开源，为大家的学习提供指导

接下来半年主要关注在数学基础、算法数据结构和深度学习机器学习三个方面

欢迎大家关注