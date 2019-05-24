# Семинар **23**

Синхронизация нитей и организация взаимодействия.

## Мьютексы

> m1_1.c

```c
#include <pthread.h>
#include <stdio.h>

pthread_mutex_t m = PTHREAD_MUTEX_INITIALIZER;

long long counter = 1000;

void func(void);

void *fthr1(void *ptr)
{
    for (int i = 0; i < 1000000; i++) {
        func();
        counter += 1;
    }
    return NULL;
}

void *fthr2(void *ptr)
{
    for (int i = 0; i < 1000000; i++) {
        func();
        counter -= 1;
    }
    return NULL;
}

int main()
{
    pthreadd_t thr1 = 0, thr2 = 0;
    pthread_create(&thr1, NULL, fthr1, NULL);
    pthread_create(&thr2, NULL, fthr2, NULL);
    pthread_join(thr1, NULL);
    pthread_join(thr2, NULL);
    printf("%ld\n", counter);
}
```

```bash
gcc -O0 -Wall -std=gnu11 -pthread m1.c -om1
```

```bash
clang -g -O2 -Wall -std=gnu11 -pthread m1.c -fsanitize=thread -om1
```

> m1.c

```c
#include <pthread.h>
#include <stdio.h>

pthread_mutex_t m = PTHREAD_MUTEX_INITIALIZER;

long long counter = 1000;

void func(void);

void *fthr1(void *ptr)
{
    for (int i = 0; i < 1000000; i++) {
        func();
        pthread_mutex_lock(&m);
        counter += 1;
        pthread_mutex_unlock(&m);
    }
    return NULL;
}

void *fthr2(void *ptr)
{
    for (int i = 0; i < 1000000; i++) {
        func();
        pthread_mutex_lock(&m);
        counter -= 1;
        pthread_mutex_unlock(&m);
    }
    return NULL;
}

int main()
{
    pthreadd_t thr1 = 0, thr2 = 0;
    pthread_create(&thr1, NULL, fthr1, NULL);
    pthread_create(&thr2, NULL, fthr2, NULL);
    pthread_join(thr1, NULL);
    pthread_join(thr2, NULL);
    printf("%ld\n", counter);
}
```

> m2_1.c

```c
#include <pthread.h>
#include <stdio.h>
#include <sched.h>

// pthread_mutex_t m = PTHREAD_MUTEX_INITIALIZER;

long long counter = 1000;

void func(void);

void *fthr1(void *ptr)
{
    for (int i = 0; i < 1000; i++) {
        pthread_mutex_lock(&m);

        printf("0 %lld\n", (long long) counter);
        ++counter;

        pthread_mutex_unlock(&m);

        // sched_yield();
    }
    return NULL;
}

void *fthr2(void *ptr)
{
    for (int i = 0; i < 1000; i++) {
        pthread_mutex_lock(&m);

        printf("1 %lld\n", (long long) counter);
        ++counter;

        pthread_mutex_unlock(&m);

        // sched_yield();
    }
    return NULL;
}

int main()
{
    pthreadd_t thr1 = 0, thr2 = 0;
    pthread_create(&thr1, NULL, fthr1, NULL);
    pthread_create(&thr2, NULL, fthr2, NULL);
    pthread_join(thr1, NULL);
    pthread_join(thr2, NULL);
    printf("%ld\n", counter);
}
```

> m2_2.c

```c
#include <pthread.h>
#include <stdio.h>
#include <sched.h>

pthread_mutex_t m = PTHREAD_MUTEX_INITIALIZER;
pthread_cond_t c0 = PTHREAD_COND_INITIALIZER;
pthread_cond_t c1 = PTHREAD_COND_INITIALIZER;

long long counter = 1000;

void func(void);

void *fthr1(void *ptr)
{
    for (int i = 0; i < 1000; i++) {
        pthread_cond_wait(&c0, &m);

        printf("0 %lld\n", (long long) counter);
        ++counter;
        pthread_cond_signal(&c1);
    }
    return NULL;
}

void *fthr2(void *ptr)
{
    for (int i = 0; i < 1000; i++) {
        pthread_cond_wait(&c1, &m);

        printf("1 %lld\n", (long long) counter);
        ++counter;
        pthread_cond_signal(&c0);
    }
    return NULL;
}

int main()
{
    pthreadd_t thr1 = 0, thr2 = 0;
    pthread_create(&thr1, NULL, fthr1, NULL);
    pthread_create(&thr2, NULL, fthr2, NULL);
    pthread_cond_signal(&c0);
    pthread_join(thr1, NULL);
    pthread_join(thr2, NULL);
    printf("%ld\n", counter);
}
```

