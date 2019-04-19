# Сокеты

Для тестирования можно использовать `nc`.

`nc` все что вы вводите отправляет в сокет, и все что сокет отправляет выводит в stdout.

> Открыть сокет в режиме клиента
```bash
nc www.ru 80

GET /
```
> Посмотреть статус подключений в сети
```bash
netstat -atn
```

> Запускаем сервер
```bash
nc -l -p 7777
```

> Подключаемся к созданному серверу
```bash
nc -6 localhost 7777
```

> 1.c

```c
#include <stdio.h>

int main()
{
    int x;
    while (scanf("%d", &x) == 1) {
        printf("%d\n", x + 1011 - 234);
        fflush(stdout);
    }
}
```

```bash
gcc 1.c -o1
./1
mkfifo fps
mkfifo fsp

# сервер
nc -l -p 7777 < fps > fsp

# программа
./1 > fps < fsp

# клиент
nc localhost 7777
```

> Более простой способ с пайпом
```bash
# сервер
nc -l -p 7777 < fps | ./1 > fps

# клиент
nc localhost 7777
```

Есть механизм локальных сокетов `man unix`, но мы будем рассматривать обычные сокеты.

`man getaddrinfo`

`man gethostname`

```c
#include <sys/types.h>
#include <sys/socket.h>
#include <netdb.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <errno.h>
#include <unistd.h>

int main(int argc, char *argv[])
{
    struct addrinfo *addrs = NULL;
    struct addrinfo h = {};

    h.ai_family =  AF_INET; # классический TCP/IP
    h.ai_socktype = SOCK_STREAM;

    int res = getaddrinfo(argv[1], argv[2], &h, &addrs);
    if (res < 0) {
        fprintf(stderr, "error: %s\n", gai_strerror(res));
        return 1;
    }

    for (struct addrinfo *p = addrs; p; p = p->ai_next) {
        printf("canon: %s\n", p->ai_canonname);
        struct sockaddr_in *si = (struct sockaddr_in*) p->ai_addr;
        printf("port: %d\n", ntohs(si->sin_port));
        printf("addr: %s\n", inet_ntoa(st->sin_addr));
    }

    int fd = socket(PF_INET, SOCK_STREAM, 0);

    if (connect(fd, addrs->ai_addr, addrs->ai_addrlen) < 0) {
        fprintf(stderr, "error: %s\n", strerror(errno));
        return 1;
    }

    write(fd, "GET /\r\n", sizeof("GET /\r\n") - 1);

    char c;
    int r;
    while ((r = read(fd, &c, 1)) == 1) {
        putchar(c);
    }

    freeaddrinfo(addrs);
}
```

`gcc -Wall -std=gnu11 resolve.c -oresolve`

`./resolve localhost http`

`./resolve google.com http`

Есть библиотеки для использования протоколов, например `libcurl`.

> server.c
```c
#include <stdio.h>
#include <errno.h>
#include <string.h>
#include <stdlib.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <netdb.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <unistd.h>

int main(int argc, char *argv[])
{
    int port = strtol(argv[1], NULL, 10);

    int fd = socket(PF_INET, SOCK_STREAM, 0);

    struct sockaddr_in s1;
    s1.sin_family = AF_INET;
    s1.sin_port = htons(port);
    s1.sin_addr.s_addr = INADDR_ANY;
    if (bind(fd, (struct sockaddr*) &s1, sizeof(s1)) < 0) {
        fprintf(stderr, "bind: %s\n", strerror(errno));
        return 1;
    }

    listen(fd, 5);

    struct sockaddr_in s2;
    socklen_t sl = sizeof(s2);

    int afd = accept(fd, (struct sockaddr*) &s2, &sl);

    if (afd < 0) {
        fprintf(stderr, "bind: %s\n", strerror(errno));
        return 1;
    }

    printf("accept: %s %d\n", inet_ntoa(s2.sin_addr), ntohs(s2.sin_port));

    char c;
    int r;

    // сделаем echo server
    while ((r = read(afd, &c, sizeof(c))) == 1) {
        write(afd, &c, sizeof(c));
    }
}
```

```bash
# сервер
./server 8899

# клиент
nc localhost 8899
```

> собственно вот простейшая реализация сервера