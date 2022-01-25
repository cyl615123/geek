##socket

`1.总结几种 socket 粘包的解包方式：fix length/delimiter based/length field based frame decoder。尝试举例其应用。`
~~~~
fix length：固定长度

delimiter based：基于分隔符
应用：
http协议事使用\r\n划分head，body
http chunk模式下使用\r\n划分chuck

length field based frame decoder：定义长度字段
应用：
http携带content-length头
protobuffer等二进制协议

~~~~

`2.实现一个从 socket connection 中解码出 goim 协议的解码器。`
~~~~

~~~~