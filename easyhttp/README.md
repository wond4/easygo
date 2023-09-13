# easyhttp

使用 golang 原生 http 包处理 web 请求太麻烦（比如各种错误处理），所以把一些常用的方法抽出来做一个公共库，方便使用

## body
对 response 的 body 做一层包装，方便转 json 、 string 等常用格式。

## request
计划实现常用的一些请求方式，比如 application/x-www-form-urlencoded 之类的请求方法