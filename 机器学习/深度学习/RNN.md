## RNN模型

### 语句表示

对于一个序列数据 x，用符号 x⟨t⟩来表示这个数据中的第 tt个元素，用 y⟨t⟩来表示第 t个标签，用 Tx和 Ty来表示输入和输出的长度。对于一段音频，元素可能是其中的几帧；对于一句话，元素可能是一到多个单词。

想要表示一个词语，需要先建立一个**词汇表（Vocabulary）**，或者叫**字典（Dictionary）**。将需要表示的所有词语变为一个列向量，可以根据字母顺序排列，然后根据单词在向量中的位置，用 **one-hot 向量（one-hot vector）**来表示该单词的标签：将每个单词编码成一个 R|V|×1向量，其中 |V|是词汇表中单词的数量。一个单词在词汇表中的索引在该向量对应的元素为 1，其余元素均为 0。one-hot 向量是最简单的词向量。它的**缺点**是，由于每个单词被表示为完全独立的个体，因此单词间的相似度无法体现。

### 正向传播

$$
W_a=[W_{aa},W_{ax}]
$$

$$
a^{<t>}=g_1(W_a[a^{<t-1>};x^{<t>}]+b_a)
$$

$$
\hat y^{<t>}=g_2(W_{ya}a^{<t>}+b_y)
$$

![Recurrent-Neural-Network](https://raw.githubusercontent.com/bighuang624/Andrew-Ng-Deep-Learning-notes/master/docs/Sequence_Models/Recurrent-Neural-Network.png)

### 反向传播

单个位置上（或者说单个时间步上）某个单词的预测值的损失函数采用**交叉熵损失函数**，将单个位置上的损失函数相加，得到整个序列的成本函数

![formula-of-RNN](https://raw.githubusercontent.com/bighuang624/Andrew-Ng-Deep-Learning-notes/master/docs/Sequence_Models/formula-of-RNN.png)

通过RNN来预测一个序列的值通常需要两步，第一步是构建一个语言模型，第二部是通过采样来构建序列

### 语言模型

建立语言模型所采用的训练集是一个大型的**语料库（Corpus）**，指数量众多的句子组成的文本。建立过程的第一步是标记化，即建立字典；然后将语料库中的每个词表示为对应的 one-hot 向量。第二步是将标志化后的训练集用于训练 RNN

![language-model-RNN-example](https://raw.githubusercontent.com/bighuang624/Andrew-Ng-Deep-Learning-notes/master/docs/Sequence_Models/language-model-RNN-example.png)

### 采样

在训练好一个语言模型后，可以通过**采样（Sample）**新的序列来了解这个模型中都学习到了一些什么

在第一个时间步输入 a⟨0⟩a⟨0⟩和 x⟨1⟩x⟨1⟩为零向量，输出预测出的字典中每个词作为第一个词出现的概率，根据 softmax 的分布进行随机采样（`np.random.choice`），将采样得到的 ŷ ⟨1⟩作为第二个时间步的输入 x⟨2⟩。以此类推，直到采样到 EOS，最后模型会自动生成一些句子，从这些句子中可以发现模型通过语料库学习到的知识。

## 改进RNN模型

传统的RNN存在梯度消失问题，不善于捕捉语句长期的依赖关系，通过GRU与LSTM模型可以解决长期依赖的问题

### GRU

GRU 单元有一个新的变量称为 c，代表记忆细胞，其作用是提供记忆的能力，记住例如前文主语是单数还是复数等信息。

更新门$\Gamma_u$代表是否要更新参数，在0-1取值；相关门$\Gamma_r$表示 c̃ ⟨t⟩和 c⟨t⟩的相关性
$$
\Gamma_u=\sigma(W_u[c^{<t-1>},x^{<t>}]+b_u)
$$

$$
\Gamma_r=\sigma(W_r[c^{<t-1>},x^{<t>}]+b_r)
$$

$$
\tilde c^{<t>}=tanh(W_c[\Gamma_r*c^{<t-1>},x^{<t>}]+b_c)
$$

$$
a^{<t>}=c^{<t>}=\Gamma_u * \tilde c^{<t-1>}+(1-\Gamma_u)*c^{<t-1>}
$$

### LSTM

相比于GRU，LSTM新引入了遗忘门$\Gamma_f$和输出门$\Gamma_o$，去掉了相关门$\Gamma_r$
$$
\Gamma_u=\sigma(W_u[a^{<t-1>},x^{<t>}]+b_u)
$$

$$
\Gamma_f=\sigma(W_f[a^{<t-1>},x^{<t>}]+b_f)
$$

$$
\Gamma_o=\sigma(W_o[a^{<t-1>},x^{<t>}]+b_o)
$$

$$
\tilde c^{<t>}=tanh(W_c[\Gamma_r*c^{<t-1>},x^{<t>}]+b_c)
$$

$$
c^{<t>}=\Gamma_u * \tilde c^{<t-1>}+\Gamma_f*c^{<t-1>}
$$

$$
a^{<t>}=\Gamma_o*tanh(c^{<t>})
$$

### BRNN双向循环神经网络

**双向循环神经网络（Bidirectional RNN，BRNN）**可以在序列的任意位置使用之前和之后的数据。**缺点**是需要完整的序列数据，才能预测任意位置的结果。

![BRNN](https://raw.githubusercontent.com/bighuang624/Andrew-Ng-Deep-Learning-notes/master/docs/Sequence_Models/BRNN.png)

### DRNN深度循环神经网络

循环神经网络的每个时间步上也可以包含多个隐藏层

![DRNN](https://raw.githubusercontent.com/bighuang624/Andrew-Ng-Deep-Learning-notes/master/docs/Sequence_Models/DRNN.png)

## word2vec









