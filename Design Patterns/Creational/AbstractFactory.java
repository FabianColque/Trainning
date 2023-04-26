package Design Patterns.Creational;

/*
 * Abstract Factory
 * 
 * It is a creational design pattern that lets you produce families or related
 * objects without specifying their concret classes
 * 
 * More info: https://refactoring.guru/design-patterns/abstract-factory
 * 
*/

interface GUIFactory{
    public Button createButtom();
    public CheckBox createCheckbox();
}

interface Button{
    public void Paint();
}

interface CheckBox{
    public void Paint();
}

class WinButton implements Button{
    public void Paint(){

    }
}

class MacButton implements Button{
    public void Paint(){
        
    }
}

class WinCheckBox implements CheckBox{
    public void Paint(){

    }
}

class MacCheckBox implements CheckBox{
    public void Paint(){
        
    }
}

class WinFactory implements GUIFactory{
    public Button createButtom(){
        return new WinButton();
    }

    public CheckBox createCheckbox(){
        return new WinCheckBox();
    }
}

class MacFactory implements GUIFactory{
    public Button createButtom(){
        return new MacButton();
    }

    public CheckBox createCheckbox(){
        return new MacCheckBox();
    }
}

public class AbstractFactory {
    public static void main(String []args){

    }
}


