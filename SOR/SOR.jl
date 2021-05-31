using LinearAlgebra
using CSV
using DataFrames

function main(args)
    for i in 1:length(args)
        println(args[i], ":")
        a, b = readCSVFile(args[i])
        solver_sor(a, b)
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
    return A, b
    
end

function solver_sor(A, b) 
    X0 = zeros(size(A, 1))
    w = 0.5

    tolerancia = Float64(0.000001) 
    iteracionMax = Int(100)
    

    maxiter = Int(100)


    n = size(A, 1)


    @time begin
        norma = norm(b - A * X0, 2)
        iteracion = 0
        while !( norma <= tolerancia ||  iteracion > iteracionMax)

            if iteracion == maxiter
                println("Maximum number of iterations reached: $(iteracion)")
                break
            end

            for i = 1:n
                sigma = Float64(0)
                for j = 1:n
                    if j != i
                        sigma = sigma + A[i, j] * X0[j]
                    end
            
                end

                X0[i] = (1 - w) * X0[i] + ((w / A[i,i]) * (b[i] - sigma))
        
            end
            iteracion += 1
            norma = norm(b - A * X0, 2)
        end
        print("EXECUTION TIME: ")
    end
    # println(X0)
    
    end


main(ARGS)