//
// main.go
// Copyright (C) 2019 Tim Hughes
//
// Distributed under terms of the MIT license.
//
package main

/*
#include <stdlib.h>
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: -L. -lperson
#include "person.h"
*/

/*
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type (
	Person C.struct_APerson
)

func GetPerson(name string, long_name string) *Person {
	return (*Person)(C.get_person(C.CString(name), C.CString(long_name)))
}

func main() {
	var f *Person
	f = GetPerson("tim", "tim hughes")
	fmt.Printf("Hello Go world: My name is %s, %s.\n", C.GoString(f.name), C.GoString(f.long_name))
	name := C.CString("Gopher")
	defer C.free(unsafe.Pointer(name))

	year := C.int(2018)

	g := C.struct_Greetee{
		name: name,
		year: year,
	}

	ptr := C.malloc(C.sizeof_char * 1024)
	defer C.free(unsafe.Pointer(ptr))

	size := C.greet(&g, (*C.char)(ptr))

	b := C.GoBytes(ptr, size)
	fmt.Println(string(b))
}
