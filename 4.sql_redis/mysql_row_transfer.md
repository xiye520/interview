## mysql行转列

建表语句
```
CREATE table test
(id int,name nvarchar(20),quarter int,number int);
insert into test values(1,N'苹果',1,1000)        ;
insert into test values(1,N'苹果',2,2000)        ;
insert into test values(1,N'苹果',3,4000)        ;
insert into test values(1,N'苹果',4,5000)        ;
insert into test values(2,N'梨子',1,3000)        ;
insert into test values(2,N'梨子',2,3500)        ;
insert into test values(2,N'梨子',3,4200)        ;
insert into test values(2,N'梨子',4,5500)        ;
select * from test                               ;

mysql> select * from test;
+------+--------+---------+--------+
| id   | name   | quarter | number |
+------+--------+---------+--------+
|    1 | 苹果   |       1 |   1000 |
|    1 | 苹果   |       2 |   2000 |
|    1 | 苹果   |       3 |   4000 |
|    1 | 苹果   |       4 |   5000 |
|    2 | 梨子   |       1 |   3000 |
|    2 | 梨子   |       2 |   3500 |
|    2 | 梨子   |       3 |   4200 |
|    2 | 梨子   |       4 |   5500 |
+------+--------+---------+--------+
```

转换：
```
SELECT 
 id,
 name,
 SUM(CASE WHEN quarter=1 THEN number ELSE 0 END) '一季度',
 SUM(CASE WHEN quarter=2 THEN number ELSE 0 END) '二季度',
 SUM(CASE WHEN quarter=3 THEN number ELSE 0 END) '三季度',
 SUM(CASE WHEN quarter=4 THEN number ELSE 0 END) '四季度'
FROM test
GROUP BY id,name;

mysql> SELECT 
    ->  id,
    ->  name,
    ->  SUM(CASE WHEN quarter=1 THEN number ELSE 0 END) '一季度',
    ->  SUM(CASE WHEN quarter=2 THEN number ELSE 0 END) '二季度',
    ->  SUM(CASE WHEN quarter=3 THEN number ELSE 0 END) '三季度',
    ->  SUM(CASE WHEN quarter=4 THEN number ELSE 0 END) '四季度'
    -> FROM test
    -> GROUP BY id,name;
+------+--------+-----------+-----------+-----------+-----------+
| id   | name   | 一季度    | 二季度    | 三季度    | 四季度    |
+------+--------+-----------+-----------+-----------+-----------+
|    1 | 苹果   |      1000 |      2000 |      4000 |      5000 |
|    2 | 梨子   |      3000 |      3500 |      4200 |      5500 |
+------+--------+-----------+-----------+-----------+-----------+
2 rows in set (0.00 sec)

```



