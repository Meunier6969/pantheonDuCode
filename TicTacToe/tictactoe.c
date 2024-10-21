#include <stdio.h>
#include <stdint.h>

void clearstdin() {
	int c;
	while ((c = getchar()) != '\n' && c != EOF);
}

int main() {
	uint16_t playerone = 0x0;
	uint16_t playertwo = 0x0;
	char play;

	uint8_t counter = 9; 
	while (counter > 0)
	{
		printf("%d> ", counter & 1);

		scanf(" %c", &play); // play-48
		printf("%c, %d\n", play, play-48);

		playerone |= 1 << (play-48);

		printf("P1 : %9b\n", playerone);
		printf("P2 : %9b\n", playertwo);

		counter--;
		clearstdin();
	}
	


	printf("\n");

	return 0;
}
/*
-------0 00010000  P1
-------0 00000010  P2
AND
-------0 00010010  Total

MASKS
-------1 11000000 = 0x1c0
-------0 00111000 = 0x038
-------0 00000111 = 0x007
-------1 00100100 = 0x124
-------0 10010010 = 0x092
-------0 01001001 = 0x049
-------1 00010001 = 0x111
-------0 01010100 = 0x054
*/