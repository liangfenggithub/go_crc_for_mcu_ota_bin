## 功能说明
为mcu编译生成的bin文件附加crc校验值,用于ota升级时进行完整性校验

## 分支描述
* main: 最简单实现
* flag: 使用flag库解析命令行参数


## 编译命令
```bash
go build ./main.go
```

## 使用方式

main分支
```bash
./main.exe `待操作的bin文件的path`
```
该命令自动生成一个新文件,文件名以crc及时间日期命令


flag分支
```bash
./main.exe -h //帮助信息

./main.exe -s `源bin文件路径` //同main分支效果

./main.exe -s `源bin文件路径` -o `目标文件路径` //可指定输出文件名

 
```


