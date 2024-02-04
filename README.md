## 友行应用级OS框架底库



- [点击查看tros使用示例](https://gitee.com/idigpower/tros-example-server.git
  )
- 目录结构：
  - |--conf: 配置文件管理
  - |--contants: 配置文件key常量 && 应用级的常量定义
  - |--context: tros版本的context封装
  - |--enums: 应用级枚举定义
    - |--country: 国家定义
    - |--currency: 货币定义
  - |--lang: 语言定义（后续可能废弃）
  - |--logx: 应用级日志管理trlogger
  - |--pkg: 应用级工具包，包含公共的工具方法和第三方
    - |--third_party: 第三方访问方法
    - |--utils: 工具类方法
      - |--encrypt: 应用级加解密方法
  - |--server: 应用级微服务启动的主入口
    - |--grpc: grpc微服务
    - |--http: http微服务
    - |--middleware: grpc和http的拦截器
  - |--sys: 系统级别的方法
    - |--cmd: 命令行的执行方法
    - |--structure: 应用级数据结构
    - |--timer: 应用级定时器
  - |--trerror: 应用级通用的错误定义
  - |--trkit: 中间件工具包，包含db、消息队列等服务器中间件的连接和使用
    - |--mongox: mongodb客户端
    - |--mysqlx: mysqldb客户端
    - |--redisx: redisdb客户端
- 说明：
  - 采用grpc-gateway技术，同时支持 restful api 和 grpc api。
  - 外部可采用 http 方式接入。
  - 内部可采用 grpc 做微服务通讯
  - 接口定义示例：
    - 通过定义 protobuf 文件来实现接口定义
    - 使用 buf 命令可以将 protobuf 文件，编译出 pb.go、pb.gw.go、_grpc.pb.go 三个文件
    - 通过“继承”的形式，来实现 service 中的 api 接口
  ```protobuf
  syntax = "proto3";
  
  package console.v1;
  
  //import "console/v1/common.proto";
  import "google/api/annotations.proto";
  //import "google/api/field_behavior.proto";
  import "google/protobuf/descriptor.proto";
  //import "google/protobuf/empty.proto";
  import "protoc-gen-openapiv2/options/annotations.proto";
  
  option go_package = "example-server/console/v1";
  
  // XXX api
  service ExampleService {
    // XXXX功能
    rpc List(ListRequest) returns (ListResponse) {
      option (google.api.http) = {
        post: "/console/v1/example/list"
        body: "*"
      };
      option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_operation) = {
        tags: ["XXXX功能"]
      };
    }
  }
  
  // ListResponse响应结构
  message ListResponse {
    // XXXX
    int64 total = 2 [(grpc.gateway.protoc_gen_openapiv2.options.openapiv2_field) = {description: "漏洞实例总数"}];
  }
  
  // 详情请求结构
  message ListRequest {
    // XXXXX
    string id = 1;
  }
  ```