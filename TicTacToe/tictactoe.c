#include <stdio.h>
#include <stdint.h>

void clearstdin() 
{
	int c;
	while ((c = getchar()) != '\n' && c != EOF);
}

void printgrid(uint16_t p1, uint16_t p2)
{
	for (uint8_t i = 0; i < 3; i++)
	{
		for (uint8_t j = 0; j < 3; j++)
		{
			if (p1 & 1 << (j + i * 3))
				printf("O ");
			else if (p2 & 1 << (j + i * 3))
				printf("X ");
			else
				printf("%d ", j + i*3);
		}
		printf("\n");
	}
}

int main() 
{
	uint16_t playerone = 0b000000000;
	uint16_t playertwo = 0b000000000;

	uint8_t counter = 9;
	int8_t whowon = 0;

	char play;
	uint16_t playbitmask;

	while (counter > 0)
	{
		printgrid(playerone, playertwo);

		for(;;)
		{
		    printf("%c> ", "XO"[counter & 1]);
			scanf(" %c", &play);
			clearstdin();

			play -= 0x30; // convert to int (only for x∈[0,9])
			playbitmask = 1 << play; // convert to bit mask

			if (play > 8 || play < 0)
				continue; // move out of bound

			if ((playbitmask & (playerone | playertwo)) > 0)
				continue; // move already played

			break; // all good
		};
		
		printf("\n");
		
		if (counter & 1) // get which players turn is it
			playerone |= playbitmask;
		else
			playertwo |= playbitmask;

		/*
		MASKS
		-------111000000 = 0x1c0
		-------000111000 = 0x038
		-------000000111 = 0x007
		-------100100100 = 0x124
		-------010010010 = 0x092
		-------001001001 = 0x049
		-------100010001 = 0x111
		-------001010100 = 0x054
		*/
		// check if someone won (alor la)
		if ((playerone & 0x1c0) == 0x1c0) whowon = 1;
		if ((playerone & 0x038) == 0x038) whowon = 1;
		if ((playerone & 0x007) == 0x007) whowon = 1;
		if ((playerone & 0x124) == 0x124) whowon = 1;
		if ((playerone & 0x092) == 0x092) whowon = 1;
		if ((playerone & 0x049) == 0x049) whowon = 1;
		if ((playerone & 0x111) == 0x111) whowon = 1;
		if ((playerone & 0x054) == 0x054) whowon = 1;

		if ((playertwo & 0x1c0) == 0x1c0) whowon = 2;
		if ((playertwo & 0x038) == 0x038) whowon = 2;
		if ((playertwo & 0x007) == 0x007) whowon = 2;
		if ((playertwo & 0x124) == 0x124) whowon = 2;
		if ((playertwo & 0x092) == 0x092) whowon = 2;
		if ((playertwo & 0x049) == 0x049) whowon = 2;
		if ((playertwo & 0x111) == 0x111) whowon = 2;
		if ((playertwo & 0x054) == 0x054) whowon = 2;

		if (whowon != 0)
			break;
		
		counter--;
	}

	printgrid(playerone, playertwo);
	
	printf("\n");

	if (whowon == 0)
		printf("Égalité !\n");
	else if (whowon == 1)
		printf("Joueur 1 a gagné !\n");
	else if (whowon == 2)
		printf("Joueur 2 a gagné !\n");

	return 0;
}
