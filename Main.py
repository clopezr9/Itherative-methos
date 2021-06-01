import os
import sys
import numpy as np
from itertools import permutations 

def main():
    method = sys.argv[1]
    args = sys.argv[2:]
    
    if (method == 'SimpleGauss'):
        commands('SimpleGauss', args)
    elif (method == 'SOR'):
        commands('SOR', args)
    elif (method == 'Jacobi'):
        commands('Jacobi', args)


def commands (file, args):

    go_com = 'go run ' + file + '/' + file + '.go '
    julia_com = 'julia ' + file + '/' + file + '.jl '
    scala_run_com = 'scala ' + file + '/' + file + '.scala '
    python_com = 'python3 ' + file + '/' + file + '.py '

    for arg in args:
        go_com += arg + ' '
        julia_com += arg + ' '
        scala_run_com += arg + ' '
        python_com += arg + ' '

    print('\nLENGUAGE: GO')
    os.system(go_com)
    print("\nLENGUAGE: JULIA")
    os.system(julia_com)
    print("\nLENGUAGE: SCALA")
    os.system('scalac ' + file + '/' + file + '.scala ')
    os.system(scala_run_com)
    print("\nLENGUAGE: PYTHON")
    os.system(python_com)

if __name__ == '__main__':
    main()
