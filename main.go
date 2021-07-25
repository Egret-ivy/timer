package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

type Watch struct {
	y    int
	m    int
	d    int
	hour int
	min  int
	sec  int
}

var now Watch

func pass() {
	now.sec++
	if now.sec >= 60 {
		now.sec -= 60
		now.min++
	}
	if now.min >= 60 {
		now.min -= 60
		now.hour++
	}
	now.hour %= 24

}

func main() {
	now.y = time.Now().Year()
	//注意月份这里很特别，不能直接得到int类型 直接出来的是July这种 要int转换
	now.m = int(time.Now().Month())
	//mon := time.Now().Month()
	now.d = time.Now().Day()
	now.hour = time.Now().Hour()
	now.min = time.Now().Minute()
	now.sec = time.Now().Second()
	fmt.Println(now.hour)
	fmt.Println(now.m)

	//得到当前时间的格式化输出
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(timeStr)

	for {
		pass()
		time.Sleep(time.Second)
		fmt.Printf("当前时间为： %d-%d-%d %02d:%02d:%02d\n", now.y, now.m, now.d, now.hour, now.min, now.sec)
		c := exec.Command("cmd.exe", "/c", "cls")
		c.Stdout = os.Stdout
		c.Run()
	}

}
