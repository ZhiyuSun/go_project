package main

import (
	"fmt"
	"strings"
)

func main()  {
	//字符串基本操作
	//1. 求解字符串的长度
	//返回的是字节的长度
	//涉及到中文问题就产生了变化
	//unicode 字符集， 存储的时候需要编码 utf8编码规则， utf8编码是一个动态的编码规则
	//utf8编码， 还能够用一个字节表示英文
	var name string = "孙志宇"
	fmt.Println(len(name)) // 9

	// 类型转换
	nameArr := []int32(name)
	fmt.Println(len(nameArr))

	var name2 string = "我的名字：\"孙志宇\""
	fmt.Println(name2)

	date := "2020\\06\\18" //转义符
	fmt.Println(date)

	date2 := `2020\06\18` //转义符
	fmt.Println(date2)

	//2. 是否包含某个子串
	var name3 string = "我的名字：\"孙志宇\""
	fmt.Println(strings.Contains(name3, "孙志宇"))
	fmt.Println(strings.Index(name3, "孙志宇"))

	//3. 统计出现的次数
	fmt.Println(strings.Count(name3, "孙"))

	//4. 前缀和后缀
	fmt.Println(strings.HasPrefix(name3, "你"))
	fmt.Println(strings.HasSuffix(name3, "\""))

	//5. 大小写转换
	fmt.Println(strings.ToUpper("love"))
	fmt.Println(strings.ToLower("LOVE"))

	//6. 字符串的比较
	fmt.Println(strings.Compare("aa", "ab")) //字符的比较就是ascii的比较 返回-1， 1， 0
	fmt.Println(strings.Compare("b", "a"))   //字符的比较就是ascii的比较 返回-1， 1， 0
	fmt.Println(strings.Compare("b", "b"))   //字符的比较就是ascii的比较 返回-1， 1， 0

	//7. 去掉空格和指定字符串
	fmt.Println(strings.TrimSpace(" sun "))
	fmt.Println(strings.TrimLeft("sunsets", "s"))
	fmt.Println(strings.Trim("sunsets", "s"))

	//8. split方法
	fmt.Println(strings.Split("love you", " "))

	//9. 合并 join方法将字符串数组连接起来
	arrs := strings.Split("I love you", " ")
	fmt.Println(strings.Join(arrs, "-"))

	//10. 字符串替换
	fmt.Println(strings.Replace("sun: 26, zhi: 26, yu: 26", "26", "27", 2))
}