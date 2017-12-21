package org.fuguang;

import org.fuguang.pub.Animal;
import org.fuguang.*;

/**
 * Hello world!
 *
 */
public class App 
{
    public static void main( String[] args )
    {
        Animal dog = (Animal)AnimalFactory.getAnimal(Dog.class);

        dog.run();

        System.out.println("My name is " + dog.getName());

        dog.setName("LittleBlack");

        System.out.println("My name is " + dog.getName());
    }
}
