// LAb1.cpp : Este archivo contiene la función "main". La ejecución del programa comienza y termina ahí.
//

#include <iostream>
#include <time.h>
using namespace std;


int main()
{

    clock_t t;
    clock_t t2;
    
    const int MAX = 1000;
    double** A = new double* [MAX];
    for (int i = 0; i < MAX; i++) {
        A[i] = new double[MAX];
    }

    double* x = new double[MAX];
    double* y = new double[MAX];
    int i, j;

    //INICIALIZACIÓN
    for (i = 0; i < MAX; i++) {
        for (j = 0; j < MAX; j++) {
            A[i][j] = 1;
        }
        x[i] = 2;
        y[i] = 0;
    }

    cout << "PRIMER LOOP\n";
    t = clock();
    
    for (i = 0; i < MAX; i++) {
        for (j = 0; j < MAX; j++) {
            y[i] += A[i][j] * x[j];
            //cout << y[i] << " ";
        }
        //cout << endl;
    }
    t = clock() - t;

    
    //asignar y=0
    for (i = 0; i < MAX; i++) {
        y[i] = 0;
    }

    cout << "SEGUNDO LOOP\n";

    t2 = clock();
    //SEGUNDO LOOP
    for (j = 0; j < MAX; j++) {
        for (i = 0; i < MAX; i++) {
            y[i] += A[i][j] * x[j];
           //cout << y[i] << " ";
        }
        //cout << endl;
    }
    t2 = clock() - t2;
    cout << "Tiempo segundo bucle: " << (1000.0 * t2) / CLOCKS_PER_SEC << " ms" << endl;
    cout << "Tiempo primer bucle: " << (1000.0 * t) / CLOCKS_PER_SEC << " ms" << endl;
}