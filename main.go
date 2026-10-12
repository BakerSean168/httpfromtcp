package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// getLinesChannel 把"读文件 + 切行"的逻辑封装起来：
// 后台 goroutine 每次产出一行裸字符串，通过 channel 发给调用方。
func getLinesChannel(f io.ReadCloser) <-chan string {
	ch := make(chan string)
	go func() {
		defer f.Close() // 读完才关文件（关文件在 goroutine 内、读完后）
		defer close(ch) // 所有行发完后，goroutine 返回时才关通道

		buf := make([]byte, 8)
		currentLine := ""
		for {
			n, err := f.Read(buf)
			if n > 0 {
				parts := strings.Split(string(buf[:n]), "\n")
				// 除最后一段，前面每段都已遇到 \n，是完整行 → 发裸行
				for i := 0; i < len(parts)-1; i++ {
					ch <- currentLine + parts[i]
					currentLine = ""
				}
				// 最后一段还没成行，攒着
				currentLine += parts[len(parts)-1]
			}
			if err == io.EOF {
				break
			}
			if err != nil {
				return // 别用 log.Fatal：os.Exit 会跳过 defer，通道不关会死锁
			}
		}
		// 文件读完，若还有没收完的最后一行，补发
		if currentLine != "" {
			ch <- currentLine
		}
	}()
	return ch
}

func main() {
	f, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal(err)
	}
	// main 不再 defer f.Close()：文件由 getLinesChannel 内部关闭
	for line := range getLinesChannel(f) {
		fmt.Printf("read: %s\n", line)
	}
}
