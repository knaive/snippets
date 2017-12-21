package org.fuguang;

import java.lang.reflect.Proxy;

public class AnimalFactory {
    private static Object getAnimalBase(Object obj) {
        return Proxy.newProxyInstance(
            obj.getClass().getClassLoader(),
            obj.getClass().getInterfaces(), 
            new AopHandler(obj)
        );
    }

    @SuppressWarnings("unchecked")
    public static Object getAnimal(Object obj) {
        return (Object) getAnimalBase(obj);
    }

    @SuppressWarnings("unchecked")
    public static Object getAnimal(String className) {
        Object obj = null;
        try {
            obj = getAnimalBase(Class.forName(className).newInstance());
        } catch (Exception e) {
            e.printStackTrace();
        }
        return (Object) obj;
    }

    @SuppressWarnings("unchecked")
    public static Object getAnimal(Class cls) {
        Object obj = null;
        try {
            obj = getAnimalBase(cls.newInstance());
        } catch (Exception e) {
            e.printStackTrace();
        }
        return (Object) obj;
    }
}