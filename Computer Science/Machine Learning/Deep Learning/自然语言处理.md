# 自然语言处理

[toc]

## 文本处理

### 文本预处理

- Stopwords and Punctuation: 去除停顿词(and is are at ...)和标点符号(, . : !)
- Handles and URLs: 去除句柄(@...)和链接(www.google.com)
- Stemming: 词根化(tune tuned tuning 属于tun词根)
- Lowercasing: 字母均改为小写
- Tokenize: 将句子分开为词语列表
- Token padding: 将所有句子长度统一，缺失的部分用<pad>代替，计算accuracy时不计算这部分

### 词频统计

- Dictionary: 构建词典，记录出现的每个词
- Frequency: 统计词频，每个词出现的次数，用python词典记录
- Sentiment Count: 统计每个词出现在正面(positive)和负面(negative)句子中的次数

### 向量空间

将词语表示为向量可以代表相对语义

- Word by Word: Number of times they occur together within a certain distance k
- Word by Doc: Number of times a word occurs within a certain category

距离度量

- Euclidean distance: $d(B,A)=\sqrt{(B_1-A_1)^2+(B_2-A_2)^2}$
- Cosine similarity: $d(B,A)=\frac{B\cdot A}{\sqrt{|B||A|}}$, cosine相似度相比欧式距离更好的度量长度差别较大的两个向量之间的相似度

## 传统模型

### Logistic Regression

- 构建特征：Postive and Negative count构建词语正负情绪词典，统计句子中词语出现在语料正负情绪的总和
- 构建模型：构建Logistic Regression model对句子进行情感分析，特征为$[1, X_1, X_2]$，分别为偏置项，词正向频率和，词负向频率和

### Naive Bayes

- 构建特征：Postive and Negative count构建词语正负情绪词典，统计句子中词语出现在语料正负情绪的总和
- 朴素贝叶斯模型：使用朴素贝叶斯模型做情感分类，$\frac{P(pos)}{P(neg)}\prod_{i=1}^m\frac{P(w_i|pos)}{P(w_i|neg)}$
- 拉普拉斯平滑：$P(w_i|class)=\frac{freq(w_i,class)+1}{N_c+V_c}$, $N_c$该类别所有词频次，$V_c$类别中不同词个数

- 对数似然：对连续乘法取对数变为加法，$\lambda(w)=\frac{P(w|pos)}{P(w|neg)}$, $logprior=\frac{D_{pos}}{D_{neg}}$

<img src="https://tva1.sinaimg.cn/large/008i3skNgy1gyu0bzfm2sj31bm0kq77y.jpg" alt="image-20220121164812133" style="zoom:50%;" />

### K Nearest Neighbor

Locality sensitive hashing: 对KNN进行加速

- 将整个空间由不同平面切分，只在hash value相同的点中找nearest neighbor
- 通过平面两侧的符号确定hash function的值，$hash=\sum_i^H2^i*h_i$

<img src="https://tva1.sinaimg.cn/large/008i3skNgy1gyu0c3fkqmj31ae0i0dhx.jpg" alt="image-20220121171408041" style="zoom:50%;" />

## 概率模型

### Hidden Markov Model

- Transition probability: $P(t_i|t_{i-1})=\frac{C(t_{i-1},t_i)+\epsilon}{\sum_{j=1}^N C(t_{i-1},t_j)+N*\epsilon}$
- Emission probability: $P(w_i|t_i)=\frac{C(t_i,w_i)+\epsilon}{\sum_{j=1}^V C(t_i,w_j)+N*\epsilon}$

- Transition matrix: transition probability between each speech tags
- Emission matrix: emission probability between speech tag and word

<img src="https://tva1.sinaimg.cn/large/008i3skNgy1gyu0c5s7gaj31c6092wfc.jpg" alt="image-20220121194512835" style="zoom:50%;" />

Viterbi Algorithm:

<img src="https://tva1.sinaimg.cn/large/008i3skNgy1gyu0c83rl1j31ci0ju76k.jpg" alt="image-20220121195316069" style="zoom:50%;" />

<img src="https://tva1.sinaimg.cn/large/008i3skNgy1gyu0cax566j31ac0is75t.jpg" alt="image-20220121195818619" style="zoom:50%;" />

### N-Gram Language Model

N-Gram probability: $P(w_N|w_1^{N-1})=\frac{w_1^{N-1}w_N}{C(w_1^{N-1})}$

Sequence Probability: $P(A,B,C,D)=P(A)P(B|A)P(C|A,B)P(D|A,B,C)$

N-Gram Language Model:

- Markov assumption: only last N words matter$P(w_N|w_1^{N-1})=P(w_N|w_{N-1})$ (Bigram)
- Start and end of sentence: start token \<s>, end token\</s>, add N-1 start tokens for N-Gram model

$$
p(A,B,C,D)=P(s)P(A|s)P(B|A)P(C|B)P(D|C)P(/s|D)
$$

<img src="https://tva1.sinaimg.cn/large/008i3skNgy1gyu0cen7olj319q0a0t9w.jpg" alt="image-20220121204605339" style="zoom: 33%;" />

Model Evaluation: 

- Perplexity: The smaller the perplexity, the better the model

$$
PP(W)=P(s_1,s_2,...,s_m)^{-1/m}
$$

