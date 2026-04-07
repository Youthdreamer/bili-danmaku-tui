# B站弹幕TUI小工具

一个简洁高效的终端用户界面（TUI）工具，用于实时接收和发送B站直播间弹幕。

## 感谢 blivedm-go

本项目基于 blivedm-go 二次开发。实现简单的弹幕展示与弹幕发送功能。非常感谢。  
[blivedm-go项目地址](https://github.com/Akegarasu/blivedm-go)

## 推荐用法

> [!TIP]
> 推荐在 `niri` 与 `hyprland` 中使用  
> 本项目是在终端中显示，所以配合上面推荐的WM可以利用其浮窗，可使本工具始终在屏幕中显示。

## ✨ 功能特性

- 🎯 **实时弹幕功能**：实时接收并显示B站直播间弹幕，还可直接在终端中发送弹幕
- 🖥️ **终端界面**：基于bubbletea的现代化TUI界面
- 🔐 **双模式支持**：支持匿名模式（仅接收）和Cookie认证模式（接收+发送）
- 📱 **轻量高效**：Go语言编写，资源占用低，启动快速
- ⌨️ **快捷键操作**：支持清屏、退出等便捷操作
- 🎨 **简洁美观**：只显示弹幕核心信息，无冗余内容
- 🔧 **完整构建**：支持本地构建、Nix开发和GitHub Actions自动化发布

## 📦 快速开始

### 前提条件

- Go 1.25.6 或更高版本

#### 从 AUR 安装

```bash
yay -S bili-danmaku-tui
```

#### Nix 使用 Flake 安装

使用 [**`Nix Manager`**](https://nixos.org/download/) 安装，并开启 `flakes` 与 `nix-command` 功能使用下面的命令安装:

```nix
nix profile add github:Youthdreamer/bili-danmaku-tui
```

Nixos 用户可以使用以下方式安装到配置中,在 `flake.nix` 非常方便的导入

```nix
{
  inputs = {
    bili-danmaku-tui.url = "github:Youthdreamer/bili-danmaku-tui";
  }
}
```

安装到系统中，可以在 `inputs` 中引入，或者为其名

```nix
{ inputs, system, ...}:
{
  # NixOS
  environment.systemPackages = [ inputs.bili-danmaku-tui.packages.${pkgs.system}.default ];
  # home-manager
  home.packages = [ inputs.bili-danmaku-tui.packages.${pkgs.system}.default ];
}
```

### 从源码安装

```bash
# 克隆项目
git clone https://github.com/Youthdreamer/bili-danmaku-tui.git
cd bili-danmaku-tui

# 构建项目
make build

# 运行程序
./bili-danmaku-tui <直播间ID>
```

### 直接下载二进制文件

从 [Releases](https://github.com/Youthdreamer/bili-danmaku-tui/releases) 页面下载对应平台的二进制文件。

## 🚀 使用方法

### 基本使用（匿名模式）

```bash
# 使用匿名模式（仅接收弹幕，无需Cookie）
bili-danmaku-tui <直播间ID>

# 示例：查看直播间1895286942的弹幕
bili-danmaku-tui 1895286942
```

### 使用Cookie认证模式

#### 先决条件(获取Cookie)

- `BLIVE_COOKIE`：B站Cookie，用于认证模式
- 如果不使用`Cookie` 会无法发送弹幕，并无法获取完整的弹幕信息

#### 获取B站Cookie

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
buvid3=XXXXXXXX;SESSDATA=XXXXXXXXX;bili_jct=XXXXXXX
```

#### 正常使用

导入环境变量`BLIVE_COOKIE`(以`zsh`为例)

```bash
export BLIVE_COOKE="buvid3=XXXXXXXX;SESSDATA=XXXXXXXXX;bili_jct=XXXXXXX"
```

之后可在任何位置使用以下命令查看弹幕与发送弹幕

```bash
bili-danmaku-tui <直播间ID>
```

#### 开发模式

创建 `.env` 文件：

```bash
echo "BLIVE_COOKIE=buvid3=XXXXXXXX;SESSDATA=XXXXXXXXX;bili_jct=XXXXXXX" > .env
```

运行程序：

```bash
./bili-danmaku-tui <直播间ID>
```

### 命令行参数

```bash
# 使用 -r 参数指定房间号
bili-danmaku-tui -r 1895286942

# 查看版本信息
bili-danmaku-tui --version

# 查看帮助信息
bili-danmaku-tui --help
```

### 快捷键

| 快捷键   | 功能                         |
| -------- | ---------------------------- |
| `Enter`  | 发送弹幕（输入框中有内容时） |
| `Ctrl+L` | 清屏（清除所有弹幕）         |
| `Ctrl+C` | 退出程序                     |

### 界面布局

```
欢迎使用 Bilibili 弹幕助手！

-> [用户A] 大家好！
-> [用户B] 这个工具真好用
-> [用户C] 支持开源项目

输入弹幕，回车发送...[0/40]
```

## 📁 项目结构

```
bili-danmaku-tui/
├── cmd/                    # 命令行接口
│   └── root.go            # Cobra命令行框架
├── config/                # 配置管理
│   └── config.go          # 环境变量加载
├── danmaku/               # 弹幕功能模块
│   ├── client.go          # 弹幕接收客户端
│   └── sender.go          # 弹幕发送功能
├── tui/                   # 终端用户界面
│   ├── app.go             # TUI程序入口
│   ├── model.go           # 数据模型
│   ├── update.go          # 消息更新处理
│   └── view.go            # 界面渲染
├── main.go                # 程序入口
├── go.mod                 # Go模块定义
├── go.sum                 # 依赖校验
├── Makefile               # 构建脚本
├── flake.nix              # Nix开发环境配置
├── .github/workflows/     # GitHub Actions
│   └── release.yml        # 自动化发布流程
└── README.md              # 项目说明文档
```

### 开发指南

#### 环境安装

```bash
# 安装依赖
go mod download

# 运行开发版本
go run main.go <直播间ID>

```

#### 构建发布版本

```bash
# 构建当前平台
make build

# 构建指定平台
make build GOOS=linux GOARCH=amd64

# 构建发布包（用于GitHub Actions）
make release GOOS=linux GOARCH=amd64
```

#### 使用Nix开发环境

```bash
# 进入Nix开发环境
nix develop

# 在开发环境中构建和运行
make build
# or
nix build
```

## 📄 许可证

本项目采用 MIT 许可证

## ⚠️ 注意事项

1. **Cookie安全**：请勿泄露你的B站Cookie
2. **使用限制**：请遵守B站用户协议
3. **网络要求**：需要稳定的网络连接
4. **兼容性**：目前主要支持Linux环境

---

**感谢使用 B站弹幕TUI客户端！** 🎉
