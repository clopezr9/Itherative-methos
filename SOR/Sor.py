import numpy as np
import time
from numpy import genfromtxt

def sor_solver(A, b, omega, initial_guess, convergence_criteria, kmax):
    step = 0
    phi = initial_guess[:]
    residual = np.linalg.norm(np.matmul(A, phi) - b)

    while (residual > convergence_criteria or kmax < 100) :
        for i in range(A.shape[0]):
            sigma = 0
            for j in range(A.shape[1]):
                if j != i:
                    sigma += A[i, j] * phi[j]
            phi[i] = (1 - omega) * phi[i] + (omega / A[i, i]) * (b[i] - sigma)
        residual = np.linalg.norm(np.matmul(A, phi) - b)
        step += 1
        kmax += 1
    return phi

def main():
    
    residual_convergence = 1e-6
    omega = 0.5

    #en test3.csv va el args    
    my_data = genfromtxt('../test/'+'3-0.csv', delimiter=',')
    n = my_data.shape[1]
    A =  np.delete(my_data, n-1, 1)
    b = my_data[:,n-1]
    initial_guess = np.zeros(A.shape[0])

    t = time.time()
    kmax = 0

    phi = sor_solver(A, b, omega, initial_guess, residual_convergence, kmax)
    
    print(phi)
    tiempo = time.time()- t 
    print("El tiempo fue de [{0:6.4f}]".format(tiempo))
    np.savetxt("../answers/3-0-res.csv", phi, delimiter = ",")

main()