### <p align="center">command-fanyi</p>
#### <p align="center">Command line translation tool</p>
#### <p align="center"><a href="https://github.com/jeffcail/command-fanyi/releases"><img src="https://img.shields.io/github/release/command-fanyi/releases.svg" alt="GitHub release"></a><a href="https://github.com/jeffcail/command-fanyi/blob/master/LICENSE"><img src="https://img.shields.io/github/license/mashape/apistatus.svg" alt="license"></a><p>
#### <p align="center"><a href="./README.md" target="_blank">Simplified Chinese</a> | <a href="./README_en.md" target="_blank">English</a> </p>

<img src="./images/img.png">

## The translation is done by Baidu Translate. You can change it to other translations.
> Set the appid and app secret of Baidu Translate
>
> export APPID="**********"
>
> export APPSECRET="**********"

## Install
1. Mac
```shell
curl -sSL https://raw.githubusercontent.com/jeffcail/command-fanyi/refs/heads/master/install.sh | bash
```
2. Windows
Download link： https://github.com/jeffcail/command-fanyi/releases/download/v1.0.2/command-fanyi-amd64-windows.exe

## Example
```shell
fanyi go mod -h 
fanyi tldr git push
fanyi tldr curl
```
