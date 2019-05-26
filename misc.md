
## asm

- [x86 Assembly Guide](http://flint.cs.yale.edu/cs421/papers/x86-asm/asm.html)
- [Опкоды](http://ref.x86asm.net/)
- [08-asm1](https://github.com/hseos/hseos-course/tree/master/2018/08-asm1)

## io

[reference](https://en.cppreference.com/w/c/io)

### Посимвольное

Пользуемся `getchar` (`fgetc`) и `putchar`.

Также multithread-unsafe версия `getchar_unlocked` и `putchar_unlocked`, которая может быть быстрее в 10 раз.

Пример посимвольного копирования:

```c
int c; // int, чтобы можно было отличать EOF
while ((c = getchar_unlocked()) != EOF) {
    putchar_unlocked(c);
}
```

### Форматированное

- [scanf](https://en.cppreference.com/w/c/io/fscanf)
- [printf](https://en.cppreference.com/w/c/io/fprintf)

## Приоритет операторов

Смотреть [здесь](https://en.cppreference.com/w/c/language/operator_precedence).

## tools

TODO: GNU Toolchain, etc...

### Течет память?

Проверить программу на утечки волгриндом:

`valgrind --leak-check=full ./a.out`

TODO: санитайзеры

### hexdump

`hexdump -C dostext.c`

### А также в ролях:

`vi, nano, tmux, mc`

## Code style

Основное [здесь](https://caos.ejudge.ru/style.html).

Главное - отступ 4 пробела, никаких табов.

```C
struct Foo
{
    int bar;
};

int *p;
char *str, **pstr;
```

_TODO: сделать конфиг для cpplint, ~~и выложить .vscode~~_