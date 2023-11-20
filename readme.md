# websocket + gin
go get -u github.com/gin-gonic/gin
go get github.com/gorilla/websocket


# 目录
- db
  - 用来放数据库里的内容
- msg
  - 用来放接口传来的类型
  - 放这里的原因是会和 const 变量名冲突。
- /
  - 方法，挤挤
  - gin相关
  - 之后还有 channal
绷不住了，目录怎么改都不好改是吧。
别动了。

- MarshalJson 里面用了 json.Marshal 会循环引用