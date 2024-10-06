#include <stdio.h>
#include <stdint.h>
#include <time.h>	// Random number
#include <stdlib.h>

// Get a random number between lb and hb (inclusive)
uint16_t getRandomNumber(uint16_t lb, uint16_t hb)
{
	return lb + (rand() % (hb - lb + 1));
}

int main() 
{
	srand(time(NULL)); // Initialization

	uint16_t lowerBound = 0;
	uint16_t higherBound = 100;

	uint8_t try = 0;

	uint16_t guess;
	uint16_t randomNumber = getRandomNumber(lowerBound, higherBound);

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
	{
		printf("Tu as trouvé.e le nombre du premier coup!\n");
	} 
	else
	{
		printf("Tu as trouvé.e le nombre en %d essais!\n", try);
	}

	return 0;
}