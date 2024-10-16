import java.util.Scanner;

public class Main {

    public static void main(String[] args) {
        // for (int i = 0; i < 69; i++) {
        //     try {
        //         Word www = new Word(0);
        //         System.out.printf(
        //             "%s %s\n",
        //             www.finalWord(),
        //             www.sanitizeWord()
        //         );
        //     } catch (Exception e) {}
        // }

        Scanner s = new Scanner(System.in);

        System.out.println("Bienvenue dans ce jeu du pendu !"); // un peu glauque dit comme ça
        System.out.println("Choisissez votre mode de jeu :");
        System.out.println("0 - Français (fr)");
        System.out.println("1 - English (en)");
        System.out.println("2 - Latinus (la)");
        System.out.println("3 - Mot personalisé");

        System.out.print("> ");
        int choice = s.nextInt();

        Word w = new Word(0);
        switch (choice) {
            case 0:
                w = new Word(0);
                break;
            case 1:
                w = new Word(1);
                break;
            case 2:
                w = new Word(2);
                break;
            case 3:
                boolean loop = true;
                while (loop) {
                    System.out.print("Entrer le mot à deviner : ");
                    w = new Word(s.next()); // TODO: lit que un mot
                    if (w.isWordFound()) System.out.println(
                        "Le mot n'est pas valide."
                    );
                    else loop = false;
                }
                break;
            default: // French by default, can be changed ("try again" ?)
        }

        Character c;

        do {
            System.out.println(w.wordHidden());
            System.out.printf("Essais restants : %d\n", w.remainingAttempt);
            System.out.print("Lettres essayées : ");
            System.out.println(w.triedLetters);
            System.out.print("> ");

            c = s.next().trim().charAt(0);

            w.tryLetter(c);
        } while (w.remainingAttempt > 0 && !w.isWordFound());

        if (w.isWordFound()) {
            System.out.println("Bravo, vous avez gagné.e !");
        } else {
            System.out.println("Dommage...");
        }

        System.out.print("Le mot était : ");
        System.out.println(w.finalWord());

        s.close();
        /**/
    }
}
