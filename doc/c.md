# Прогаем на си

## Посимвольный ввод

Пользуемся `getchar` и `putchar`.

Также multithread-unsafe версия `getchar_unlocked` и `putchar_unlocked`, которая может быть быстере в 10 раз.

Пример посимвольного копирования:

```c
int c; // int, чтобы можно было отличать EOF
while ((c = getchar_unlocked()) != EOF) {
    putchar_unlocked(c);
}
```

## Форматированный ввод-вывод

TODO: написать здесь про всякие флаги