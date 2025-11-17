## 项目目标

* 构建一个可运行的 Go 后端服务，用于“合同审查”场景的演示。

* 明确且可被主流代码扫描工具识别的常见漏洞点，覆盖多类安全问题。

* 保持实现简洁：不生成测试文件、仅在复杂逻辑处保留必要注释、避免页面预览与自动执行命令。

## 技术栈

* 语言：Go 1.21+

* 标准库：`net/http`、`database/sql`、`html/template`、`crypto/md5`、`os/exec`、`io`、`os`、`path/filepath`、`math/rand`、`crypto/tls` 等

* 依赖（仅 1 个）：`github.com/mattn/go-sqlite3`（用于演示 SQL 注入的经典触发模式）

## 目录结构

* `go.mod`：模块与依赖

* `cmd/server/main.go`：入口与路由注册

* `internal/server/server.go`：HTTP 服务器初始化、CORS 中间件（含误配）

* `internal/db/db.go`：SQLite 初始化与示例数据；故意拼接 SQL

* `internal/handlers/`

  * `auth.go`：`/login`（弱加密、硬编码密钥、弱随机）

  * `users.go`：`/users`（SQL 注入）

  * `files.go`：`/file`、`/upload`（路径遍历/不安全写入）

  * `exec.go`：`/exec`（命令注入；Windows 使用 `cmd /C`，Linux 使用 `sh -c`）

  * `proxy.go`：`/proxy`（SSRF + 不安全 TLS）

  * `render.go`：`/render`（XSS：不安全模板转义）

  * `debug.go`：`/debug`（信息泄露/栈追踪）

  * `redirect.go`：`/redirect`（开放重定向）

  * `cors.go`：`/cors`（CORS 误配：`*` + 凭据）

## 端点与刻意漏洞

* `POST /login`

  * 硬编码凭据：如 `admin:admin123`

  * 弱加密：`md5` 存储/校验

  * 弱随机令牌：`rand.Intn`

  * 直接在响应体中回显敏感信息

* `GET /users?name=<value>`

  * 经典 SQL 注入：`db.Query("SELECT * FROM users WHERE name = '" + name + "'")`

* `GET /file?path=<value>` / `POST /upload`

  * 路径遍历：`./uploads/` + 用户输入

  * 无白名单/规范化检查，覆盖读取与写入场景

* `POST /exec`

  * 命令注入：不做校验即拼接到系统命令

  * Windows：`exec.Command("cmd", "/C", userCmd)`；Linux：`exec.Command("sh", "-c", userCmd)`

* `GET /proxy?url=<value>`

  * SSRF：直接 `http.Get(url)`

  * 不安全 TLS：`Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}`

* `GET /render?txt=<value>`

  * XSS：使用 `template.HTML(userInput)` 禁用转义

* `GET /debug`

  * 信息泄露：输出环境变量、内部错误详情

* `GET /redirect?to=<url>`

  * 开放重定向：直接跳转到用户提供的 URL

* `GET /cors`

  * 误配：`Access-Control-Allow-Origin: *` 与 `Access-Control-Allow-Credentials: true`

## 实现要点

* 仅在复杂或容易误解的逻辑处保留简短注释，其余不加注释。

* 路由使用 `net/http` 与简单 `http.HandleFunc`，避免额外框架。

* 可选初始化示例数据表 `users`，便于 SQL 语句演示。

* 按操作系统选择命令调用（运行时判断 `runtime.GOOS`）。

* 所有端点默认无鉴权，并在部分响应头/体故意泄露信息。

## 依赖与配置

* `go.mod` 初始化模块名：例如 `module contractAudit`

* 引入 `github.com/mattn/go-sqlite3` 作为唯一外部依赖

* 可选 `.env` 读取但故意硬编码部分“密钥”（不创建额外文件）

## 验证方式（不运行，仅静态检查）

* 通过代码扫描工具进行静态分析验证（SQLi/Command Injection/Path Traversal/SSRF/XSS/CORS/Weak Crypto/Hardcoded Secret 等）。

* 代码层面保留典型模式，确保扫描规则易于匹配。

<br />

引用有历史漏洞的SCA包

## 交付物

* 一组可编译的 Go 源码文件，包含上述目录与端点。

* 不包含测试文件与页面预览；不自动执行任何命令。

## 下一步

* 若您确认该方案，我将开始编写代码，严格按上述结构与漏洞点实现，避免过度注释与额外文件。

