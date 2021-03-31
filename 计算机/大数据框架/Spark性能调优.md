## 基础调优

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