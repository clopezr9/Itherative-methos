using LinearAlgebra
using CSV
using DataFrames
@time begin
    

df = CSV.read("../test/3-0.csv", DataFrame)
i = names(df)
l = nrow(df)+1
A = zeros(Float64, l, l)
b = zeros(Float64, l, 1)
for j in 1:l+1
    if j == l+1
        b[1,1] = parse(Float64,i[j])
    else
        A[1,j] = parse(Float64,i[j])
    end
end
df1 = Matrix(df)

for i in 1:l-1
    for j in 1:l+1
        if j == l+1
            b[i+1,1] = df1[i,j]
        else 
            A[i+1,j] = df1[i,j]
        end
    end    
end
    
X0 = zeros(size(A,1))   # initial guess
w = 1.5#1.25

tolerancia = Float64(0.000001) 
iteracionMax = Int(100)

tol = Float64(0.00001)
maxiter = Int(100)

x = copy(X0)

iter = Int(0)

n = size(A,1)

norma =norm(b - A * X0,2)


    
diferencia = [1. 1. 1.]
    
errado = tolerancia * 2
iteracion = Int(0)


while !( norma <= tolerancia || iteracion > iteracionMax)
    if iter == maxiter
        println("Maximum number of iterations reached: $(iter)")
         break
    end

    for i = 1:n
        sigma = Float64(0)
        for j = 1:n
            if j != i
                sigma = sigma + A[i, j]*X0[j]
            end
            
        end
        # ALTERNATIVE IMPLEMENTATION
        global X0[i] = (1-w)*X0[i] + ((w/A[i,i]) * (b[i] - sigma))
        
    end
    global iter += 1
    global norma = norm(b - A * X0,2)
end

X = X0
#println("Error: ", errado, "--- Iteraciones hechas: ", iteracion)
println("")
println("Respuesta X: ", X)
println("")
#println("Verificar A*X=B: ", revision)
println("")
print("Tiempo de ejecución: ")
end