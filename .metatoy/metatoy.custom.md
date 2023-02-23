## metatoy.custom.jsonc 配置文档

### 1. 介绍
metatoy.custom.jsonc 用于储存自定义接口，其中包含了自定义接口的配置信息，以及自定义接口的数据，只需要配置好配置文件，metatoy即可自动生成接口，其中包含了鉴权、参数校验、数据返回等功能。

### 2. 结构
```json
[
    {
        "route" : "/api/custom", // 接口路由
        "method" : "get", // 接口方法
        "args" : { // 接口参数
            "参数名" : {
                "type" : "string", // 参数类型，支持Golang原生类型与master_type.PrimaryId，master_type.PrimaryId用于表示主键类型
                "validate" : {
                    "required" : true, // 是否必填
                    "min" : 1, // 最小值，支持数字与字符串
                    "max" : 10, // 最大值，支持数字与字符串
                }
            },
        },
        "require_login" : true, // 是否需要登录
        "service" : {
            "generate" : true, // 是否生成服务层，如果为false，则需要自行实现服务层，metatoy只会生成路由层和控制层
        },
        "permission" : [ // 权限配置，用户只要满足其中一个权限即可访问
            {
                "mode" : "admin" // 只有admin权限的用户才能访问这个服务
            },
            {
                "mode" : "model", // 基于数据模型的权限，当指定数据存在时，用户才能访问这个服务
                "model" : "Manager", // 数据模型名
                "model_info" : { // 数据模型信息
                    "name" : "xxxx", // 可以填写args中的参数名，但需要使用大驼峰命名法，metatoy后续可能会更新为小驼峰命名法
                    "time" : "$time", // $time会被替换为当前时间戳
                    "owner" : "$owner" // $owner会被替换为当前用户的id
                }
            }
        ],
        "actions" : [ //这个服务会进行的数据库IO操作，metatoy会按顺序执行
            {
                "type" : "create", // 创建数据
                "model" : "Manager", // 数据模型名
                "model_info" : { // 数据模型信息
                    "name" : "xxxx", // 可以填写args中的参数名，但需要使用大驼峰命名法，metatoy后续可能会更新为小驼峰命名法
                    "time" : "$time", // $time会被替换为当前时间戳
                    "owner" : "$owner", // $owner会被替换为当前用户的id
                    "xxx" : "$xxx", // 由$开头的参数会被替换为context中的值，$time与$owner等例外
                }
            },
            {
                "type" : "list", // 查询数据
                "model" : "Manager", // 数据模型名
                "model_info" : { // 数据模型信息
                    "name" : "xxxx", // 可以填写args中的参数名，但需要使用大驼峰命名法，metatoy后续可能会更新为小驼峰命名法
                    "time" : "$time", // $time会被替换为当前时间戳
                    "owner" : "$owner", // $owner会被替换为当前用户的id
                    "$page" : "xxx", //可以填写args中的参数名，但需要使用下划线命名法
                    "$page_size" : "xxx", //可以填写args中的参数名，但需要使用下划线命名法，
                    "$sort" : {
                        "type" : "const", // const, form，const表示常量，form表示从args中获取
                        "field" : "_id", // 排序字段
                        "value" : "-1" // asc: 1, desc: -1
                    },
                    "xxx" : "$xxx", // 由$开头的参数会被替换为context中的值，$time与$owner等例外
                },
                "dst" : "xxx", // 用于存储查询结果的变量名，当前结果会被存储在context中，可以在后续的action中使用
                "set" : "xxx" // 用于存储查询结果的变量名，当前结果会被设置到返回值中
            },
            {
                "type" : "get", // 查询单条数据
                "model" : "Manager", // 数据模型名
                "model_info" : { // 数据模型信息
                    "name" : "xxxx", // 可以填写args中的参数名，但需要使用大驼峰命名法，metatoy后续可能会更新为小驼峰命名法
                    "time" : "$time", // $time会被替换为当前时间戳
                    "owner" : "$owner", // $owner会被替换为当前用户的id
                    "xxx" : "$xxx", // 由$开头的参数会被替换为context中的值，$time与$owner等例外
                },
                "dst" : "xxx", // 用于存储查询结果的变量名，当前结果会被存储在context中，可以在后续的action中使用
                "set" : "xxx" // 用于存储查询结果的变量名，当前结果会被设置到返回值中
            }
        ],
        "response" : {
            "key" : "type" // 返回值的key，对应其类型，类型支持数组
        },
        "description" : "this is a custom api", // 接口描述
        "description_cn" : "这是一个自定义接口" // 接口描述（中文）
    }
]
```