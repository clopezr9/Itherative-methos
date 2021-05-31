import numpy as np
import time
from numpy import genfromtxt
from itertools import permutations 

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

def jacobi(A, b, N, ig, rc):
    
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
    i=0
    while (i < N or residual > rc):
        ig = (b - np.dot(R,ig)) / D
        residual = np.linalg.norm(np.matmul(A,ig)-b)
        i = i + 1
    return ig


residual_converg = 1e-8

my_data = genfromtxt('test.csv', delimiter=',')
n = my_data.shape[1]
A =  np.delete(my_data, n-1, 1)
b = my_data[:,n-1]

#A = [[2.0,10.0,-8.0], [3.0, 8.0, 13.0], [5.0,2.0, -3]]
#b = [4.0, 7.0, 1.0]
ig = [0.0,0.0,0.0]#initial guess
t = time.time()
N = 25 #number of iterations

sol = jacobi(A, b, N, ig, residual_converg)

print ("A:", A)

print ("b:", b)

print ("x:", sol)

print("time: ", (time.time())-t)