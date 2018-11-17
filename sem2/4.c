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