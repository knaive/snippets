package demo.util

object ObjectDemo {
    var count = 0

    def currentCount(): Unit = {
        count += 1
        count
    }

    def printCurrentCount(): Unit = {
        println("Current count: " + ObjectDemo.currentCount()) 
    }
}

object StringUtil {
    def joiner(strings: List[String], separator: String = ", "): String = strings.mkString(separator)
}

object Maths {
    def factorial(i: Int): Int = {
        if (i==1) i 
        else i * factorial(i-1)
    }
}