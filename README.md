# InvFileLocker

<a href="LICENSE"><img src="https://img.shields.io/badge/协议-GPL v3.0-blue" alt=""/></a>&nbsp;
<a href=""><img src="https://img.shields.io/badge/语言-golang-blue" alt=""/></a>&nbsp;

InvFileLocker 是一个专门用于保护个人隐私的加密/解密器的生成器，旨在用于保护文件被黑客窃取，一键通过GUI界面一键可生成加密器解密器. 

![](/resources/generate_show_small.gif)

## 为什么使用 InvFileLocker

1. 一键生成加密解密器，一键加密/解密
2. 逻辑简单，特征简单，不会被杀软误判为病毒
3. 后期更新及维护
4. 当文本大小较低时，加密使用非对称算法，防止黑客窃取敏感文件

加密器演示: 

![](/resources/encryptor_show.gif)

解密器演示: 

![](/resources/decryptor_show.gif)

## 手动生成

由于本项目较为简单，因此仅仅对 go 文件简单分文件夹管理，只需要运行 /Generator/GeneratorScript.py 即可生成生成器(请确保在生成前安装了 go1.19.1 windows/amd64 或以上版本、python3.6 或以上版本)

```bash
cd Generator
python GeneratorScript.py
```

## 下载

Release 中最新版本 生成器.exe

## 声明

InvFileLocker 的部分功能可能与某些加密器解密器类似，但实际上有本质区别，InvFileLocker 开发者团队坚决抵制任何犯罪行为，任何利用该软件谋取不正当利益的行为均与开发者无关. 
