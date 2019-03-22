# Семинар 17. Подпрограммы

> exec1.c
```c
#include <stdio.h>
#include <unistd.h>

#include <sys/wait.h>

int main()
{
    pid_t pid = fork();
    if (!pid) {
        printf("son\n"); fflush(stdout);
        execlp("busybox", "ls", "-la", "/", NULL);
        perror("execlp");
        _exit(1);
    }

    wait(NULL);
}
```

```bash
ldd `which busybox`
```

> exec2_0.c
```c
#include <stdio.h>
#include <unistd.h>

#include <sys/wait.h>

int main()
{
    pid_t pid = fork();
    if (!pid) {

        char *args[] = { "ls", "-l", "/", NULL };
        char *envs[] = { "TERM=", NULL };
        execve("/bin/ls", args, envs);

        perror("execve");
        _exit(1);
    }

    wait(NULL);
}

> exec2_1.c
```c
#include <stdio.h>
#include <unistd.h>

#include <sys/wait.h>

extern char **environ;

int main()
{
    pid_t pid = fork();
    if (!pid) {

        char *args[] = { "env", NULL };
        char *envs[] = { "TERM=", NULL };
        execve("/bin/env", args, environ);

        perror("execve");
        _exit(1);
    }

    wait(NULL);
}
```

> exec3.c
```c
#include <stdio.h>
#include <unistd.h>

#include <sys/wait.h>
#include <fcntl.h>

int main()
{
    pid_t pid = fork();
    if (!pid) {
        int fd = open("out", O_WRONLY | O_CREAT | O_TRUNC, 0600);
        if (fd < 0) {
            perror("open");
            _exit(1);
        }

        dup2(fd, 1); close(fd);

        execlp("ls", "ls", "-la", "/proc/self/fd", NULL);
        perror("execlp");
        _exit(1);
    }

    wait(NULL);
}
```

> exec4.c
```c
#include <stdio.h>
#include <unistd.h>

#include <sys/wait.h>
#include <fcntl.h>

int main()
{
    pid_t pid = fork();
    if (!pid) {
        int fd = open("out", O_WRONLY | O_CREAT | O_TRUNC | O_CLOEXEC, 0600);
        if (fd < 0) {
            perror("open");
            _exit(1);
        }

        dup2(fd, 1); //close(fd);
        dup3(fd, 2, O_CLOEXEC);
        close(0);

        execlp("ls", "ls", "-la", "/proc/self/fd", NULL);
        perror("execlp");
        _exit(1);
    }

    wait(NULL);
}
```

`-D_GNU_SOURCE`

`man chdir`

`man system`

```bash
{ ls ; }
( ls ; )

{ A=X; } ; echo $A
( A=Y; ) ; echo $A
```

> py1.py
```python
#! /usr/bin/python3

print ("yeah")
```

> print.c
```c
#include <stdio.h>

int main(int argc, char **argv)
{
    printf("%d\n", geteuid());
    for (int i = 0; i < argc; i++) {
        printf("[%d] %s\n", i, argv[i]);
    }
}
```

> script
```
#! ...../print -x -y

a
b
c
```

```bash
./script 1 2 3 4 5 6

# result:
# print
# -x -y
# ./script
# 1
# 2
# 3
# 4
```