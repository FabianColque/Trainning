/*
 * Factory Method
 * 
 * It is a creational design pattern that provides an interface for creating objects
 * in a superclass, but allows subclasses to alter the type of objects that will be
 * created
 * 
 * More info: https://refactoring.guru/design-patterns/factory-method
 * 
*/

interface Dialog{
    public int createButtom();
}

class WindowDialog implements Dialog{
    public int createButtom(){
        return 1;
    }
}

class WebDialog implements Dialog{
    public int createButtom(){
        return 2;
    }
}


//Main
class FactoryMethod{
    public static void main(String []args){

    }
}