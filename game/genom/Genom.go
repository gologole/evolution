package genom

// #cgo LDFLAGS: -L. -lmylib
//#include <stdlib.h>
//#include <time.h>
//int** InitGenom(int ROWS, int COLS, int maxnum, unsigned int seed);
//void freeMemory(void** genom, int rows);
//void PrintMatrix(int** matrix, int ROWS, int COLS);
import "C"
import (
	"time"
	"unsafe"
)

func InitGenom(genomGo [][]int, rows, cols, maxnum int) **C.int {
	//rows := C.int(2)
	//cols := C.int(2)
	//maxnum := C.int(100)

	genom := C.InitGenom(C.int(rows), C.int(cols), C.int(maxnum), C.uint(time.Now().UnixNano()))
	//	genomGo := make([][]int, int(rows))

	for i := 0; i < int(rows); i++ {
		row := (**C.int)(unsafe.Pointer(uintptr(unsafe.Pointer(genom)) + uintptr(i)*unsafe.Sizeof(*genom)))
		genomGo[i] = make([]int, int(cols))
		for j := 0; j < int(cols); j++ {
			genomGo[i][j] = int(*(*C.int)(unsafe.Pointer(uintptr(unsafe.Pointer(*row)) + uintptr(j)*unsafe.Sizeof(**row))))
		}
	}
	return (**C.int)(unsafe.Pointer(genom))
}

func FreeMemoryByC(genom unsafe.Pointer, rows int) { //end
	cgenom := (*unsafe.Pointer)(unsafe.Pointer(genom))
	C.freeMemory(cgenom, C.int(rows))
}
