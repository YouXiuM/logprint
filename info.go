/*
要求：
1. package 后面的名字最好和目录有关联，比如你的logprint目录名，package最好logprint
2. 最好是小写字母，并且最好不要带下划线，见名知义
3. 函数名字的首字母大写就可以给其他人用，小写字母只能给当前文件使用
*/
package logprint

import (
	"fmt"
	"time"
)

func Info(msg interface{}) {
	t := time.Now()
	fmt.Printf("[Info] %s: %s \n", t.Format("2006-01-02 15:04:05.000"), msg)
}
