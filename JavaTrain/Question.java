import javax.lang.model.util.ElementScanner6;

public class Question{
    public static void main(String []args){
        Integer x1 = new Integer(16);
        Integer x2 = new Integer(32);
        x2 = x1 * 2;

        if(x1 == x2)
            System.out.println("x1 igual x2");
        else
        System.out.println("x1 no igual x2");
    }
}