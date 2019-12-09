# Server

该仓库是博客的后端服务器项目，实现了项目文档中的API功能，同时包含一个已经装载好真实数据的boltDB实例数据库。

## 文件说明

```
| -- smo.db //由boltDB生成的数据库实例，已经装载好真实数据
| -- main.go //启动服务器的主程序
| -- addData.go //将真实数据装入数据库的程序(server不需要使用到)
| -- LICENSE 
| -- README.md 
| -- go // 基于swagger实现的所有API的go代码
| -- data // 爬去CSDN获取的部分真实数据 (json格式存储)，被装入到smo.db中
| -- images
```

## API接口测试

由于API的设计是利用swagger的，因此服务端的API测试采用Easy-mock来测试。Easy-mock是一款测试服务端API的在线mock测试工具。可以通过导入swagger的OAS yaml文件来生成对应的测试接口，快速进行API的测试。

