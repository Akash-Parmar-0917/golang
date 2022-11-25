package mymath

import "testing"


func TestCenteredAvg(t *testing.T) {
	type test struct{
		values []int
		answer float64
	}

	t1:=[]test{
		test{values:[]int{1,2,3,4,5},answer:3},
		test{values:[]int{5,6,7,8,9},answer:7},
		test{values:[]int{10,11,12,13,14},answer:12},
	}

	for _,v:=range(t1){
		a:=CenteredAvg(v.values)
		if(a!=v.answer){
			t.Error("expectd",v.answer,"got",a)
		}
	}
	

}


func BenchmarkCenteredAvg(b *testing.B){
	for i:=0;i<b.N;i++{
		CenteredAvg([]int{1,2,3,4,5,6,7})
	}
}