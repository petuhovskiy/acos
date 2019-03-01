# Файловая система

* одно дерево  (в unix, в windows не совсем (буквы устройств))
* новое устройство монтируется в фс
* получить инфо о файле: `stat` или `access` (вызовы)

## Работа в каталогами

* стандартизированного вызова системного нет, но есть в либе C

```C
#include <dirent.h> // для каталогов
#include <stdlib.h>
#include <errno.h>

#include <limits.h>


int main()
{
    DIR *d = opendir("."); // текущий каталог
    if (!d) {
        fprintf(stderr, "opendir: %s: %s\n", ".", strerror(errno));
        exit(1);
    }

    printf("NAME_MAX:", NAME_MAX);
    struct dirent *dd;
    while ((dd = readdir(d))) { // указатели в буфер *d, поэтому для дальнейшего использования надо копировать в свои структуры
        printf("%lu %s\n", dd->d_ino, d->d_name); // номер индексного дескриптора (уникальный номер на устройстве, можно увидеть через `ls -li`) и имя
    }

    closedir(d);
}
```

```sh
$ touch `perl -e 'print "a" x 255'` -- норм
$ touch `perl -e 'print "a" x 256'` -- ошибка
```

Индексным дескриптором пользоваться нельзя там, где файловая система "склеивается" (поскольку разные устройства)
```sh
# делаем файловую систему
$ dd if=/dev/zero of=/tmp/fs bs=1M count=16
$ mkfs.ext2 /tmp/fs
$ mount /tmp/fs /mnt/loop -o loop
$ touch 1 2 3
$ ls -ali
# ..., а в opendir будет другая информация о номерах
```

### filename c буфером
```C
#include <dirent.h> // для каталогов
#include <stdlib.h>
#include <errno.h>

#include <limits.h>

#include <sys/types.h>
#include <sys/stat.h>


int main()
{
    DIR *d = opendir("."); // текущий каталог
    if (!d) {
        fprintf(stderr, "opendir: %s: %s\n", ".", strerror(errno));
        exit(1);
    }

    char buf[PATH_MAX];

    printf("NAME_MAX:", NAME_MAX);
    struct dirent *dd;
    while ((dd = readdir(d))) {
        printf("%lu %s\n", dd->d_ino, d->d_name);
        struct stat st;
        int res = snprintf(buf, sizeof(buf), "%s/%s", dirname, dd->d_name);
        if (res >= sizeof(buf)) {
            fprintf(stderr, "file name too long\n");
            continue;
        }
        if (stat(buf, &st) < 0) {
            fprintf(stderr, "stat failed\n");
            continue;
        }
    }

    closedir(d);
}
```



### filename c динамической памятью
```C
#include <dirent.h> // для каталогов
#include <stdlib.h>
#include <errno.h>

#include <limits.h>

#include <sys/types.h>
#include <sys/stat.h>


int main()
{
    DIR *d = opendir("."); // текущий каталог
    if (!d) {
        fprintf(stderr, "opendir: %s: %s\n", ".", strerror(errno));
        exit(1);
    }

    printf("NAME_MAX:", NAME_MAX);
    struct dirent *dd;
    while ((dd = readdir(d))) {
        printf("%lu %s\n", dd->d_ino, d->d_name);
        struct stat st;
        char buf* = NULL;
        asprintf(&buf, "%s/%s", dirname, dd->d_name); // -D_GNU_SOURCE
        if (res >= sizeof(buf)) {
            fprintf(stderr, "file name too long\n");
        } else if (stat(buf, &st) < 0) {
            fprintf(stderr, "stat failed\n");
            continue;
        }
        free(buf);
    }

    closedir(d);
}
```

```
if (S_ISDIR(st.st_mode)) {
    ...
}
```

* при рекурсии надо пропустить `.` и `..`
* не забываем про симлинки (для игнорирования симлинков можно lstat заюзать)
