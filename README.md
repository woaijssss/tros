<h2 style="text-align: center;">
    Golang Application Level OS Framework Base Library
</h2>

---

## Table of Contents
- [Quickstart](#quickstart)
- [Directory structure](#directorystructure)
- [Explain for use](#explain)
- [DatabaseAccessLayer](#databaseaccesslayer)

## Quickstart
Browse [this tros example](https://github.com/woaijssss/tros-example-server.git) to quickly build your Golang web application.

## DirectoryStructure：
  - |--**conf**: config file management
  - |--**contants**: configuration file key constants && application level constant definitions
  - |--**context**: the context encapsulation of the TROS version
  - |--**enums**: application level enumeration definition
    - |--**country**: national definition
    - |--**currency**: money definition
  - |--**lang**: language definition (may be abandoned in the future)
  - |--**logx**: application level log management trlogger
  - |--**pkg**: application level toolkit, including common tool methods and third-party tools
    - |--**third_party**: third party access methods
    - |--**utils**: tool based methods
      - |--**encrypt**: application level encryption and decryption methods
  - |--**server**: main entrance for launching application level microservices
    - |--**grpc**: GRPC microservice
    - |--**http**: HTTP microservice
    - |--**middleware**: interceptors for GRPC and HTTP
  - |--**sys**: system level approach
    - |--**cmd**: command line execution method
    - |--**structure**: application level data structure
    - |--**timer**: application level timer
  - |--**trerror**: application level universal error definition
  - |--**trkit**: middleware toolkit, including connections and usage of server middleware such as DB and message queues
    - |--**mongox**: mongodb client
    - |--**mysqlx**: mysqldb client
    - |--**redisx**: redisdb client

## Explain
  - Adopting grpc gateway technology, supporting both RESTful API and grpc API
  - External access can be done via HTTP
  - GRPC can be used internally for microservice communication
  - Interface Definition Example：
    - Implement interface definition by defining protobuf file
    - The buf command can be used to compile protobuf files into three files: pb.go, pb.gw.go, and _grpc.pb.go
    - Implementing API interfaces in services through the form of 'inheritance'
```protobuf
  syntax = "proto3";
  
  package console.v1;
  
  import "google/api/annotations.proto";
  import "google/protobuf/descriptor.proto";
  import "protoc-gen-openapiv2/options/annotations.proto";
  
  option go_package = "example-server/console/v1";
  
  // XXX api
  service ExampleService {
    // XXXX function
    rpc List(ListRequest) returns (ListResponse) {
      option (google.api.http) = {
        post: "/console/v1/example/list"
        body: "*"
      };
      option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_operation) = {
        tags: ["XXXX desc"]
      };
    }
  }
  
  // ListResponse response
  message ListResponse {
    // XXXX
    int64 total = 2 [(grpc.gateway.protoc_gen_openapiv2.options.openapiv2_field) = {description: "total desc"}];
  }
  
  // ListRequest request
  message ListRequest {
    // XXXXX
    string id = 1;
  }
```

## DatabaseAccessLayer
The trkit directory provides fast read and write access methods for MySQL, MongoDB and Redis
Before use, you need to do the following steps(**Taking MySQL as an example**):
- Step 1:Define your. sql file and data table
```sql
CREATE TABLE `t_user`
(
    `id`        bigint      NOT NULL AUTO_INCREMENT COMMENT 'primary key',
    `user_name` varchar(64) NOT NULL COMMENT 'username',

) ENGINE=InnoDB DEFAULT COMMENT='user table';
```

- Step 2: Compile the CREATE TABLE statement file into Golang code using [GTC Tool(Go Table Compiler）](https://github.com/woaijssss/gtc)
- Step 3: Copy the generated database access Golang code to your project
  - Put the [dependency repo](https://github.com/woaijssss/godbx) warehouse into your project
- Step 4: Add MySQL connection parameters to your project configuration file and load them through Golang code
  - such as:
```yaml
mysql:
  url: root:123456@tcp(127.0.0.1:3306)/example
  maxIdleCons: 5
  maxLife: 3600
  poolSize: 15
  maxIdleTime: 1200
```

- Step 5: At this point, you can efficiently start the web server service and read/write databases