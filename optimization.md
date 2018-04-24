https://signalfx.com/blog/a-pattern-for-optimizing-go-2/

https://github.com/golang/go/wiki/Performance

https://github.com/golang/go/wiki/CompilerOptimizations

边界检查
查看哪些检查： go build -gcflags="-d=ssa/check_bce/debug=1" main.go
关闭边界检查： go build -gcflags=-B main.go
循环使用range可消除检查，或者从大到小的idx可消除非首次的检查

inline
查看哪些可以inline：go build -gcflags -m main.go 

gc
用局部结构体存放多个变量，减少gc及内存占用
map的k、v不存在指针时，gc不对k、v进行扫描，所以持有小结构体可以直接存，不放指针以提高gc效率
