package Str

func buildNext(str string) []int {
	n:=make([]int,len(str)) //新建一个int数组存储表
	n[0]=-1 //将首元素设置为-1
	t:=n[0]
	j:=0
	for j<len(str)-1{
		if t<0 || str[j]==str[t]{
			t++
			j++
			if str[j]!=str[t]{
				n[j]=t
			}else{
				n[j]=n[t]
			}
		}else{
			t=n[t]
		}
	}
	return n
}

func KMP(text,pattern string) int {
	next:=buildNext(pattern)
	t:=0
	p:=0
	for t<len(text) && p<len(pattern){
		if p<0 || text[t]==pattern[p]{
			t++
			p++
		}else{
			p=next[p]
		}
	}
	if (t-p) <= (len(text)-len(pattern)){
		return t-p
	}else{
		return -1
	}
}