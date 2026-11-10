# grpc-demo

用 Go 和 `buf` 实现的 gRPC 学习示例，包含服务提供方和调用方。

## 项目结构

| 目录 | 内容 |
| --- | --- |
| `greeter-api/` | 契约：`proto/` 中的 `.proto` 文件，`gen/` 中生成的 Go 代码 |
| `greeter-server/` | 服务提供方，默认监听 `:50051` |
| `greeter-client/` | 调用方 |
| `docs/` | 学习笔记 |

`greeter-server` 和 `greeter-client` 用 `go.mod` 中的 `replace` 指令引用本地的 `greeter-api`。

## 运行

1. 启动 server：

   ```bash
   cd greeter-server && go run .
   ```

2. 在另一个终端运行 client：

   ```bash
   cd greeter-client && go run . -name tonghs
   ```

   结果：client 输出 `Hello, tonghs!`。

修改 `.proto` 文件后，在 `greeter-api/` 下执行 `buf lint` 和 `buf generate`。

## 文档

- [一元 RPC 调用链路复盘](docs/01-unary-rpc-walkthrough.md)
