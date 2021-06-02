import numpy as np
import time
import sys
from itertools import permutations 

def main():
    args = sys.argv[1:]
    N = 100 #number of iterations
    residual_converg = 1e-8
    for arg in args:
        print("\n" + arg)
        a,b = read_csv_file(arg)
        jacobi(a, b, N, residual_converg)

def read_csv_file (file_name):
    my_data = np.genfromtxt(file_name, delimiter=',')
    n = my_data.shape[1]
    A =  np.delete(my_data, n-1, 1)
    b = my_data[:,n-1]
    return A, b

def jacobi(A, b, N, rc):
    
    ig = [0] * len(A)
    residual = np.linalg.norm(np.matmul(A,ig)-b)

    # Create a vector of the diagonal elements of A                                                                                                                                                
    # and subtract them from A                                                                                                                                                                     
    D = np.diag(A)
    R = A - np.diagflat(D)
    # Iterate for N times
    i = 0
    start_time = time.time_ns()
    while (i < N and residual > rc):
        ig = (b - np.dot(R,ig)) / D
        residual = np.linalg.norm(np.matmul(A,ig)-b)
        i = i + 1

    print("EXECUTION TIME:" + " %s nano seconds " % (time.time_ns() - start_time))


if __name__ == '__main__':
    main()