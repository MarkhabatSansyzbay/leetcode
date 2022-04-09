package cinema;
import java.util.Scanner;

public class Cinema {

    public static void showTheSeats(String[][] room,int rows, int seats) {
        System.out.println("Cinema:");
        System.out.println();
        for (int i = 0; i < rows + 1; i++) {
            for (int j = 0; j < seats + 1; j++) {
                System.out.print(room[i][j] + " ");
            }
            System.out.println();
        }
    }
    public static void fillTheArray(String[][] room,int rows, int seats){
        for (int i = 0; i < rows + 1; i++) {
            for (int j = 0; j < seats + 1; j++) {
                if (i == 0 && j == 0) {
                    room[i][j] = " ";
                } else if (i == 0) {
                    room[i][j] = String.valueOf(j);
                } else if (j == 0) {
                    room[i][j] = String.valueOf(i);
                } else {
                    room[i][j] = "S";
                }
            }
        }
    }
    public static void buyATicket(String[][] room, int rows, int seats) {
        Scanner sc = new Scanner(System.in);
        int price;
        System.out.println("Enter a row number: ");
        int row = sc.nextInt();
        System.out.println("Enter a seat number in that row: ");
        int seat = sc.nextInt();
        if(rows * seats <= 60){
            price = 10;
        } else {
            if (row <= rows/2) {
                price = 10;
            } else {
                price = 8;
            }
        }
        System.out.println("\nTicket price: $" + price);
        room[row][seat] = "B";
    }
    public static int income (String[][] room, int rows, int seats) {

    }
    public static void menu (String[][] room, int rows, int seats){
        Scanner sc = new Scanner(System.in);
        int x;
        do {
            System.out.println("\n1. Show the seats\n" +
                    "2. Buy a ticket\n" +
                    "0. Exit");
            int n = sc.nextInt();
            x = n;
            switch (n) {
                case 1:
                    showTheSeats(room, rows, seats);
                    break;
                case 2:
                    buyATicket(room, rows, seats);
                    break;
                case 0:
                    break;
            }
        } while (x > 0);
    }
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        System.out.println("Enter the number of rows: ");
        int rows = sc.nextInt();
        System.out.println("Enter the number of seats in each row: ");
        int seats = sc.nextInt();
        String[][] room = new String[rows + 1][seats + 1];
        fillTheArray(room, rows, seats);
        menu(room, rows, seats);
    }
}