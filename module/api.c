// START OMIT
#include "string.h"
#include "stdlib.h"

#include "_cgo_export.h"

#include <security/pam_modules.h>
#include <security/pam_appl.h>

// utility to convert args list to go string slice
GoSlice cArgsToGoSlice(int, const char**);

PAM_EXTERN int pam_sm_authenticate(pam_handle_t* pamh, int flags, int argc, const char** argv) {
    return goAuthenticate(pamh, flags, cArgsToGoSlice(argc, argv));
}
//END OMIT

PAM_EXTERN int pam_sm_setcred(pam_handle_t* pamh, int flags, int argc, const char** argv) {
    return PAM_SUCCESS;
}

GoSlice cArgsToGoSlice(int argc, const char** argv) {
    GoString* strs = malloc(sizeof(GoString) * argc);
    GoSlice ret;
    ret.cap = argc;
    ret.len = argc;
    ret.data = (void*)strs;

    int i;
    for(i=0; i<argc; i++){
        strs[i] = *((GoString*)malloc(sizeof(GoString)));
        strs[i].p = (char*)argv[i];
        strs[i].n = strlen(argv[i]);
    }

    return ret;
}
