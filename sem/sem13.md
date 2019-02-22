# Немного про кр

```assembly 
    // esp & 0xf
    // call process

process:
    push    ebp
    mov     esp, ebp

    push    esi
    push    edi

    push    ebp
    sub     $12, esp
```

> Вот даже если мы не знаем инструкцию чтобы сконвертировать int в long long, то что мы можем сделать?
>
> Правильно, загуглить!

> Вскрытие показало, что две идеи в одну задачу это слишком много

> Как проверять что стек выравнен?
>
> Забабахайте туда SSE инструкцию, которая падает если стек не выравнен.

# Работа с бинарными файлами

`man ascii`

`hexdump -C /usr/bin/cat | less`

`file /usr/bin/cat`

ejudge.ru/study/3sem/elf.html

`man 2 lseek`

```c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <errno.h>
#include <sys/types.h>
#include <unistd.h>
#include <fcntl.h>

int main(int argc, char *argv[])
{
    int fd = open(argv[1], O_RDONLY, 0);

    off_t pos = lseek(fd, 0, SEEK_END);
    // pos - смещение относительно начала файла
    if (pos < 0) {
        printf(stderr, "error: %s\n", strerror(errno));
        exit(1);
    }

    printf("%zu\n", sizeof(pos));
    printf("%lld\n", (long long) pos);
}
```

Специальный дефайн чтобы `off_t` стал 8 байт и можно было читать файлы >2ГБ `-D-FILE_OFFSET_BITS=64`

```c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <errno.h>
#include <sys/types.h>
#include <unistd.h>
#include <fcntl.h>

int main(int argc, char *argv[])
{
    int fd = open(argv[1], O_RDONLY, 0);

    unsigned value;
    off_t pos = lseek(fd, -sizeof(value), SEEK_END);
    // pos - смещение относительно начала файла
    if (pos < 0) {
        printf(stderr, "error: %s\n", strerror(errno));
        exit(1);
    }

    int r = read(fd, &value, sizeof(value));
    printf("%zu\n", sizeof(pos));
    printf("%lld\n", (long long) pos);
}
```

```c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <errno.h>
#include <sys/types.h>
#include <unistd.h>
#include <fcntl.h>

int main(int argc, char *argv[])
{
    int fd = open(argv[1], O_RDONLY | O_CREAT, 0600);

    char magic[4] = { 'A', 'B', 'C', 'D' };
    unsigned value;
    off_t pos = lseek(fd, -sizeof(magic), SEEK_END);
    // pos - смещение относительно начала файла
    if (pos < 0) {
        printf(stderr, "error: %s\n", strerror(errno));
        exit(1);
    }

    lseek(fd, 0, SEEK_SET);

    int r = write(fd, &magic, sizeof(magic));
    printf("%zu\n", sizeof(pos));
    printf("%lld\n", (long long) pos);
}
```

# Про время

> А все программисты на коболе либо на пенсии,
>
> либо еще где-то.

`man gettimeofday`

`man stroftime`

```c
#include <time.h>
#include <stdio.h>
#include <sys/time.h>

int main()
{
    time_t cur = time(NULL);
    printf("%ld\n", cur);

    struct timeval tv;
    gettimeofday(&tv, NULL);
    long long cur2 = tv.tv_sec * 1000000LL + tv.tv_usec;
    printf("%lld\n", cur2);

    struct tm *ptm = localtime(&cur);
    printf("%04d-%02d-%02d %02d:%02d:%02d\n", 
        ptm->tm_year + 1900,
        ptm->tm_mon + 1,
        ptm->tm_mday,
        ptm->tm_hour,
        ptm->tm_min,
        ptm->tm_sec);
}
```