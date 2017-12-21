package org.fuguang;

import org.fuguang.pub.Animal;

public class Dog implements Animal {
    private String name = "sheepdog";

    public Dog(){}

    @Override
    public void setName(String name) {
        this.name = name;
    }

    @Override
    public String getName() {
        return this.name;
    }

    @Override
    public void run() {
        System.out.println("Doggy run...");
    }

    @Override
    public void getProperty() {
        System.out.println("Dogs can swim");
    }
}