* 初始文本：
```
20170102 admin,password Open
20170801 nmask,nmask close
20180902 nm4k,test filter
```

# awk
AWK是一种处理文本文件的语言，是一个强大的文本分析工具;awk是以列为划分计数的，$0表示所有列，$1表示第一列，$2表示第二列。

#### awk参数
* -F 指定输入文件折分隔符，如-F:
* -v 赋值一个用户定义变量，如-va=1
* -f 从脚本文件中读取awk命令
```
## 自定义分隔符

使用”,”进行分割，参数用-F
awk -F, '{print $1,$2}' test.log

使用多个分隔符，先使用空格分割，然后对分割结果再使用”,”分割
$ awk -F '[ ,]'  '{print $1,$2,$3}'  test.log  #注意逗号前面有一个空格


## awk逻辑判断

* 输出第一列为20170801的记录
# root @ xiye in ~/sed_awk_grep [15:39:10]
$ cat test.log| awk '$1==20170102 {print}'
20170102 admin,password Open

* 输出第二列不是nmask,nmask的记录
$ cat test.log| awk '$2!="nmask,nmask" {print}'
20170102 admin,password Open
20180902 nm4k,test filter

设置变量
设置awk自定义变量，用参数-v
例子：设置变量a为1
$ cat test.log | awk -va=1 '{print $1,$1+a}'
20170102 20170103
20170801 20170802
20180902 20180903


## awk内建函数

* NR，打印行号
$ cat test.log| awk '{print NR,$1}'
1 20170102
2 20170801
3 20180902

* substr字符串截取 截取第一列的第一到第四个字符
$ cat test.log| awk '{print substr($1,1,4)}'
2017
2017
2018

* split切分字符串 以逗号分隔第2列的数据，并输出分别输出第2列的内容
$ cat test.log| awk '{split($2,a,",");print a[1],a[2]}'
admin password
nmask nmask
nm4k test

* gsub替换 将第2列中的nmask替换成nMask
$ cat test.log| awk '{gsub("nmask","nMask",$2);print}'
20170102 admin,password Open
20170801 nMask,nMask close
20180902 nm4k,test filter

```


# grep
* Linux grep命令用于查找文件里符合条件的字符串。
* 用法： grep [选项]... PATTERN [FILE]...
  grep -cniv ‘关键字（正则）’ 文件路径
  选项
  -i 不区分大小写，**默认情况下grep不区分关键字大小写**
  -c 行数
  -n 显示关键词所在行号
  -v取反（使用较多）
  -r 遍历所有层级子目录
  -A后面跟数字，过滤出符合要求的行以及下面的n行
  -B后面跟数字，过滤符合要求的行以及上面的n行
  -C后面跟数字，过滤符合要求的行以及上下各n行
```
递归查询
grep -r nmask /etc/  #查看/etc目录下内容包含nmask的文件

查询取反
grep -v test  test.log

1.This is my cat, my cat's name is betty
2.This is my dog, my dog's name is frank
3.This is my fish, my fish's name is george
4.This is my goat, my goat's name is adam

# root @ xiye in ~/shell [22:23:48] 
$ grep -rn "fish" . -A 1 -B 1
./my.txt-2-2.This is my dog, my dog's name is frank
./my.txt:3:3.This is my fish, my fish's name is george
./my.txt-4-4.This is my goat, my goat's name is adam
```


# sed
Linux sed命令是利用script来处理文本文件。

#### 参数
* -e 以选项中指定的script来处理输入的文本文件。
* -f 以选项中指定的script文件来处理输入的文本文件。
* -h 显示帮助。
* -n 仅显示script处理后的结果。
* -V 显示版本信息。

#### 动作
* a ：新增， a 的后面可以接字串，而这些字串会在下一行出现
* i ：插入， i 的后面可以接字串，而这些字串会在上一行出现
* c ：取代， c 的后面可以接字串，这些字串可以取代 n1,n2 之间的行
* d ：删除
* s ：取代，通常这个s的动作可以搭配正规表示法！如 s/old/new/g

```
插入操作
在test.log文件的第4行后插入一行，内容为nmask
sed -e 4a\nmask test.log

删除操作
删除test.log的第2行、第3行数据	
cat test.log | sed '2,3d'

匹配删除，删除行中有nmask字符串的
nl test.log | sed '/nmask/d'

替换操作
sed -i 's/要被取代的字串/新的字串/g'

只要删除第 2 行
nl /etc/passwd | sed '2d' 

要删除第 3 到最后一行
nl /etc/passwd | sed '3,$d' 

```

