n = 6
augmented_matrix = [[-1.2 0.0 0.0 4.0 -1.0 0.0 61.0],
         [0.0 -1.0 0.0 -1.0 4.0 -1.0 14.0],
         [4.0 -1.0 1.0 -1.0 0.0 0.0 8.0],
         [-1.0 5.0 -1.0 0.0 -1.0 0.0 5.0],
         [0.0 -1.0 5.0 0.0 0.0 -1.0 9.0],
         [0.0 0.0 -1.0 0.0 -1.0 6.0 23.0]]

x = zeros((n))

function gaussSimple(augmented_matrix, n)
    @time begin
        for i in 1:n 
            if augmented_matrix[i][i] == 0.0
                return 
            end
            j = 1
            aux_i = i + 1
            for j in aux_i:n
                ratio = augmented_matrix[j][i] / augmented_matrix[i][i]
                for k in 1:n + 1
                    augmented_matrix[j][k] = augmented_matrix[j][k] - ratio * augmented_matrix[i][k]
                end
            end    
        end
        x[n] = augmented_matrix[n][n + 1] / augmented_matrix[n][n]

        aux_n = n - 1
        for i = aux_n:-1:1
            x[i] = augmented_matrix[i][n + 1]
            aux_i = i + 1
            for j in aux_i:n
                aux_xi = x[i]
                x[i] = aux_xi - augmented_matrix[i][j] * x[j]
    
            end
            x[i] = x[i] / augmented_matrix[i][i]
        end
    end
    println(x) 
end

gaussSimple(augmented_matrix, n)
