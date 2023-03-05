# 功能
一个方便的单词，句子翻译工具，基本功能有：
- 基本的cobra命令行框架 ✅
- 从命令行读取要翻译的内容，同时输出其他2种语言的翻译（中/英/日）✅
- 翻译engine的聚合，保证请求的稳定程度
  - google ✅
  - mojo ✅
  - baidu
  - hujiang
  - deepl
- 指定翻译engine和翻译的模式（单词/句子）✅
- 控制是否从剪贴板读取要翻译的内容
- 自动判断输入语言的种类
- 设置可以把翻译结果缓存在本地，并且每次命令行Prompt可提示
- check单词的拼写正确，如果错误则可以提示类似的单词

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
7. 判断参数输入文本是中文，英文还是日文
8. [golang 中结构体匿名嵌套时的初始化问题](https://juejin.cn/post/7138428171224875038)

# 感想
1. go中的接口是**一组方法**的集合，不是某个单个方法
2. 这个接口是一种实现和使用相分离的机制，比如我想在上层用某个工具（因为这个工具能够提供某种功能），但是此时我还不知道他的实现
   1. eg: 使用xxx（这个时候**主体还不明确**，可以是翻译机提供的，也可以是翻译人员提供的）的翻译方法，那这个时候就可以这么用：`能够翻译的东西（接口名字）.翻译()`
   2. 等把上层的逻辑实现完成了后，再去写能够实现上面那个接口的主体和具体实现就好了，比如`谷歌翻译.翻译()`
   3. 这里还有个上层使用时具体主体如何实例化的问题，接受变量可以是模糊的借口，但是具体做事的那个主体需要是具体实例，可以用工厂方法在这里再分离逻辑领域一次
   4. 逻辑领域实体要素：工具的使用 ---- 工具的模糊定义 ---- 工具的实例化（设置工具的初始参数）---- 工具的实现