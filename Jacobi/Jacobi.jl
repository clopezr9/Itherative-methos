using Base: func_for_method_checked
using LinearAlgebra
using CSV
using DataFrames

function main(args)

    for i in 1:length(args)
        println(args[i], ":")
        N = 100
        residual_converge = 1e-8
        A, b = readCSVFile(args[i])
        jacobi(A, b, N, residual_converge)
    end
end

function readCSVFile(file_name) 
    df = CSV.read(file_name, DataFrame)
    i = names(df)
    l = nrow(df) + 1
    A = zeros(Float64, l, l)
    b = zeros(Float64, l, 1)

    for j in 1:l + 1
        if j == l + 1
            b[1,1] = parse(Float64, i[j])
        else 
            A[1,j] = parse(Float64, i[j])
        end
    end

    df1 = Matrix(df)
    
    for i in 1:l - 1
        for j in 1:l + 1
            if j == l + 1
                b[i + 1,1] = df1[i,j]
            else 
                A[i + 1,j] = df1[i,j]
            end
        end
    end

    return A,b 

end

function dominant(A)
    l = length(A)
    for i in 1:l
        dom = abs(A[i][i])
        coun = 0.0
        for j in 1:l
            coun = coun + abs(A[i][j])
        end
        coun = coun - dom
        if dom < coun
            return false
        end
    end
    return true
end

function division(a, D, l)
    r = zeros(Float64, l, 1)
    for j in 1:l
        for i in 1:l
            if i == j
                r[j] = a[j] / D[j,i]
            end
        end
    end
    return r
end

function dotProduct(A, b, l)
    r = zeros(Float64, l, 1)
    for i in 1:l
        aux = 0.0
        for j in 1:l 
            aux = aux + (b[j] * A[i,j])
        end
        r[i] = aux
    end
    return r
end

function jacobi(A, b, N, rc)
    @time begin

        ig = zeros(size(A, 1))
        l = size(A)[1]

        residual = norm(b - A * ig, 2)
        # println("residual: ", residual)

        D = Diagonal(A)
        # println("D:", D)

        R = A - D
        # println("R:", R)

        i = 0
        while (i < N || residual > rc)
            residual = norm(b - A * ig, 2)
            aux = (b - (dotProduct(R, ig, l)))
            ig = division(aux, D, l)
            i = i + 1
        end
        print("EXECUTION TIME: ")
    end
    println(ig)
end

main(ARGS)