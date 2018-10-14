package main

/*
#include <stdio.h>
void print(const char *str) {
   printf("%s", str);
}
*/
import "C"

func main() {
    str := C.CString("Hello, World\n")
    C.print(str)
}
