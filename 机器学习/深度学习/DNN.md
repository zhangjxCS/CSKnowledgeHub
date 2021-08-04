### 神经网络结构

1.DNN为什么要有bias term, bias term的intuition是什么

bias term可以使得我们分类的直线平移，增加了一个自由度，更加灵活，也可以视作激活函数的阈值

2.什么是Back Propagation

神经网络中根据链式法则，误差项反向传播，计算最终结果的误差对各个节点的导数，从而对整个网络的所有节点的参数值进行优化

3.梯度消失和梯度爆炸是什么，怎么解决

深度神经网络在反向传播过程中，多层的导数相乘，若每层的导数较大，则会发生梯度爆炸，最终结果很大。若每层导数较小，则会发生梯度消失，最终结果很小

常用的解决方法有更换激活函数、batchnorm、采用残差结构等

Relu激活函数的导数恒为1，不存在梯度消失和梯度爆炸的问题

Batchnorm对每一层的输出进行标准化，保证均值和方差一致

残差网络对上一层预测的残差进行训练，解决了梯度消失和爆炸的问题

4.神经网络初始化能不能把weights都initialize成0

不能，因为如果将权值初始化为 0 ，或者其他统一的常量，会导致后面的激活单元具有相同的值，所有的单元相同意味着它们都在计算同一特征，网络变得跟只有一个隐含层节点一样，这使得神经网络失去了学习不同特征的能力

5.DNN和Logistic Regression的区别

逻辑回归可以看作只有一个神经元的神经网络，深度神经网络通过增加神经元的个数获取更多隐藏在变量当中的信息，拟合能力强

6.common activation functions （sigmoid, tanh, relu, leaky relu） 是什么以及每个的优缺点

sigmoid 函数 ：$a=\frac{1}{1+e^{-z}}$

常用的激活函数，也是逻辑回归中的激活函数，当z趋近于无穷时，梯度趋近于0，可能会有梯度消失问题，在梯度下降时，收敛速度很慢

tanh 函数（the hyperbolic tangent function，双曲正切函数）：$a = \frac{e^z - e^{-z}}{e^z + e^{-z}}$

效果通常好于sigmoid函数，值域在-1到1之间，有数据中心化的效果，当z趋近于无穷时，梯度趋近于0，可能会有梯度消失问题，在梯度下降时，收敛速度很慢

ReLU 函数（the rectified linear unit，修正线性单元）：$a=max(0,z)$

当 z > 0 时，梯度始终为 1，从而提高神经网络基于梯度算法的运算速度，收敛速度远大于 sigmoid 和 tanh。然而当 z < 0 时，梯度一直为 0

Leaky ReLU（带泄漏的 ReLU）：$a=max(0.01z,z)$

Leaky ReLU 保证在 z < 0 的时候，梯度仍然不为 0。理论上来说，Leaky ReLU 有 ReLU 的所有优点，但在实际操作中没有证明总是好于 ReLU，因此不常用

