## Storage Hierarchy

- Volatile Devices: if pull the power from the machine, then the data is lost. Volatile storage supports fast random access with byte-addressable locations.
- Non-Volatile Devices: Non-volatile means that the storage device does not require continuous power in order for the device to retain the bits that it is storing; It is also block/page addressable; Non-volatile storage is traditionally better at sequential access.

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h05hcqi3zjj21ha0u079f.jpg" alt="image-20220310161239124" style="zoom: 33%;" />

## DBMS

A high-level design goal of the DBMS is to support databases that exceed the amount of memory available.

- File Organization: the data in database files is organized into pages, with the first page being the directory page.

- Buffer Pool: manages the data movement back and forth between disk and memory.
- Storage Manager: responsible for managing a database’s files. It represents the files as a collection of pages. It also keeps track of what data has been read and written to pages as well how much free space there is in these pages.

- Execution Engine: The execution engine will ask the buffer pool for a specific page, and the buffer pool will take care of bringing that page into memory and giving the execution engine a pointer to that page in memory.

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h05he3ow9kj21hc0u0jw2.jpg" alt="image-20220310161410774" style="zoom:33%;" />

## Data Representation

- Integers(INTEGER, BIGINT): “native” C/C++ types
- Variable Precision Numbers(FLOAT, REAL):  inexact, variable-precision numeric types, and fixed length;
- Fixed-Point Precision Numbers(NUMERIC, DECIMAL): These are numeric data types with arbitrary precision and scale. These data types are used when rounding errors are unacceptable, but the DBMS pays a performance penalty to get this accuracy.
- Variable-Lenth Data(VARCHAR, VARBINARY, TEXT): These represent data types of arbitrary length. They are typically stored with a header that keeps track of the length of the string to make it easy to jump to the next value. 
- Dates(TIME, DATE): Representations for date/time vary for different systems.
- System Catalogs: In order for the DBMS to be able to deciphter the contents of tuples, it maintains an internal catalog to tell it meta-data about the databases.

## Workloads

OLTP: Online Transaction Processing

- An OLTP workload is characterized by fast, short running operations, simple queries that operate on single entity at a time, and repetitive operations.
- An OLTP workload will typically handle more writes than reads.

OLAP: Online Analytical Processing

- An OLAP workload is characterized by long running, complex queries, reads on large portions of the database.
- The database system is analyzing and deriving new data from existing data collected on the OLTP side.

HTAP: Hybrid Transaction and Analytical Processing

- A combination which tries to do OLTP and OLAP together on the same database.

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h05ipr43b1j21800oudhi.jpg" alt="image-20220310165958850" style="zoom: 33%;" />

## File Storage

The DBMS organizes the database across one or more files in fixed-size blocks of data called *pages*.

A hardware page is the largest block of data that the storage device can guarantee failsafe writes.

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h05jav0qlrj20hs0gk402.jpg" alt="image-20220310172016067" style="zoom:50%;" />

### Heap

A *heap file* is an unordered collection of pages where tuples are stored in random order.

- Linked List: Header page holds pointers to a list of free pages and a list of data pages.
- Page Directory: DBMS maintains special pages that track locations of data pages along with the amount of free space on each page.

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h05jcytg6yj20iy0hkgn9.jpg" alt="image-20220310172217205" style="zoom:50%;" />

## Page Layout

- Page size
- Checksum
- DBMS version
- Transaction visibility
- Self-containment. (Some systems like Oracle require this.)

### Tuple Oriented

Keep track of the number of tuples in a page and then just append a new tuple to the end.

- Tuple Header: Contains meta-data about the tuple.
- Tuple Data: Actual data for attributes.
- Unique Identifier: Each tuple in the database is assigned a unique identifier.

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h05jhb6l0wj20e00d4gm0.jpg" alt="image-20220310172628091" style="zoom: 50%;" />

### Slotted Pages

- The slot array maps "slots" to the tuples' starting position offsets.
- The header keeps track of the number of used slots and the offset of the starting location of the last slot used.

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h05jkbrbf4j20dy0hi3zp.jpg" alt="image-20220310172921908" style="zoom:50%;" />

### Log Structured

