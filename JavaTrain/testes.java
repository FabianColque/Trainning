import java.lang.Runtime;
import java.util.Scanner;

abstract class Person{  
    abstract void eat();  
  }  
  class testes{  
   public static void main(String args[]) throws Exception{
    for(int i=0;i<args.length;i++){  
        System.out.println(args[i]);  
    }  

    Integer uno = 1;
    Integer dos = 1;

    if(uno == dos){
        System.out.println("equal");
    }else{
        System.out.println("not equal");
    }
    
    //Scanner sc = new Scanner(System.in);
    //Runtime.getRuntime().exec("sublime");  
    /*{Person p=new Person(){  
    void eat(){System.out.println("nice fruits");}  
    };  
    p.eat();*/  
   }  
  } 

/*final class Employee{  
    final String pancardNumber;  
      
    public Employee(String pancardNumber){  
        this.pancardNumber=pancardNumber;  
    }  

    public void setPancardNumber(String newpancard){
        this.pancardNumber = newpancard;
    }
      
    public String getPancardNumber(){  
        return pancardNumber;  
    }        
}  

public class testes {
    public static void main(String []args){
        Employee oi = new Employee("boli");
        System.out.println(oi.getPancardNumber());        
    }
}*/
