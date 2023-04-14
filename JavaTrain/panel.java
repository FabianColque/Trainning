import java.awt.*;
import java.awt.event.*;
import javax.swing.*;

public class App extends JFrame{
    public App() {
        super("fundatec");

        JPanel panel = new JPanel();
        panel.setLayout(new FlowLayout());

        JButton btn = new JButton("Ok");
        btn.addMouseListener(new MouseListener() {
            public void mouseEntered(MouseEvent e){
                System.out.print("A");
            }

            public void mouseExited(MouseEvent e){
                System.out.print("B");
            }

            public void mouseReleased(MouseEvent e){
                System.out.print("C");
            }

            public void mousePressed(MouseEvent e){
                System.out.print("D");
            }

            public void mouseClicked(MouseEvent e){
                System.out.print("E");
            }
        });

        panel.add(btn);
        add(panel);
    };

    public static void main(String[] args){
        App app = new App();
        app.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        app.setSize(200, 75);
        app.setVisible(true);
    }
}