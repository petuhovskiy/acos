# acos

Навигатор по АКОСу.

- [Main page](https://caos.ejudge.ru/)
- [Лекции 2017 года](https://github.com/hseos/hseos-course/tree/master/2017/00-lectures)
- ~~Лекции 2018 года~~
- [Семинары 2017 года](https://github.com/hseos/hseos-course/tree/master/2017)
- [Семинары 2018 года](https://github.com/hseos/hseos-course/tree/master/2018)
- [Живые семинары](https://github.com/petuhovskiy/acos/tree/master/sem#Живые-семинары)
- [Утилита для тестирования](https://github.com/petuhovskiy/acos/tree/master/tool#acos-tool)

## Семинары с гитхаба [2018](https://github.com/hseos/hseos-course/tree/master/2018)

- [01-linux intro](https://github.com/hseos/hseos-course/blob/master/2018/01-intro/01-intro.md)
- [01-terminal commands](https://github.com/hseos/hseos-course/blob/master/2018/01-intro/02-cmdline-part1.md)
- [01-bash scripts](https://github.com/hseos/hseos-course/blob/master/2018/01-intro/03-cmdline-part2.md)

- [02-stdio](https://github.com/hseos/hseos-course/tree/master/2018/02-stdio)

- [03-integers](https://github.com/hseos/hseos-course/tree/master/2018/03-integers)

- [04-floating](https://github.com/hseos/hseos-course/tree/master/2018/04-floating-point)

- [05-arrays](https://github.com/hseos/hseos-course/tree/master/2018/05-arrays)

- [06-pointers (aka strings)](https://github.com/hseos/hseos-course/tree/master/2018/06-pointers)

- [07-function-pointers](https://github.com/hseos/hseos-course/tree/master/2018/07-function-pointers)

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