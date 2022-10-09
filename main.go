package main

import (
	"fmt"
	"hash/crc32"
	"io"
	"os"
	"time"
)

var file_name string = "./bapp.bin"
var read_cnt int = 0

func getNowTimeStr() string {

	timeStr := time.Now().Format("20060102_150405") //当前时间的字符串，2006-01-02 15:04:05据说是golang的诞生时间，        固定写法
	// fmt.Println(timeStr)                                //打印结果：2017-04-11 13:24:04
	return timeStr
}

func main() {

	// 获取命令行参数
	if len(os.Args) < 2 { //判断命令行参数数量
		fmt.Println("please input ota file path")
		return
	}
	//遍历显示命令行参数
	// for k, v := range os.Args {
	// 	fmt.Printf("args[%v]=[%v]\n", k, v)
	// }

	file_name = os.Args[1]
	fmt.Printf("ota file path is %v\n", file_name)

	//判断文件是否存在
	stat, err := os.Stat(file_name)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("bin file size is %v byte\n", stat.Size())

	// 打开读取的文件
	file, err := os.Open(file_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// 每次读取的内容缓存
	context := make([]byte, stat.Size())
	read_cnt = 0
	count, err := file.Read(context)
	// 判断是否读到文件尾部
	if err != io.EOF && count == int(stat.Size()) {
		read_cnt = count
	} else {
		fmt.Printf("can not read all byte one time\n")
		os.Exit(0)
	}

	// 以下功能为循环读取文件内容，直到读完
	// var context []byte
	// for {
	// 	// 读取文件内容
	// 	count, err := file.Read(buf)

	// 	read_cnt += count
	// 	// 判断是否读到文件尾部
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	curByte := buf[:count]
	// 	// 追加内容
	// 	context = append(context, curByte...)
	// }
	fmt.Printf("read file byte count is %v\n", read_cnt)

	//crc 计算
	crcRes := crc32.Checksum(context, crc32.IEEETable)
	fmt.Printf("calculate crc32 res is 0x%x\n", crcRes)

	crcBuff := make([]byte, 4)
	crcBuff[0] = byte(crcRes)
	crcBuff[1] = byte(crcRes >> 8)
	crcBuff[2] = byte(crcRes >> 16)
	crcBuff[3] = byte(crcRes >> 24)

	//写入新文件
	newFileName := fmt.Sprintf("%x_%s.bin", crcRes, getNowTimeStr())
	fmt.Printf("new file name :%s\n", newFileName)

	file_new, err := os.Create(newFileName)
	if err != nil {
		fmt.Println("create new file fail") //create函数在创建文件时，首先会判断要创建的文件是否存在，如果不存在，则创建，如果存在，会先将文件中已有的数据清空。同时，当文件创建成功后，该文件会默认的打开，所以不用在执行打开操作，可以直接向该文件中写入数据。
		return
	}
	fmt.Printf("crate new file:%s success\n", newFileName)
	defer file_new.Close()
	//先写入源文件
	write_cnt, err := file_new.Write(context)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("write new file bin part byte cnt is %d\n", write_cnt)
	}

	//再写crc值
	write_cnt, err = file_new.Write(crcBuff)
	if err != nil {
		fmt.Println(err)
		fmt.Println("write new file crc fail")
		return
	} else {
		fmt.Printf("write new file crc part byte cnt is %d\n", write_cnt)
	}

	fmt.Printf("congratulation, crc append success, targe file is [ %s ]\n", newFileName)

}
