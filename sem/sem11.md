# Семинар 11

Чтобы выравнивание даблов было правильно, не забываем писать `align`.

```assembly
        .att_syntax     noprefix
        .section        .rodata
        .align  8
pi:     .double 4.0
        .text
        .align  16
        .global main
main:
        push    ebp
        mov     esp, ebp

        sub     $8, esp

        mov     esp, eax
        push    eax
        push    eax
        push    eax
        push    $f1
        call    scanf

        movsd   -8(ebp), xmm0
        movsd   xmm0, (esp)
        call    func

        fstpl   4(esp)
        movl    $f2, (esp)
        call    printf

        add     $16 + 8, esp

        pop ebp
        ret
func:
        push    ebp
        mov     esp, ebp

        sub     $8, esp

        movsd   8(ebp), xmm0
        movsd   pi, xmm1
        mulsd   xmm1, xmm0

        movsd   xmm0, (esp)
        fldl    (esp)

        add     $8, esp
        pop     ebp
        ret
        
        .section        rodata
f1:     .asciz  "%lf"
f2:     .asciz  "%.10g\n"
```

> здесь были кеки про `%eiz`

## А теперь пишем тоже самое для x64

```assembly
        .att_syntax     noprefix
        .section        .rodata
        .align  8
pi:     .double 3.14
        .text
        .align  16
        .global main
main:
        push    rbp
        mov     rsp, rbp

        sub     $16, rsp

        lea     f1(rip), rdi
        mov     rsp, rsi
        xor     rax, rax
        call    scanf

        movsd   (rsp), xmm0
        call    func

        lea     f2(rip), rdi
        mov     $1, rax
        call    printf

        add     $16, rsp
        xor     rax, rax
        pop     rbp
        ret
func:
        movsd   pi(rip), xmm1
        mulsd   xmm1, xmm0
        ret
        .section        .rodata
f1:     .asciz  "%lf"
f2:     .asciz  "%.10g\n"

```

> `cat /proc/cpuinfo`