## Facade
![GitHub](https://img.shields.io/github/license/kushao1267/facade.svg)
![GitHub repo size](https://img.shields.io/github/repo-size/kushao1267/facade.svg)
![Build](https://travis-ci.org/kushao1267/Facade.svg?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/kushao1267/Facade)](https://goreportcard.com/report/github.com/kushao1267/Facade)
[![codecov](https://codecov.io/gh/kushao1267/Facade/branch/master/graph/badge.svg)](https://codecov.io/gh/kushao1267/Facade)


## About

Facade是一个全网通用的链接预览服务，它的功能如下:

* 支持REST API方式获取链接预览信息，API Server支持热加载
* 支持自定义预览信息的字段（已能够支持视频，feed，音频，标题，描述等字段）
* 支持提取图片url以及图片解析
* 有很强的可扩展性，能够自定义支持许多网站，可将technique自由搭配使用
* 使用Golang实现，有良好的性能
* 使用docker-compose，一键启动API和缓存服务


## Installation

```
go get -u github.com/kushao1267/Facade
```

## Module

* controllers模块
提供链接预览的控制器，gonic-gin框架，[gin文档](https://gin-gonic.com/docs/)

* services模块
提供相应实体的服务，例如LinkPreviewService，提供所有跟此实体(LinkPreview)相关的操作。

* config模块
加载toml文件配置，使用github.com/BurntSushi/toml库, [详见](https://github.com/BurntSushi/toml)

* db模块
数据库封装模块，目前只使用redis，用的是go-redis/redis库, [详见](https://github.com/go-redis/redis)

* extractors模块
使用者能够自定义extractor，来组合使用已有的technique，也可以调节使用technique的优先级，从而保证输出预览信息的完善和精确。例如使用technique调用链: WeiboTechnique -> HeadTagsTechnique -> SemanticTagsTechnique (指定->为调用优先级)，能够在抓取weibo feed链接的预览信息失败时，调用相应的通用technique来兜底。

* techniques模块
techniques中每个technique都提供了针对特定网站的多字段提取方法。此外还提供了通用的technique: HeadTagsTechnique、HTML5SemanticTagsTechnique、SemanticTagsTechnique，在其他特定technique提取信息失败时，可以用来兜底。使用者能够加入更多网站的technique，来完善该项目，欢迎提PR :)

* utils
工具模块，包含加密相关工具，http网络请求工具，图片解析工具，时间相关工具等方法


## Usage

1.复制环境变量文件.env，环境变量视情况自行更改
```
$ cp .env.tpl .env
```

2.使用docker-compose一键启动API和Redis缓存服务，只需运行：
```
make prod
```
仅使用alpine镜像，严格控制制作的镜像大小在10M内.可以在docker-compose.yml内设置Redis相关配置.

3.调试API server, 需要安装[gin - A live reload utility for Go web applications.](https://github.com/codegangsta/gin)来实现web服务的热加载.
```
make dev
```
其对应点命令是`gin -a 8080 -p 3000`, -a参数是WebApp所监听的端口，-p是gin为该WebApp服务代理的端口；最终访问API是从-p所指定的端口访问。

4.调用预览接口
```
$ curl http://127.0.0.1:3000/api/v1/ping

{"message":"pong"}
```

```
$ curl http://127.0.0.1:3000/api/v1/preview -F "url=https://blog.csdn.net/hugejihu9/article/details/83992009"

{
    "code":"1",
    "msg":"success"
    "data":{
        "description":"文章来自：源码在线https://www.shengli.me/php/209.html 注：",
        "image":"https://img-blog.csdnimg.cn/20181112144936680.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2h1Z2VqaWh1OQ==,size_16,color_FFFFFF,t_70",
        "title":"CURL模拟表单的post提交 - hugejihu9的专栏 - CSDN博客"
    },
}
```

5.调用接口console打印出详细的technique调用日志，并通过颜色区分
![api-log](https://github.com/kushao1267/Facade/blob/master/gin_api_log.jpg)

6.更多使用参见Makefile


## Test
目前测试所有定义的technique
```
$ make test
```

## LICENSE
[MIT License](https://github.com/kushao1267/facade/blob/master/LICENSE)
