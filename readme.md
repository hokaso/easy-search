# Go Service Development Practices (Go服务开发实践)

## 前言

没有写Best Practice，仅仅只是个人实践，有待在项目中持续改进。规范网上一抓一把，没有见过两个项目用一个规范的，sample application也都有很多不同的规范，所以还是根据团队在一起做项目的时候落实下来的实践，慢慢形成这个共识规范。

## 路径

```text
/cmd             // 所有的服务都需要从命令行启动
/config          // 所有的配置文件
/api             // RESTful API handler
/router          // RESTful API router
/services        // Service Object
/db              // DB 连接管理库
/pkgs            // 可供内外部使用的pkgs
/models          // 数据库model层
/serializers     // 项目共享的 Request,Response serializer
/README.md
```

## 参考

#### 命名规范
https://juejin.cn/post/6844903779175759885

#### API JSON structure standard
https://github.com/omniti-labs/jsend
http://www.ruanyifeng.com/blog/2014/05/restful_api.html

#### HTTP Status Code standard
https://restfulapi.net/http-status-codes/

#### Project Structure
http://www.zhangjiee.com/wiki/programming/go/project-layout.html
https://github.com/golang-standards/project-layout

#### Request Model 和 Response Model
当前端与后端交互的时候，比如API交互的时候，接受 Request，和返回 Response 的时候，如果结构复杂参数有很多的情况，特别是JSON Request的时候，必须设置 Request Response Model. (有的时候DB Model 和 Response 基本一样的情况下，也可考虑复用。)
![](https://z3.ax1x.com/2021/06/28/RNRoaq.png)

#### Serializers
> 为什么要？
1. Model 和 Response Model 不是完全对等，会有出入的时候。
2. Response的值跟DB的值可能不一样，比如DB是1，2，3而Response里面需要转换成 a, b, c。
3. 有独立的serializer层方便swagger生成结果。
4. Request的格式跟Model的格式也不一定一样，而且还需要validate。

> Serializer 会被用在2个地方（如上图描述的）
1. Request 内容 Serialize
2. Response 内容的 Serialize

> 下面以一个Station Model举例。假如你有个db model 叫做 station，那么你就会有一下 serializer。
1. CreateStationRequest => 创建 station 的时候需要
2. Station => 与 model station 对应的 serialize 后的结构
3. StationsResponse => List result
4. StationResponse => Single result
5. func SerializeStation(model.Station) serializers.Station => 用于把model序列化。

#### Config 规范
- 默认都放在 /configs 文件夹下。
- 不许把 .yml, .conf, ini 文件直接上传到git。（因为每个人的本地开发环境不一样，提交上来会让开发变得混乱。）
- 每个配置文件都有一个默认模板叫做 .yml.example 这样方便新的 developer 快速配置。
- 可以将 config.go 配置文件解析代码放在 configs 下面。
- 可以使用 https://github.com/jinzhu/configor 来读取和解析项目的 yml 配置。
- 原则上配置文件要在读取的时候全部加载好，不能再运行时 crash

#### Model Methods 命名规范
以 Product 为例你要实现，下面是CRUD默认方法的命名方法
```
QueryProduct(ID uint) *Product
QueryProducts(limit, offset int) []*Product
CreateProduct(Name string, ...) (*Product, error)
DeleteProduct(ID uint) error
UpdateProduct(ID uint, Name string, ...) error
```

#### Service Object
1. 跨 Model 的操作 （不能被放到单个Model。）
2. 放在API层代码太多，而且不能重用。

这样就可以将这个代码抽取出来变成一个service。下面是一个 service 的使用概念。
![](https://z3.ax1x.com/2021/07/02/RciJSO.png)

#### 开发技巧
##### Request Validation
表单和Query String，可以用Gin Validator来做，这样可以更容易理解和规范代码。
https://segmentfault.com/a/1190000022527284

##### Request Binding
JSON和Query String的Binding可以避免很多自己手动检查，这样让代码量减少很多。
- JSON

https://chenyitian.gitbooks.io/gin-web-framework/content/docs/17.html

- Query String

https://www.bookstack.cn/read/gin-en/only-bind-query-string

##### BindJSON, ShouldBindJSON, ShouldBindWith的区别 (BindQuery, ShouldBindQuery)

- BindJSON == MustBindWith 他在bind错误的时候会返回400。会调用AbortWithError。
- ShouldBindJSON 只会返回错误信息，并不会返回400.

##### Pagination Simple Solution
https://gorm.io/docs/scopes.html#Pagination（这个版本不支持 Total Count）

> 需求

1. 根据Page Size (Limit) 和 Page Index (Offset) 找到 Page 数据。
2. 获取 Total count，Current page size 用户可以知道是否已经到了最后一页。

##### Gorm First vs Find
First 当查不到结果的时候回返回一个 gorm.ErrRecordNotFound 错误，你可以通过 errors.Is(err, gorm.ErrRecordNotFound) 去处理这个 404 的错误。而 Find 则不会返回任何错误。所以根据需求使用。

##### 什么情况需要把代码放到 model 层
1. 方法很大很长，不如有的创建代码需要由很多逻辑，那么 CreateXXXXX 方法就可以放到model层。
2. 当有些方法需要反复的重用。你们可以抽象出去。
结论就是： 两三代码的，不需要重用的代码，就放在 handler 层吧，以后改起来也方便。

##### Sample Application

https://github.com/gothinkster/golang-gin-realworld-example-app

https://github.com/eddycjy/go-gin-example

https://github.com/GoogleCloudPlatform/microservices-demo

##### References

https://www.toptal.com/ruby-on-rails/rails-service-objects-tutorial

## 用法
localhost:3000/api/qq?qq=?

localhost:3000/api/phone?phone=?