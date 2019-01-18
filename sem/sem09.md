# Семинар 9

## `1.S`

```assembly
        .att_syntax     noprefix
        .text
        .global main
main:
        // scanf("%d", &x)
        push    $x
        push    $str1
        call    scanf
        add     $8, esp

        // printf("%d\n", x)
        push    x
        push    $str2
        call    printf
        add     $8, esp



        xor     eax, eax
        ret
        .section .rodata
str1:   .asciz  "%d"
str2:   .asciz  "%d\n"
        .data
x:      .int     0
```

## `2.S`

```assembly
        .att_syntax     noprefix
        .text
        .global main
main:
        push    ebp
        mov,    esp, ebp
        and     $-16, esp

        // scanf("%d", &x)
        sub     $16, esp
        movl    $str1, (esp)
        movl    $x, 4(esp)
        call    scanf
        // add     $16, esp = плохой вариант, но тут можно даже не чистить

        // printf("%d\n", x)
        push    x
        push    $str2
        call    printf
        // add     $8, esp = плохой вариант, но тут можно даже не чистить

        mov     ebp, esp
        pop     ebp

        xor     eax, eax
        ret
        .section .rodata
str1:   .asciz  "%d"
str2:   .asciz  "%d\n"
        .data
x:      .int     0
```

## `3.S`

```assembly
        .att_syntax     noprefix
        .text
        .global main
main:
        push    ebp
        mov     esp, ebp
        sub     $8, esp
        // x: (esp), или -8(ebp)

        // scanf("%d", &x)
        sub     $16, esp
        movl    $str1, (esp)
        //lea     -8(ebp), eax
        //mov     eax, 4(esp)
        mov     ebp, 4(esp)
        sub     $8, 4(esp)
        call    scanf
        // add     $16, esp // плохой вариант, но тут можно даже не чистить

        // printf("%d\n", x)
        movl    $str, (esp)
        mov     -8(ebp), eax
        mov     eax, 4(esp)
        call    printf
        add     $16, esp

        mov     ebp, esp
        pop     ebp

        xor     eax, eax
        ret
        .section .rodata
str1:   .asciz  "%d"
str2:   .asciz  "%d\n"
```

## `main.C`

```c
#include <stdio.h>

int main()
{
        int x;

        scanf("%d", &x);
        printf("%d\n", x);
}
```


## `4.S`

```assembly
        .att_syntax     noprefix
        .text
        .global main
main:
        push    ebp
        mov     esp, ebp
        push    ebx
        push    esi
        push    edi

        sub     $8, esp
        // x: (esp), или -8(ebp)

        // scanf("%d", &x)
        sub     $16, esp
        movl    $str1, (esp)
        //lea     -8(ebp), eax
        //mov     eax, 4(esp)
        mov     ebp, 4(esp)
        sub     $8, 4(esp)
        call    scanf
        // add     $16, esp // плохой вариант, но тут можно даже не чистить

        // printf("%d\n", x)
        movl    $str, (esp)
        mov     -8(ebp), eax
        mov     eax, 4(esp)
        call    printf
        add     $16, esp

        add     $12, esp
        pop     edi
        pop     esi
        pop     ebx
        pop     ebp

        xor     eax, eax
        ret
        .section .rodata
str1:   .asciz  "%d"
str2:   .asciz  "%d\n"
```