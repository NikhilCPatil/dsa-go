type OrderedStream struct {
	data []string
	ptr int
 }
 
 
 func Constructor(n int) OrderedStream {
	 return OrderedStream {
		 data : make([]string, n+1),
		 ptr : 0,
	 }
 }
 
 
 func (this *OrderedStream) Insert(idKey int, value string) []string {
	 this.data[idKey-1] = value
	 res := []string{}
	 for this.data[this.ptr] != "" &&  this.ptr < len(this.data)-1 {
		 res = append(res, this.data[this.ptr])
		 this.ptr++
	 }
	 return res
 }
 
