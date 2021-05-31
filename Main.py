import os
import sys

def main():
    args = sys.argv[1:]
    print(args)
    
    if (args[0] == 'SimpleGauss'):
        commands('SimpleGauss', args)
    elif (args[0] == 'SOR'):
        commands('SOR', args)
    elif (args[0] == 'Jacobi'):
        commands('Jacobi', args)

def commands (file, args):
    go_com = 'go run ' + file + '/' + file + '.go '
    julia_com = 'julia ' + file + '/' + file + '.jl '
    scala_run_com = 'scala ' + file + '/' + file + '.scala '
    python_com = 'python ' + file + '/' + file + '.py '

    for arg in args:
        go_com += arg + ' '
        julia_com += arg + ' '
        scala_run_com += arg + ' '
        python_com += arg + ' '

        print('LENGUAGE: GO')
        os.system(go_com)
        print("LENGUAGE: JULIA")
        os.system(julia_com)
        print("LENGUAGE: SCALA")
        os.system('scalac' + file + '/' + file + '.scala ')
        os.system(scala_run_com)
        print("LENGUAGE: PYTHON")
        os.system(python_com)

if __name__ == '__main__':
    main()
