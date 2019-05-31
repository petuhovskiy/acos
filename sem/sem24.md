# Пишем на **С++**

> m1.cpp

```c++
#include <pthread.h>
#include <stdio.h>
#include <thread>

pthread_mutex_t m = PTHREAD_MUTEX_INITIALIZER;

long long counter = 1000;

void func(void);

void *fthr1()
{
    for (int i = 0; i < 1000000; i++) {
        pthread_mutex_lock(&m);
        counter += 1;
        pthread_mutex_unlock(&m);
    }
    return NULL;
}

void fthr2(int &rval)
{
    for (int i = 0; i < 1000000; i++) {
        pthread_mutex_lock(&m);
        counter -= 1;
        pthread_mutex_unlock(&m);
    }
    rval = 100500;
}

int main()
{
    std::thread thr1(fthr1);

    int val{};
    std::thread thr2(fthr2, std::ref(val));

    thr1.join();
    thr2.join();

    printf("%ld\n", counter);
}
```

```bash
g++ -std=gnu++17 -Wall -Werror -pthread m1.cpp -o m1
```

> m2.cpp

```c++
#include <pthread.h>
#include <stdio.h>
#include <thread>
#include <vector>
#include <mutex>

pthread_mutex_t m = PTHREAD_MUTEX_INITIALIZER;

long long counter = 1000;

void func(void);

void *fthr1()
{
    for (int i = 0; i < 1000000; i++) {
        pthread_mutex_lock(&m);
        counter += 1;
        pthread_mutex_unlock(&m);
    }
    return NULL;
}

void fthr2(int &rval)
{
    for (int i = 0; i < 1000000; i++) {
        pthread_mutex_lock(&m);
        counter -= 1;
        pthread_mutex_unlock(&m);
    }
    rval = 100500;
}

int main()
{
    std::vector<std::thread> thrs;

    thrs.emplace_back(fthr1);
    // thrs.push_back(std::thread(fthr1));

    int val = 100;
    auto thr2 = std::thread(fthr2, std::ref(val));
    thrs.push_back(std::move(thr2));

    for (auto &t : thrs) {
        t.join();
    }

    printf("%ld\n", counter);
}
```

> m3.cpp

```c++
#include <stdio.h>
#include <thread>
#include <vector>
#include <mutex>

std::mutex m;
//std::recurive_mutex m

long long counter = 1000;

void *fthr1()
{
    for (int i = 0; i < 1000000; i++) {
        std::lock_guard lg{m};
        counter += 1;
    }
    return NULL;
}

void fthr2(int &rval)
{
    for (int i = 0; i < 1000000; i++) {
        // std::scoped_lock lg(m);
        std::unique_lock lg(m);
        counter -= 1;
    }
    rval = 100500;
}

int main()
{
    std::vector<std::thread> thrs;

    thrs.emplace_back(fthr1);
    // thrs.push_back(std::thread(fthr1));

    int val = 100;
    auto thr2 = std::thread(fthr2, std::ref(val));
    thrs.push_back(std::move(thr2));

    for (auto &t : thrs) {
        t.join();
    }

    printf("%ld\n", counter);
}
```

> m5.cpp

```c++
#include <stdio.h>
#include <thread>
#include <vector>
#include <mutex>
#include <atomic>

std::atomic<long long> counter = 1000;

void *fthr1()
{
    for (int i = 0; i < 1000000; i++) {
        counter += 1;
    }
    return NULL;
}

void fthr2(int &rval)
{
    for (int i = 0; i < 1000000; i++) {
        counter -= 1;
    }
    rval = 100500;
}

int main()
{
    std::vector<std::thread> thrs;

    thrs.emplace_back(fthr1);

    int val = 100;
    auto thr2 = std::thread(fthr2, std::ref(val));
    thrs.push_back(std::move(thr2));

    for (auto &t : thrs) {
        t.join();
    }

    printf("%ld\n", counter);
}
```

> m5_2.cpp

```c++
#include <stdio.h>
#include <thread>
#include <vector>
#include <mutex>
#include <atomic>

std::atomic<long long> counter = 1000;
// long long _Atomic counter = 1000;
// 

void *fthr1()
{
    for (int i = 0; i < 1000000; i++) {
        counter.fetch_add(1, std::memory_order_relaxed);
    }
    return NULL;
}

void fthr2(int &rval)
{
    for (int i = 0; i < 1000000; i++) {
        counter -= 1;
    }
    rval = 100500;
}

int main()
{
    std::vector<std::thread> thrs;

    thrs.emplace_back(fthr1);

    int val = 100;
    auto thr2 = std::thread(fthr2, std::ref(val));
    thrs.push_back(std::move(thr2));

    for (auto &t : thrs) {
        t.join();
    }

    printf("%ld\n", counter);
}
```

> m6.cpp

```c++
#include <thread>
#include <future>
#include <iostream>

int main()
{
    std::promise<int> p1;

    std::thread thr1([&f1]() {
        int val;

        std::cin >> val;
        p1.set_value(val);
    });
    thr1.detach();

    std::future<int> f1 = p1.get_future();
    int val = f1.get();
    std::cout << val << std::endl;
}
```

> m6_2.cpp

```c++
#include <thread>
#include <future>
#include <iostream>

int main()
{
    auto f = std::async(std::launch::async, []() {
        int val;

        std::cin >> val;
        p1.set_value(val);
    });

    int val = f.get();
    std::cout << val << std::endl;
}
```