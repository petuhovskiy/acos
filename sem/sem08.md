# Семинар 8

```bash
gcc apb.S -oapb
./apb
echo $?
```

## `apb.S`

```assembly
        .text
        .global main
main:

    call    readi32     // eax = readi32(), ecx = ?, edx = ?

    mov     %eax, %ebx

    call    readi32

    add     %ebx, %eax

    call    writei32    // writei32(eax), eax = ?, ecx = ?, edx = ?

    call    nl

    xor     %eax, %eax  // eax = eax ^ eax
    ret
// int main() { return 0; }
```

Первое время используем `simpleio_$(uname -m)`

## `abcd.S`

```assembly
        .data
        .align  4
A:      .int    0
        .global A
B:      .space  4
        .global B
        .bss
        .align  4
C:      .space  4
D:      .space  4
R:      .space  4
        .global C, D, R

        .text
        .global process
process:
        mov     A, %eax
        add     B, %eax
        mov     %eax, R
        ret
```

```c
#include <stdio.h>

void process(void);

extern int A, B, C, D, R;

int main()
{
    scanf("%d%d%d%d%d", &A, &B, &C, &D, &R);
    process();
    printf("%d\n", R);
}
```