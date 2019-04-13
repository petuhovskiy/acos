# Signals

Приходят в любой момент времени.

Список
```sh
$ kill -l
```

`SIGINT` (ctrl-C) и `SIGTERM` (из других процессов приходит) обычно стоит обрабатывать.

`SIGHUP` -- разорвано соединение с терминалом
(нормальная реакция -- завершиться).
Еще может значить рестарт (например, для всяких daemons).

```c
int main() {
    signal(SIGINT, SIG_IGN); // игнорирование
}
```

Завершение программы (его игнорировать нельзя)
```sh
$ kill -KILL <pid>
$ kill -9 <pid>
```

`SIGPIPE` -- часто нужно игнорить.
Если игнорить `SIGCHLD`, то зомби-процессы чистятся сами без `wait`.

Для обработки сигналов `signal` (легаси, лучше не юзать).
Следует использовать 
```c
void handler(int s) {
    // ...
}

int main() {
    struct sigaction sa = {
        .sa_flex = SA_RESTART, // без этого флага завершается системный вызов, а если флаг установлен, то системный вызов отработает
        .sa_handler = handler
    };
    sigaction(SIGINT, &sa, NULL);
    sigaction(SIGINT, &(struct sigaction) {...}, NULL);
}
```

### Асинхронно-безопасные (signal-safety) функции. 
Можно вызывать из хендлера.

`volatile sig_atomic_t flag = 0;` // можно юзать в хендлере.

`errno = EINTR` // Системный вызов завершился плохо (без SA_RESTART)


Обязательно надо все, что юзается в хендлерах, делать `volatile`.


Вот так плохо:
```c
volatile sig_atomic_t flag = 0;
void handler(int s) {
    flag = 1;
}

int main() {
    struct sigaction sa = {
        //.sa_flex = SA_RESTART,
        .sa_handler = handler
    };

    while (1) {
        int r = scanf(...);
        if (flag) {
            print("signal\n");
            sleep(10);
            // race condition! Потеряются сигналы.
            flag = 0;
        }

        if (!r) break;
        if (r == -1 && errno == EINTR) {
            continue;
        }
        if (r == -1) break;
        ...
    }
}
```

А вот так хорошо:
```c
volatile sig_atomic_t flag = 0;
void handler(int s) {
    flag = 1;
}

int main() {
    sigset_t ss;
    sigemptyset(&ss);
    sigaddset(&ss, SIGINT);
    sigaddset(&ss, SIGTERM);
    sigprocmask(SIG_BLOCK, &ss, NULL); // сигналы заблокируются, но сохраняются для послудующей обработки

    struct sigaction sa = {
        //.sa_flex = SA_RESTART,
        .sa_handler = handler
    };

    while (1) {
        sigprocmask(SIG_UNBLOCK, &ss, NULL);
        if (flag) {
            print("signal\n");
            // все еще остался race condition здесь
            flag = 0;
            continue;
        }
        int r = scanf(...);
        sigprocmask(SIG_BLOCK, &ss, NULL);
        if (flag) {
            print("signal\n");
            sleep(10);
            flag = 0;
        }

        if (!r) break;
        if (r == -1 && errno == EINTR) {
            continue;
        }
        if (r == -1) break;
        ...
    }
}
```
