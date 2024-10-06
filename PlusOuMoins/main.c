// TODO: Error checking

#include <stdio.h>
#include <time.h>	// Random number
#include <stdlib.h>

// Get a random number between lb and hb (inclusive)
int getRandomNumber(int lb, int hb)
{
	return lb + (rand() % (hb - lb + 1));
}

int main() 
{
	// Initialization
	srand(time(NULL)); 

	int choice;

	int lowerBound, higherBound;

	int try;

	int guess;
	int randomNumber;
	
	BEGINNING:

	// hi :D
	printf("Bienvenue de ce jeu de Plus ou Moins !\n");

	printf("Enter la borne minimale : ");
	scanf("%d", &lowerBound);
	printf("Enter la borne maximale : ");
	scanf("%d", &higherBound);
	
	try = 0;
	randomNumber = getRandomNumber(lowerBound, higherBound);

	// main loop
	do
	{
		try++;
		printf("Essai #%d > ", try);
		scanf("%d", &guess);
		
		if (guess > randomNumber) 
			printf("Le nombre est plus petit\n");
		if (guess < randomNumber)
			printf("Le nombre est plus grand\n");

	} while (guess != randomNumber);

	// winner
	if (try == 1)
		printf("Tu as trouvé.e le nombre du premier coup!\n");
	else
		printf("Tu as trouvé.e le nombre en %d essais!\n", try);

	// restart
	printf("Veux tu rejouer ? (y/N) ");
	scanf(" %c", &choice);
	if (choice == 'y')
		goto BEGINNING;

	return 0;
}