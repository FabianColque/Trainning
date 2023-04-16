import java.io.*;
import java.util.*;

public class Priority {

    public static class Student{
        int id; 
        String name;
        double cgpa;
        
        Student before;
        Student next;
        
        public Student(){
            
        }
        
        public Student(int id, String name, double cgpa){
            this.id   = id; 
            this.name = name;
            this.cgpa = cgpa;
            
            next = null;
            before = null;    
        }    
        
        int getID(){
            return id;
        }
        
        String getName(){
            return name;
        }
        
        double getCGPA(){
            return cgpa;
        }
        
        boolean hasPriority(Student student){
            if(cgpa < student.getCGPA()) return false;
            if(cgpa > student.getCGPA()) return true;            
            if(name.compareTo(student.getName()) > 0) return false;                
            if(name.compareTo(student.getName()) < 0) return true;
            if(id > student.getID()) return true;                
                    
            return false;
        }
        
        public String toString(){
            return name + " " + id + " " + cgpa;
        }
    };

    public static class PriorQueue{
        Student root;
        
        public PriorQueue(){
           root = null;  
        }
        
        void enter(Student student){
            if(root == null){
                root = student;
                return;
            }
            
            Student temp = root;
            
            while(temp != null){
                if(temp.hasPriority(student)){                    
                    if(temp.next == null){
                        temp.next = student;
                        student.before = temp;
                        break;
                    }                    
                    temp = temp.next;
                }else{
                    if(temp.before == null){
                        student.next = temp;
                        temp.before  = student;
                        root         = student;
                        
                        //System.out.println(student.toString() + " -->" + temp.toString());
                    }else{ 
                        //System.out.println(student.toString() + " -->" + temp.toString());                   
                        student.before = temp.before;
                        temp.before.next = student;
                        student.next = temp;
                        temp.before  = student;                    
                    }
                    break;
                }
            }
            
            //print();
            //System.out.println("----------");
        }
        
        Student served(){
            Student result = root;
            
            if(root != null){                
                root = root.next;
                
                if(root != null)
                    root.before = null;
            }
            
            //System.out.println("**************");
            //System.out.println(result.toString());            
            //print();
            //System.out.println("**************");
            return result; 
        }
        
        List<Student> getStudents(List<String> events){
            for(String event : events){
                String [] instruction = event.split(" ");
                //System.out.println(instruction.length);
                if(instruction[0].equals("ENTER")){
                    Student newStudent = new Student(Integer.parseInt(instruction[3]), instruction[1], Double.parseDouble(instruction[2]));
                    enter(newStudent);    
                }else{                    
                    served();
                }
            }
                        
            return toList();
        }
        
        public void print(){
            Student temp = root;
            while(temp != null){
                //System.out.println(temp.toString());
                temp = temp.next;
            }
        }
        
        List<Student> toList(){
            List<Student> result = new ArrayList<>();
            Student temp = root;
            while(temp != null){
                //System.out.println(temp.toString());
                result.add(temp);
                temp = temp.next;
            }
            
            return result;
        }         
    };

    public static void main(String[] args) {
        /* Enter your code here. Read input from STDIN. Print output to STDOUT. Your class should be named Solution. */
        
        Scanner sc = new Scanner(System.in);
        int sz = sc.nextInt();
        List<String> lines = new ArrayList<>();
        sc.nextLine();
        while(sz-- > 0){
            String aux = sc.nextLine();
            lines.add(aux);
            //System.out.println(aux);
        }
        
        PriorQueue holi = new PriorQueue();
        List<Student> result = holi.getStudents(lines);        
        if(result.size() == 0)
            System.out.println("EMPTY");        
        else
            for(Student rata : result)
                System.out.println(rata.getName());
    }
}