- Stores records to file of how the database was modified. To read a record, the DBMS scans the log file backwards and “recreates” the tuple.
- Build indexes to allow it to jump to locations in the log.
- Periodically compact the log. Compaction coalesces larger log files into smaller files by removing unnecessary records.

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h05jog9v7dj217g0ccad6.jpg" alt="image-20220310173319785" style="zoom:50%;" />

## Storage Models

### N-Ary Storage Model

In the n-ary storage model, the DBMS stores all of the attributes for a single tuple contiguously in a single page, so NSM is also known as a “row store.”

This approach is ideal for OLTP workloads where requests are insert-heavy and transactions tend to operate only an individual entity.

- Advantage: Fast inserts, updates, and deletes; Good for queries that need the entire tuple
- Disadvantage: Not good for scanning large portions of the table and a subset of the attributes.

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h05iztge9yj21mc0k6gov.jpg" alt="image-20220310170925611" style="zoom: 33%;" />

### Decomposition Storage Model

The DBMS stores a single attribute (column) for all tuples contiguously in a block of data. Thus, it is also known as a “column store.” 

This model is ideal for OLAP workloads with many read-only queries that perform large scans over a subset of the table’s attributes.

- Advantage: Reduces the amount of wasted work during query execution; Enables better compression
- Disadvantage: Slow for point queries, inserts, updates, and deletes because of tuple splitting/stitching

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h05j2kaoh4j21ky0l00w1.jpg" alt="image-20220310171217363" style="zoom:33%;" />

## Buffer Pool

The *buffer pool* is an in-memory cache of pages read from disk. It is essentially a large memory region allocated inside of the database to store pages that are fetched from disk.

The buffer pool’s region of memory organized as an array of fixed size pages. Each array entry is called a *frame*.

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h05rkge5l5j20sc0igq4e.jpg" alt="image-20220310220617051" style="zoom:50%;" />

### Meta Data

*page table* is an in-memory hash table that keeps track of pages that are currently in memory. The page table also maintains additional meta-data per page, a dirty-flag and a pin/reference counter.

- Page table maps page ids to frame locations in the buffer pool.

- The *dirty-flag* is set by a thread whenever it modifies a page. This indicates to storage manager that the page must be written back to disk.
- The *pin/reference Counter* tracks the number of threads that are currently accessing that page. A thread has to increment the counter before they access the page. If a page’s count is greater than zero, then the storage manager is not allowed to evict that page from memory.

### Memory Allocation

Memory in the database is allocated for the buffer pool according to two policies.

- *Global policies* deal with decisions that the DBMS should make to benefit the entire workload that is being executed. It considers all active transactions to find an optimal decision for allocating memory.
- *local policies* makes decisions that will make a single query or transaction run faster, even if it isn’t good for the entire workload. Local policies allocate frames to a specific transactions without considering the behavior of concurrent transactions.

### Buffer Pool Optimization

- Multiple Buffer Pool: The DBMS can maintain multiple buffer pools for different purposes
  - *Object IDs* involve extending the record IDs to include meta-data about what database objects each buffer pool is managing. 
  - *hashing* where the DBMS hashes the page id to select which buffer pool to access.
- Pre-fetching: The DBMS can also optimize by pre-fetching pages based on the query plan.
- Scan Sharing: Query cursors can reuse data retrieved from storage or operator computations. This allows multiple queries to attach to a single cursor that scans a table.
- Buffer Pool Bypass: The sequential scan operator will not store fetched pages in the buffer pool to avoid overhead. Instead, memory is local to the running query.

### Buffer Relacement

A replacement policy is an algorithm that the DBMS implements that makes a decision on which pages to evict from buffer pool when it needs space.

- Least Recently Used(LRU): The Least Recently Used replacement policy maintains a timestamp of when each page was last accessed. The DBMS picks to evict the page with the oldest timestamp.
- Clock: The CLOCK policy is an approximation of LRU without needing a separate timestamp per page. In the CLOCK policy, each page is given a reference bit. When a page is accessed, set to 1.

- LRU-K: tracks the history of the last K references as timestamps and computes the interval between subsequent accesses.
- Localization per query: The DBMS chooses which pages to evict on a per transaction/query basis. This minimizes the pollution of the buffer pool from each query.