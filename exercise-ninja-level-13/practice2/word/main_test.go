package word

import "testing"


func TestCount(t *testing.T) {
	a:=Count("My Name is Akash")
	if(a!=4){
		t.Error("expected 4 but got",a)
	}

}

func TestUseCount(t *testing.T){
	a:=UseCount("one and two two three three one")
	for k,v:=range a{
		switch k{
		case "one":
			if(v!=2){
				t.Error("kjenkjn")
			}
		case "and":
			if(v!=1){
				t.Error("kjnskcjn")
			}
		case "two":
			if(v!=2){
				t.Error("Janki weds Akash <3")
			}
		case "three":
			if(v!=2){
				t.Error("skdjcbknj")
			}
		}
	}
}

func BenchmarkCount(b *testing.B){
	for i:=0;i<b.N;i++{
		Count("My Name is Akash")
	}
}