#include <stdio.h>

int main()
{
    char buf[32];
    while (fgets(buf, sizeof(buf), stdin)) {
        printf("|%s|\n", buf);
    }
}