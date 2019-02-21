# Семинар 12

```assembly
#include <asm/unistd_32.h>
    .global _start
_start:
        push    %ebp
        mov     %esp, %ebp
        and     $-16, %esp

        sub     $16, %esp
        mov     $__NR_read, %eax
        ..
        ..
```

Пример простейшего `_start`:

```c
asm(
    ".global _start\n"
    "_start:\n"
    "   call    main\n"
    "   mov %eax, %ebx\n"
    "   mov $1, %eax\n"
    "   int 0x80\n"
);
```

Здесь было немного по асм вставкам, очень вовремя.

## Файловые дескрипторы и файлы

Выводим процессы: `ps ax`

Смотрим открытые файлы: `cat /proc/7635/fd`

`man 2 open` - документация по системному вызову `open`.

```c
#include <stdio.h>
#include <stdlib.h>
#include <errno.h>
#include <string.h>

// #include <sys/types.h>
// #include <sys/stat.h>
#include <fcntl.h>

int main(int argc, char **argv)
{
    int fd = open(argv[1], O_RDONLY, 0);
    if (fd < 0) {
        fprintf(stderr, "open: %s %s\n", argv[1], strerror(errno));
        exit(1);
    }
}
```

```c
#include <stdio.h>
#include <stdlib.h>
#include <errno.h>
#include <string.h>

// #include <sys/types.h>
// #include <sys/stat.h>
#include <fcntl.h>
#include <unistd.h>

int main(int argc, char **argv)
{
    close(0);

    int fd = open(argv[1], O_RDONLY, 0);
    if (fd < 0) {
        fprintf(stderr, "open: %s %s\n", argv[1], strerror(errno));
        exit(1);
    }
    printf("%d\n", fd);

    int c;
    while ((c = getchar()) != EOF) {
        putchar(c);
    }
}
```

```bash
strace ./1 1.c
ulimit -a
ulimit -n 10    # Сделать макс. 10 открытых файлов
umask
```

```c
#include <stdio.h>
#include <stdlib.h>
#include <errno.h>
#include <string.h>

// #include <sys/types.h>
// #include <sys/stat.h>
#include <fcntl.h>
#include <unistd.h>

int main(int argc, char **argv)
{
    int fd = open(argv[1], O_RDWR, 0); // O_RDWR | O_APPEND, O_WRONLY, O_WRONLY | O_TRUNC
    // int fd = open(argv[1], O_WRONLY | O_TRUNC | O_CREAT, 0666);
    // int fd = open(argv[1], O_WRONLY | O_TRUNC | O_CREAT | O_EXCL, 0666);
    if (fd < 0) {
        fprintf(stderr, "open: %s %s\n", argv[1], strerror(errno));
        exit(1);
    }
    printf("%d\n", fd);

    char buf[16];
    read(fd, buf, 16);

    write(fd, "Hello", 5);
    write(fd, " world\n", 7);
}
```

`lseek`