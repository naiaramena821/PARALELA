#include <iostream>
#include <time.h>
using namespace std;

int main()
{
    clock_t t;
    int MAX = 100;
    double** A = new double* [MAX];
    double** B = new double* [MAX];
    double** resultado = new double* [MAX];
    int i, j, k;
    int l, m, n;
    int BLOCK = 30;

    for (i = 0; i < MAX; i++) {
        A[i] = new double[MAX];
        B[i] = new double[MAX];
        resultado[i] = new double[MAX];
    }
    // INICIALIZACIÓN
    for (i = 0; i < MAX; i++) {
        for (j = 0; j < MAX; j++) {
            A[i][j] = 2;
            B[i][j] = 4;
            resultado[i][j] = 0;
        }

    }

    t = clock();
    
    for (i = 0; i < MAX; i += BLOCK) {
        for (j = 0; j < MAX; j += BLOCK) {
            for (k = 0; k < MAX; k += BLOCK) {
                for (l = i; l < min(i + BLOCK, MAX); l++) {
                    for (m = j; m < min(j + BLOCK, MAX); m++) {
                        for (n = k; n < min(k + BLOCK, MAX); n++) {
                            resultado[l][m] += A[l][n] * B[n][m];
                        }
                    }
                }
            }
        }
    }

    t = clock() - t;

    cout << "Multiplicación de matrices: " << (1000.0 * t) / CLOCKS_PER_SEC << " ms" << endl;
}