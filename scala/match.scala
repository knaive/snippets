package demo.matcher

case class Phone(feature: String, brand: String)

class MatchIt {
    def getType[T](o: T): String = o match {
        case i: Int => "Int"
        case i: Double => "Double"
        case i: String => "String"
        case i: Boolean => "Boolean"
        case _ => "Unknown"
    }

    def getManufacturer(o: Phone): String = o match {
        case Phone(_, "Iphone") => "Apple"
        case Phone(_, "Mate") => "Huawei"
        case Phone("explosion", _) => "Samsung"
        case Phone("bullet proof", _) => "Nokia"
        case Phone(_, _) => "Unknown"
    }
}