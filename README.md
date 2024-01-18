# Zin

### 基于Gin上封装的Http接口接口开发框架

### Dependencies
1. wire
> go install github.com/google/wire/cmd/wire@latest

### Install
> git clone https://github.com/jiebutech/zin.git youProjectname

### 初始化
### 生成实例化构建文件
> cd wire && wire gen  
> 或者 make wire
#### 安装依赖包
1. 删除原有go.mod (可选, 不重置的化， 对应项目的module为`github.com/jiebutech/zin`)
> rm -f go.*

2. 拉取依赖
> go mod init  
> go mod tidy

#### 在main.go中修改对应路由和配置文件的引用
> rename github.com/jiebutech/zin/conf => youConfigPkg  
> rename github.com/jiebutech/zin/route => youRoutePkg


#### 配置文件初始化
> make config env(dev/test/prod)

#### 启动
> make run

### 相关目录模块概述
+ wire 相关实例的依赖注入和实例化操作 
+ route 业务路由注册
+ conf 配置文件及其数据结构定义
+ app 相关业务接口实现
  + api 控制器层接口实现， 即Http handler的实现
  + dto api接口出入参的数据结构定义
  + entity 数据库结构的定义
  + repo 数据库相关操作的定义, (所有的repo都要用interface封装一层， 以免需要更换数据存储层)
  + service 业务逻辑的定义

### 相关规范
+ 业务必须分层
+ controller实例化时对service， repo的依赖， 在wire中完成注入
+ controller层
  + controller这块负责参数的校验， 调用对应service层的逻辑处理相关业务
  + 一个模块一个controller， 在对应方法上实现相关业务
  + controller只依赖于service， 不直接调用repo层
  + controller层不相互调用
+ service层
  + service层实现具体业务逻辑
  + 一个模块一个service
  + service依赖于repo， 不能调用controller
  + service层之间可以相互调用
+ repo 层
  + repo层实现数据的相关curd操作
  + 所有repo必须要用interface封装， 对外提供interface就行
  + repo直接不相互调用， 有相关性的数据操作在service中完成调用

