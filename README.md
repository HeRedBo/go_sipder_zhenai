# go_sipder_zhenai
golang 珍爱网爬虫

>`资深工程师深度讲解Go语言` 下教程爬虫部分代码，因珍爱网后续对用户详情页做反爬虫机制 
 所有本项目针对次问题做了相关的优化 ，用户详情数据 在列表页面获取，详情页面数据 使用了`goquery` 包获取 。数据目前能正常返回 为方便查看代码 本项目项目已经将
>单任务版本爬虫与多任务版本爬虫分开了，方便大家对比查阅
   因为 珍爱网详情的反爬虫机制，但是由于数据在列表数据是可以有获取所有的用户信息的，对于原本的教程已做修改 通过获取页面的`window.__INITIAL_STATE__=` 对象获取即可获取所有的用户数据，在在 json 数据转换一下即可获取整个数据结构 



## 目录结构

```
├── ConcurrenceTask  // 并发爬虫项目
└── SingleTask   # 单任务版本爬虫
```

## 项目架构明
### 单任务爬虫架构
![示例图片](https://github.com/HeRedBo/go_sipder_zhenai/blob/master/SingleTask/code_images/%E5%8D%95%E4%BB%BB%E5%8A%A1%E7%89%88.PNG)

### 并发爬虫项目爬虫架构图
![示例图片](https://github.com/HeRedBo/go_sipder_zhenai/blob/master/SingleTask/code_images/%E5%B9%B6%E5%8F%91%E7%89%88%E7%88%AC%E8%99%AB%2B%E7%B2%97%E6%95%B0%E6%8D%AE%E5%AD%98%E5%82%A8.png)



### 环境要求
- golang >=1.7
- elasticsearch 5.8

### 项目运行 

#### 下载至本地，SingleTask、ConcurrenceTask 都是通过 mod 创建的独立项目 需要项目的初始化 步骤如下


```
# 下载项目到本地：
 git clone git@github.com:HeRedBo/go_sipder_zhenai.git

# 进入对应的项目 如 SingleTask  一次 执行如下命令：
go  mod tidy 
go mod vendor

```







