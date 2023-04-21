
public class Heranca {

    public static void main(String []args){
        C2 v=new C3(2);
        v.opC2(2);
    }
}

class C1{
    int x=1,y=1;

    public C1(){
        System.out.println("C1");
        x++;
        y--;
    }

    public C1(int f){
        System.out.println("C1 f");
        x*=f;
        y+=f;
    }
}

class C2 extends C1{
    public void opC2(int f){
        System.out.println(x+y);
    }

    public C2(int f){
        System.out.println("C2 f");
        y*=f;
        x+=f;
    }

    public C2(){
        System.out.println("C2");
        y*=2;
        x+=3;
    }
}

class C3 extends C2{
    public C3(){        
        super(5);
        System.out.println("C3");
    }

    public C3(int f){
        super(6);
        System.out.println("C3 f");
    }
}
