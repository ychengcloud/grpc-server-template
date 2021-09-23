## 特性

- 符合 Golang 设计哲学的工程框架，包括基础功能(JWT, OpenTracing, ZapLog, Promtheus)
- 支持查询过滤、排序、分页、指定返回字段、批量创建、批量查询、批量更新、批量删除
- 根据数据库生成对应 CRUD 方法,支持字段 BelongTo 、HasOne 、HasMany 、ManyToMany 关联配置
- 支持 Swagger 文档自动生成
- 框架代码自主可控，初始生成后可根据业务需要灵活修改

- 安全特性
    - JWT
    - CSRF
    - XSS
    - CORS


## 项目生成

参见 [Heidou](https://docs.ycheng.pro/heidou)


## 运行项目

项目生成后，根据业务情况修改项目配置文件，即可构建运行

### 安装依赖

- [Make](setup-make.md)
- [Golang 1.16+](https://golang.org/doc/install)
- [Mysql 8.0+](https://dev.mysql.com/doc/refman/8.0/en/installing.html)
- [Heidou 0.1.10+](setup-local.md) 
- [gowatch](https://github.com/silenceper/gowatch) (可选)

### 修改项目配置文件

生成代码后，会生成样例项目配置文件 config/server-example.yaml，根据业务修改

```bash
cp config/server-example.yaml config/server.yaml
```

### 生成 mock 代码

    make mock

### 编译

    make build

### 启动

    make run

## 运行配置项说明

参见配置文件注释
