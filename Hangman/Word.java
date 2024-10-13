public class Word {
	private String word;
	private boolean[] foundLetters; // list of letters maybe

	public int remainingAttempt;

	public Word() {
		this.word = "Cephalaspis";
		this.word = this.word.toUpperCase();
		
		this.foundLetters = new boolean[this.word.length()];
		this.remainingAttempt = 5;
	}

	public Word(int lang) {
		this.word = switch (lang) {
			case 0 -> "Racontée"; // French
			case 1 -> "Cephalaspis"; // English (yes)
			case 2 -> "Mercatores"; // Latin
			default -> "Word";
		};
		this.word = this.word.toUpperCase();
		
		this.foundLetters = new boolean[this.word.length()];
		this.remainingAttempt = 5;
	}
	
	public Word(String word) {
		this.word = word;
		this.word = this.word.toUpperCase();

		this.foundLetters = new boolean[this.word.length()];
		this.remainingAttempt = 5;
	}

	public void tryLetter(char letter) {
		letter = Character.toUpperCase(letter);

		int attemptPenalty = 1;

		for (int i = 0; i < this.word.length(); i++) {
			if (this.word.charAt(i) == letter) {
				if (!this.foundLetters[i]) // Still remove an attempt if the letter was already found 
					attemptPenalty = 0;
				this.foundLetters[i] = true;
			}
		}

		this.remainingAttempt -= attemptPenalty;
	}

	public String wordHidden() {
		StringBuilder sb = new StringBuilder(this.word.length());

		for (int i = 0; i < word.length(); i++) {
			if (this.foundLetters[i]) sb.append(word.charAt(i));
			else sb.append("_");
		}

		return sb.toString();
	}

	public boolean isWordFound() {
		for (boolean b : foundLetters) {
			if (b == false)
				return false;
		}
		return true;
	}

}
