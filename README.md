## 功能说明
为mcu编译生成的bin文件附加crc校验值,用于ota升级时进行完整性校验

## 分支描述
* main: 最简单实现
* flag: 使用flag库解析命令行参数
* watch 为手环工程生成固件，支持指定输出文件路径

## 编译命令
```bash
go build ./main.go
```

## 使用方式
```bash
./main.exe `待操作的bin文件的path` `要输出的目标path（不包含文件名）`
```

