# 根据模板自动生成文件工具

注意，目前仅支持替换一处

#### 兼容环境

MacOS

#### 如何使用

```
$ scg
usage: scg [-c classname] [-o OutputPath] [-t TemplatePath] [-o OUTPUT]

Usage of scg:
  -c string
    	class name  
  -o string
    	output path
  -t string
    	template path
```

#### 示例

模板文件存放在 `/Users/AAA/Desktop/scg/template`，想将文件名和内容里面的`<%ClassName%>`替换成`abc`，输出到 `/Users/AAA/Desktop/scg/template/outPath`



`scg -t /Users/AAA/Desktop/scg/template -o /Users/AAA/Desktop/scg/template/outPath -c abc`
