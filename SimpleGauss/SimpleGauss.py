import numpy as np
import time
import sys

def main():
    args = sys.argv[1:]
    for arg in args:
        print("\n" + arg)
        simple_gauss(read_csv_file(arg))

def read_csv_file (file_name):
    my_data = np.genfromtxt(file_name, delimiter=',')
    n = my_data.shape[1]
    return n-1, my_data


def simple_gauss(df):
    n = df[0]
    augmented_matrix = df[1]
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
        
    print("EXECUTION TIME:" + " %s seconds " % (time.time() - start_time))
    # Displaying solution
    print('Required solution is: ')
    for i in range(n):
        print('X%d = %f' %(i,x[i]), end = '\t')


if __name__ == '__main__':
    main()
