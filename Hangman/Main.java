import java.util.Scanner;

public class Main {
	public static void main(String[] args) {
		Word w = new Word(1);

		Scanner s = new Scanner(System.in);

		Character c;

		do {
			System.out.println(w.wordHidden());
			System.out.printf("Essais restants : %d\n", w.remainingAttempt);
			System.out.print("> ");
			
			c = s.next().trim().charAt(0);
			
			w.tryLetter(c);
			
		} while (w.remainingAttempt > 0 && !w.isWordFound());
		
		System.out.print("Le mot était : ");
		System.out.println(w.wordHidden());
		
		s.close();
	}
}
