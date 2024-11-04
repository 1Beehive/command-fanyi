### <p align="center">command-fanyi</p>
#### <p align="center">命令行翻译工具,将命令行执行的结果翻译成中文</p>
#### <p align="center"><a href="https://github.com/jeffcail/command-fanyi/releases"><img src="https://img.shields.io/github/release/command-fanyi/releases.svg" alt="GitHub release"></a><a href="https://github.com/jeffcail/command-fanyi/blob/master/LICENSE"><img src="https://img.shields.io/github/license/mashape/apistatus.svg" alt="license"></a><p>
#### <p align="center"><a href="./README.md" target="_blank">简体中文</a> | <a href="./README_en.md" target="_blank">English</a> </p>

<img src="./images/img.png">

## 翻译采用的百度翻译，你可以自行修改为其他翻译
> 设置百度翻译的appid 和 app secret
> 
> export APPID="**********"
> 
> export APPSECRET="**********"

## 安装
```shell
brew install fanyi
```

## 示例
```shell
go mod -h fanyi
tldr git push fanyi
tldr curl fanyi
```

## homebrew formula
```shell
1.使用cd $(brew --repository homebrew/homebrew-core)切换到本地的homebrew-core目录；
```
报错:
Error: Tapping homebrew/core is no longer typically necessary.
Add --force if you are sure you need it for contributing to Homebrew.
解决：
```shell
brew tap --force homebrew/core
```

```shell
2.使用git commit提交自己的修改；
3.把https://github.com/Homebrew/homebrew-core fork一份；
4.使用git remote add命令添加自己的fork的homebrew-core库；
5.使用git push推送将本地提交推送到自己的fork的homebrew-core库中；
6.在GitHub网页上发起Pull request。
```
