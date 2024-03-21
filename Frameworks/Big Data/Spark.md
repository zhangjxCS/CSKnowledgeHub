# Spark

![What is Spark? - Introduction to Apache Spark and Analytics - AWS](../../Images/what-is-apache-spark.b3a3099296936df595d9a7d3610f1a77ff0749df.png)

## Components

### Spark Driver

- **Role**: The Spark Driver is the heart of a Spark application. It converts the user's program into tasks, schedules the tasks on the executors, and orchestrates the execution by distributing the tasks among the executors.
- Components
  - **SparkContext**: Establishes the connection to the Spark execution environment and acts as the client through which the Spark application is executed. It's responsible for making RDDs resilient and distributed across the cluster.
  - **DAGScheduler**: Translates RDDs into execution graphs (DAGs) and stages (groups of tasks).
  - **TaskScheduler**: Responsible for sending tasks to the cluster, tracking their execution, and retrying if necessary.

### Spark Executor

- **Role**: Executors are responsible for executing the tasks assigned to them by the driver. Each executor runs in its own JVM on a worker node.
- **Functionality**: Executors perform computations on the tasks, read and write data to external sources, and store intermediate data in memory or disk. Executors report the state of the computation back to the driver node.

### Cluster Manager

- **Role**: The Cluster Manager is an external service for acquiring resources on the cluster (e.g., nodes in the cluster where executors should be launched). It's responsible for managing and allocating resources for Spark applications.
- Types
  - **Standalone**: A simple cluster manager included with Spark that makes it easy to set up a cluster.
  - **Apache Mesos**: A general cluster manager that can also run Hadoop MapReduce and service applications.
  - **Hadoop YARN**: The resource manager in Hadoop 2.
  - **Kubernetes**: An open-source system for automating deployment, scaling, and management of containerized applications.

## Workflow

1. **Initialization**: When a Spark application starts, the Driver program creates a SparkContext, which then connects to a Cluster Manager to ask for resources to launch Executors.
2. **Task Distribution**: The Driver program converts the application into a set of stages that are broken down into tasks. These tasks are bundled and sent to the appropriate Executors for execution.
3. **Task Execution**: Executors run the tasks and save the results. If any computations require shuffling data across the executors (for operations like `reduceByKey`), Spark will manage the data transfer.
4. **Results**: Once all tasks are executed, the results are sent back to the Driver for aggregation or further processing. The final results can be collected at the Driver or written to an external data source.

## Data Model

- **RDD (Resilient Distributed Datasets)**: The fundamental data structure of Spark. It represents an immutable, distributed collection of objects that can be processed in parallel. RDDs support two types of operations: transformations and actions.
- **DAG Execution**: Unlike traditional MapReduce models which execute tasks in a linear sequence, Spark uses a DAG (Directed Acyclic Graph) to represent a sequence of computations. Each node represents an RDD partition, and each edge represents a transformation that leads to a new RDD.
- **DataFrame:** The DataFrame is a key component of Apache Spark’s data processing model, introduced as an extension of the RDD (Resilient Distributed Dataset) concept to provide a higher-level abstraction. DataFrames offer a more structured and intuitive interface for data manipulation and analysis, making Spark more accessible for users familiar with SQL and data analysis tools.

## Functions

### Transformation Functions
Transformation functions create a new dataset from an existing one. They are lazy operations, meaning they don't compute their results right away. Instead, Spark remembers the set of transformations applied to some base dataset. The transformed dataset is only computed when an action is called. Examples include:

- `map`: Applies a function to each element in the dataset.
- `filter`: Returns a new dataset formed by selecting those elements of the source on which the function returns true.
- `flatMap`: Similar to map, but each input item can be mapped to 0 or more output items.
- `reduceByKey`: When called on a dataset of (K, V) pairs, returns a dataset of (K, V) pairs where the values for each key are aggregated using the given reduce function.
- `join`: Joins two datasets based on the keys.

### Action Functions
Action functions trigger the computation and return values. Actions are the operations that produce a result from a Spark computation. Unlike transformations, actions force the computation of the transformations required for the result. Examples include:

- `collect`: Returns all the elements of the dataset as an array.
- `count`: Returns the number of elements in the dataset.
- `first`: Returns the first element in the dataset.
- `take`: Returns an array with the first n elements of the dataset.
- `reduce`: Aggregates the elements of the dataset using a function, returning a single value.

