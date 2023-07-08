# 使用
`fy trans word/sentence -t=google -sl=en -m=word/sentence`
# 功能
- 从命令行读取要翻译的内容，同时输出其他2种语言的翻译（中/英/日）✅
- 翻译engine的聚合，保证请求的稳定程度
  - zh -> en: google, zh -> ja: mojo
  - en -> zh: cambridge, en -> ja: google
  - ja -> zh: mojo, ja -> en: ?
- 指定翻译engine和翻译的模式（单词/句子）✅
- 自动判断输入语言的种类 ✅
- 设置可以把翻译结果缓存在本地，随机显示一条 ✅
- 支持配置
- 把日志换成zap，test换成testify
- 制作全屏幕划词翻译的工具

# 备注
1. 在项目根目录运行 `go build .`会在当前根目录下编译出可执行的文件（目标机器的可执行二进制文件）
2. `go clean -cache -i`
3. log.Fatal()&klog.Fatal()
4. 在go中如何一个返回值为多层嵌套的slice的请求结果中拿到想要的内容
   ```go
   func main() {
	str := `[[["but","但是",null,null,10]],null,"zh-CN",null,null,null,null,[]]`
	var data []interface{}
	if err := json.Unmarshal([]byte(str), &data); err != nil {
		fmt.Println("error:", err)
		return
	}
	data1 := data[0].([]interface{})[0].([]interface{})
	fmt.Println(data1[1].(string))
    }
   ```
5. [在 Golang 中使用 Cobra 创建 CLI 应用](https://www.qikqiak.com/post/create-cli-app-with-cobra/)
6. http请求含有中文参数导致乱码：[golang常用的http请求操作](https://cloud.tencent.com/developer/article/1515297)
8. [golang 中结构体匿名嵌套时的初始化问题](https://juejin.cn/post/7138428171224875038)
9. 把代码工具化，然后本地命令行工具式的使用
   - 不要编辑zsh里的那个PATH，那个每次source后会重复添加已有的PATH，会导致PATH出现很多重复的
   - 全局的话，建议修改`/etc/paths` or `/etc/bashrc` => https://www.jianshu.com/p/acb1f062a925
   - 设置好了PATH能找到`go/bin`后直接`go instal`就可以了
1.  但是;but;しかし%, 命令行的输出后面有%号的话，表示最后没有换行
2.  `type LangType string`给基本类型起别名的好处是1.让参数更有语义，2.可以限制传参数返参的类型，要不然的话只要是string都可以，相当于变相缩小domain了
    - 程序提供的基本数据类型相当于基本工具，编程实现具体业务需求的时候最好尽可能的包装成具体业务领域相关的字段，达到语义化和类型限定的作用

# 感想
1. go中的接口是**一组方法**的集合，不是某个单个方法
2. 这个接口是一种实现和使用相分离的机制，比如我想在上层用某个工具（因为这个工具能够提供某种功能），但是此时我还不知道他的实现
   - eg: 使用xxx（这个时候**主体还不明确**，可以是翻译机提供的，也可以是翻译人员提供的）的翻译方法，那这个时候就可以这么用：`能够翻译的东西（接口名字）.翻译()`
   - 等把上层的逻辑实现完成了后，再去写能够实现上面那个接口的主体和具体实现就好了，比如`谷歌翻译.翻译()`
   - 这里还有个上层使用时具体主体如何实例化的问题，接受变量可以是模糊的借口，但是具体做事的那个主体需要是具体实例，可以用工厂方法在这里再分离逻辑领域一次
   - 逻辑领域实体要素：工具的使用 ---- 工具的模糊定义 ---- 工具的实例化（设置工具的初始参数）---- 工具的实现
3. 可以利用接口实现部分继承机制，当需要指定一个类型为某几个类型之一时可以使用，可以使代码不会到处都是interface{}
4. 从业务外部过来的数据都是散的（流向ui接口层的数据），如何在传入之后的业务domain处理中把它包装成rich data是个要考虑的事情
   1. 数据流的逐步rich（类型 + 包装）的方式
   