$$
PP(W)=m\sqrt{\prod_{i=1}^m\prod_{j=1}^{|s_i|}\frac{1}{P(w_j^{(i)}|w_{j-1}^{(i)})}},for\ bigram\ model
$$

$$
PP(W)=m\sqrt{\prod_{i=1}^m\frac{1}{P(w_i|w_{i-1})}},concatenate\ all\ sentences
$$

Out of Vocabulary Words:

- Unknown words can be represented by UNK

Smoothing:

- Laplacian smoothing: $P(w_n|w_{n-1})=\frac{C(w_{n-1},w_n)+1}{C(w_{n-1})+V}$
- Add-k smoothing: $P(w_n|w_{n-1})=\frac{C(w_{n-1},w_n)+k}{C(w_{n-1})+k*V}$

- Backoff: if the higher n-gram probability is missing, then use the lower-order probability

- Interpolation: $\hat P(w_n|w_{n-2},w_{n-1})=\lambda_1P(w_n|w_{n-2},w_{n-1})+\lambda_2P(w_n|w_{n-1})+\lambda_3P(w_n)$

## 词嵌入

### 词表征

- 整数：使用不同整数来对应不同词，没有语义含义
- One-hot vector：使用One-hot vector来代表词，向量维度高，向量无含义
- 词向量：One-hot vector的低维空间表征，含有语义

### CBOW

- 构造滑动窗口，使用周围词来预测中间词
- 词向量可以是$W_1$ $W_2$或者$W_1$, $W_2$的平均值

<img src="https://tva1.sinaimg.cn/large/008i3skNgy1gyu0cjro3kj30uu0e4dh8.jpg" alt="image-20220121210712817" style="zoom: 33%;" />
$$
\bold Z_1=\bold W_1\bold X+\bold B_1,\ \bold H=Relu(\bold Z_1)
$$

$$
\bold Z_2=\bold W_2\bold+\bold B_2,\ \hat Y=softmax(\bold Z_2)
$$

<img src="https://tva1.sinaimg.cn/large/008i3skNgy1gyu0cmj6g9j31bi0mmq6k.jpg" alt="image-20220121210848349" style="zoom: 33%;" />

## 序列模型

### RNN

- RNN模型利用了句子的序列性，构建序列神经网络来处理自然语言

- 改进RNN模型包括GRU LSTM等，减少了长句子中依赖性和gradient vanishing等问题
- Seq2Seq模型可以用于机器翻译等任务，采用了两个RNN模型，一个作为encoder，另一个作为decoder，但两者间的信息传输有bottleneck
- Teacher forcing可用于机器翻译任务，将真实单词输入decoder用于训练

### Siamese Network

Siamese Network用于识别两个句子的相似性

- 输入网络的特征是一个本体Anchor，一个正例Positive，一个负例Negative，作为Triplets
- Loss function: $diff=s(A,N)-s(A,P)$, $L(A,P,N)=max(diff+\alpha,0)$
- Hard Negative Mining: 构造diff时，采用mean negative（所有非对角线元素和）或者closest negative（与对角线元素最近但小于的元素）
- One Shot Learning: 无需重复学习

<img src="https://tva1.sinaimg.cn/large/008i3skNgy1gyu0covfbyj317o0i8tbl.jpg" alt="image-20220122100105624" style="zoom:50%;" />

<img src="https://tva1.sinaimg.cn/large/008i3skNgy1gyu0cqen54j31au0k0q65.jpg" alt="image-20220122100903254" style="zoom:50%;" />

### 机器翻译

模型评估：

- BLEU Score: 将机器翻译结果与reference进行对比，距离1越近，模型效果越好，但BLEU Score未考虑语义

<img src="https://tva1.sinaimg.cn/large/008i3skNgy1gyu0cwh9j5j31800ewq4h.jpg" alt="image-20220122101823497" style="zoom:50%;" />

- ROUGE-N Score:  将机器翻译结果与reference进行对比，计算时除以reference长度，距离1越近，模型效果越好

<img src="https://tva1.sinaimg.cn/large/008i3skNgy1gyu0cxuqy7j31660eodha.jpg" alt="image-20220122102054979" style="zoom:50%;" />

- F1 Score: $F_1=2*\frac{BLEU*ROUGE-N}{BLEU+ROUGE-N}$

下一个词预测：

- Greedy decoding: Selects the most probable word at each step
- Random sampling: 通过预测概率，构建多项式分布，随机采样，通过temperature控制随机性
- Beam search: Calculate probability of multiple possible sequences
- Minimum Bayes Risk: 计算不同candidate的 ROUGE-N Score，选择分数最高的一组

## Transformer

### Transformer

- 包括Encoder Decoder两个模块，在Encoder中包括多头注意力和一个前向神经网络，Decoder包括带掩码的多头注意力和前向神经网络
- 模型输入先过embedding层，并与位置编码相结合
- Attention机制采用scaled dot-product attention
- Transformer可用于摘要生成，机器翻译，自动补全，命名实体识别，对话，问答等任务

### Transfer Learning

- Feature-based: word embedding, etc.
- Model-based: BERT, T5, etc.

<img src="https://tva1.sinaimg.cn/large/008i3skNgy1gyu0d0o9paj31by0kiju9.jpg" alt="image-20220122104120662" style="zoom:50%;" />

<img src="https://tva1.sinaimg.cn/large/008i3skNgy1gyu0d2hi0aj311y0heac7.jpg" alt="image-20220122104505074" style="zoom:50%;" />

