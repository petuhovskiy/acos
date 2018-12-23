# Семинар 6

## Используемая память

`cat /proc/self/status`

## Выделения

Сначала просто выделим:

```c
#include <stdio.h>
#include <stdlib.h>

int main()
{
    int n;
    scanf("%d", &n);

    printf("%zu\n", n * sizeof(int));
    int *ptr = malloc(n * sizeof(int));
    if (!ptr) {
        fprintf(stderr, "malloc failed\n");
        exit(1);
    }
    // ptr = realloc(ptr, (n * 2) * sizeof(*ptr2));
    free(ptr);
}
```

Самая простая реализация вектора: (и неправильная, нету проверок на ошибки)

```c
#include <stdio.h>
#include <stdlib.h>

int main()
{
    int x;
    int *ptr = NULL;
    size_t a = 0, u = 0;

    while (scanf("%d", &x) == 1) {
        if (u == a) {
            if (!(a *= 2)) a = 4; // CHECK for OVF
            ptr = realloc(ptr, a * sizeof(*ptr)); // CHECK for NULL
        }
    }
}
```

## Односвязный список

Какой-то интрузив:

```c
struct Data
{
};

struct Item
{
    struct Item *prev, *next;
    struct Data *data;
};
```

И инклузив (норм):

```c
struct Item
{
    struct Item *next;
    int data;
};

int main()
{
    int x;
    struct Item *head = NULL;

    while (scanf("%d", &x) == 1) {
        struct Item *ptr = calloc(1, sizeof(*ptr));
        // FIXME: check for NULL
        ptr->data = x;
        ptr->next = head;
        head = ptr;
    }

    for (struct Item *ptr = head; ptr; ptr = ptr->next) {
        printf("%d\n", ptr->data);
    }

    for (struct Item *ptr = head; ptr; ) {
        struct Item *tmp = ptr->next;
        free(ptr);
        ptr = tmp;
    }
}
```

## Двусвязный список

Очень лень.

```c
struct Item
{
    struct Item *prev, *next;
    int data;
};

struct List
{
    struct Item *head, *tail;
};

void insert(struct List *list, struct Item *place, struct Item *ptr)
{
    // здесь мне нужно сделать много проверок
    ptr->next = place;
    if (place) {
        ptr->prev = place->prev;
    } else {
        ptr->prev = list->tail;
        list->tail = ptr;
    }
    
    if (ptr->prev) {
        ptr->prev->next = ptr;
    } else {
        list->head = ptr;
    }

    if (ptr->next) {
        ptr->next->prev = ptr;
    } else {
        list->tail = ptr;
    }
}

int main()
{
    int x;
    struct Item *head = NULL;
    struct Item *tail = NULL;

    while (scanf("%d", &x) == 1) {
        struct Item *ptr = calloc(1, sizeof(*ptr));
        // FIXME: check for NULL
        ptr->data = x;
        ptr->next = head;
        if (head) {
            head->prev = ptr;
        } else {
            tail = ptr;
        }
        head = ptr;
    }

    for (struct Item *ptr = head; ptr; ptr = ptr->next) {
        printf("%d\n", ptr->data);
    }

    for (struct Item *ptr = head; ptr; ) {
        struct Item *tmp = ptr->next;
        free(ptr);
        ptr = tmp;
    }
}
```

## Что-то с деревьями

```c
struct Node
{
    struct Node *left, *right;
    int data;
}

struct Node *
insert(struct Node *root, int data)
{
    if (!root) {
        root = calloc(1, sizeof(*root));
        root->data = data;
    } else if (root->data == data) {
    } else if (root->data < data) {
        root-> right = insert(root->right, data);
    } else {
        root->left = insert(root->left, data);
    }
    return root;
}

int main()
{
    struct Node *root = NULL;
    int x;

    while (scanf("%d", &x) == 1) {
        root = insert(root, x);
    }
}
```

## И немного с юникодом

```c
#include <stdio.h>
#include <wchar.h>
#include <locale.h>

int main()
{
    wchar_t c;

    setlocale(LC_ALL, "");
    while ((c = getwc(stdin)) != WEOF) {
        putwc(c, stdout);
    }
}
```