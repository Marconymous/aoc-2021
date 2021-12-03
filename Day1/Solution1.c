#include "stdio.h"
#include "stdlib.h"

#define IGNORE -1543

int main()
{
    FILE *file = fopen("input.txt", "r");
    char line[256];
    int previous = IGNORE;
    int increased = 0;

    while (fgets(line, sizeof(line), file))
    {
        int val = atoi(line);
        printf("Value => %d\n", val);

        if (previous != IGNORE && val > previous)
        {
            ++increased;
        }

        previous = val;
    }

    printf("Increased %d times", increased);

    return 0;
}