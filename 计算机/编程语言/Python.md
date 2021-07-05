## Python语法

Python语言中单引号和双引号没有区别，都可以表示字符串，为了跟C语言保持一致，单个字符用单引号，字符串用双引号，三引号可以表示多行字符串，在嵌入式写SQL语句时非常有用，同时三引号也可以用来实现多行注释

Python语言字符串不能直接相减，这点与C语言不同，可以用ord()函数返回字符对应的ascii码值之后再相减

### Python赋值、引用、拷贝、作用域

python中没有赋值，赋值语句建立对象的引用值，而不是复制对象

values[:]  生成对象的拷贝或者是复制序列，不再是引用和共享变量，但此法只能顶层复制，相当于浅复制

python中对象的复制可以使用copy模块，浅复制为copy.copy()，深复制为copy.deepcopy()

copy使用场景：列表或字典，且内部元素为数字，字符串或元组

deepcopy使用场景：列表或字典，且内部元素包含列表或字典

可变对象：包括list，set，dict等。可变类型数据对对象操作的时候，不需要再在其他地方申请内存，只需要在此对象后面连续申请(+/-)即可，也就是它的内存地址会保持不变，但区域会变长或者变短。

不可变对象：包括int，float，long，str，tuple等。更改变量时，创建一个新值，把变量绑定到新值上，而旧值如果没有被引用就等待垃圾回收。

### 常用语法糖

列表表达式：`[my_func(i) for i in range(5)]`

条件赋值：`value = a if condition else b`

匿名函数：`lambda x: my_func(x)`

map方法：使用某个函数，对可迭代对象进行处理`map(lambda x: 2*x, range(5))`

zip方法：把多个可迭代对象打包成一个元组构成的可迭代对象`zip(range(1,5), range(2,6))`，常用于循环迭代和建立字典，使用*操作符可以对zip方法作用后的元组列表进行解压

enumerate方法：`for index, value in enumerate(list)`

## Python输入输出

Python中的标准输入方式有两种，一种是利用sys库，另一种是使用input()函数

```python
import sys
# 按行读取，末尾有'\n'符号
for line in sys.stdin:
		print(line)
# 按行读取，末尾有'\n'符号
sys.stdin.readline()
# 从标准输入中读入一行，以换行作为输入结束，也就是说raw_input()读入的东西结尾没有换行符'\n'，并且默认为字符串格式
raw_input()
# 读取标准输入中的一行，以换行作为输入结束，将标准输入当作一个表达式，并且计算出这个表达式的值
input()
```

与标准输入对应，Python标准输出也有两种方式

```python
import sys
# 标准输出，末尾无'\n'
sys.stdout.write()
# print()调用标准输出，并在末尾加'\n'
print()
```

Python格式化输出有两种方法，一种是%，用于基础格式输出，另一种是format()函数，实现更复杂的功能

```python
print('%o' %20) # oct 八进制
print('%d' %20) # dec 十进制
print('%x' %20) # hex 十六进制
print('.3f' % 1.11) # 保留3位小数位
print('.3e' % 1.11) # 保留3位小数位，使用科学计数法
print('.3g' % 1.11) # 在保证六位有效数字的前提下，使用小数方式，否则使用科学计数法
round(1.1135,3)  # 取3位小数，由于3为奇数，则向下“舍”
round(1.1125,3)  # 取3位小数，由于2为偶数，则向上“入”
print('%s' % 'hello world')  # 字符串输出
print('%20s' % 'hello world')  # 右对齐，取20位，不够则补位
print('%-20s' % 'hello world')  # 左对齐，取20位，不够则补位

print('{} {}'.format('hello','world'))  # 不带字段
print('{0} {1}'.format('hello','world'))  # 带数字编号
print('{0} {1} {0}'.format('hello','world'))  # 打乱顺序
'Coordinates: {latitude}, {longitude}'.format(latitude='37.24N', longitude='-115.81W') # 通过位置匹配
print('{:d}'.format(20)) # 十进制整数。将数字以10为基数进行输出，其他进制类似
```

## Python函数

1.函数参数引用

传值：被调函数局部变量改变不会影响主调函数局部变量

传址：被调函数局部变量改变会影响主调函数局部变量

Python参数传递方式：传递对象引用（传值和传址的混合方式），如果是数字，字符串，元组则传值；如果是列表，字典则传址；

传值就是传入一个参数的值，传址就是传入一个参数的地址，也就是内存的地址（相当于指针）。他们的区别是如果函数里面对传入的参数重新赋值，函数外的全局变量是否相应改变，用传值传入的参数是不会改变的，用传址传入就会改变。

在函数的默认参数中，不要使用可变对象，否则同一个变量在函数调用的每次过程都会被反复引用

## Python面向对象

