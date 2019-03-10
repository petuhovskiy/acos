# Семинар 4. Вещественные числа.

Хотим получить битовое представление вещественного типа `float`.

```c
float x;
unsigned ux;
scanf("%f", &x);

// UB version
// also bad because of
// strict aliasing rule (optimization)
ux = *(unsigned*)(&x);

// Nice version w/ memcpy
#include <string.h>
memcpy(&ux, &x, sizeof(ux));

// it's binary operations time
(ux >> 7) & 0x3F
          & 077
          & 0b111111
          & ((1u << 6)- 1)
```

Еще одна версия.

```c
union FU
{
    float fv;
    unsigned uv;
}
union FU u;
```

И еще один вариант.

```c
union FU
{
    float fv;
    struct {
        unsigned m:23;
        unsigned p:8;
        unsigned s:1;
    };
};
union FU u;
```

**Концептуально**.
