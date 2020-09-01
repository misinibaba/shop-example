# 介绍
```
 每个文件夹都是一个微服务，依赖k8s和istio，这里主要是一个下单到展示的demo
 由于我这里使用的minikube作为实验环境，没有外部负载均衡，未使用istio-ingressgateway，用的原生ingress
 数据库缓存等由于实验环境单机服务器资源限制，所以选择了外部搭建
```
- 配置管理:configMap
- 熔断限流:istio
- 网关:ingress
- 服务发现:service
- 日志中心:<暂缺>
- 服务编排:deployment
- 监控：istio自带的grafana,promethus,zipklin等

## 依赖软件
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
