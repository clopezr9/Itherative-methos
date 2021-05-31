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


def dominant(A):
    for i in range(len(A)):
        dom = abs(A[i][i])
        count = 0
        for j in range(len(A)):
            count = count + abs(A[i][j])
        count = count - dom
        if dom < count:
            return False
    return True
            
def make_dominant(A, b):
    #number of permutations
    C = [0,0]
    perm = list(permutations(A))
    #print(perm)
    permb = list(permutations(b))
    #print(permb)
    for i in range(len(perm)):
        if (dominant(perm[i])):
            C[0] = perm[i]
            C[1] = permb[i]
            return C
    return "Not possible"

def jacobi(A, b, N, rc):
    
    ig = [0] * len(A)
    residual = np.linalg.norm(np.matmul(A,ig)-b)
    C=[]
    if (dominant(A) != True):
        C = make_dominant(A, b)
        if C == "Not possible":
            return C
        else:
            A = C[0]
            b = C[1]
            

    # Create a vector of the diagonal elements of A                                                                                                                                                
    # and subtract them from A                                                                                                                                                                     
    D = np.diag(A)
    R = A - np.diagflat(D)
    # Iterate for N times
    i = 0
    start_time = time.time()
    while (i < N or residual > rc):
        ig = (b - np.dot(R,ig)) / D
        residual = np.linalg.norm(np.matmul(A,ig)-b)
        i = i + 1

    print("EXECUTION TIME:" + " %s seconds " % (time.time() - start_time))
    # Displaying solution
    print('Required solution is: ')
    for i in range(len(ig)):
        print('X%d = %f' %(i,ig[i]), end = '\t')


if __name__ == '__main__':
    main()