> m2.c

```c
#include <pthread.h>
#include <stdio.h>
#include <sched.h>

pthread_mutex_t m = PTHREAD_MUTEX_INITIALIZER;

pthread_cond_t c0 = PTHREAD_COND_INITIALIZER;
int f0;

pthread_cond_t c1 = PTHREAD_COND_INITIALIZER;
int f1;

long long counter = 1000;

void func(void);

void *fthr1(void *ptr)
{
    for (int i = 0; i < 1000; i++) {
        pthread_mutex_lock(&m);
        while (!f0) {
            pthread_cond_wait(&c0, &m);
        }
        f0 = 0;
        pthread_mutex_unlock(&m);

        printf("0 %lld\n", (long long) counter);
        ++counter;

        pthread_mutex_lock(&m);
        f1 = 1;
        pthread_cond_signal(&c1);
        pthread_mutex_unlock(&m);
    }
    return NULL;
}

void *fthr2(void *ptr)
{
    for (int i = 0; i < 1000; i++) {
        pthread_mutex_lock(&m);
        while (!f1) {
            pthread_cond_wait(&c1, &m);
        }
        f1 = 0;
        pthread_mutex_unlock(&m);

        printf("1 %lld\n", (long long) counter);
        ++counter;

        pthread_mutex_lock(&m);
        f0 = 1;
        pthread_cond_signal(&c0);
        pthread_mutex_unlock(&m);
    }
    return NULL;
}

int main()
{
    pthreadd_t thr1 = 0, thr2 = 0;
    pthread_create(&thr1, NULL, fthr1, NULL);
    pthread_create(&thr2, NULL, fthr2, NULL);
    
    pthread_mutex_lock(&m);
    pthread_cond_signal(&c0);
    f0 = 1;
    pthread_mutex_unlock(&m);
    
    pthread_join(thr1, NULL);
    pthread_join(thr2, NULL);
    printf("%ld\n", counter);
}
```

> m2.c

```c
#include <pthread.h>
#include <stdio.h>
#include <sched.h>

pthread_mutex_t m = PTHREAD_MUTEX_INITIALIZER;

pthread_cond_t c0 = PTHREAD_COND_INITIALIZER;
int f0;

pthread_cond_t c1 = PTHREAD_COND_INITIALIZER;
int f1;

long long counter = 1000;

void func(void);

void *fthr1(void *ptr)
{
    for (int i = 0; i < 1000; i++) {
        pthread_mutex_lock(&m);
        while (!f0) {
            pthread_cond_wait(&c0, &m);
        }
        f0 = 0;
        pthread_mutex_unlock(&m);

        printf("0 %lld\n", (long long) counter);
        ++counter;

        pthread_mutex_lock(&m);
        f1 = 1;
        pthread_cond_signal(&c1);
        pthread_mutex_unlock(&m);
    }
    return NULL;
}

void *fthr2(void *ptr)
{
    for (int i = 0; i < 500; i++) {
        pthread_mutex_lock(&m);
        while (!f1) {
            pthread_cond_wait(&c1, &m);
        }
        f1 = 0;
        pthread_mutex_unlock(&m);

        printf("%d %lld\n", *(int*)ptr, (long long) counter);
        ++counter;

        pthread_mutex_lock(&m);
        f0 = 1;
        pthread_cond_signal(&c0);
        pthread_mutex_unlock(&m);
    }
    return NULL;
}

int main()
{
    pthreadd_t thr1 = 0, thr2 = 0;
    pthread_create(&thr1, NULL, fthr1, NULL);
    pthread_create(&thr2, NULL, fthr2, &v1);
    pthread_create(&thr3, NULL, fthr2, &v2);
    
    pthread_mutex_lock(&m);
    pthread_cond_signal(&c0);
    f0 = 1;
    pthread_mutex_unlock(&m);
    
    pthread_join(thr1, NULL);
    pthread_join(thr2, NULL);
    pthread_join(thr3, NULL);
    printf("%ld\n", counter);
}
```