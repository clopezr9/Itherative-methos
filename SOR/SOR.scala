import scala.io.Source
import javax.print.DocFlavor.INPUT_STREAM
import javax.print.attribute.standard.PrinterURI
import math._
import java.io.File
object SOR {
	def main(args: Array[String]): Unit = {
        for(arg<-args){
            var param = readCSVFile(arg)
            var (aug_matrix) = param
            sor_solver(aug_matrix)
        }
    }
    def readCSVFile(file: String): (Array[Array[Double]]) = {
        var C = io.Source.fromFile(file).getLines().map(_.split(",").map(_.trim.toDouble)).toArray
        var n = C.length
        var m = C(0).length
        var aug_matrix = Array.ofDim[Double](n, m)

        for(i <- 0 to n-1){
           for (j <- 0 to n){
              aug_matrix(i)(j) = C(i)(j)
           }
        }
        return (aug_matrix)
    }
    def sor_solver(aug_matrix: Array[Array[Double]]): Unit = {
        
        var B = Array.ofDim[Double](aug_matrix.length,1)
        var BtoArray = Array.fill(aug_matrix.length){0.0}
       
          for(i <- 0 to B.length-1){
                  B(i)(0) = aug_matrix(i)(aug_matrix(1).length-1)
                  BtoArray.update(i,B(i)(0))
          }
        var A = Array.ofDim[Double](aug_matrix.length,aug_matrix(1).length-1)

        for(i <- 0 to aug_matrix.length-1){
             for(j <- 0 to aug_matrix(1).length-2){
                A(i)(j) = aug_matrix(i)(j)
             }
          }
          for(i <- 0 to aug_matrix.length-1){
             for(j <- 0 to aug_matrix(1).length-2){
                A(i)(j) = aug_matrix(i)(j)
             }
          }

	    var tol = Array.ofDim[Double](aug_matrix.length,aug_matrix.length)
        var X0  = Array.ofDim[Double](aug_matrix.length)
        
        var matMul = A.zip(tol) map (_.zipped map (_ * _)) map (_.sum)
        
        var aiuda = distance(matMul,BtoArray)

        var tolerancia: Double = 0.000001
        var iteracionMax: Int = 100
        var Omega: Double = 0.5
        
        var n: Int = aug_matrix.size    
        var m: Int = aug_matrix.size

        var X = X0
        var iteracion : Int = 0
        
        var time_ini = System.nanoTime()

        while ((aiuda > tolerancia || iteracion > 100)){
            for(fila <- 0 to n - 1){
                var suma: Double = 0
                var sigma: Double = 0
                for(columna <- 0 to m - 1){
                    if(fila != columna){
                        sigma = sigma + ((A(fila)(columna)) * (X(columna))) 
                    }
                }
                var phi = (1-Omega) * X(fila) + (Omega / A(fila)(fila)) * (B(fila)(0)-sigma)

                X.update(fila, phi)                
                tol.update(fila,X)
            }
            var matMul = A.zip(tol) map (_.zipped map (_ * _)) map (_.sum)

            aiuda = distance(matMul,BtoArray)
            iteracion = iteracion + 1
        }     
        
        var Xt = Array(Array(X(0)))
        for(i <- 1 to X.length-1){
            Xt ++= Array(Array(X(i)))
        }
        
        var duration = (System.nanoTime - time_ini) / 1e9d

        println("EXECUTION TIME:" + duration + " nanoseconds")

     }

    def distance(xs: Array[Double], ys: Array[Double]) = {
        sqrt((xs zip ys).map { case (x,y) => pow(y - x, 2) }.sum)
    }   

}
