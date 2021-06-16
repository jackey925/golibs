package golibs

import (
	"testing"
	"unsafe"
)

const name = "jackysdasdasdasdadasdasdasdadasdas"


func test() {

 	s :=string([]byte(name))
	if s == ""{

	}
}

func test2() {

BytesToString(str2bytes(name))

}



func str2bytes(s string) []byte {
   x := (*[2]uintptr)(unsafe.Pointer(&s))
  h := [3]uintptr{x[0], x[1], x[1]}
  return *(*[]byte)(unsafe.Pointer(&h))
 }
func BenchmarkBytes2(b *testing.B)  {
	for i:=0;i< b.N;i++{
		test2()
	}
}

func BenchmarkBytes1(b *testing.B)  {
	for i:=0;i< b.N;i++{
		test()
	}
}

func BenchmarkBytesForUnsafe2(b *testing.B)  {
	for i:=0;i< b.N;i++{
		str2bytes(name)
	}
}

func BenchmarkBytesForUnsafe(b *testing.B)  {
	for i:=0;i< b.N;i++{
		StringToBytes(name)
	}
}

func BenchmarkBytes(b *testing.B)  {
	for i:=0;i< b.N;i++{
		s := []byte(name)
		if len(s) > 0{

		}
	}
}
