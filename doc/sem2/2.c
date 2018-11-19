#include <stdio.h>
#include <stdlib.h>

int main()
{
    char *buf = NULL;

    while (scanf("%ms", &buf) == 1) {
        printf("|%s|\n", buf);
        free(buf);
    }
}