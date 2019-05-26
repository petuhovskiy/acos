# Нити, как их создавать и что с ними делать

У всех нитей в рамках одного процесса общая память, а на каждую нить выделяются свои только регистры. Поэтому работа с нитями, отладка и так далее чрезвычайно сложна, даже по сравнению с многопроцессорными случаями.

Библиотека pthread - стандарт для нитей в Си и UNIX, и большинстве программ.

> thr1.c

```c
#include <pthread.h>
#include <stdio.h>
#include <unistd.h>
#include <string.h>

void *thrfunc(void *ptr)
{
    printf("in thread\n");
    pthread_exit(strdup("exited"));
    sleep(100);
    return strdup("success");
}

int main()
{
    pthread_t tid = 0;

    int err = pthread_create(&tid, NULL, thrfunc, NULL);
    printf("in main\n");

    void *res = NULL;
    pthread_join(tid, &res);
    printf("result: %s\n", res);
}
```

```bash
gcc thr1.c -Wall -O2 -std=gnu11 -pthread -othr1
```

```bash
ps      # показывает одну строку на процесс
ps -Ta  # показывает одну строку на нить
```

> thr2.c

```c
#include <pthread.h>
#include <stdio.h>
#include <unistd.h>
#include <string.h>

void *thrfunc(void *ptr)
{
    int *pserial = (int*) ptr;
    printf("%d\n", *pserial);
    pthread_t tid = 0;
    ++*pserial;
    if (*pserial == 4) return NULL;
    pthread_create(&tid, NULL, thrfunc, pserial);
    pthread_join(tid, NULL);
    return NULL;
}

int main()
{
    pthread_t tid = 0;

    int serial = 0;
    int err = pthread_create(&tid, NULL, thrfunc, &serial);
    pthread_join(tid, NULL);
}
```

> thr3.c

```c
#include <pthread.h>
#include <stdio.h>
#include <unistd.h>
#include <string.h>

void *thrfunc(void *ptr)
{
    int *pserial = (int*) ptr;
    printf("%d\n", *pserial);
    pthread_t tid = 0;
    ++*pserial;
    if (*pserial == 4) return NULL;
    pthread_create(&tid, NULL, thrfunc, pserial);
    pthread_join(tid, NULL);
    return NULL;
;}

int main()
{
    pthread_t tid = 0;

    int serial = 0;
    pthread_attr_t attr;
    pthread_attr_init(&attr);
    // здесь можно задать аттрибуты для задания минимального размера стека
    // может быть полезно в задаче 22-3
    int err = pthread_create(&tid, &attr, thrfunc, &serial);
    pthread_attr_destroy(&attr);
    pthread_join(tid, NULL);
}
```