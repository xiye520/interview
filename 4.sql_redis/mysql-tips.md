# Tips-mysql

* [MYSQL 设置无密码登录](https://blog.csdn.net/qq_42142258/article/details/96319590)
* [MySQL指定IP用户访问数据库](https://blog.csdn.net/u010814849/article/details/52813361)
* [mysql-行锁的实现](https://blog.csdn.net/alexdamiao/article/details/52049993?utm_medium=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-1.nonecase&depth_1-utm_source=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-1.nonecase)
* []()
* []()
* []()

```
GRANT ALL PRIVILEGES ON *.* TO 'xiye'@'127.0.0.1' IDENTIFIED BY '123456' WITH GRANT OPTION;
flush privileges;

GRANT ALL PRIVILEGES ON *.* TO 'root'@'172.17.0.1' IDENTIFIED BY '123456' WITH GRANT OPTION;
flush privileges;


GRANT ALL PRIVILEGES ON *.* TO 'xiye'@'127.0.0.1' IDENTIFIED BY '123456' WITH GRANT OPTION;
flush privileges;


update mysql.user set authentication_string = password("123456") where user = "root" and host = "172.17.0.1";
flush privileges;
```

## innodb行锁实现方式

**InnoDB行锁是通过给索引上的索引项加锁来实现的，这一点MySQL与Oracle不同，后者是通过在数据块中对相应数据行加锁来实现的。** InnoDB这种行锁实现特点意味着：只有通过索引条件检索数据，InnoDB才使用行级锁，否则，InnoDB将使用表锁！