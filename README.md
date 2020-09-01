# 介绍
### 每个文件夹都是一个微服务，依赖k8s和istio，这里主要是一个下单到展示的demo

#依赖软件
- kubernetes
- istio
- docker
- golang
- grpc
- hbase
- es
- mysql
- redis
- kafka

### 部署方式
#### 一、
解压jobs.tar.gz,然后jenkins导入解压后的目录

#### 二、
修改Jenkins脚本中的安装位置

#### 三、
根据eddy-ingress.yaml中的主机名配置，在服务器的中配置相应的映射

#### 四、
开启istio自动注入

#### 五、
运行Jenkins部署脚本
