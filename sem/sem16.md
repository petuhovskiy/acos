# Семинар 16

> fork1.c
```c
#include <stdio.h>
#include <stdlib.h>

#include <unistd.h>

int main()
{
    int pid = fork();
    printf("%d\n", getpid());
}
```

> fork2.S
```assembly
#include <asm/unistd_32.h>

    .global main
main:

    mov $__NR_fork, %eax
    int $0x80

    push    %eax
    push    $str
    call    printf
    add     $8, %esp

    xor     %eax, %eax
    call    exit

str:
    .asciz  "%d\n"
```

> fork3.c
```c
#include <stdio.h>
#include <stdlib.h>

#include <unistd.h>

int main()
{
    int pid = fork();
    printf("%d %d %s\n", getpid(), pid, pid ? "father" : "son");
}
```

> fork4.c
```c
#include <stdio.h>
#include <stdlib.h>

#include <unistd.h>

int main()
{
    int pid1 = fork();
    int pid2 = fork();
    printf("%d %d %d\n", getpid(), pid1, pid2);
}
```

> fork5.c
```c
#include <stdio.h>
#include <stdlib.h>

#include <unistd.h>

int main()
{
    printf("hello");
    int pid = fork();
    printf(" world\n");
    // выводит
    // hello world
    // hello world
}
```

> fork6.c
```c
#include <stdio.h>
#include <stdlib.h>

#include <unistd.h>

int main()
{
    printf("hello");
    int pid = fork();
    printf("world");
    if (!pid) {
        _exit(0);
    }
}
```

> fork7.c
```c
#include <stdio.h>
#include <stdlib.h>
#include <errno.h>
#include <string.h>

#include <unistd.h>
#include <sys/types.h>
#include <sys/wait.h>

int main()
{
    printf("hello");
    fflush(stdout);
    int pid = fork();
    printf(" world\n");
    int r = wait(NULL);
    if (r < 0) {
        printf("%d %s\n", getpid(), strerror(errno));
    } else {
        printf("%d %d\n", getpid(), r);
    }
}
```

> fork8.c
```c
#include <stdio.h>
#include <stdlib.h>
#include <errno.h>
#include <string.h>

#include <unistd.h>
#include <sys/types.h>
#include <sys/wait.h>

enum { COUNT = 100000 };

int main()
{
    for (int i = 0; i < COUNT; ++i) {
        int pid = fork();
        if (pid < 0) {
            fprintf(stderr, "fork() failed: %s\n", strerror(errno));
            exit(1);
        } else {
            // son
            printf("son %d\n", getpid());
            fflush(stdout);
            _exit(0);
        }
    }

    getchar();

    int pid;
    while ((pid = wait(NULL)) > 0) {
        printf("%d\n", pid);
        fflush(stdout);
    }
}
```