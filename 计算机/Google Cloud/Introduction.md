## 简介

![image-20210528161127462](/Users/shawnzhang/Library/Application Support/typora-user-images/image-20210528161127462.png)

组成：计算、存储、网络、安全四个部分，在四个部分之上构建大数据与机器学习产品

• compute engine: a virtual machine, use ssh to connect

• storage: use gsutil to get access to storage

• network: different machines are connected

• security: communications are encrypted in transit

Google在大数据存储、计算、处理的进展如下，GFS是Google File System，是一种分布式文件系统；MapReduce是一个分布式计算的框架，实现了快速的并行计算；BigTable是一种用来处理海量数据的非关系型数据库；Dremel和BigQuery是构建在存储与计算框架上，用于大数据查询的工具

![image-20210528164718756](/Users/shawnzhang/Library/Application Support/typora-user-images/image-20210528164718756.png)

