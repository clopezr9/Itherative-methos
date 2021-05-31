import scala.io.Source
import javax.print.DocFlavor.INPUT_STREAM

object GaussSimple {
    
    def main(args: Array[String]): Unit = {
        var augmented_matrix = Array(Array(-1.2, 0.0, 0.0, 4.0, -1.0, 0.0, 61.0),
                                Array(0.0, -1.0, 0.0, -1.0, 4.0, -1.0, 14.0),
                                Array(4.0, -1.0, 1.0, -1.0, 0.0, 0.0, 8.0),
                                Array(-1.0, 5.0, -1.0, 0.0, -1.0, 0.0, 5.0),
                                Array(0.0, -1.0, 5.0, 0.0, 0.0, -1.0, 9.0),
                                Array(0.0, 0.0, -1.0, 0.0, -1.0, 6.0, 23.0))
        var n = 6
        gaussSimple(augmented_matrix, n)
    }

    def gaussSimple(aug_matrix: Array[Array[Double]], n : Int): Unit = {
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

        var duration = (System.nanoTime - t1) / 1e9d

        println("Execution Time:" + duration + "seconds")

        for ( i <- 0 to n-1){
            println("x" + i + "= " + x(i))
        }
    }
}
