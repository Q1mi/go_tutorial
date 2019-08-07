package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 文件操作

func readFromFile() {
	// 打开文件
	fileObj, err := os.Open("./xx.txt") // 相对路径，相对执行的程序同目录下的xx.txt
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}

	defer fileObj.Close() // 关闭文件
	// 读取文件的内容
	var tmp = make([]byte, 128)
	n, err := fileObj.Read(tmp)
	if err != nil {
		fmt.Printf("read from file failed, err:%v\n", err)
		return
	}
	fmt.Printf("read %d bytes from file.\n", n)
	fmt.Println(string(tmp))
}

func readAll() {
	// 打开文件
	fileObj, err := os.Open("./xx.txt") // 相对路径，相对执行的程序同目录下的xx.txt
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	defer fileObj.Close() // 关闭文件

	// 读取文件的内容
	for {
		var tmp = make([]byte, 128)
		n, err := fileObj.Read(tmp)
		if err == io.EOF { // End Of File
			// 把当前读了多少个字节的数据打印出来，然后退出
			fmt.Println(string(tmp[:n]))
			return
		}
		if err != nil {
			fmt.Printf("read from file failed, err:%v\n", err)
			return
		}
		fmt.Printf("read %d bytes from file.\n", n)
		fmt.Println(string(tmp[:n]))
	}
}

// read by bufio
func readByBufio() {
	// 打开文件
	fileObj, err := os.Open("./xx.txt") // 相对路径，相对执行的程序同目录下的xx.txt
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	defer fileObj.Close() // 关闭文件
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Print(line)
			return
		}
		if err != nil {
			fmt.Printf("read file by bufio failed,err:%v\n", err)
			return
		}
		fmt.Print(line)
	}
}

// ioutil
func readByIoutil() {
	content, err := ioutil.ReadFile("./xx.txt")
	if err != nil {
		fmt.Printf("read file by ioutil failed,err:%v\n", err)
		return
	}
	fmt.Println(string(content))
}

// write
func write() {
	fileObj, err := os.OpenFile("./xxx.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	defer fileObj.Close()
	str := "沙河小王子"
	fileObj.Write([]byte(str))      // []byte
	fileObj.WriteString("hello 沙河") // string
}

func writeByBufio() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	defer fileObj.Close()
	writer := bufio.NewWriter(fileObj)
	writer.WriteString("沙河娜扎") // 将内容写入缓冲区
	writer.Flush()             // 将缓冲区的内容写入磁盘
}

func writeByIoutil() {
	str := "我命由我不由天！"
	err := ioutil.WriteFile("./xx.txt", []byte(str), 0644)
	if err != nil {
		fmt.Printf("write file failed, err:%v\n", err)
		return
	}
}

func main() {
	//readAll()
	//readByBufio()
	//readByIoutil()
	//write()
	//writeByBufio()
	writeByIoutil()
}
