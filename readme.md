# 根据模板自动生成文件工具

注意，目前仅支持替换一处

#### 兼容环境

MacOS

#### 如何安装

- 方法一 使用brew安装scg`brew install chyugo/tap/scg`

- 方法二 clone 当前仓库

#### 如何使用

```
$ ./scg --help
usage: scg  <-t TemplatePath> <-o OutputPath> <-c classname>

Usage of scg:
  -c string
    	class name  
  -o string
    	output path
  -t string
    	template path
```

- 注意：三个参数都须输入

#### 示例

终端cd到当前scg所在的目录，模板文件存放在 `/Users/AAA/Desktop/scg/template`，想将文件名和内容里面的`<%ClassName%>`替换成`abc`，输出到 `/Users/AAA/Desktop/scg/template/outPath`



`./scg -t /Users/AAA/Desktop/scg/template -o /Users/AAA/Desktop/scg/template/outPath -c abc`
