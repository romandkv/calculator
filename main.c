#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>


int get_type(char symbol) {
    if (symbol >= 'A' && symbol <= 'Z') {
        return 2;
    }
    if (symbol >= 'a' && symbol <= 'z') {
        return 1;
    }
    return 0;
}

char *str_to_lower(char *str) {
    int i;

    i = -1;
    while (str[++i] != '\0') {
        if (get_type(str[i]) != 2) {
            continue;
        }
        str[i] = str[i] + 32;
    }
    return str;
}


char *str_capitalize(char *str) {
    int i;

    i = 0;
    str = str_to_lower(str);
    while (str[i] != '\0') {
        if (get_type(str[i]) == 0) {
            i++;
            continue;
        }
        str[i] -= 32;
        while (str[i] != '\0' && get_type(str[i]) != 0) {
            i++;
        }
    }
    return str;
}

int main() {
    char str[] = {"ajbfsj nvcAnsv 1mfsTmkf fasfl;asfRsf;a"};
    char *str2 = str_capitalize(str);
    printf("str2 = \"%s\"", str2);
}
