<h1 align="center"><img src="https://raw.githubusercontent.com/cvelab/notify/main/static/logo.png" alt="Logo"/></h1>

[![license](https://img.shields.io/github/license/cvelab/notify?style=flat-square)](https://github.com/cvelab/notify/LIENCES)
[![Go Report Card](https://goreportcard.com/badge/github.com/cvelab/notify)](https://goreportcard.com/report/github.com/cvelab/notify)
[![CodeFactor](https://www.codefactor.io/repository/github/cvelab/notify/badge)](https://www.codefactor.io/repository/github/cvelab/notify)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fcvelab%2Fnotify.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fcvelab%2Fnotify?ref=badge_shield)

## Usage

```shell
> ./notify -h       
Usage:
  notify [OPTIONS]

Application Options:
  -s, --show        显示配置信息
  -c, --config=     指定配置文件 (default: $HOME/.config/notify-config.yaml)
  -i, --id=         指定通知id
  -l, --level=      指定通告等级 (default: 0)
  -a, --aboveLevel= 指定最低通告等级
  -n, --noticer=    指定通知模块
  -v, --version     版本信息

Help Options:
  -h, --help        Show this help message
```

## 安装

```shell
go install -v github.com/cvelab/notify/cmd/notify@latest
```

## 配置

1. 配置文件默认路径`$HOME/.config/notify-config.yaml`，或通过`-c|--config`指定
2. 具体每个推送配置、样例见相关路径下 README.md

### 基本设置

每个配置内容下均可以设置三个参数，均可在`-s`中可见

- notifyLevel 指定当前配置等级，默认`0`
- notifyFormatter 指定当前配置内容占位符，默认`[]`
- notifyDescription 当前配置内容描述内容，默认`""`

## 支持

- 多参输入
  - `echo -e "part1\npart2" | notify -n bark`
  - `notify part1 part22 -n bark`
- 指定id、level、模块，或最低level广播

## Licenses

[GPL v3.0](https://github.com/cvelab/nofity/blob/main/LICENSE)

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fcvelab%2Fnotify.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fcvelab%2Fnotify?ref=badge_large)
