import scala.io.Source
import javax.print.DocFlavor.INPUT_STREAM
import javax.print.attribute.standard.PrinterURI
import math._
import java.io.File
object Jacobi {
	def main(args: Array[String]): Unit = {
   
        var N = 100
        for(arg<-args){
            var ab = readCSVFile(arg)
            var (a, b, l) = ab
            jacobi(a,b,l, N)
        }

    }

   def jacobi(A: Array[Array[Double]], B:Array[Array[Double]], l:Int, N:Int) {

      var tol = Array.ofDim[Double](l,l)
      var BtoArray = Array.fill(l){0.0}
      //for (i <- 0 to l-1){
      //   for (j <- 0 to l-1){
      //      println(tol(i)(j))
      //   }
      //}
      var ig  = Array.ofDim[Double](l)
      //var matMul = A.zip(tol) map (_.zipped map (_ * _)) map (_.sum) 

      //var tolerancia: Double = 0.00001
      //var aiuda = distance(matMul,BtoArray)

      var D = Array.fill(l)(0.0)
         for(i <- 0 to l-1){
            for(j <- 0 to l-1){
               if(i == j){
                  D(i) = A(i)(j)
                  //println("Position of D: ", i)
                  //println("D: " , D(i))
               }
            }
         }

      var R = Array.ofDim[Double](l, l) 
      for(i <- 0 to l-1){
         for(j <- 0 to l-1){
            if(i != j){
               R(i)(j) = A(i)(j)
            }
         }
      }

      var iteracion = 0
      var t1 = System.nanoTime
      while((iteracion < N )){ //|| aiuda > tolerancia)){
         
         ig = division(substraction(B, dotProduct(R, ig, l),l), D, l)
         //for (i<-0 to l-1){
         //   tol.update(i,ig)
         //}

         iteracion = iteracion + 1
         //var matMul = A.zip(tol) map (_.zipped map (_ * _)) map (_.sum)
         //aiuda = distance(matMul,BtoArray)

      }

      var duration = (System.nanoTime - t1) / 1e9d
      println("EXECUTION TIME:" + duration + " nanoseconds")
   }

   def readCSVFile(file: String) : (Array[Array[Double]], Array[Array[Double]], Int) = {

      var C = io.Source.fromFile(file).getLines().map(_.split(",").map(_.trim.toDouble)).toArray
      var l = C.length
      var A = Array.ofDim[Double](l, l) 
      for(i <- 0 to (l-1)){
         for (j <- 0 to (l-1)){
            A(i)(j) = C(i)(j)
         }
      }

      var B = Array.ofDim[Double](l, l) 
      for(i <- 0 to l-1){
            B(i)(0) = C(i)(3)
            //println("Position: ",i,0)
            //println("B: ",B(i)(0))
         }
      return (A,B,l)
   }

   def dotProduct(R: Array[Array[Double]], b:Array[Double], l:Int): Array[Double] = {
      var r = Array.fill(l)(0.0)
      for(i <- 0 to l-1){
         var aux = 0.0
         for(j <- 0 to l-1){
            aux = aux + (b(j)*R(i)(j))
         }
         r(i) = aux
      }
      return r
   }

   def substraction(a:Array[Array[Double]] , b:Array[Double], l:Int): Array[Double] = {
      var r = Array.fill(l)(0.0)
      for(i <- 0 to l-1){
         r(i) = a(i)(0) - b(i)
      }
      return r
   }
   
   def division(a:Array[Double] , b:Array[Double], l:Int): Array[Double] = {
      var r = Array.fill(l)(0.0)
      for(i <- 0 to l-1){
         r(i) = a(i) / b(i)
      }
      return r
   }

   def distance(xs: Array[Double], ys: Array[Double]) = {
      sqrt((xs zip ys).map { case (x,y) => pow(y - x, 2) }.sum)
   }
}