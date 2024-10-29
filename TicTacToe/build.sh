if [ "$1" == "-c" ] 
then
	echo "cleaning"
	rm *.o *.out
else 
	echo "compiling"
	nasm -felf64 source.asm
	ld -o a.out source.o
fi

