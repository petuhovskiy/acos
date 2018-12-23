# Семинар 7

```c
#include <stdio.h>

int *v1;

double **v2[10];

// postfix
// []
// ()
// prefix
// *

double *(*v3)[10];

double (* (*v4)[15])(int (*par)[10]);

double func(int (*par)[10])
{
    printf("in func\n");
    return 0;
}

double (*v5[15])(int (*par)[10]);

int main()
{
    v4 = &v5;
    v5[0] =func;

    (*(*v4)[0])(NULL);
    (*v4)[0](NULL);
}
```

```c
#include <stdio.h>

int add1(int x)
{
    return x + 1;
}

int sub2(int x)
{
    return x - 2;
}

int main()
{
    int (*func)(int) = NULL;

    func = &add1;
    printf("%d\n), (*func)(100)); // 101

    func = sub2;
    printf("%d\n", func(44)); // 42
}
```

```c
#include <stdio.h>
#include <stdlib.h>

int arr[] = { 5, 2, 1, 10, -100, 400 };

int cmpfunc(const void *p1, const void *p2)
{
    int val1 = *(const int*) p1;
    int val2 = *(const int*) p2;

    // return val1 - val2; // WRONG!
    if (val1 < val2) {
        return -1;
    } else if (val1 > val2) {
        return 1;
    }
    return 0;
}

int main()
{
    qsort(arr, sizeof(arr) / sizeof(arr[0]), sizeof(arr[0]), &cmpfunc);
    for (int i = 0; i < sizeof(arr) / sizeof(arr[0]); ++i) {
        printf("%d\n", arr[i]);
    }
}
```

```c++
#include <iostream>

using namespace std;

struct LambdaImpl
{
    int ia, &ib;
    LambdaImpl(int a, int &b) : ia(a), ib(b) {}
    int operator() (int c, int d) {
        this->ib = c + d;
        return this->ia + this->ib + c + d;
    }
};

int main()
{
    int a = 10, b = 20;

    auto func = [a, &b](int c, int d) { b = c + d; return a + b + c + d; };

    cout << "func: " << func(39, 40) << endl;
    cout << a << ", " << b << endl;

    LambdaImpl func2{a, b};
    cout << "func: " << func2(30, 40) << endl;
    cout << a << "," << b << endl;
}
```

Теперь делаю низкоуровневый хак.

```c++
#include <iostream>

using namespace std;

struct F
{
    int a, b;
    F(int a, int b) : a(a), b(b) {}
    int sum() const { return a + b; }
};

// nm
extern "C" int _ZNK1F3sumEv(const struct F*);

int main()
{
    F f(33, 53);

    cout << f.sum() << endl;

    // int (*pfunc)(const struct F *) = NULL;
    // pfunc = (void*) &F::sum; // если как-то компилятор нагнуть // нагнуть не получилось
    int (*pfunc)(const struct F *) = _ZNK1F3sumEv; // вот так мы нагнули плюсы
    cout << pfunc(&f) << endl; 
}
```