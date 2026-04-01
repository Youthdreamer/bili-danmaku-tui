# B站弹幕TUI工具

一个简洁美观的终端用户界面（TUI）工具，用于实时显示B站直播间弹幕。

## 感谢 blivedm-go

本项目基于 blivedm-go 二次开发。实现简单的弹幕展示功能。非常感谢。  
[blivedm-go项目地址](https://deepwiki.com/Akegarasu/blivedm-go)

## ✨ 功能特性

- 🎯 **实时弹幕显示**：实时接收并显示B站直播间弹幕
- 🖥️ **终端界面**：基于TUI的简洁界面，无需浏览器
- 🔐 **双模式支持**：支持匿名模式和Cookie认证模式
- 📱 **轻量高效**：Go语言编写，资源占用低
- ⌨️ **快捷键操作**：支持清屏、退出等操作
- 🎨 **简洁美观**：只显示弹幕核心信息，无冗余内容

## 📦 安装方法

### 前提条件

- Go 1.25.6 或更高版本

### 从源码安装

```bash
# 克隆项目
git clone https://github.com/Youthdreamer/bili-danmaku-tui.git
cd bili-danmaku-tui

# 构建项目
go build -o bili-danmaku-tui main.go

# 将可执行文件移动到PATH目录（可选）
sudo mv bili-danmaku-tui /usr/local/bin/
```

### 直接下载二进制文件

从 [Releases](https://github.com/Youthdreamer/bili-danmaku-tui/releases) 页面下载对应平台的二进制文件。

## 🚀 使用方法

### 基本使用

```bash
# 使用匿名模式（无需Cookie）
./bili-danmaku-tui <直播间ID>

# 示例：查看直播间22222的弹幕
./bili-danmaku-tui 22222
```

### 使用Cookie认证模式

1. 创建 `.env` 文件：

```bash
echo "BLIVE_COOKIE=你的B站Cookie" > .env
```

2. 运行程序：

```bash
./bili-danmaku-tui <直播间ID>
```

### 获取B站Cookie

1. 在浏览器中登录B站
2. 打开开发者工具（F12）
3. 访问任意B站页面
4. 在Network标签中找到请求，复制`Cookie`请求头的值

**必要的Cookie字段：**

- `buvid3`：设备标识
- `SESSDATA`：会话数据
- `bili_jct`：CSRF令牌

**Cookie格式示例：**

```
buvid3=XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX; SESSDATA=XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX; bili_jct=XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```

## ⚙️ 配置说明

### 环境变量

- `BLIVE_COOKIE`：B站Cookie，用于认证模式（可选）
- 如果没有 Cookie 弹幕功能会受限，建议使用

### 环境变量使用方式

项目支持两种方式设置环境变量：

1. **开发环境**：使用 `.env` 文件（推荐）
   ```bash
   echo "BLIVE_COOKIE=你的B站Cookie" > .env
   ```

2. **生产环境**：使用终端环境变量
   ```bash
   # Linux/macOS
   export BLIVE_COOKIE="你的B站Cookie"
   ./bili-danmaku-tui <直播间ID>

   # Windows (PowerShell)
   $env:BLIVE_COOKIE="你的B站Cookie"
   .\bili-danmaku-tui <直播间ID>
   ```

### 配置文件

项目使用 `.env` 文件管理环境变量，该文件已被添加到 `.gitignore` 中。

## ⌨️ 快捷键

| 快捷键                 | 功能                 |
| ---------------------- | -------------------- |
| `q` / `ctrl+c` / `esc` | 退出程序             |
| `c`                    | 清屏（清除所有弹幕） |

## 📁 项目结构

```
bili-danmaku-tui/
├── main.go          # 主程序文件
├── go.mod          # Go模块定义
├── go.sum          # 依赖校验
├── .gitignore      # Git忽略文件
└── README.md       # 项目说明文档
```

## 🔧 技术栈

- **Go**：主要编程语言
- **blivedm-go**：B站弹幕协议库
- **bubbletea**：TUI框架
- **godotenv**：环境变量管理

## 🛠️ 开发指南

### 环境设置

```bash
# 安装依赖
go mod download

# 运行开发版本
go run main.go <直播间ID>
```

### 构建发布版本

```bash
# 构建当前平台
go build -o bili-danmaku-tui main.go

# 构建多平台
GOOS=linux GOARCH=amd64 go build -o bili-danmaku-tui-linux-amd64 main.go
GOOS=darwin GOARCH=arm64 go build -o bili-danmaku-tui-darwin-arm64 main.go
GOOS=windows GOARCH=amd64 go build -o bili-danmaku-tui-windows-amd64.exe main.go
```

## 📄 许可证

本项目采用 MIT 许可证

## ⚠️ 注意事项

1. **Cookie安全**：请勿泄露你的B站Cookie
2. **使用限制**：请遵守B站用户协议
3. **网络要求**：需要稳定的网络连接
4. **兼容性**：目前仅支持终端环境

---

**感谢使用 B站弹幕TUI工具！** 🎉