- **类(Class):** 用来描述具有相同的属性和方法的对象的集合。它定义了该集合中每个对象所共有的属性和方法。对象是类的实例。
- **对象：**通过类定义的数据结构实例。对象包括两个数据成员（类变量和实例变量）和方法。
- **数据成员：**类变量或者实例变量, 用于处理类及其实例对象的相关的数据。
- **方法：**类中定义的函数。

Python的数据成员有两种，分别是类变量和实例变量

- **类变量：**类变量在整个实例化的对象中是公用的。类变量定义在类中且在函数体之外。类变量通常不作为实例变量使用。
- **实例变量**：实例化之后，每个实例单独拥有的变量。

Python的方法有3种,即静态方法(staticmethod),类方法(classmethod)和实例方法

类的方法与普通的函数只有一个特别的区别——它们必须有一个额外的**第一个参数名称**, 按照惯例它的名称是 self

```python
def foo(x):
    print "executing foo(%s)"%(x)

class A(object):
    def foo(self,x):
        print "executing foo(%s,%s)"%(self,x)

    @classmethod
    def class_foo(cls,x):
        print "executing class_foo(%s,%s)"%(cls,x)

    @staticmethod
    def static_foo(x):
        print "executing static_foo(%s)"%x

a=A()
```

## Numpy科学计算

### 构造np数组

等差数列：`np.linespace(1,5,11)# 最后一个值代表样本数目`，`np.arange(1,5,2)# 最后一个值代表步长`

特殊矩阵：`np.zeros((2,3))# 零矩阵`，`np.eye(3) # 单位矩阵`，`np.full((2,3),10)# 维度为2*3，元素全为10的矩阵`

随机矩阵：`np.random.rand(3)# 生成3个0-1均匀分布随机数`，`np.random.uniform(5, 15, 3)# 生成3个5-15均匀分布随机数`，`np.random.randn(2,2)# 生成2*2维度标准正态分布矩阵`，`np.random.normal(3,2.5,3)# 生成均值3，标准差2.5的正态分布`，`np.random.randint(low, high, size)# 生成随机整数`，`np.random.choice(my_list, (3,3))# 从my_list中有放回随机采样`，`np.random.permutation(my_list)# 随机打乱列表`

### np数组变形

转置：`A.T`

合并：`np.r_[A,B] # A与B上下合并`，`np.c_[A,B] # A与B左右合并`

维度变换：`A.reshape(3,2)`，`A.reshape(-1,1)# -1代表维度空缺，由python自行决定维度`

### np数组索引切片

`A[:-1, [0,2]] # 传入列表指定维度的索引切片`

### 常用函数

`np.where(a>0,a,0)# 条件为True填充a，否则填充0`

`np.nonzero(a) # 返回非零值的索引`

`a.argmax() # 求最大值索引`

`a.argmin() # 求最小值索引`

`(a>b).any() # 序列至少一个True时返回True`

`(a>b).all() # 序列全为True时返回True`

统计函数：max, min, mean, median, std, var, sum, quantile(全局方法，通过np.quantile调用)

### 矩阵运算

`np.dot(A, B) # 矩阵A与B相乘`

`np.linalg.norm # 向量范数与矩阵范数`

`A@B # 矩阵乘法`

## Pandas数据处理

### 文件读写

`pd.read_csv('data.csv', header=None, sep=',', nrows=10, usecols=['col1', 'col2'])# 指定行数、列名、分隔符读取`

`df.to_csv('data.csv', sep='\t', index=False)# 指定分隔符，不输出列标`

### 数据结构

Series构造：`s = pd.Series(data=['a', 'b'], index=pd.Index([1, 2], name='my_index'))`

获取Series信息：`s.values`, `s.index`, `s.dtype`, `s.name`, `s.shape`

DataFrame构造：`df = pd.DataFrame(data={'col1':[1, 2, 3], 'col2':['a', 'b', 'c']})`

获取DataFrame信息：`df.values`, `df.index`, `df.columns`, `df.dtypes`, `df.shape`

### 常用函数

汇总函数：`df.head()`, `df.tail()` `df.info()`, `df.describe()`

统计函数：`df.sum()`, `df.mean()`, `df.median()`, `df.var()`, `df.std()`, `df.max()`, `df.min()`, `df.quantile(0.8)`, `df.count()` 

唯一值函数：`df['item'].unique()`, `df['item'].nunique()`, `df['item'].value_counts()`, `df.drop_duplicates(['item', 'date'], keep='first')`, `df['item'].duplicated() # 返回是否为唯一值的列表`

替换函数：`df['gender'].replace({'female':0, 'male':1})`, `s.where(s<0)# 替换False行`, `s.mask(s<0) #替换True的行`

排序函数：`df.sort_values('gender')`, `df.sort_index(level=['Name'])`

自定义函数：`df.apply(lambda x: x.mean())`

### 窗口对象

rolling对象：构建滑动窗口`df.rolling(window=3)`

shift偏移对象：`df.shift(2)`

diff对象：`df.diff(3)`

