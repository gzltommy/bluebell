package middleware

//import (
//	"bluebell/render"
//	"github.com/gin-gonic/gin"
//	"net/http"
//)
//
//// ErrorHandle 错误处理中间件
//// Context.go 的结构体提供了 一个 Errors 字段用以传递错误，并且作者也在 Error() 方法注释上建议用此方法来存储错误并调用中间件处理错误。
//// 我们要想用中间件来完成全局的错误处理，就应该将错误处理中间件放到执行链的最顶端；并且还要将错误处理的逻辑放到 c.Next() 后执行。
//// 使用参考：https://juejin.cn/post/7064770224515448840
//func ErrorHandle() func(c *gin.Context) {
//	return func(c *gin.Context) {
//		c.Next() // 先调用 c.Next() 执行后面的中间件
//		// 所有中间件及 router 处理完毕后从这里开始执行
//		// 检查 c.Errors 中是否有错误
//		for _, e := range c.Errors {
//			err := e.Err
//			// 若是自定义的错误则将 code、msg 返回
//			if reEr, ok := err.(*render.Error); ok {
//				c.JSON(http.StatusOK, reEr)
//			} else {
//				// 若非自定义错误则返回详细错误信息 err.Error()
//				// 比如 save session 出错时设置的 err
//				c.JSON(http.StatusOK, gin.H{
//					"code": 500,
//					"msg":  err.Error(),
//				})
//			}
//			return // 检查一个错误就行
//		}
//	}
//}
