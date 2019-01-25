# Семинар 10

## Флаги процессора

**Z** - знак нуля

**S** - флаг знака

**C** - флаг переноса

**O** - флаг переполнения

## Поставить флаг?

`SET??  R/M`

```c
bool x = a > b;
```

## Условное присваивание

```assembly
cmov ?      SRC, DST
```

## код

```assembly
        .att_syntax noprefix
        .text
        .global main
main:
        mov     $12332, eax
        mov     $-1211, ebx
        add     ebx, eax

        pushf
        pop     ebx

        push    ebx
        push    eax
        push    $str
        call    printf
        add     $12, esp

        xor     eax, eax
        ret
str:    .asciz  "%d %x\n"
```

```assembly
        .att_syntax noprefix
        .text
        .global main
main:
        mov     $12332, eax
        mov     $-1211, ebx
        add     ebx, eax

        setz    bl
        movzb   bl, ebx
#        sets    cl
#        seto    dl

        push    ebx
        push    eax
        push    $str
        call    printf
        add     $12, esp

        xor     eax, eax
        ret
str:    .asciz  "%d %x\n"
```


```assembly
        .att_syntax noprefix
        .text
        .global main
main:
        mov     $12332, eax
        mov     $-1211, ebx
        add     ebx, eax

        mov     $2, ecx
        mov     $0, ebx
        cmovc   ecx, ebx

        push    ebx
        push    eax
        push    $str
        call    printf
        add     $12, esp

        xor     eax, eax
        ret
str:    .asciz  "%d %x\n"
```

## 100 20 8 битные числа


```assembly
arg1:   .int    1, 2, 3, 4
arg2:   .int    5, 6, 7, 8

sum:
        xor     %esi, %esi

        clc
LOOP:
        mov     arg1(, %esi, 4), %eax
        add     arg2(, %esi, 4), %eax
        lea     1(%esi), %esi

        cmp     $4, %esie
        jb      LOOP

        ret

        .global main
main:
        ret
```

```assebmly
        .att_syntax noprefix
        .global process
        .text
process:
        push    ebp
        mov     esp, ebp
        push    ebx

        mov     A, eax
        mov     A + 4, edx

        mov     B, ebx
        mov     B + 4, ecx

        shl     $1, eax
        rcl     $1, edx

        clc
        rcr     ecx
        rcr     ebx

        clc
        rcr     ecx
        rcr     ebx

        sub     ebx, eax
        sbb     ecx, edx

        push    edx
        push    eax
        push    $str
        call    printf
        add     $12, esp

        pop     ebx
        pop     ebp
        ret
```

```c
struct List
{
        struct List *next;
        int data;
};
```

```assembly
        mov (ebx), ebx  ;; next pointer
        mov 4(ebx), eax ;; data
```

Инструкция `SHLD` для 64-битных