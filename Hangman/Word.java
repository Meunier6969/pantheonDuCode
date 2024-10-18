import java.io.RandomAccessFile;
import java.util.ArrayList;

public class Word {

    private String word;
    private boolean[] foundLetters;

    public ArrayList<Character> triedLetters;

    public Word() {
        this.word = "Cephalaspis";

        this.triedLetters = new ArrayList<>();
        this.foundLetters = new boolean[this.word.length()];
        this.word = this.sanitizeWord();
    }

    public Word(int lang) {
        // raaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
        try {
            this.word = switch (lang) {
                case 0 -> Word.getRandomWordFromFile("Mots/mots_fr.txt"); // French
                case 1 -> "Cephalaspis"; // English (yes)
                case 2 -> "Mercatores"; // Latin
                default -> "Word";
            };
        } catch (Exception e) {
            System.err.println("Couldn't get word : " + e);
        }

        this.triedLetters = new ArrayList<>();
        this.foundLetters = new boolean[this.word.length()];
        this.word = this.sanitizeWord();
    }

    public Word(String word) {
        this.word = word;

        this.triedLetters = new ArrayList<>();
        this.foundLetters = new boolean[this.word.length()];
        this.word = this.sanitizeWord();
    }

	// return false if the letter was already tried, or isn't in the word
    public boolean tryLetter(char letter) {
        letter = Character.toUpperCase(letter);

        if (this.triedLetters.contains(letter)) {
            return false;
        }

        this.triedLetters.add(letter);

        boolean attemptPenalty = false;

        for (int i = 0; i < this.word.length(); i++) {
            if (this.word.charAt(i) == letter) {
                this.foundLetters[i] = true;
                attemptPenalty = true;
            }
        }

        return attemptPenalty;
    }

    public String wordHidden() {
        StringBuilder sb = new StringBuilder(this.word.length());

        for (int i = 0; i < word.length(); i++) {
            if (this.foundLetters[i]) sb.append(word.charAt(i));
            else sb.append("_");
        }

        return sb.toString();
    }

    public String finalWord() {
        return this.word;
    }

    public boolean isWordFound() {
        for (boolean b : foundLetters) {
            if (b == false) return false;
        }
        return true;
    }

    public String sanitizeWord() {
        StringBuilder newString = new StringBuilder(this.word.length());

        for (int i = 0; i < this.word.length(); i++) {
            char c = this.word.charAt(i);
            if (Character.isLetter(c)) {
                newString.append(c);
            } else {
                this.foundLetters[i] = true;
                if (c != '_') newString.append(c);
                else newString.append(' ');
            }
        }

        return newString.toString().toUpperCase();
    }

    public static String getRandomWordFromFile(String filePath)
        throws Exception {
        final RandomAccessFile f = new RandomAccessFile(filePath, "r");
        final long randomLocation = (long) (Math.random() * f.length());

        f.seek(randomLocation);
        f.readLine();

        String randomLine = f.readLine();

        f.close();
        return randomLine;
    }
}
