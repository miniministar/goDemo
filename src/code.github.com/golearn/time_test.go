package main

import (
	"fmt"
	"testing"
	"time"
)

func timeDemo() {
	now := time.Now()
	fmt.Printf("current time: %v\n", now)

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("	year:%d\n	month:%v\n	day:%v\n	hour:%v\n	minute:%v\n	second:%v\n", year, month, day, hour, minute, second)
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

func TestTime(t *testing.T) {
	timeDemo()
	timestampDemo()
	timeCalculate()
	formatDemo()
	decodeTimeStr()
	tickDemo()
}

func timestampDemo() {
	now := time.Now()            //获取时间
	timestamp1 := now.Unix()     //获取秒级时间戳
	timestamp2 := now.UnixNano() //获取纳秒时间戳
	fmt.Printf("current timestamp1: %v\n", timestamp1)
	fmt.Printf("current timestamp2: %v\n", timestamp2)

	time1 := time.Unix(timestamp1, 0)
	fmt.Println(time1)
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", time1.Year(), time1.Month(), time1.Day(), time1.Hour(), time1.Minute(), time1.Second())
}

func timeCalculate() {
	now := time.Now()
	later := now.Add(time.Hour)
	fmt.Printf("current:%v\n", now)
	fmt.Printf("one hour later:%v\n", later)
	fmt.Printf("two time sub value:%v\n", later.Sub(now))
	fmt.Printf("two time equal :%v\n", later.Equal(now))
	fmt.Printf("later before now :%v\n", later.Before(now))
	fmt.Printf("later after now :%v\n", later.After(now))
}

func tickDemo() {
	ticker := time.Tick(time.Second)
	start := time.Now()
	time.Sleep(100 * time.Millisecond)
	end := time.Now()
	fmt.Println(end.Sub(start))
	for i := range ticker {
		fmt.Println(i)
	}
}

func formatDemo() {
	now := time.Now()
	// 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05"))
	fmt.Println(now.Format("2006/01/02 15:04"))
}

func decodeTimeStr() {
	now := time.Now()
	fmt.Println(now)
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	location, err := time.ParseInLocation("2006-01-02 15:04:05", "2020-01-04 20:20:20", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(location)
	fmt.Println(location.Sub(now))
}
