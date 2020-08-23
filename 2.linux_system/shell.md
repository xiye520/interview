#排查问题常用shell

* [Linux下进程信息/proc/pid/status的深入分析](https://blog.csdn.net/beckdon/article/details/48491909)
* []()
* []()
* []()
* []()
* []()
* []()

通过： ps aux | sort -k4,4nr | head -n 10 查看内存占用前10名的程序

查看进程的status文件： `cat /proc/2913/status` VmRSS对应的值就是物理内存占用


