// package main

// import (
// 	api "./api" // 这里使用的是相对路径
// 	"github.com/labstack/echo"
// )

// func main() {
// 	e := echo.New()
// 	e.GET("/", api.HelloWorld)
// 	e.Logger.Fatal(e.Start(":1323"))
// }

package main

import (
	api "github.com/pingf/learn_go/api" // 这是更新后的引入方法

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", api.HelloWorld)
	e.Logger.Fatal(e.Start(":1323"))
}
