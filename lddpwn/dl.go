package main

/*
#include <stdlib.h>
*/
import "C"

func Random() int {
    return int(C.random())
}

