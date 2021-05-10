## 基础概念

### 1.Spark的前世今生

Spark计算机集群是2009年由UC Berkeley AMP lab开发的一个集群计算的框架，目的是让数据分析更加快速。在大数据的场景下，数据处理的流程变慢，针对这种情况通常有两种解决办法，一个是设计出更先进的计算机，快速处理数据，另一个是利用多台计算机，协同处理数据，Spark就是采用后一种方法来提高计算速度。Spark作为一个计算引擎，本身并不储存和计算。对于一组计算机，每个计算机节点作为独立的计算资源，而Spark计算引擎确定了各个机器之间应该如何合作处理数据。

### 2.Spark核心组件

![img](https://pic3.zhimg.com/v2-c082dd3bc32544d46788453cebdc70fe_b.png)

Spark使用Scala语言进行实现，它是一种面向对象、函数式编程语言，能够像操作本地集合对象一样轻松地操作分布式数据集。Spark接受的语法有SQL, Python, Scala, Java, R，读入的数据格式可以是JSON、CSV、HDFS（Hadoop分布式文件）、MySQL数据表等等。Spark引擎主要由下述的五个部分组成。

- Spark Streaming：支持高吞吐量、支持容错的实时流数据处理
- Spark SQL: 结构化数据查询
- MLLib：Spark 生态系统里用来解决大数据机器学习问题的模块
- GraphX：构建于Spark上的图计算模型
- SparkR：一个R语言包，它提供了轻量级的方式使得可以在R语言中使用 Spark

### 3.Spark作业方式

![img](https://pic4.zhimg.com/v2-73734f1009ecaa6d1c77728797789f9b_b.png)

我们使用spark-submit提交一个Spark作业之后，这个作业就会启动一个对应的Driver进程。根据你使用的部署模式（deploy-mode）不同，Driver进程可能在本地启动，也可能在集群中某个工作节点上启动。Driver进程本身会根据我们设置的参数，占有一定数量的内存和CPU core。而Driver进程要做的第一件事情，就是向集群管理器Cluster Manager（可以是Spark Standalone集群，也可以是其他的资源管理集群，常用的集群管理器还有Yarn Mesos等等）申请运行Spark作业需要使用的资源，这里的资源指的就是Executor进程。YARN集群管理器会根据我们为Spark作业设置的资源参数，在各个工作节点上，启动一定数量的Executor进程，每个Executor进程都占有一定数量的内存和CPU core。

下面讲述了上图中的一些模块的概念和作用

- **Client：**客户端进程，负责提交作业到Master。
- **Cluster Master:** 主控节点，负责接收Client提交的作业，管理Worker，并命令Worker启动分配Driver的资源和启动Executor的资源。
- **Worker：**守护进程，负责管理本节点的资源，定期向Master汇报，接收Master的命令，启动Driver和Executor。
- **Driver：** 一个Spark作业运行时包括一个Driver进程，也是作业的主进程，负责作业的解析、生成Stage并调度Task到Executor上。包括DAGScheduler，TaskScheduler。
- **Executor：**即真正执行作业的地方，一个集群一般包含多个Executor，每个Executor接收Driver的命令Launch Task，一个Executor可以执行一到多个Task。

![img](https://pic3.zhimg.com/v2-19e4098433ba55e57c73cfa2881206a2_b.png)

Driver与Executor之间的关系

### 4.Spark数据结构

在Spark中，数据以RDD或者DataFrame的格式储存。

**Resilient Distributed Datasets(RDD)：**意为容错的、并行的数据结构，可以让用户显式地将数据存储到磁盘和内存中，并能控制数据的分区。同时，RDD还提供了一组丰富的操作来操作这些数据。

**DataFrame：**DataFrame是一种以RDD为基础的分布式数据集，类似于传统数据库中的二维表格。在Spark1.3版本开始使用。DataFrame带有schema元信息，即DataFrame所表示的二维表数据集的每一列都带有名称和类型。这使得Spark得以洞察更多的结构信息，从而对藏于DataFrame背后的数据源以及作用于DataFrame之上的变换进行了针对性的优化，最终达到大幅提升运行时效率的目标。

### 5.Spark 数据处理流程

在Spark数据处理过程中，通常是将外部的数据，例如HDFS文件、Hive表导入成Spark RDD，之后通过对已有的RDD中的数据执行计算（筛选、排序、映射等）进行转换，而产生新的符合需求的RDD，并对已有的RDD中的数据执行计算产生结果，将结果返回Driver程序或写入到外部物理存储，至此实现数据的ETL（提取、转换、加载）过程。这个数据处理过程通常可以概括为以下三个步骤。

- **创建：**通过加载外部物理存储（如HDFS）中的数据集，或Application中定义的对象集合（如List）来创建。RDD在创建后不可被改变，只可以对其执行下面两种操作。
- **转换（Transformation）**：对已有的RDD中的数据执行计算进行转换，而产生新的RDD，在这个过程中有时会产生中间RDD。Spark对于Transformation采用惰性计算机制，遇到Transformation时并不会立即计算结果，而是要等遇到Action时一起执行。
- **行动（Action）：**对已有的RDD中的数据执行计算产生结果，将结果返回Driver程序或写入到外部物理存储。在Action过程中同样有可能生成中间RDD。

为了实现多台计算机之间的协同工作，Spark中还定义了一些概念如下。

- **Partition（分区）：**一个RDD在物理上被切分为多个Partition，即数据分区，这些Partition可以分布在不同的节点上。Partition是Spark计算任务的基本处理单位，决定了并行计算的粒度，而Partition中的每一条Record为基本处理对象。例如对某个RDD进行map操作，在具体执行时是由多个并行的Task对各自分区的每一条记录进行map映射。
- **Dependency（依赖）：**对RDD的Transformation或Action操作，让RDD产生了父子依赖关系（事实上，Transformation或Action操作生成的中间RDD也存在依赖关系），这种依赖分为宽依赖和窄依赖两种

![img](https://pic2.zhimg.com/v2-f63fcdb5b1a53135bd507738b0b4cbc5_b.png)

Spark宽依赖与窄依赖

- **窄依赖：**parent RDD中的每个Partition最多被child RDD中的一个Partition使用。让RDD产生窄依赖的操作可以称为窄依赖操作，如map、union。
- **宽依赖：**parent RDD中的每个Partition被child RDD中的多个Partition使用，这时会依据Record的key进行数据重组，这个过程即为Shuffle（洗牌）。让RDD产生宽依赖的操作可以称为宽依赖操作，如reduceByKey, groupByKey。
- **Shuffle：**有一部分Transformation或Action会让RDD产生宽依赖，这样过程就像是将父RDD中所有分区的Record进行了“洗牌”（Shuffle），数据被打散重组，如属于Transformation操作的join，以及属于Action操作的reduce等，都会产生Shuffle。

正如我们之间所讲到的，在申请到了作业执行所需的资源之后，Driver进程就会开始调度和执行我们编写的作业代码了。Driver进程会将我们编写的Spark作业代码分拆为多个stage，每个stage执行一部分代码片段，并为每个stage创建一批task，然后将这些task分配到各个Executor进程中执行。task是最小的计算单元，负责执行一模一样的计算逻辑（也就是我们自己编写的某个代码片段），只是每个task处理的数据不同而已。一个stage的所有task都执行完毕之后，会在各个节点本地的磁盘文件中写入计算中间结果，然后Driver就会调度运行下一个stage。下一个stage的task的输入数据就是上一个stage输出的中间结果。如此循环往复，直到将我们自己编写的代码逻辑全部执行完，并且计算完所有的数据，得到我们想要的结果为止。

Spark是根据shuffle类算子来进行stage的划分。如果我们的代码中执行了某个shuffle类算子（比如reduceByKey、join等），那么就会在该算子处，划分出一个stage界限来。可以大致理解为，shuffle算子执行之前的代码会被划分为一个stage，shuffle算子执行以及之后的代码会被划分为下一个stage。因此一个stage刚开始执行的时候，它的每个task可能都会从上一个stage的task所在的节点，去通过网络传输拉取需要自己处理的所有key，然后对拉取到的所有相同的key使用我们自己编写的算子函数执行聚合操作（比如reduceByKey()算子接收的函数）。这个过程就是shuffle。

Spark作业根据上述的流程可以画出一个有向无环图(DAG)。Spark根据用户Application中的RDD的转换和行动，生成RDD之间的依赖关系，RDD之间的计算链构成了RDD的血统（Lineage），同时也生成了逻辑上的DAG（有向无环图）。每一个RDD都可以根据其依赖关系一级一级向前回溯重新计算，这便是Spark实现容错的一种手段

![img](https://pic3.zhimg.com/v2-0ddbb35cb65aae25dbb3eb5cbb050856_b.png)

### 6.Spark作业过程

下面我们简单总结一下Spark的作业过程

1. 用户提交spark作业，Driver构建DAG图并分解成Stage
2. Driver向Cluster Manager申请资源(Executor)
3. Cluster Manager向某些Work Node发送征召信号
4. 被征召的Work Node启动Executor进程响应征召，并向Driver申请任务
5. Driver分配Task给Work Node
6. Executor以Stage为单位执行Task，期间Driver进行监控
7. Driver收到Executor任务完成的信号后向Cluster Manager发送注销信号
8. Cluster Manager向Work Node发送释放资源信号
9. Work Node对应Executor停止运行
10. 返回运行结果

## Spark调优

### 1.开发调优

1.避免创建重复的RDD：在开发过程中要注意：对于同一份数据，只应该创建一个RDD，不能创建多个RDD来代表同一份数据

2.尽可能复用同一个RDD：对于多个RDD的数据有重叠或者包含的情况，我们应该尽量复用一个RDD，这样可以尽可能地减少RDD的数量，从而尽可能减少算子执行的次数

3.对多次使用的RDD进行持久化：Spark会根据你的持久化策略，将RDD中的数据保存到内存或者磁盘中。以后每次对这个RDD进行算子操作时，都会直接从内存或磁盘中提取持久化的RDD数据，然后执行算子，而不会从源头处重新计算一遍这个RDD，再执行算子操作。

cache()方法表示：使用非序列化的方式将RDD中的数据全部尝试持久化到内存中。

persist()方法表示：手动选择持久化级别，并使用指定的方式进行持久化。在持久化过程中，按照内存是否溢出依次尝试MEMORY_ONLY, MEMORY_ONLY_SER, MEMORY_AND_DISK_SER

![img](https://pic3.zhimg.com/v2-1599a76d45d6e4542154f6e5454566ea_b.png)

持久化级别分类

4.尽量避免使用shuffle类算子：在shuffle过程中，可能会发生大量的磁盘文件读写的IO操作，以及数据的网络传输操作。磁盘IO和网络数据传输也是shuffle性能较差的主要原因。在我们的开发过程中，能避免则尽可能避免使用reduceByKey、join、distinct、repartition等会进行shuffle的算子，尽量使用map类的非shuffle算子。在两个数据集的大小相差比较大时，使用广播机制可以加速join的过程；将数据量较小的RDD作为广播变量，相当于在每个executor中创建一份被广播RDD的数据副本

5.使用map-side预聚合的shuffle操作：在可能的情况下，建议使用reduceByKey或者aggregateByKey算子来替代掉groupByKey算子。因为reduceByKey和aggregateByKey算子都会使用用户自定义的函数对每个节点本地的相同key进行预聚合。而groupByKey算子是不会进行预聚合的，全量的数据会在集群的各个节点之间分发和传输，性能相对来说比较差。

6.使用高性能算子：

- 使用reduceByKey/aggregateByKey替代groupByKey
- 使用mapPartitions替代普通map
- 使用foreachPartitions替代foreach
- 使用filter之后进行coalesce操作
- 使用repartitionAndSortWithinPartitions替代repartition与sort类操作

7.使用Kryo优化序列化性能

8.优化数据结构：尽量使用字符串替代对象，使用原始类型（比如Int、Long）替代字符串，使用数组替代集合类型，这样尽可能地减少内存占用，从而降低GC频率，提升性能

### 2.资源参数调优

1.num-executors：该参数用于设置Spark作业总共要用多少个Executor进程来执行。每个Spark作业的运行一般设置50~100个左右的Executor进程比较合适，设置太少或太多的Executor进程都不好。设置的太少，无法充分利用集群资源；设置的太多的话，大部分队列可能无法给予充分的资源。

2.executor-memory：该参数用于设置每个Executor进程的内存。每个Executor进程的内存设置4G~8G较为合适。

3.executor-cores：该参数用于设置每个Executor进程的CPU      core数量。Executor的CPU core数量设置为2~4个较为合适。

4.driver-memory：该参数用于设置Driver进程的内存。Driver的内存通常来说不设置，或者设置1G左右应该就够了。唯一需要注意的一点是，如果需要使用collect算子将RDD的数据全部拉取到Driver上进行处理，那么必须确保Driver的内存足够大，否则会出现OOM内存溢出的问题。

5.spark.storage.memoryFraction：该参数用于设置RDD持久化数据在Executor内存中能占的比例，默认是0.6。

6.spark.shuffle.memoryFraction：该参数用于设置shuffle过程中一个task拉取到上个stage的task的输出后，进行聚合操作时能够使用的Executor内存的比例，默认是0.2。如果Spark作业中的RDD持久化操作较少，shuffle操作较多时，建议降低持久化操作的内存占比，提高shuffle操作的内存占比比例，避免shuffle过程中数据过多时内存不够用，必须溢写到磁盘上，降低了性能。此外，如果发现作业由于频繁的gc导致运行缓慢，意味着task执行用户代码的内存不够用，那么同样建议调低这个参数的值。