import scala.io.Source
import javax.print.DocFlavor.INPUT_STREAM

object GaussSimple {
    
    def main(args: Array[String]): Unit = {
        for(arg<-args){
            println(arg + ":")
            var param = readCSVFile(arg)
            var (aug_matrix, n) = param
            gaussSimple(aug_matrix, n)
        }
    }

    def readCSVFile(file: String): (Array[Array[Double]], Int) = {
        var C = io.Source.fromFile(file).getLines().map(_.split(",").map(_.trim.toDouble)).toArray
        var n = C.length
        var m = C(0).length
        var aug_matrix = Array.ofDim[Double](n, m)

        for(i <- 0 to n-1){
           for (j <- 0 to n){
              aug_matrix(i)(j) = C(i)(j)
           }
        }
        return (aug_matrix , n)
    }

    def gaussSimple(aug_matrix: Array[Array[Double]], n: Int): Unit = {
        var x = Array.fill(n)(0.0)
        var t1 = System.nanoTime

        for (i <- 0 to n-1){
            if (aug_matrix(i)(i) == 0.0){
                return
            }

            for (j <- i+1 to n-1){
                var ratio = aug_matrix(j)(i)/aug_matrix(i)(i)

                for (k <- 0 to n){
                    aug_matrix(j)(k) = aug_matrix(j)(k) - ratio * aug_matrix(i)(k)
                }
            }
        }

        x(n-1) = aug_matrix(n-1)(n) / aug_matrix(n-1)(n-1)

        for (i <- n-2 to 0 by -1){
            x(i) = aug_matrix(i)(n)

            for(j <- i+1 to n-1){
                x(i) = x(i) - aug_matrix(i)(j) * x(j)
            }

            x(i) = x(i)/aug_matrix(i)(i)
        }

        var duration = (System.nanoTime - t1)
        println("EXECUTION TIME:" + duration + " nanoseconds")
    }
}