![w600](https://i.loli.net/2021/08/02/GEeuqFxMY9noD2b.jpg)

7.为什么需要non-linear activation functions

使用线性激活函数和不使用激活函数、直接使用 Logistic 回归没有区别，那么无论神经网络有多少层，输出都是输入的线性组合，与**没有隐藏层**效果相当

### 优化问题

1.how to do hyperparameter tuning in DL

Random search：随机搜索，在参数空间中随机选择参数组合，计算效率相比网格搜索更高

Grid search：网格搜索，M个参数，每个参数有N个候选值，尝试$M^N$种不同的参数组合

2.Deep Learning有哪些预防overfitting的办法

适用于传统机器学习算法解决过拟合问题的方法同样适用于深度学习

除此之外深层神经网络还可以采用Dropout、Early stopping、Data Augmentation等方法解决过拟合

Dropout: 神经网络的每个单元都被赋予在计算中被暂时忽略的概率p。超参数p称为丢失率，通常将其默认值设置为0.5。然后，在每次迭代中，根据指定的概率随机选择丢弃的神经元。因此，每次训练会使用较小的神经网络，这种方法减小了模型的复杂度。

训练模型阶段：在训练网络的每个单元都要添加一道概率流程，以一定概率舍弃某个神经元，同时对权重进行缩放，乘概率 $\frac{1}{1-p}$

测试模型阶段：用所有神经元参与计算

Early stopping：在用梯度下降法的过程中，在模型开始过拟合之前就中断学习过程

数据扩增（Data Augmentation）：通过图片的一些变换（翻转，局部放大后切割等），得到更多的训练集和验证集

3.常用的优化方法

Batch Gradient Descend：最常用的梯度下降形式，即同时处理整个训练集；每一步梯度下降法需要对整个训练集进行一次处理，如果训练数据集很大的时候，处理速度就会比较慢

Mini-batch Gradient Descend：每次同时处理单个的 mini-batch，其他与 batch 梯度下降法一致。对整个训练集的一次遍历（称为一个 epoch）能做 mini-batch 个数个梯度下降。之后，可以一直遍历训练集，直到最后收敛到一个合适的精度。

Stochastic Gradient Descend：对每一个训练样本执行一次梯度下降，训练速度快

![image-20210320131646991](https://i.loli.net/2021/08/02/NZgAcEaq74b3r1e.png)

Batch梯度下降对所有样本做一次梯度下降，每次的迭代时间较长；Mini-batch梯度下降使用部分样本做梯度下降，兼容了模型的训练速度和梯度下降方向的准确性；随机梯度下降对每个样本进行一次梯度下降，速度最快，但噪声较大，且永远不会收敛

4.梯度下降的优化器

动量梯度下降（Gradient Descent with Momentum）：计算梯度的指数加权平均数，并利用该值来更新参数值

$$v_{dW^{[l]}} = \beta v_{dW^{[l]}} + (1 - \beta) dW^{[l]}$$

$$v_{db^{[l]}} = \beta v_{db^{[l]}} + (1 - \beta) db^{[l]}$$

$$W^{[l]} := W^{[l]} - \alpha v_{dW^{[l]}}$$ 

$$b^{[l]} := b^{[l]} - \alpha v_{db^{[l]}}$$

![img](http://www.ai-start.com/dl2017/images/cc2d415b8ccda9fdaba12c575d4d3c4b.png)

使用动量梯度下降时，通过累加过去的梯度值来减少抵达最小值路径上的波动，加速了收敛，因此在横轴方向下降得更快，从而得到图中红色的曲线

RMSProp（Root Mean Square Propagation，均方根传播）：在对梯度进行指数加权平均的基础上，引入平方和平方根。RMSProp 有助于减少抵达最小值路径上的摆动，并允许使用一个更大的学习率 α，从而加快算法学习速度

$$s_{dw} = \beta s_{dw} + (1 - \beta)(dw)^2$$

$$s_{db} = \beta s_{db} + (1 - \beta)(db)^2$$ 

$$w := w - \alpha \frac{dw}{\sqrt{s_{dw} + \epsilon}}$$

$$b := b - \alpha \frac{db}{\sqrt{s_{db} + \epsilon}}$$

![img](http://www.ai-start.com/dl2017/images/d43cf7898bd88adff4aaac607c1bd5a1.png)

Adam 优化算法（Adaptive Moment Estimation，自适应矩估计）：将 Momentum 和 RMSProp 算法结合在一起，通常有超越二者单独时的效果。

$$v_{dW} = \beta_1 v_{dW} + (1 - \beta_1) dW$$ 

$$v_{db} = \beta_1 v_{db} + (1 - \beta_1) db$$

$$s_{dW} = \beta_2 s_{dW} + (1 - \beta_2) {(dW)}^2$$ 

$$s_{db} = \beta_2 s_{db} + (1 - \beta_2) {(db)}^2$$

$$v^{corrected}_{dW} = \frac{v_{dW}}{1-{\beta_1}^t}$$ 

$$v^{corrected}_{db} = \frac{v_{db}}{1-{\beta_1}^t}$$

$$s^{corrected}_{dW} = \frac{s_{dW}}{1-{\beta_2}^t}$$ 

$$s^{corrected}_{db} = \frac{s_{db}}{1-{\beta_2}^t}$$

$$W := W - \alpha \frac{v^{corrected}_{dW}}{{\sqrt{s^{corrected}_{dW}} + \epsilon}}$$

$$b := b - \alpha \frac{v^{corrected}_{db}}{{\sqrt{s^{corrected}_{db}} + \epsilon}}$$

5.learning rate过大过小对于模型的影响

学习率过小，学习速度较慢，较多时间才能获得最优解

学习率过大，可能会错过最优解，梯度下降过程不收敛

6.Problem of Plateau, saddle point

**鞍点（saddle）**是函数上的导数为零，但不是轴上局部极值的点。当我们建立一个神经网络时，通常梯度为零的点是下图所示的鞍点，而非局部最小值。减少损失的难度也来自误差曲面中的鞍点，而不是局部最低点。因为在一个具有高维度空间的成本函数中，如果梯度为 0，那么在每个方向，成本函数或是凸函数，或是凹函数。而所有维度均需要是凹函数的概率极小，因此在低维度的局部最优点的情况并不适用于高维度。

![img](http://www.ai-start.com/dl2017/images/a8c3dfdc238762a9f0edf26e6037ee09.png)