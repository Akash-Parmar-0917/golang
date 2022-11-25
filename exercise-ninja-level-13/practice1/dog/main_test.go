package dog

import (
"testing"
"fmt"
)




func TestYears(t *testing.T) {
	a:=Years(7)
	if(a!=49){
		t.Error("not valid")
	}
}

func ExampleYears(){
	fmt.Println(Years(7))
	//Output: 49
}
func TestYearsTwo(t *testing.T){
	a:=YearsTwo(7)
	if(a!=49){
		t.Error("expected 49 but got",a)
	}
}


func BenchmarkYears(b *testing.B){
	for i:=0;i<b.N;i++{
		Years(100)
	}
}

func BenchmarkYearsTwo(b *testing.B){
	for i:=0;i<b.N;i++{
		YearsTwo(100)
	}
}