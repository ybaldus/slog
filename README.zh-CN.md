# slog

Go 实现的简单、开箱即用的日志库

> 项目实现参考了 [Seldaek/monolog](https://github.com/Seldaek/monolog) and [sirupsen/logrus](https://github.com/sirupsen/logrus) ，非常感谢它们。

## [English](README.md)

English instructions please read [README](README.md)

## 功能特色

- 简单，无需配置，开箱即用
- 可以同时添加多个 `Handler` 日志处理器，输出日志到不同的地方
- 可以任意扩展自己需要的 `Handler` `Formatter` 
- 支持支持自定义 `Handler` 处理器
- 支持支持自定义 `Formatter` 格式化处理

## GoDoc

- [Godoc for github](https://pkg.go.dev/github.com/gookit/slog?tab=doc)

## Install

```bash
go get github.com/gookit/slog
```

## Usage

`slog` 使用非常简单，无需任何配置即可使用

## Quick Start

```go
package main

import (
	"github.com/gookit/slog"
)

func main() {
	slog.Info("info log message")
	slog.Warn("warning log message")
	slog.Infof("info log %s", "message")
	slog.Debugf("debug %s", "message")
}
```

**Output:**

```text
[2020/07/16 12:19:33] [application] [INFO] info log message  
[2020/07/16 12:19:33] [application] [WARNING] warning log message  
[2020/07/16 12:19:33] [application] [INFO] info log message  
[2020/07/16 12:19:33] [application] [DEBUG] debug message  
```

### Use JSON Format

```go
package main

import (
	"github.com/gookit/slog"
)

func main() {
	// use JSON formatter
	slog.SetFormatter(slog.NewJSONFormatter())

	slog.Info("info log message")
	slog.Warn("warning log message")
	slog.WithData(slog.M{
		"key0": 134,
		"key1": "abc",
	}).Infof("info log %s", "message")

	r := slog.WithFields(slog.M{
		"category": "service",
		"IP": "127.0.0.1",
	})
	r.Infof("info %s", "message")
	r.Debugf("debug %s", "message")
}
```

**Output:**

```text
{"channel":"application","data":{},"datetime":"2020/07/16 13:23:33","extra":{},"level":"INFO","message":"info log message"}
{"channel":"application","data":{},"datetime":"2020/07/16 13:23:33","extra":{},"level":"WARNING","message":"warning log message"}
{"channel":"application","data":{"key0":134,"key1":"abc"},"datetime":"2020/07/16 13:23:33","extra":{},"level":"INFO","message":"info log message"}
{"IP":"127.0.0.1","category":"service","channel":"application","datetime":"2020/07/16 13:23:33","extra":{},"level":"INFO","message":"info message"}
{"IP":"127.0.0.1","category":"service","channel":"application","datetime":"2020/07/16 13:23:33","extra":{},"level":"DEBUG","message":"debug message"}
```

## Workflow

```text
         Processors
Logger -{
         Handlers -{ With Formatters
```

## Processor

## Handler

## Formatter

## Refer

- https://github.com/golang/glog
- https://github.com/Seldaek/monolog

## Related

- https://github.com/sirupsen/logrus
- https://github.com/uber-go/zap
- https://github.com/rs/zerolog
- https://github.com/syyongx/llog

## LICENSE

[MIT](LICENSE)