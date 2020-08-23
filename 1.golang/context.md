#go context
#### 参考链接
* [Golang中的context包详解](https://blog.csdn.net/why444216978/article/details/105209237)
* [深度解密Go语言之context](https://www.cnblogs.com/qcrao-2018/p/11007503.html)
* [源码分析context的超时及关闭实现](http://xiaorui.cc/archives/5604)
* []()

## golang控制并发有两种经典的方式，一种是 WaitGroup，另外一种就是 Context

#### 1.WaitGroup
```
func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("1号完成")
		wg.Done()
	}()
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("2号完成")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("好了，大家都干完了，放工")
}
```
一个很简单的例子，一定要例子中的 2 个 goroutine 同时做完，才算是完成，先做好的就要等着其他未完成的，所有的 goroutine 要都全部完成才可以。

这是一种控制并发的方式，这种尤其适用于，好多个 goroutine 协同做一件事情的时候，因为每个 goroutine 做的都是这件事情的一部分，只有全部的 goroutine 都完成，这件事情才算是完成，这是等待的方式。

#### 2.chan 通知
我们都知道一个 goroutine 启动后，我们是无法控制它的，大部分情况是等待它自己结束，那么如果这个 goroutine 是一个不会自己结束的后台 goroutine 呢？比如监控等，会一直运行的。

这种情况化，一直傻瓜式的办法是全局变量，其它地方通过修改这个变量完成结束通知，然后后台 goroutine 不停的检查这个变量，如果发现被通知关闭了，就自我结束。

这种方式也可以，但是首先我们要保证这个变量在多线程下的安全，基于此，有一种更好的方式：chan + select 。
```
func main() {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("监控退出，停止了...")
				return
			default:
				fmt.Println("goroutine监控中...")
				time.Sleep(2 * time.Second)
			}
		}
	}()

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	stop <- true
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}
```

#### 3.Context控制多个goroutine(context内部也是通过channel实现的goroutine通信功能)
使用 Context 控制一个 goroutine 的例子如上，非常简单，下面我们看看控制多个 goroutine的例子，其实也比较简单。

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go watch(ctx, "【监控1】")
	go watch(ctx, "【监控2】")
	go watch(ctx, "【监控3】")

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "监控退出，停止了...")
			return
		default:
			fmt.Println(name, "goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}
输出结果



在官方博客里，对于使用 context 提出了几点建议：

1. Do not store Contexts inside a struct type; instead, pass a Context explicitly to each function that needs it. The Context should be the first parameter, typically named ctx.
2. Do not pass a nil Context, even if a function permits it. Pass context.TODO if you are unsure about which Context to use.
3. Use context Values only for request-scoped data that transits processes and APIs, not for passing optional parameters to functions.
4. The same Context may be passed to functions running in different goroutines; Contexts are safe for simultaneous use by multiple goroutines.

我翻译一下：

1. 不要将 Context 塞到结构体里。直接将 Context 类型作为函数的第一参数，而且一般都命名为 ctx。
2. 不要向函数传入一个 nil 的 context，如果你实在不知道传什么，标准库给你准备好了一个 context：todo。
3. 不要把本应该作为函数参数的类型塞到 context 中，context 存储的应该是一些共同的数据。例如：登陆的 session、cookie 等。
4. 同一个 context 可能会被传递到多个 goroutine，别担心，context 是并发安全的。