### Aggregate Functions
Aggregate functions compute a summary statistic (or statistics) about a dataset. Spark SQL provides built-in methods for aggregating data, such as counting, summing, or averaging. These functions are often used in conjunction with groupBy operations. Examples include:

- `sum`: Calculates the sum of a numeric column.
- `avg`: Computes the average of a numeric column.
- `max`: Finds the maximum value in a numeric column.
- `min`: Finds the minimum value in a numeric column.
- `groupBy`: Groups the DataFrame using the specified columns, then, we can run aggregation on them.

### Window Functions
Window functions perform calculations across a set of rows related to the current row. This is similar to aggregate functions, but, instead of returning a single value for the entire table, a separate value is returned for each row. Examples include:

- `rank`: Assigns a rank to each row within a partition of a result set.
- `dense_rank`: Similar to `rank`, but without gaps in the ranking sequence when there are ties.
- `row_number`: Assigns a unique sequential integer to rows within a partition of a result set, starting at 1 for the first row in each partition.

### UDFs (User-Defined Functions)
Spark allows you to define custom transformations using UDFs. A UDF is a function created by the user to extend the capabilities of Spark's built-in functions. UDFs can be used in Spark SQL to perform operations that are not available through standard Spark functions. 

### Column Functions
In Spark DataFrames, column functions are used to perform operations on DataFrame columns. They can be used for simple transformations like adding or subtracting values, or for more complex operations like applying regular expressions. Examples include:

- `lit`: Creates a Column of literal value.
- `col` or `column`: Returns a Column based on the given column name.
- `asc` or `desc`: Returns a Column sorted in ascending or descending order, respectively.

## HDFS

### Architecture

- **NameNode**: The master server that manages the file system namespace and regulates access to files by clients. It maintains the file system tree and the metadata for all the files and directories. This information is stored in memory for fast access.
- **DataNode**: These are the slave servers responsible for storing the actual data. A HDFS cluster may consist of hundreds or thousands of DataNodes. The DataNodes handle read and write requests from the file system’s clients, and they also manage block creation, deletion, and replication based on instructions from the NameNode.
- **Secondary NameNode**: Despite its name, it does not serve as a backup to the NameNode but rather performs housekeeping tasks for it, such as periodically merging changes with the main metadata.

### File Type

- **Text Files**

  - The simplest form of data storage, where data is stored in plain text format.

  - Easy to read and write but not the most efficient in terms of storage and processing, especially for large datasets.

- **Sequence Files**

  - A binary file format that stores key-value pairs.

  - It's a flat file consisting of binary key/value pairs, which is splittable and thus suitable for MapReduce jobs.

  - Offers a higher performance than plain text files and is suitable for storing intermediate data in Hadoop computations.

- **Avro**

  - A row-based format developed within the Apache Hadoop project.

  - It supports schema evolution, allowing serialized data to be stored with its schema for easy data interpretation.

  - Offers efficient data compression and serialization, making it suitable for both storage and network transmission.

- **Parquet**

  - A columnar storage file format optimized for use in the Hadoop ecosystem.

  - It's highly efficient for both storage and performance, especially for complex nested data structures.

  - Supports schema evolution and is ideal for queries that fetch specific columns from large datasets, significantly reducing I/O operations.

## Spark Run

- Application Properties

  - **`--class`**: The entry point for your application (e.g., `org.apache.spark.examples.SparkPi`).

  - **`--name`**: Name of your Spark application.

  - **`--master`**: The master URL for the cluster (e.g., `spark://host:port`, `yarn`, `local[4]`).

- Deploying Modes
  - **`--deploy-mode`**: Whether to deploy your application in `client` or `cluster` mode.

- Resource Allocation

  - **`--executor-memory`**: Amount of memory to use per executor process (e.g., `1000M`, `2G`).

  - **`--driver-memory`**: Amount of memory to use for the driver process.

  - **`--executor-cores`**: The number of cores to use on each executor.

  - **`--num-executors`**: The number of executors to launch for YARN or Kubernetes modes. For standalone mode, use `--total-executor-cores`.

  - **`--total-executor-cores`**: For Spark standalone and Mesos only, the total cores for all executors.

```shell
spark-submit \
  --class org.apache.spark.examples.SparkPi \
  --master spark://host:port \
  --deploy-mode cluster \
  --executor-memory 4G \
  --driver-memory 2G \
  --executor-cores 2 \
  --num-executors 10 \
  /path/to/spark-examples_*.jar \
  1000
```