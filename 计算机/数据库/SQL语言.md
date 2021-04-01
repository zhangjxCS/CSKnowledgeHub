## SQL语法

### 查询规范

1. 充分利用表上已经存在的索引,避免使用双 % 号的查询条件。如 a like '%123%'，（如果无前置 %，只有后置 %，是可以用到列上的索引的）一个 SQL 只能利用到复合索引中的一列进行范围查询，如：有 a,b,c 列的联合索引，在查询条件中有 a 列的范围查询，则在 b,c 列上的索引将不会被用到，在定义联合索引时，如果 a 列要用到范围查找的话，就要把 a 列放到联合索引的右侧；使用 left join 或 not exists 来优化 not in 操作， 因为 not in 也通常会使用索引失效；
2. 禁止使用 SELECT * 必须使用 SELECT <字段列表> 查询；
3. 避免使用子查询，可以把子查询优化为 JOIN 操作；
4. 避免使用 JOIN 关联太多的表。

### join连接

1. INNER JOIN（内连接,或等值连接）：获取两个表中字段匹配关系的记录。
2. LEFT JOIN（左连接）：获取左表所有记录，即使右表没有对应匹配的记录。
3. RIGHT JOIN（右连接）：与 LEFT JOIN 相反，用于获取右表所有记录，即使左表没有对应匹配的记录。

MySQL 使用了嵌套循环（Nested-Loop Join）的实现方式。Nested-Loop Join 需要区分驱动表和被驱动表，先访问驱动表，筛选出结果集，然后将这个结果集作为循环的基础，访问被驱动表过滤出需要的数据。Nested-Loop Join 分下面几种类型：

1.SNLJ，简单嵌套循环。这是最简单的方案，性能也一般。实际上就是通过驱动表的结果集作为循环基础数据，然后一条一条的通过该结果集中的数据作为过滤条件到下一个表中查询数据，然后合并结果。如果还有第三个参与 Join，则再通过前两个表的 Join 结果集作为循环基础数据，再一次通过循环查询条件到第三个表中查询数据，如此往复。

![图片](https://mmbiz.qpic.cn/mmbiz_jpg/j3gficicyOvaufKISYMVA6xk1p9XpneMNjQfBpDaMUiadkYynibYCats5l27FC4nn8PQBGM52td3MCqb1xnokvJsoQ/640?wx_fmt=jpeg&wxfrom=5&wx_lazy=1&wx_co=1)

这个算法相对来说就是很简单了，从驱动表中取出 R1 匹配 S 表所有列，然后 R2，R3,直到将 R 表中的所有数据匹配完，然后合并数据，可以看到这种算法要对 S 表进行 RN 次访问，虽然简单，但是相对来说开销还是太大了

2.INLJ，索引嵌套循环。索引嵌套联系由于非驱动表上有索引，所以比较的时候不再需要一条条记录进行比较，而可以通过索引来减少比较，从而加速查询。这也就是平时我们在做关联查询的时候必须要求关联字段有索引的一个主要原因。

![图片](https://mmbiz.qpic.cn/mmbiz_jpg/j3gficicyOvaufKISYMVA6xk1p9XpneMNjIoXSJee6o9WBibHQrkvsyHtOkSM9VXf9xZmjxiawJgEnPEr3aicFRHWTQ/640?wx_fmt=jpeg&wxfrom=5&wx_lazy=1&wx_co=1)

3.BNLJ，块嵌套循环。如果 join 字段没索引，被驱动表需要进行扫描。这里 MySQL 并不会简单粗暴的应用前面算法，而是加入了 buffer 缓冲区，降低了内循环的个数，也就是被驱动表的扫描次数。

![图片](https://mmbiz.qpic.cn/mmbiz_jpg/j3gficicyOvaufKISYMVA6xk1p9XpneMNjutXPdYZukU6K4R14OHCLU5cTPp15DcaxJ86DINdddhvIBceAzgn6Vw/640?wx_fmt=jpeg&wxfrom=5&wx_lazy=1&wx_co=1)

这个 buffer 被称为 join buffer，顾名思义，就是用来缓存 join 需要的字段。MySQL 默认 buffer 大小 256K，如果有 n 个 join 操作，会生成 n-1 个 join buffer。

### join的优化

1. 小结果集驱动大结果集。用数据量小的表去驱动数据量大的表，这样可以减少内循环个数，也就是被驱动表的扫描次数。
2. 用来进行 join 的字段要加索引，会触发 INLJ 算法，如果是主键的聚簇索引，性能最优。例子：第一个子查询是 72075 条数据，join 的第二条子查询是 50w 数据，主要的优化还是驱动表是小表，后面的是大表，on 的条件加上了唯一索引。
3. 如果无法使用索引，那么注意调整 join buffer 大小，适当调大些
4. 减少不必要的字段查询（字段越少，join buffer 所缓存的数据就越多）

## SQL执行过程

1. 客户端发送一条查询给服务器。服务器先检查查询缓存，如果命中了缓存，则立刻返回存储在缓存中的结果。否则进入下一阶段；
2. 在解析一个查询语句之前，如果查询缓存是打开的，那么 MYSQL 会优先检查这个查询是否命中查询缓存中的数据；
3. 这个检查是通过一个对大小写敏感的哈希查找的。查询和缓存中的查询即使只有一个不同，也不会匹配缓存结果；
4. 如果命中缓存，那么在但会结果前 MySQL 会检查一次用户权限，有权限则跳过其他步骤直接返回数据；
5. 服务器端进行 SQL 解析、预处理，再由优化器生成对应的执行计划。MySQL 解析器将使用 MySQL 语法规则验证和解析查询。例如验证是否使用错误的关键字、关键字顺序、引号前后是否匹配等；预处理器则根据一些 MySQL 规则进一步解析树是否合法，例如检查数据表和数据列是否存在，解析名字和别名是否有歧义等；
6. MySQL 根据优化器生成的执行计划，再调用存储引擎的 API 来执行查询。

