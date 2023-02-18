# tiktok_server
简易的抖音服务器

## 如何加入新的微服务：
1. 在**IDL**文件夹下创建新的服务文件夹，存放thrift文件
2. 在**cmd**文件夹下创建新的服务文件夹，存放可运行程序
3. 编写thrift文件
4. 在 第2步 中创建的文件夹下执行kitex工具生成代码(例如 ```kitex -module "github.com/PCBismarck/tiktok_server/cmd/user" -service a.b.c ../../IDL/userServer/user.thrift``` )
5. 在生成的main.go文件中编辑服务器地址和端口(提供本服务的地址端口)
6. 在生成的handler.go文件中完事业务逻辑

## 如何在网关中使用服务：
1. 在 **cmd/api/biz/rpc/** 下创建自己服务的rpc文件
2. 在该文件中需要初始化对应服务的 **service.client**
3. 通过对应的**client**访问微服务
4. 在 **cmd/api/biz/handler/** 中使用rpc文件中的函数
