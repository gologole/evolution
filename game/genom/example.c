#include <stdlib.h>
#include <time.h>
#include<pthread.h>
// void helloFromC() {
//     printf("Hello from C!\n");
// }

// int** InitGenom(int ROWS, int COLS, int maxnum) {
//     srand(time(NULL));
//     int** matrix = (int**)malloc(ROWS * sizeof(int*));
//     for (int y = 0; y < ROWS; y++) {
//         matrix[y] = (int*)malloc(COLS * sizeof(int));
//         for (int x = 0; x < COLS; x++) {
//             int random =rand() % maxnum;
//             matrix[y][x] = random;
//         }
//     }
//     return matrix;
// }

// void freeMemory(void** genom, int rows) {
//     for (int i = 0; i < rows; i++) {
//         free(genom[i]);
//     }
//     free(genom);
// }
#include <stdio.h>
#include <stdlib.h>
#include <time.h>

void PrintMatrix(int** matrix, int ROWS, int COLS);

static unsigned int seedCounter = 0;

int** InitGenom(int ROWS, int COLS, int maxnum) {
    srand(time(NULL)+seedCounter++);
    int** matrix = (int**)malloc(ROWS * sizeof(int*));
    for (int y = 0; y < ROWS; y++) {
        matrix[y] = (int*)malloc(COLS * sizeof(int));
        for (int x = 0; x < COLS; x++) {
            int random = rand() % maxnum;
            matrix[y][x] = random;
        }
    }
    return matrix;
}

void PrintMatrix(int** matrix, int ROWS, int COLS) {
    for (int y = 0; y < ROWS; y++) {
        for (int x = 0; x < COLS; x++) {
            printf("%d ", matrix[y][x]);
        }
        printf("\n");
    }
}


void freeMemory(void** genom, int rows) {
    for (int i = 0; i < rows; i++) {
        free(genom[i]);
    }
    free(genom);
}
