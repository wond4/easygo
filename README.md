# easygo

使用 golang 原生语法需要大量的错误处理，为了减少代码量，计划把常用的一些功能抽出来，让 gopher 更 easy

## easyhttp
让原生 http 包更 easy

### body
对 response 的 body 做一层包装，方便转 json 、 string 等常用格式。

### request
计划实现常用的一些请求方式，比如 FormData 、 application/x-www-form-urlencoded 之类的请求方法