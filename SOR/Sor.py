import numpy as np
import time
from numpy import genfromtxt
import sys

def main():
    args = sys.argv[1:]
    for arg in args:
        print("\n"+ arg)
        phi = solver_sor(read_csv_file(arg))
    np.savetxt("../answers/3-0-res.csv", phi, delimiter = ",")

def read_csv_file(file_name):
    my_data = genfromtxt(file_name, delimiter=',')
    return my_data

def solver_sor(data):
    n = data.shape[1]
    A =  np.delete(data, n-1, 1)
    b = data[:,n-1]
    initial_guess = np.zeros(A.shape[0])
    residual_convergence = 1e-6
    omega = 1.5
    kmax = 0
    step = 0
    phi = initial_guess[:]
    t = time.time()

    residual = np.linalg.norm(np.matmul(A, phi) - b)
    while (residual > residual_convergence or kmax < 100) :
        for i in range(A.shape[0]):
            sigma = 0
            for j in range(A.shape[1]):
                if j != i:
                    sigma += A[i, j] * phi[j]
            phi[i] = (1 - omega) * phi[i] + (omega / A[i, i]) * (b[i] - sigma)
        residual = np.linalg.norm(np.matmul(A, phi) - b)
        step += 1
        kmax += 1
    print("EXECUTION TIME:" + " %s seconds " % (time.time() - t))
    print(phi)
    return phi

main()