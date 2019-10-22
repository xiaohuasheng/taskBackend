package function

import (
	"fmt"
	"time"
)

//{ 不能单独一行
func Afunction(str string) {
	fmt.Println("function.go in function package.")
	fmt.Println(str)
	//声明并赋值 :=
	//声明了一定要使用, 引入的话可以引到 _ 里

	//var a = 5
	//a := 5
}

//首字母小写不能在包外引用
func inPackageFunc() {
	fmt.Println("I'm in package")
}

func OutPackageFunc() {
	fmt.Println("I can be used out of package")
}

func SliceDemo() {
	// 切片，可变长的数组
	//var numbers = []int{0,1,2,3,4,5,6,7,8}
	// int a = 5
	// int *a, b
	// a,b int *
	numbers := []int{0,1,2,3,4,5,6,7,8}
	printSlice(numbers)
	fmt.Println("numbers[1:4] ==", numbers[1:4])
	numbers = append(numbers, 2,3,4)
	printSlice(numbers)
}

func RangeDemo() {
	//这是我们使用range去求一个slice的和。使用数组跟这个很类似
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	//在数组上使用range将传入index和值两个变量。上面那个例子我们不需要使用该元素的序号，所以我们使用空白符"_"省略了。有时侯我们确实需要知道它的索引。
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	//range也可以用在map的键值对上。
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	//range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func GoDemo() {
	//goroutine
	go say("world")
	say("hello")
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

func ChannelDemo()  {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从通道 c 中接收

	fmt.Println(x, y, x+y)
}

