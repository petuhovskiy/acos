# Семинар 3

```c
const int f = 10; // read-only variable
```

```c
enum { F = 10 }; // true const
```

```c
// >= 8
char
signed char
unsigned char

// >= 16
short
unsigned short

// >= 16
int
unsigned int

// >= 32
long
unsigned long

// >= 64
long long
unsigned long long
```

```c
POSIX/LINUX/opengronp

char = 8
int = 32
```

```c
#include <inttype.h>
#include <stdint.h>

int8_t, int16_t
uint8_t, uint16_t
```

```c
// проверка на переполнение
unsigned a, b;

c = a + b // c < a
c = a - b // c > a

c = a * b // ?
```

```c
int a, b;

(a ^ b) < 0 // проверка на числа разных знаков
((unsigned)a + b)^a // проверка на переполнение
```

```c
__builtin_add_overflow
__builtin_sub_overflow
__builtin_mul_overflow
__builtin_div_overflow
```

## Минимальное и максимальное значение

```c
unsigned
-1u // max value

signed
(int) (~0u >> 1) // max value
(int) ((~0u >> 1) + 1) // min value

INT_MIN
```

## bitcount

```c
__builtin_popcount (плохо, не использовать)
```

Генри С. Уоррен мл. "Алгоритмические трюки для программистов"

```c
a&(a-1) // младший бит

```