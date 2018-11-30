# Семинар 5. Указатели и массивы

## Массивы

```c
int *p;
int a[32];

1) sizeof отличается
2) можем p = ; не можем a = ;

int x[] = {1, 2, 3};
```

```c
// НЕ ПИСАТЬ
int n;
int x[n]; // VLA
void func(int n, double x[n]); // тоже неоч
```

```c
void func(int n, const double *x);

double *b[3]; // массив из трех указателей, никаких гарантий
double a[3][2]; // классический массив

void func(int n, double a[][2]);
void func(int n, double a[][n]); // норм, это не VLA
```

## Строки

```c
char *s;
char sz[] = "a";

// нам нужен unsigned char, но у нас обычно char!
const char *s;
int x[256];
++x[*s++];

isalpha(*s) // UB
isalpha((unsigned char) *c) // нужно писать так
```

```c
strcopy(dst, src)
    while (*dst++ = *src++) {}

int strcmp(const char *c1, const char *s2) {
    while (*s1 == *s2 && *s1) {
        ++s1;
        ++s2;
    }
    return (unsigned char)*s1 - (unsigned char)*s2;
}
```

## Строку в число

Забудьте про `atoi`.

Нужно обрабатывать все ошибки.

```c
long strtol(char const *p, char **eptr, int base);
```

Если число переполняется, то число полностью считается до конца и вернет корректную ошибку в `errno`.

```c
char *ep = NULL; // указатель на первый некорректный символ
errno = 0;
long x = strtol(s, &ep, 10);
if (errno || *ep || ep == s) {
    
}
```