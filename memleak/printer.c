#include <math.h>
#include <stdio.h>

#include "_cgo_export.h"

void printPiJSON() 
{
    printf("%s\n", EncodeF64("pi", M_PI));
}