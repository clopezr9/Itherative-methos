import numpy as np
import sys
import time

def simple_gauss(n, augmented_matrix):
    start_time = time.time()
    # Applying Gauss Elimination
    x = np.zeros(n)

    for i in range(n):
        if augmented_matrix[i][i] == 0.0:
            return None

        for j in range(i+1, n):
            ratio = augmented_matrix[j][i]/augmented_matrix[i][i]
            for k in range(n+1):
                augmented_matrix[j][k] = augmented_matrix[j][k] - ratio * augmented_matrix[i][k]
                
    # Back Substitution
    x[n-1] = augmented_matrix[n-1][n]/augmented_matrix[n-1][n-1]

    for i in range(n-2,-1,-1):
        x[i] = augmented_matrix[i][n]

        for j in range(i+1,n):
            x[i] = x[i] - augmented_matrix[i][j]*x[j]

        x[i] = x[i]/augmented_matrix[i][i]
        
    print("\n" + "--- %s seconds ---" % (time.time() - start_time))
    # Displaying solution
    print('\nRequired solution is: ')
    for i in range(n):
        print('X%d = %f' %(i,x[i]), end = '\t')


sistema_1 = [[-1.2, 0, 0, 4, -1, 0, 61],
             [0, -1, 0, -1, 4, -1, 14],
             [4, -1, 1, -1, 0, 0, 8],
             [-1, 5, -1, 0, -1, 0, 5],
             [0, -1, 5, 0, 0, -1, 9],
             [0, 0, -1, 0, -1, 6, 23]]
simple_gauss(6, sistema_1)

