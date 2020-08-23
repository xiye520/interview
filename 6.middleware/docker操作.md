DOCKER其他参数说明:



Attach--将终端依附到容器上

Build--通过Dockerfile创建镜像

Commit--通过容器创建本地镜像

Cp--在宿主机和容器之间相互COPY文件

Create--创建一个新的容器，注意，此时，容器的status只是Created

Diff--查看容器内发生改变的文件，以我的mysql容器为例

Events--实时输出Docker服务器端的事件，包括容器的创建，启动，关闭等。

Exec--用于容器启动之后，执行其它的任务,也可以创建容器,例如:docker exec -ti cc /bin/bash

Export--将容器的文件系统打包成tar文件

History--显示镜像制作的过程，相当于dockfile

Images--列出本机的所有镜像

Import--根据tar文件的内容新建一个镜像，与之前的export命令相对应

Info--查看docker的系统信息

Inspect--用于查看容器的配置信息，包含容器名、环境变量、运行命令、主机配置、网络配置和数据卷配置等。

Kill--强制终止容器

load --与下面的save命令相对应，将下面sava命令打包的镜像通过load命令导入

Login--登录到自己的Docker register，需有Docker Hub的注册账号

Logout--退出登录

Logs--用于查看容器的日志，它将输出到标准输出的数据作为日志输出到docker logs命令的终端上。常用于后台型容器

Pause--暂停容器内的所有进程，

Port--输出容器端口与宿主机端口的映射情况

Ps--列出所有容器，其中docker ps用于查看正在运行的容器，ps -a则用于查看所有容器。

Pull--从docker hub中下载镜像

Push--将本地的镜像上传到docker hub中

Rename--更改容器的名字

Restart--重启容器

rm --删除容器，注意，不可以删除一个运行中的容器，必须先用docker stop或docker kill使其停止。

Rmi--删除镜像

Run--让创建的容器立刻进入运行状态，该命令等同于docker create创建容器后再使用docker start启动容器

Save--将镜像打包，与上面的load命令相对应

Search--从Docker Hub中搜索镜像

Start--启动容器

Stats--动态显示容器的资源消耗情况，包括：CPU、内存、网络I/O

Stop--停止一个运行的容器

Tag--对镜像进行重命名

Top--查看容器中正在运行的进程

Unpause--恢复容器内暂停的进程，与pause参数相对应

Version--查看docker的版本

Wait--捕捉容器停止时的退出码



 





六、image 文件

\# 列出本机的所有 image 文件。

$ docker image ls



\# 删除 image 文件

$ docker image rm [imageName]



七、实例：hello world

$ docker container run hello-world

docker container run命令会从 image 文件，生成一个正在运行的容器实例。

注意，docker container run命令具有自动抓取 image 文件的功能。如果发现本地没有指定的 image 文件，就会从仓库自动抓取



有些容器不会自动终止，因为提供的是服务。比如，安装运行 Ubuntu 的 image，就可以在命令行体验 Ubuntu 系统。

$ docker container run -it ubuntu bash

对于那些不会自动终止的容器，必须使用docker container kill 命令手动终止。

$ docker container kill [containID]



八、容器文件

\# 列出本机正在运行的容器

$ docker container ls



\# 列出本机所有容器，包括终止运行的容器

$ docker container ls --all



终止运行的容器文件，依然会占据硬盘空间，可以使用docker container rm命令删除。

$ docker container rm [containerID]



10.3 生成容器

docker container run命令会从 image 文件生成容器。





$ docker container run -p 8000:3000 -it koa-demo /bin/bash

\# 或者

$ docker container run -p 8000:3000 -it koa-demo:0.0.1 /bin/bash

上面命令的各个参数含义如下：



-p参数：容器的 3000 端口映射到本机的 8000 端口。

-it参数：容器的 Shell 映射到当前的 Shell，然后你在本机窗口输入的命令，就会传入容器。

koa-demo:0.0.1：image 文件的名字（如果有标签，还需要提供标签，默认是 latest 标签）。

/bin/bash：容器启动以后，内部第一个执行的命令。这里是启动 Bash，保证用户可以使用 Shell。



docker container run命令的--rm参数，在容器终止运行后自动删除容器文件。



十一、其他有用的命令

（1）docker container start

前面的docker container run命令是新建容器，每运行一次，就会新建一个容器。同样的命令运行两次，就会生成两个一模一样的容器文件。如果希望重复使用容器，就要使用docker container start命令，它用来启动已经生成、已经停止运行的容器文件。

$ docker container start [containerID]



（2）docker container stop

前面的docker container kill命令终止容器运行，相当于向容器里面的主进程发出 SIGKILL 信号。而docker container stop命令也是用来终止容器运行，相当于向容器里面的主进程发出 SIGTERM 信号，然后过一段时间再发出 SIGKILL 信号。

$ bash container stop [containerID]

这两个信号的差别是，应用程序收到 SIGTERM 信号以后，可以自行进行收尾清理工作，但也可以不理会这个信号。如果收到 SIGKILL 信号，就会强行立即终止，那些正在进行中的操作会全部丢失。



（3）docker container logs

docker container logs命令用来查看 docker 容器的输出，即容器里面 Shell 的标准输出。如果docker run命令运行容器的时候，没有使用-it参数，就要用这个命令查看输出。

$ docker container logs [containerID]



（4）docker container exec

docker container exec命令用于进入一个正在运行的 docker 容器。如果docker run命令运行容器的时候，没有使用-it参数，就要用这个命令进入容器。一旦进入了容器，就可以在容器的 Shell 执行命令了。

$ docker container exec -it [containerID] /bin/bash



（5）docker container cp

docker container cp命令用于从正在运行的 Docker 容器里面，将文件拷贝到本机。下面是拷贝到当前目录的写法。

$ docker container cp [containID]:[/path/to/file] .





查看正在运行的所有容器



[root@localhost ~]# docker ps



2，docker exec -it 容器ID /bin/bash 进入某个容器

[root@localhost ~]# docker exec -it 63be1d22a99e /bin/bash

[root@63be1d22a99e /]# 





# docker挂载卷启动一个nginx服务

```
# root @ xiye in ~/workspace [17:51:15] 
$ docker run -it -v /root/workspace:/usr/Downloads --name volume ubuntu  /bin/bash 

进入容器内部：
root@2cbf51634913:/# pwd
/
第一步 更新apt-get
root@2cbf51634913:/# apt-get update

第二步 安装nginx：
root@2cbf51634913:/# apt-get install nginx

第三步  安装vim：
root@2cbf51634913:/# apt-get install -y vim

第四步修改nginx配置
root@2cbf51634913:/# whereis nginx
nginx: /usr/sbin/nginx /etc/nginx /usr/share/nginx
root@2cbf51634913:/# ls /etc/nginx/
conf.d  fastcgi.conf  fastcgi_params  koi-utf  koi-win  mime.types  nginx.conf  proxy_params  scgi_params  sites-available  sites-enabled  snippets  uwsgi_params  win-utf
root@2cbf51634913:/# vim /etc/nginx/sites-enabled/default 
编辑修改index.html的地址，配置端口

root@2cbf51634913:/# nginx
查看进程
root@2cbf51634913:/# ps -ef
UID        PID  PPID  C STIME TTY          TIME CMD
root         1     0  0 09:51 ?        00:00:00 /bin/bash
root       872     1  0 10:15 ?        00:00:00 nginx: master process nginx
www-data   873   872  0 10:15 ?        00:00:00 nginx: worker process
www-data   874   872  0 10:15 ?        00:00:00 nginx: worker process
root      3501     1  0 10:18 ?        00:00:00 ps -ef
```