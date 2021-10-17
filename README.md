# go-rest-api

## DB

### 生成新的 ent Schema

User 示例

```shell
go run entgo.io/ent/cmd/ent init User
```

### 更新 ent 代码

```shell
go generate ./ent
```

## API 文档

### 初始化或者更新文档

```shell
swag init -g cmd/server/main.go
```

## 参考

- [ribice/gorsk](https://github.com/ribice/gorsk)
- [nixsolutions/golang-echo-boilerplate](https://github.com/nixsolutions/golang-echo-boilerplate)
