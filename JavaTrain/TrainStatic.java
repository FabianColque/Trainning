
class Student{
    protected String name;
    public int age;
    public static String college;
    public final String specie = "Human";
    //public final String planet;
    //planet = "Tierra";

    public Student(){

    }

    public Student(String name, int age, String col){
        final String planet;
        planet = "Tierra";

        this.name = name;
        this.age = age;
        college = col;
    }

    public String toString(){
        return name + " " + age + college;
    }

    public String toWelcome(){
        return "Welcome " + name;
    }
}

class Boy extends Student{

    public Boy(){

    }

    public Boy(String name, int age, String col){
        super(name, age, col);
    }

    @Override
    public String toWelcome(){
        return "Habla " + name;
    }

    public void setSpecie(String newSpecie){
        //specie = newSpecie;
    }
}

public class TrainStatic {
    static{System.out.println("static block is invoked");} 
    public static void main(String []args){
        Student student1 = new Student("Fabian", 13, "Bryce");
        Student student2 = new Student("Ernesto", 15, "Sagrada");

        System.out.println(student1.toString());
        System.out.println(student2.toString());
        //System.out.println(student1.name);

        Boy luffy = new Boy("Luffy", 20, "East Blue");
        System.out.println(luffy.toString());
        System.out.println(luffy.toWelcome());
        System.out.println(luffy.specie);

        luffy.setSpecie("dibujo");
    }
}
