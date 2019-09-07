package main

/*
#cgo LDFLAGS: -lpam
#include <stdlib.h>
#include <security/pam_modules.h>
#include <security/pam_appl.h>
*/
import "C"
import (
	"fmt"
	"os"
	"time"
	"unsafe"
)

func log(s string, args ...interface{}) {
	f, err := os.OpenFile("/tmp/pam_log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	msg := fmt.Sprintf(s, args...)
	fmt.Fprintf(f, "%s: %s\n", time.Now(), msg)
}

// START OMIT
//export goAuthenticate
func goAuthenticate(handle *C.pam_handle_t, flags C.int, args []string) C.int {
	username, err := getItem(handle, C.PAM_USER)
	if err != nil {
		return C.PAM_AUTH_ERR
	}

	if username == "test" {
		return C.PAM_SUCCESS
	}

	return C.PAM_AUTH_ERR
}

// END OMIT

func getItem(handle *C.pam_handle_t, itemType C.int) (string, error) {
	item := C.CString("")
	defer C.free(unsafe.Pointer(item))

	pamCode := C.pam_get_item(handle, itemType, (*unsafe.Pointer)(unsafe.Pointer(&item)))
	if pamCode != C.PAM_SUCCESS {
		return C.GoString(item), fmt.Errorf("unable to obtain item: %d", pamCode)
	}

	return C.GoString(item), nil
}

func main() {}
