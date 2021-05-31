using CSV
using DataFrames

function main(args)
    for i in 1:length(args)
        println(args[i], ":")
        a, b = readCSVFile(args[i])
        gaussSimple(a, b)
    end
end

function readCSVFile(file_name) 
    df = CSV.read(file_name, DataFrame)
    i = names(df)
    l = nrow(df)+1
    augmented_matrix = zeros(Float64, l, l+1)

    for j in 1:l+1
        augmented_matrix[1,j] = parse(Float64,i[j])
    end

    df1 = Matrix(df)
    
    for i in 1:l-1
        for j in 1:l+1
            augmented_matrix[i+1,j] = df1[i,j]
        end
    end
    return augmented_matrix, l
end

function gaussSimple(augmented_matrix, n)
    x = zeros((n))
    @time begin
        for i in 1:n 
            if augmented_matrix[i,i] == 0.0
                return 
            end
            j = 1
            aux_i = i + 1
            for j in aux_i:n
                ratio = augmented_matrix[j, i] / augmented_matrix[i, i]
                for k in 1:n + 1
                    augmented_matrix[j, k] = augmented_matrix[j, k] - ratio * augmented_matrix[i, k]
                end
            end    
        end
        x[n] = augmented_matrix[n, n + 1] / augmented_matrix[n, n]

        aux_n = n - 1
        for i = aux_n:-1:1
            x[i] = augmented_matrix[i, n + 1]
            aux_i = i + 1
            for j in aux_i:n
                aux_xi = x[i]
                x[i] = aux_xi - augmented_matrix[i, j] * x[j]
    
            end
            x[i] = x[i] / augmented_matrix[i, i]
        end
        print("EXECUTION TIME: ")
    end
end

main(ARGS)
