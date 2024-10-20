import java.util.Scanner;

public class Game {
	// Game state
	private Word word;
	public int remainingAttempts;

	public Game() {}

	public void run() {
		Scanner s = new Scanner(System.in);

		Game.printWordSelect();

		String strChoice;
		int choice;
		do { // Make sure to loop and check if the numbers are correct 
			System.out.print("> ");
			strChoice = s.nextLine();
		
			try {
				choice = Integer.parseInt(strChoice);
			} catch (Exception e) {
				choice = -1;
			}
		
		} while ( choice < 0 || choice > 3 );

		this.word = getWordLanguage(choice, s);
		this.remainingAttempts = 5;
		
		do {
			// "What char?"
			this.printMainGameLoop();
			
			// Get char
			String strc;
			Character c;
			do {
				System.out.print("> ");
				strc = s.nextLine().trim();
			} while (strc.length() == 0);

			c = strc.charAt(0);

			// Try char
			boolean success = this.word.tryLetter(c);

			// Update remainingAttempts
			if (success == false) remainingAttempts--;
		} while (this.remainingAttempts > 0 && !this.word.isWordFound());

		if (this.word.isWordFound()) {
			System.out.println("Bravo, vous avez gagné.e !");
		} else {
			System.out.println("Dommage...");
		}

		System.out.print("Le mot était : ");
		System.out.println(this.word.finalWord());

		s.close();
	}

	public static void printWordSelect() {
		System.out.println("Bienvenue dans ce jeu du pendu !"); // un peu glauque dit comme ça
		System.out.println("Choisissez votre mode de jeu :");
		System.out.println("0 - Français (fr)");
		System.out.println("1 - English (en)");
		System.out.println("2 - Latinus (la)");
		System.out.println("3 - Mot personalisé");
	}

	public void printMainGameLoop() {
		System.out.println(this.word.wordHidden());
		System.out.printf("Essais restants : %d\n", this.remainingAttempts);
		System.out.print("Lettres essayées : ");
		System.out.println(this.word.triedLetters);
	}

	public Word getWordLanguage(int choice, Scanner s) {
		Word w = new Word(0);
		switch (choice) {
			case 0: // French
				w = new Word(0);
				break;
			case 1: // English
				w = new Word(1);
				break;
			case 2: // Latin
				w = new Word(2);
				break;
			case 3: // Custom word
				do {
					System.out.print("Entrer le mot à deviner : ");
					w = new Word(s.nextLine());
					if (w.isWordFound())
						System.out.println("Le mot n'est pas valide.");
				} while (w.isWordFound());
				break;
			default: // French by default, can be changed ("try again" ?)
		}

		return w;
	}

}
