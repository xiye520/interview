DML（data manipulation language）：bai 它们是SELECT、UPDATE、INSERT、DELETE，就象它的名字一du样，这4条命令是用来zhi对数据库里的数据进行dao操作的语言 

DDL（data definition language）： DDL比DML要多，主要的命令有CREATE、ALTER、DROP等，DDL主要是用在定义或改变表（TABLE）的结构，数据类型，表之间的链接和约束等初始化工作上，他们大多在建立表时使用 

DCL（Data Control Language）： 是数据库控制功能。是用来设置或更改数据库用户或角色权限的语句，包括（grant,deny,revoke等）语句。在默认状态下，只有sysadmin,dbcreator,db_owner或db_securityadmin等人员才有权力执行DCL

TCL - Transaction Control Language：事务控制语言，COMMIT - 保存已完成的工作，SAVEPOINT - 在事务中设置保存点，可以回滚到此处，ROLLBACK - 回滚，SET TRANSACTION - 改变事务选项