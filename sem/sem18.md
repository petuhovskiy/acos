# IPC: pipe


Пример
```
$ ls | cat | cat | cat | cat | cat
$ ps -A | grep cat # все одновременно работают
$ ls -l /proc/<pid>/fd # можно увидеть пайпы
```

Можно делать так (нужен общий предок, который пайп сделает)
```c
#include <stdio.h>
#include <unistd.h>

int main() {
    int fd[2];
    pipe(fd); // fd[0] -- чтение, fd[1] -- запись. (пайп однонаправленный)

    if (!fork()) {
        dup2(fd[1], 1);
        close(fd[0]); close(fd[1]);
        execlp(argv[1], argv[1], NULL);
        _exit(1);
    }
    close(fd[1]);

    if (!fork()) {
        dup2(fd[0], 0);
        close(fd[0]); // чтобы не было дедлока
        execlp(argv[2], argv[2], NULL);
        _exit(1);
    }
    close(fd[0]);

    wait(NULL);
    wait(NULL);

    return 0;
}
```

`FD_CLOEXEC` -- флаг закрытия при `exec*`.
Такое работает везде.
```c
#include <stdio.h>
#include <unistd.h>

int main() {
    int fd[2];
    pipe(fd);
    fcntl(fd[0], F_SETFD, fcntl(fd[0], F_GETFD, 0) | O_CLOEXEC);
    fcntl(fd[1], F_SETFD, fcntl(fd[1], F_GETFD, 0) | O_CLOEXEC);

    if (!fork()) {
        dup2(fd[1], 1);
        execlp(argv[1], argv[1], NULL);
        _exit(1);
    }
    close(fd[1]);

    if (!fork()) {
        dup2(fd[0], 0);
        execlp(argv[2], argv[2], NULL);
        _exit(1);
    }

    close(fd[0]);

    wait(NULL);
    wait(NULL);

    return 0;
}
```

Есть вариант проще `pipe2`. Штука нестандартная, может не везде работать (под linux должно быть все хорошо). Надо добавить `-D_GNU_SOURCE`.
```c
#include <stdio.h>
#include <unistd.h>

int main() {
    int fd[2];
    pipe2(fd, O_CLOEXEC);

    if (!fork()) {
        dup2(fd[1], 1);
        execlp(argv[1], argv[1], NULL);
        _exit(1);
    }
    close(fd[1]);

    if (!fork()) {
        dup2(fd[0], 0);
        execlp(argv[2], argv[2], NULL);
        _exit(1);
    }
    close(fd[0]);

    wait(NULL);
    wait(NULL);
    return 0;
}
```

Кольцо из процессов.
```c
int main() {
    int fd[2], fd1[2];

    pipe(fd);
    pipe(fd1);

    if (!fork()) {
        dup2(fd[1], 1); close(fd[0]); close(fd[1]);
        dup2(fd1[0], 0); close(fd1[0]); close(fd1[1]);
        ...
    }
    close(fd[1]);
    close(fd1[0]);

    if (!fork()) {
        dup2(fd1[1], 1); close(fd1[0]); close(fd1[1]);
        dup2(fd[0], 0); close(fd[0]); close(fd[1]);
        ...
    }
    close(fd[0]);
    close(fd1[1]);
    
    return 0;
}
```

Не забывать про буферизацию (вспоминаем `fflush` в интерактивках).


Именованые пайпы
```bash
$ mkfifo pipe1 # именованый пайп (как файл)
$ mkfifo pipe2
$ ./w < ./pipe1 > ./pipe2 & ./w1 > ./pipe1 < ./pipe2 &
$ cat > ./pipe1
$ cat < ./pipe2
```
Открываются `open`, при этом `open` на `r/w` блокируется, пока кто-нибудь не откроет пайп на `w/r`.


