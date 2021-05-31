import scala.io.Source
import javax.print.DocFlavor.INPUT_STREAM
object Jacobi {
	def main(args: Array[String]): Unit = {
      
        var ab = readCSV()
        var (a,b,l) = ab
        var N = 500
        var tolerancia: Double = 0.00001
        jacobi(a,b,l,N)

    }

   def jacobi(A: Array[Array[Double]], B:Array[Array[Double]], l:Int, N:Int) {

      var ig = Array.fill(l)(0.0)

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
               //println("Position of R: ",i,j)
               //println("R: " , R(i)(j))
            }
         }
      }

      var iteracion = 0
      while(iteracion < N){
         ig = division(substraction(B, dotProduct(R, ig, l),l), D, l)
         iteracion = iteracion + 1
      }
      for(i<-0 to l-1){
         println(ig(i))
      }
   }

   def readCSV() : (Array[Array[Double]], Array[Array[Double]], Int) = {
      var C = io.Source.fromFile("test.csv").getLines().map(_.split(",").map(_.trim.toDouble)).toArray
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
}