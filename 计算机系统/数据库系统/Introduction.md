## Databases

A *database* is an organized collection of inter-related data that models some aspect of the real-world.

A database management system (DBMS) is the software that manages a database.

A general-purpose DBMS is designed to allow the definition, creation, querying, updation, and administration of databases.

## Ralational Model

A *data model* is a collection of concepts for describing the data in a database. The relational model is an example of a data model.

A *schema* is a description of a particular collection of data, using a given data model.

A *relation* is an unordered set that contains the relationship of attributes that represent entities.

A *tuple* is a set of attribute values (also known as its *domain*) in the relation.

Relational Model:

- Store database in simple data structures (relations). 

- Access data through high-level language.
- Physical storage left up to implementation.

## Ralational Algebra

*Relational Algebra* is a set of fundamental operations to retrieve and manipulate tuples in a relation.

- Select: Select takes in a relation and outputs a subset of the tuples from that relation that satisfy a selection predicate.
- Projection: Projection takes in a relation and outputs a relation with tuples that contain only specified attributes.
- Union: Union takes in two relations and outputs a relation that contains all tuples that appear in at least one of the input relations.
- Intersection: Intersection takes in two relations and outputs a relation that contains all tuples that appear in both of the input relations.
- Difference: Difference takes in two relations and outputs a relation that contains all tuples that appear in the first relation but not the second relation.
- Product: Product takes in two relations and outputs a relation that contains all possible combinations for tuples from the input relations.
- Join: Join takes in two relations and outputs a relation that contains all the tuples that are a combination of two tuples where for each attribute that the two relations share, the values for that attribute of both tuples is the same.

## Structured Query Language

- Data Manipulation Language (DML): SELECT, INSERT, UPDATE, and DELETE statements.
- Data Definition Language (DDL): Schema definitions for tables, indexes, views, and other objects.
- Data Control Language (DCL): Security, access controls.

### Joins

- INNER JOIN（内连接,或等值连接）：获取两个表中字段匹配关系的记录。

- LEFT JOIN（左连接）：获取左表所有记录，即使右表没有对应匹配的记录。

- RIGHT JOIN（右连接）：与 LEFT JOIN 相反，用于获取右表所有记录，即使左表没有对应匹配的记录。

数据库在通过连接两张或多张表来返回记录时，都会生成一张中间的临时表，然后再将这张临时表返回给用户。

在使用left jion时，on和where条件的区别如下：

- on条件是在生成临时表时使用的条件，它不管on中的条件是否为真，都会返回左边表中的记录。


- where条件是在临时表生成好后，再对临时表进行过滤的条件。这时已经没有left join的含义（必须返回左边表的记录）了，条件不为真的就全部过滤掉。


- inner join 与 where查询效果和查询效率相同，如果在on中加入其他过滤信息，那么inner join效率更高，因为提前过滤了信息


- left join与where查询效果不同，先用on过滤，之后将过滤掉左表的信息添加回来，之后执行where，如果将条件写在where里面，就变成了inner join，损失了部分左表的信息

### Aggregates

Types of Aggregates Functions:

- AVG(COL): The average of the values in COL
- MIN(COL): The minimum value in COL
- MAX(COL): The maximum value in COL
- COUNT(COL): The number of tuples in the relation

DISTINCT keyword: select different results

Non-aggregated values in SELECT output clause must appear in GROUP BY clause.

The HAVING clause filters output results based on aggregation computation. This make HAVING behave like a WHERE clause for a GROUP BY.

### String

The SQL standard says that strings are case sensitive and single-quotes only.

- Pattern Matching: The LIKE keyword is used for string matching in predicates. '%' matches any substrings, '_' matches any one character
- Concatenation: Two vertical bars (“||”) will concatenate two or more strings together into a single string.
- String Functions: Many database systems implement other functions in addition to those in the standard. Examples of standard string functions include SUBSTRING(S, B, E) and UPPER(S).

### Output Control

- ORDER BY: impose a sort on tuples, ASC for ascending results, DESC for descending results
- LIMIT: restrict the number of result tuples, use OFFSET to return a range in the results

### Nested Query

- ALL: Must satisfy expression for all rows in sub-query.
- ANY: Must satisfy expression for at least one row in sub-query. 
- IN: Equivalent to =ANY().
- EXISTS: At least one row is returned.

### Window Function

- ROW NUMBER: assigns a number to each row in order, regardless of any ties
- RANK: assigns same rank to ties, and skips ahead to next number in position
- DENSE_RANK: assigns same rank to ties, but continues with the sequence
- PARTITION BY: Use PARTITION BY to specify group.
- ORDER BY: ORDER BY within OVER to ensure a deterministic ordering of results even if database changes internally.

### Common Table Expressions

CTEs can be thought of as a temporary table that is scoped to a single query.

WITH clause binds the output of the inner query to a temporary result with that name.

Adding the RECURSIVE keyword after WITH allows a CTE to reference itself.