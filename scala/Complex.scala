package demo.maths
import Math.sqrt

class Complex(m: Int, n: Int) {
    // a + bi
    val a = m
    val b = n
    def *(x: Complex): Complex = new Complex(x.a*x.a-x.b*x.b, x.a*x.b*2)
}

object Complex {
    def apply(m: Int, n: Int) = new Complex(m, n)
    def unapply(input: Complex) = Some((input.a, input.b))
}

object plusOne extends Function1[Int, Int] {
    def apply(m: Int) = m + 1
}

object double extends (Int => Int) {
    def apply(m: Int) = m*2
}

object square extends ((Int, Int) => Int) {
    def apply(x: Int, y: Int) = x*x + y*y
}

// object square extends (Double => Double) {
//     def apply(x: Double, y: Double) = x*x + y*y
// }

// object squareRoot extends (Double => Double) {
//     def apply(x: Double): Double = sqrt(x)
// }