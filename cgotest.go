package main

// 読み込むCコード片を、'import "C"'の真上にコメントの形で書く必要がある
/*
#include <stdlib.h>
#include <stdio.h>

void hello(char *str)
{
  printf("hello, %s\n", str);
}
*/
import "C"
import "unsafe"

func main() {
  greetingTo := C.CString("cgo")
  // CStringで得たCのstring型は必ず解放しなければならない
  defer C.free(unsafe.Pointer(greetingTo))
  C.hello(greetingTo)
}
