# Файлы отображаемые в память

> mmap1.c
```c
#include <stdio.h>
#include <stdlib.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <unistd.h>
#include <sys/mman.h>


int main(int argc, char **argv)
{
    const char *path = argv[1];

    int fd = open(path, O_RDONLY, 0);
    if (fd < 0) abort();

    // проверить что размер файла > 0 перед отображением

    struct stat stb;
    fstat(fd, &stb);

    if (!S_ISREG(stb.st_mode)) {
        // insert your error message here
        abort();
    }

    // не отображаем нулевой размер в память
    if (stb.st_size == 0) {
        return 0;
    }

    // size_t off_t
    size_t sz = stb.st_size;
    if (sz != stb.st_size) abort();

    void *ptr = mmap(
        NULL,
        stb.st_size,
        PROT_READ,
        MAP_PRIVATE,
        fd,
        0
    );
    if (ptr == MAP_FAILED) abort();

    close(fd):

    const int64_t *eptr = (const int64_t *)((char*) ptr + sz);
    const int64_t *cptr = (const int64_t *) ptr;
    int64_t minval = *cptr++;
    for(; cptr < eptr; ++cptr) {
        if (*cptr < minval) {
            minval = *cptr;
        }
    }


    munmap(ptr, stb.st_size);

    printf("%lld\n", (long long) minval);
}
```

```bash
gcc -Wall -Werror -std=gnu11 -D_FILE_OFFSET_BITS=64 -D_GNU_SOURCE mmap1.c

dd if=/dev/urandom of=file1 bs=16M count=1

./mmap1 file1
```

> kr05-5.c
```c
#include <stdio.h>
#include <stdint.h>
#include <stdlib.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <unistd.h>
#include <sys/mman.h>


int main(int argc, char **argv)
{
    const char *path = argv[1];

    int fd = open(path, O_RDWR, 0);
    if (fd < 0) abort();

    // проверить что размер файла > 0 перед отображением

    struct stat stb;
    fstat(fd, &stb);

    if (!S_ISREG(stb.st_mode)) {
        // insert your error message here
        abort();
    }

    // не отображаем нулевой размер в память
    if (stb.st_size == 0) {
        return 0;
    }

    // size_t off_t
    size_t sz = stb.st_size;
    if (sz != stb.st_size) abort();

    void *ptr = mmap(
        NULL,
        stb.st_size,
        PROT_READ | PROT_WRITE,
        MAP_SHARED,
        fd,
        0
    );
    if (ptr == MAP_FAILED) abort();

    close(fd):

    int64_t *eptr = (int64_t *)((char*) ptr + sz);
    int64_t *cptr = (int64_t *) ptr;
    int64_t *mptr = cptr++;
    for(; cptr < eptr; ++cptr) {
        if (*cptr < *mptr) {
            mptr = cptr;
        }
    }

    *mptr = -(uint64_t) *mptr;

    munmap(ptr, stb.st_size);

    printf("%lld\n", (long long) *mptr);
}
```

```bash
cmp file1 file2
```

> fib.c
```c
#include <stdio.h>
#include <stdint.h>
#include <stdlib.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <unistd.h>
#include <sys/mman.h>


int main(int argc, char **argv)
{
    int count = strtol(argv[1], NULL, 10);
    long long val0 = strtoll(argv[2], NULL, 10);
    long long val1 = strtoll(argv[3], NULL, 10);
    int fd = open("/dev/stdout", O_RDWR | O_CREAT | O_TRUNC, 0600);
    if (fd < 0) abort();

    if (count <= 2) abort();

    // проверить что размер файла > 0 перед отображением

    struct stat stb;
    fstat(fd, &stb);

    if (!S_ISREG(stb.st_mode)) {
        // insert your error message here
        abort();
    }

    if (ftruncate(fd, sz) < 0) abort();

    // size_t off_t
    size_t sz = count * sizeof(long long);

    void *ptr = mmap(
        NULL,
        stb.st_size,
        PROT_READ | PROT_WRITE,
        MAP_SHARED,
        fd,
        0
    );
    if (ptr == MAP_FAILED) {
        fprintf(stderr, "%s\n", strerror(errno));
        exit(1);
    }

    close(fd):

    long long *data = (long long *)ptr;
    data[0] = val0;
    data[1] = val1;

    for (int i = 2; i < count; i++) {
        data[i] = data[i - 1] + data[i - 2];
    }


    munmap(ptr, stb.st_size);
}
```

```bash
man mremap
man msync
man mincore
cat /proc/self/maps
```