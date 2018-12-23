# Семинар 2

## 

```c
#include <stdio.h>
#include <stdlib.h>

int main()
{
    char *buf = NULL;

    while (scanf("%ms", &buf) == 1) {
        printf("|%s|\n", buf);
        free(buf);
    }
}
```

```c
#include <stdio.h>

int main()
{
    char buf[32];
    while (fgets(buf, sizeof(buf), stdin)) {
        printf("|%s|\n", buf);
    }
}
```

```c
#include <stdio.h>

int main()
{
    double val, sum = 0;
    
    // scanf возвращает кол-во успешно считанных
    // читаем пока успешно читается
    while (scanf("%lf", &val) == 1) {
        sum += val;
    }
    printf("%La\n", sum);
    printf("%Le\n", sum);
    printf("%Lf\n", sum);
    printf("%Lg\n", sum);

    printf("%20.3s\n", "12345");
}
```