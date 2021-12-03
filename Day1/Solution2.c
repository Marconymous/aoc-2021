#include "stdio.h"
#include "stdlib.h"

#define IGNORE -1543

int main()
{
    FILE *file = fopen("input.txt", "r");
    char line[256];
    int previous = IGNORE;
    int previousValues[3] = {IGNORE, IGNORE, IGNORE};
    int increased = 0;

    while (fgets(line, sizeof(line), file))
    {
        int newVal = atoi(line);

        previousValues[0] = previousValues[1];
        previousValues[1] = previousValues[2];
        previousValues[2] = newVal;

        int val = 0;
        int ignoreIter = 0;
        for (int i = 0; i < 3; ++i)
        {
            val += previousValues[i];
            if (previousValues[i] == IGNORE && ignoreIter == 0)
            {
                ignoreIter = 1;
            }
        }

        printf("Value => %d + %d\n", val, ignoreIter);

        if (previous != IGNORE && val > previous && !ignoreIter)
        {
            printf("INC++\n");
            ++increased;
        }

        if (!ignoreIter)
        {
            previous = val;
        }
    }

    printf("Increased %d times", increased);

    return 0;
}