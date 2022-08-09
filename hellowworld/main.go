package main

// package是最基本的分发单位，是工程管理中依赖关系的体现
// 要生成可执行文件，必须建立main包，同时必须有main函数

import "fmt"

// 导入依赖，不能打包没有用到的包，否则报错，与c的define有点类似

func main() {
	// 可执行程序起点，入口函数
	// main方法不能带参数也不能定义返回值，跟c有点类似
	// 命令行参数在os.Args中，命令行开关需要import flag包
	fmt.Println("hello world")
}

// func 函数名(参数列表)(返回值列表){
// 		函数体（支持多返回值）
//}
