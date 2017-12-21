package org.fuguang;

import java.lang.reflect.InvocationHandler;
import java.lang.reflect.Method;

public class AopHandler implements InvocationHandler {
    private Object o;

    public AopHandler(Object o) {
        this.o = o;
    }

    @Override
    public Object invoke(Object proxy, Method method, Object[] args) throws Throwable {
        Object ret = null;

        System.err.println("method: " + method.getName() + "\n type of parameters");

        for(Class type : method.getParameterTypes())
            System.err.println(type.getName());

        System.err.println("Return data type:" + method.getReturnType().getName());

        ret = method.invoke(o, args);

        System.err.println("End\n");
        return ret;
    }
}
