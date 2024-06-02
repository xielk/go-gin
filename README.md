实现API 日志的记录 
```

{
    "level":"info",                // 日志级别
    "time":"2021-11-21 16:04:50",  // 时间
    "caller":"core/core.go:474",   // 调用位置
    "msg":"core-interceptor",      // 日志信息
    "domain":"go-gin-api[fat]",    // 域名或服务名
    "method":"GET",                // HTTP 请求方式
    "path":"/api/admin",           // HTTP 请求路径
    "http_code":200,               // HTTP 状态码
    "business_code":0,             // 业务状态码
    "success":true,                // 状态，成功或失败
    "cost_seconds":0.001468745,    // 耗费时间，单位：秒
    "trace_id":"3c9a49f7d8bc0c9f1833",    // 当前请求的唯一ID
    "trace_info":{
        "trace_id":"3c9a49f7d8bc0c9f1833",
        "request":{
            "ttl":"un-limit",
            "method":"GET",               // HTTP 请求方法
            "decoded_url":"/api/admin",   // HTTP 请求路径
            "header":{                    // HTTP Request Header 信息
                ...
            },
            "body":""                     // HTTP Request Body 信息
        },
        "response":{
            "header":{                    // HTTP Response Header 信息
                ...
            },
            "body":{                      // HTTP Response Body 信息
                ...
            },
            "http_code":200,              // HTTP 状态码
            "http_code_msg":"OK",         // HTTP 状态码信息
            "cost_seconds":0.001468488    // 耗费时间，单位：秒
        },
        "third_party_requests":null,      // 请求第三方接口日志
        "debugs":null,                    // Debug 调试信息
        "sqls":[                          // SQL 执行信息
            ...
        ],
        "redis":[                         // Redis 执行信息
            ...
        ],
        "success":true,                   // 状态，成功或失败
        "cost_seconds":0.001468745        // 总耗时，单位：秒
    }
}
```
