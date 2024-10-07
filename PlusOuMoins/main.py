#################################################
# https://music.youtube.com/watch?v=WJIWINuwBgM #
#################################################

from random import randint

def chooseBounds() -> tuple:
	choice = 0

	print("Quelle difficulté voulez-vous jouer ?")
	print("1 Facile (0-100)")
	print("2 Moyen (0-1000)")
	print("3 Difficile (0-10000)")
	print("4 Au choix")

	while choice not in [1, 2, 3, 4]:
		choice = int(input("> "))

	match choice:
		case 1:
			return (0, 100)
		case 2:
			return (0, 1000)
		case 3:
			return (0, 10000)
		case 4:
			lb = int(input("Entrer la borne minimale > "))
			hb = int(input("Entrer la borne maximale > "))
			return (lb, hb)

	return (0, 100)

def game(lb: int, hb: int) -> int:
	attempt: int = 0
	guess: int = -1

	randomNumber: int = randint(lb, hb + 1)

	# Main loop
	while guess != randomNumber:
		attempt += 1

		guess = int(input(f"Essai #{attempt} > "))

		if guess > randomNumber:
			print("Le nombre est plus petit")
		elif guess < randomNumber:
			print("Le nombre est plus grand")

	return attempt

def main():
	# Hi
	print("Bienvenue dans ce jeu de Plus ou Moins!")

	# Game
	running: bool = True
	while running:
		lowerBound, higherBound = chooseBounds()

		attempts = game(lowerBound, higherBound)
		# Win
		print(f"Bravo! Tu as gagné.e en {attempts} essai(s)!")

		print("Veux tu continuer ? (y/N)")
		if input("> ").lower() != 'y': running = False

	print("Au revoir !")

if __name__=="__main__":
	main()