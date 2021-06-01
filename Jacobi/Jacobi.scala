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

   def jacobi(A: Array[Array[Double]], B:Array[Double], l:Int, N:Int) {

      var ig  = Array.ofDim[Double](l)

      var tol = distance(dotProduct(A,ig, l), B)
      var tolerancia: Double = 0.00001

      var D = Array.fill(l)(0.0)
         for(i <- 0 to l-1){
            for(j <- 0 to l-1){
               if(i == j){
                  D(i) = A(i)(j)
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
      while((iteracion < N || tol > tolerancia)){
         
         ig = division(substraction(B, dotProduct(R, ig, l),l), D, l)
         tol = distance(dotProduct(A,ig, l), B)
         iteracion = iteracion + 1
      }

      var duration = (System.nanoTime - t1) / 1e9d
      println("EXECUTION TIME:" + duration + " nanoseconds")

   }

   def readCSVFile(file: String) : (Array[Array[Double]], Array[Double], Int) = {

      var C = io.Source.fromFile(file).getLines().map(_.split(",").map(_.trim.toDouble)).toArray
      var l = C.length
      var A = Array.ofDim[Double](l, l) 
      for(i <- 0 to (l-1)){
         for (j <- 0 to (l-1)){
            A(i)(j) = C(i)(j)
         }
      }

      var B = Array.fill(l)(0.0) 
      for(i <- 0 to l-1){
            B(i) = C(i)(C(1).length-1)
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

   def substraction(a:Array[Double] , b:Array[Double], l:Int): Array[Double] = {
      var r = Array.fill(l)(0.0)
      for(i <- 0 to l-1){
         r(i) = a(i) - b(i)
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