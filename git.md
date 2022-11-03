## tips

* [Git新手教程-查看仓库的历史记录（四）](https://juejin.im/post/5da0a4b4e51d45784840b6b3)
* [Git 基础 - 查看提交历史]([https://git-scm.com/book/zh/v2/Git-%E5%9F%BA%E7%A1%80-%E6%9F%A5%E7%9C%8B%E6%8F%90%E4%BA%A4%E5%8E%86%E5%8F%B2](https://git-scm.com/book/zh/v2/Git-基础-查看提交历史))
* [Github进行fork后如何与原仓库同步](https://github.com/selfteaching/the-craft-of-selfteaching/issues/67)
* 



```
github上查看一个人评论过哪些issue 
commenter:username

git查看本地分支关联（跟踪）的远程分支之间的对应关系，本地分支对应哪个远程分支
git branch -vv
```

## 设置生成ssh地址
```
git config --global user.name "xiye"
git config --global user.email "xxx@gmail.com"
ssh-keygen -t rsa -C "xxx@gmail.com"
eval "$(ssh-agent -s)"
ssh-add ~/.ssh/id_rsa
```


## 一个项目配置多个git仓库
    查看当前remote的信息
    git remote -v 
    
    --删除某个远程地址
    git remote remove name
    
    添加一个新的远程地址（即一个项目设置多个远程地址）
    git remote add github git@github.com:xxx/leetcode.git
    
    git push -u github master


## 自动pull所有分支脚本

    #!/bin/sh
    # Usage: fetchall.sh branch ...
    
    set -x
    git fetch --all
    for branch in "$@"; do
        git checkout "$branch"      || exit 1
        git pull "origin/$branch" || exit 1
    done
    
    也可以用下面的：
    git branch | awk 'BEGIN{print "echo ****Update all local branch...@daimon***"}{if($1=="*"){current=substr($0,3)};print a"git checkout "substr($0,3);print "git pull --all";}END{print "git checkout " current}' |sh

## git log查看日志

```
# root @ xiye in ~/gopath/src/interview on git:master x [15:37:08] 
$ git log --oneline
2368a0f 补充红黑树、redis相关文档
b05958c add restfulApi、channel 文档
f66b42b 补充mysql、https相关文档

展示某个commit更改了哪个或哪些文件，stat 是单词 statistics ,为统计的意思。
$ git log --stat

其中一个比较有用的选项是 -p 或 --patch ，它会显示每次提交所引入的差异（按 补丁 的格式输出）。 你也可以限制显示的日志条目数量，例如使用 -2 选项来只显示最近的两次提交：

$ git log -p -2

```

`git log` 的常用选项:

| 选项              | 说明                                                         |
| :---------------- | :----------------------------------------------------------- |
| `-p`              | 按补丁格式显示每个提交引入的差异。                           |
| `--stat`          | 显示每次提交的文件修改统计信息。                             |
| `--shortstat`     | 只显示 --stat 中最后的行数修改添加移除统计。                 |
| `--name-only`     | 仅在提交信息后显示已修改的文件清单。                         |
| `--name-status`   | 显示新增、修改、删除的文件清单。                             |
| `--abbrev-commit` | 仅显示 SHA-1 校验和所有 40 个字符中的前几个字符。            |
| `--relative-date` | 使用较短的相对时间而不是完整格式显示日期（比如“2 weeks ago”）。 |
| `--graph`         | 在日志旁以 ASCII 图形显示分支与合并历史。                    |
| `--pretty`        | 使用其他格式显示历史提交信息。可用的选项包括 oneline、short、full、fuller 和 format（用来定义自己的格式）。 |
| `--oneline`       | `--pretty=oneline --abbrev-commit` 合用的简写。              |





## git diff命令

查看缓存区和上次提交的不同

查看已暂存起来的变化 git diff --cached

git diff --staged 可以显示暂存区（git add 指令作用的地方）和上一条提交之间的不同

对比工作目录和缓存区

git diff 可是显示工作目录和暂存区之间的不同

对比工作目录和上一条提交

git diff HEAD 可以显示工作目录和上一条提交之间不同，就是说：如果你现在把所有的文件都 add 然后 git commit，你将会提交什么
