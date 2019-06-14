# snowflake

一个简单的分布式`ID`生成实现

目前是总共用了`53`位，前面`32`位用于时间戳，中间`8`位用于机器ID，最后的`13`位用于增长序列

所以最多可以有`256`台机器，一台机器一秒最多可以生成`8192`个`ID`

使用`53`位是为了配合前端使用的时候不需要转成字符串，因为`js`整数精度只能表示到`53`位，再多就会丢失精度，只能转换成字符串了。
大部分都是用在`web`上，并且也没有那么大的并发量，所以这样的设置也是够用的。

## 安装

```sh
> go get github.com/lujin123/snowflake
```
## 使用

```go
import (
    "github.com/lujin123/snowflake"
)

machineId := 1
snow := snowflake.New(machineId)
id := snow.NextID()
```

