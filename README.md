# go_sipder_zhenai
golang 珍爱网爬虫

>`资深工程师深度讲解Go语言` 下教程爬虫部分代码，因珍爱网后续对用户详情页做反爬虫机制 
 所有本项目针对次问题做了相关的优化 ，用户详情数据 在列表页面获取，详情页面数据 使用了`goquery`  包获取 。数据目前能正常返回 为方便查看代码 本项目项目已经将
>单任务版本爬虫与多任务版本爬虫分开了，方便大家对比查阅


## 目录结构

```
├── ConcurrenceTask  // 并发爬虫项目
└── SingleTask   # 单任务版本爬虫
```

## 项目架构明
### 单任务爬虫架构
![示例图片](https://github.com/HeRedBo/go_sipder_zhenai/blob/master/SingleTask/code_images/%E5%8D%95%E4%BB%BB%E5%8A%A1%E7%89%88.PNG)


