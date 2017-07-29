import demo.maths._
import demo.matcher._

object Runner {
    def main(args: Array[String]) = {
        //deco
        var Complex(a, b) = Complex(2, 1) * Complex(2, 1)
        println("%d + %di".format(a, b))

        val increment = plusOne
        val n = 2
        println("%d + 1 = %d".format(n, increment(n)))
        val c = square(a, b)
        // val root = squareRoot(c)
        println("square(%d, %d) = %d".format(a, b, c))
        // println("squareRoot of %d and %d is %d".format(a, b, root))
        val matcher = new MatchIt
        val getType = matcher.getType _
        val getProducer = matcher.getManufacturer _

        println("Typeof('%d'): %s".format(2, getType(2)))
        println("Typeof('%f'): %s".format(2.1, getType(2.1)))
        println("Typeof('%s'): %s".format("string", getType("string")))

        val phone = Phone("ugly", "Mate")
        val phone1 = Phone("explosion", "unkown")
        val phone2 = Phone("none", "none")

        println("The manufacturer of the phone %s is %s".format(phone.toString(), getProducer(phone)))
        println("The manufacturer of the phone %s is %s".format(phone1, getProducer(phone1)))
        println("The manufacturer of the phone %s is %s".format(phone2, getProducer(phone2)))
    }
}