package mybase

import (
	"fmt"
	"os"
	"testing"
)

func TestRandInt(t *testing.T) {
	myRand := NewMyRand()

	//检查100亿次的随机数1~1亿的分布情况，理论上是1~1亿每个数各100左右
	const n = 100000000
	m := make([]int, n)
	for i := 0; i < n*100; i++ {
		if n > 10000000 && i%10000000 == 0 {
			t.Log("scan", i)
		}
		rnd := myRand.Intn(n)
		m[rnd]++
	}

	mm := make(map[int][]int)
	low := []int{1, 10, 20, 30, 40, 50, 60, 70, 80, 90}           //检查某个数随机次数小于指定数的
	up := []int{110, 120, 130, 140, 150, 160, 170, 180, 190, 200} //检查某个随机数次数大于指定数的
	for i := range m {
		v := m[i]
		find := false
		for j := range low {
			vv := low[j]
			if v <= vv {
				mm[vv] = append(mm[vv], i)
				find = true
			}
		}
		for j := len(up) - 1; j >= 0; j-- {
			vv := up[j]
			if v >= vv {
				mm[vv] = append(mm[vv], i)
				find = true
			}
		}
		if !find {
			mm[100] = append(mm[100], i)
		}
	}

	f, err := os.Create(fmt.Sprintf("D:\\myRand-out%d.csv", n))
	if err == nil {
		for i := range mm {
			fmt.Fprintf(f, "%d,%d\n", i, len(mm[i]))
		}
		f.Close()
	}
}
