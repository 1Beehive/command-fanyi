### <p align="center">command-fanyi</p>
#### <p align="center">命令行翻译工具,将命令行执行的结果翻译成中文</p>
#### <p align="center"><a href="https://github.com/mazezen/command-fanyi/releases"><img src="https://img.shields.io/github/release/command-fanyi/releases.svg" alt="GitHub release"></a><a href="https://github.com/mazezen/command-fanyi/blob/master/LICENSE"><img src="https://img.shields.io/github/license/mashape/apistatus.svg" alt="license"></a><p>
#### <p align="center"><a href="./README.md" target="_blank">简体中文</a> | <a href="./README_en.md" target="_blank">English</a> </p>

<img src="./images/img.png">

## 翻译采用的百度翻译，你可以自行修改为其他翻译
> 设置百度翻译的appid 和 app secret
> 
> export APPID="**********"
> 
> export SECRETKEY="**********"

## 安装 
1. Mac
```shell
curl -sSL https://raw.githubusercontent.com/mazezen/command-fanyi/refs/heads/master/install.sh | bash
```
2. Windows
下载地址： https://github.com/mazezen/command-fanyi/releases/download/v1.0.2/command-fanyi-amd64-windows.exe
自行配置系统变量

## 示例
```shell
fanyi go mod -h 
fanyi tldr git push
fanyi tldr curl
```
