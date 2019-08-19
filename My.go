package main

func main() {
	
	var x =[]int {3,4,1,8,7,2323,21,2,9}
	sort(x)
	for e,j := range x {
		println(e,j)
	}
	//t1 := time.Now().Unix()
	//var ii = factor(2, 12)
	//println(ii)
	//t2 := time.Now().Unix()
	//println(t2 - t1)
}


/**
插入排序精要
 */
func sort(x []int){
	for i:= 1; i< len(x);i++  {
		for j := i; j > 0; j-- {
			if x[j-1] > x[j] {
				x[j-1]^=x[j]
				x[j]^=x[j-1]
				x[j-1]^=x[j]
			}
		}
	}
}
/**
怎么算乘方
 */
func factor(x int, n int) int {
	if n == 0 {
		return 1
	} else if n&1 == 1 {
		tt := factor(x, (n-1)/2)
		return tt * tt * x
	} else {
		tt := factor(x, n/2)
		return tt * tt
	}

}
