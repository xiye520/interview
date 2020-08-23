常用shell收集

* [学习linux命令，看这篇2w多字的命令详解就够了](https://mp.weixin.qq.com/s/7bSwKiPmtJbs7FtRWZZqpA)
* [linux常用命令](http://man.linuxde.net/cp)
* 

go在windows系统中交叉编译linux可执行文件

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o test test.go

## 1.显示服务器上所有的每个IP多少个连接数
```
netstat -ntu | awk '{print $5}' | cut -d: -f1 | sort | uniq -c | sort -n
```
可编写脚本自动提取攻击ip然后自动屏蔽：
```
*/2 * * * * /usr/local/nginx/var/log/drop.sh
#!/bin/sh
cd /usr/local/nginx/var/log
tail access.log -n 1000 |grep vote.php | |sort |uniq -c |sort -nr |awk '{if ($2!=null && $1>50)}' > drop_ip.txt
for i in `cat drop_ip.txt`
do
/sbin/iptables -I INPUT -s $i -j DROP;
done
```

## 2.Screen命令

创建会话：              screen -S

离开当前session：       Ctrl -a d

恢复会话：              screen -r

将指定作业离线：        screen -d  　

列出当前所有的session： screen -ls



## 3.查看宿主端口监听状态：

$ netstat -tunpl

(No info could be read for "-p": geteuid()=500 but you should be root.)

Active Internet connections (only servers)

Proto Recv-Q Send-Q Local Address               Foreign Address             State       PID/Program name

tcp        0      0 0.0.0.0:29173               0.0.0.0:*                   LISTEN      -

tcp        0      0 127.0.0.1:25                0.0.0.0:*                   LISTEN      -

tcp        0      0 :::29173                    :::*                        LISTEN      -

tcp        0      0 :::9527                     :::*                        LISTEN      -

tcp        0      0 ::1:25                      :::*                        LISTEN      -

udp        0      0 0.0.0.0:52744               0.0.0.0:*                               -

udp        0      0 0.0.0.0:68                  0.0.0.0:*                               -

udp        0      0 0.0.0.0:45421               0.0.0.0:*                               -

udp        0      0 :::9527                     :::*                                    -


netstat命令可以显示网络接口的很多统计信息，包括打开的socket和路由表。无选项运行命令显示打开的socket

这条命令还有很多功能。比如，netstat -p命令可以显示打开的socket对应的程序。

netstat -s则显示所有端口的详细统计信息。

## 4.tar 

### 4.1 tar使用的选项有：

-c — 创建一个新归档。

-f — 当与 -c 选项一起使用时，创建的 tar 文件使用该选项指定的文件名；当与 -x 选项一起使用时，则解除该选项指定的归档。

-t — 显示包括在 tar 文件中的文件列表。

-v — 显示文件的归档进度。

-x — 从归档中抽取文件。

-z — 使用 gzip 来压缩 tar 文件。

-j — 使用 bzip2 来压缩 tar 文件。

-r：向压缩归档文件末尾追加文件

-u：更新原压缩包中的文件

\# tar -cf all.tar *.jpg

这条命令是将所有.jpg的文件打成一个名为all.tar的包。-c是表示产生新的包，-f指定包的文件名。



\# tar -rf all.tar *.gif

这条命令是将所有.gif的文件增加到all.tar的包里面去。-r是表示增加文件的意思。



\# tar -uf all.tar logo.gif

这条命令是更新原来tar包all.tar中logo.gif文件，-u是表示更新文件的意思。



\# tar -tf all.tar

这条命令是列出all.tar包中所有文件，-t是列出文件的意思



\# tar -xf all.tar

这条命令是解出all.tar包中所有文件，-x是解开的意思



### 4.2 压缩

tar –cvf jpg.tar *.jpg  将目录里所有jpg文件打包成jpg.tar

tar –czf jpg.tar.gz *.jpg   将目录里所有jpg文件打包成jpg.tar后，并且将其用gzip压缩，生成一个gzip压缩过的包，命名为jpg.tar.gz

tar –cjf jpg.tar.bz2 *.jpg 将目录里所有jpg文件打包成jpg.tar后，并且将其用bzip2压缩，生成一个bzip2压缩过的包，命名为jpg.tar.bz2

tar –cZf jpg.tar.Z *.jpg   将目录里所有jpg文件打包成jpg.tar后，并且将其用compress压缩，生成一个umcompress压缩过的包，命名为jpg.tar.Z

rar a jpg.rar *.jpg rar格式的压缩，需要先下载rar for linux

zip jpg.zip *.jpg   zip格式的压缩，需要先下载zip for linux

### 4.3 解压

tar –xvf file.tar  解压 tar包

tar -xzvf file.tar.gz 解压tar.gz

tar -xjvf file.tar.bz2   解压 tar.bz2

tar –xZvf file.tar.Z   解压tar.Z

unrar e file.rar 解压rar

unzip file.zip 解压zip



### 4.4 总结

*.tar 用 tar –xvf 解压

*.gz 用 gzip -d或者gunzip 解压

*.tar.gz和*.tgz 用 tar –xzf 解压

*.bz2 用 bzip2 -d或者用bunzip2 解压

*.tar.bz2用tar –xjf 解压

*.Z 用 uncompress 解压

*.tar.Z 用tar –xZf 解压

*.rar 用 unrar e解压

*.zip 用 unzip 解压



## 5.ps 

ps 的参数非常多, 在此仅列出几个常用的参数并大略介绍含义

-A 列出所有的行程

-w 显示加宽可以显示较多的资讯

-au 显示较详细的资讯

-aux 显示所有包含其他使用者的行程



## 6.vim

VIM退出命令是，按ESC键 跳到命令模式，然后输入:q（不保存）或者:wq（保存） 退出。

更多退出命令：

:w 保存文件但不退出vi

:w file 将修改另外保存到file中，不退出vi

:w! 强制保存，不推出vi

:wq 保存文件并退出vi

:wq! 强制保存文件，并退出vi

:q 不保存文件，退出vi

:q! 不保存文件，强制退出vi

:e! 放弃所有修改，从上次保存文件开始再编辑命令历史



## 7.telnet端口探测


```
npm部署安装包：
npm -i telnet-0.17-59.el7.x86_64.rpm

telnet测试端口情况：
telnet ip port

退出telnet 命令：
正确的命令 ctrl+]  然后在telnet 命令行输入 quit  就可以退出了
```

查看某一进程的资源占用：

```
top -p pid
```

