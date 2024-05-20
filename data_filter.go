package main

import (
	"fmt"
	"os"
)

func saveFundInfo(path string, data []byte) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// 将字节切片写入文件
	bytesWritten, err := file.Write(data)
	if err != nil {
		// 处理写入文件时的错误
		panic(err)
	}
	// 打印写入的字节数
	fmt.Println("Bytes written:", bytesWritten)

	// 关闭文件（defer语句已经确保了这一点，但这里显式调用以演示）
	err = file.Close()
	if err != nil {
		// 处理关闭文件时的错误
		panic(err)
	}
}
