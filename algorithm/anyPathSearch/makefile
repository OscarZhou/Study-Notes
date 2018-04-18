Main.exe	: Main.o graphics.o puzzle.o algorithm.o customQueue.o
	g++ -O2 -std=c++11 -o Main.exe Main.o graphics.o puzzle.o algorithm.o customQueue.o -l gdi32 -static-libgcc -static-libstdc++
			
Main.o	: Main.cpp graphics.h puzzle.h algorithm.h customQueue.h 
	g++ -O2 -std=c++11 -c -Wno-write-strings Main.cpp

# queue.o : queue.cpp queue.h
# 	g++ -O2 -std=c++11 -c -Wno-write-strings  queue.cpp

# stack.o : stack.cpp stack.h
# 	g++ -O2 -std=c++11 -c -Wno-write-strings  stack.cpp

customQueue.o : customQueue.cpp customQueue.h
	g++ -O2 -std=c++11 -c -Wno-write-strings  customQueue.cpp


puzzle.o : puzzle.cpp puzzle.h
	g++ -O2 -std=c++11 -c -Wno-write-strings  puzzle.cpp
	
algorithm.o	: algorithm.cpp algorithm.h 
	g++ -O2 -std=c++11 -c -Wno-write-strings algorithm.cpp

graphics.o  : graphics.cpp graphics.h
	g++ -O2 -std=c++11 -c -Wno-write-strings graphics.cpp
	
clean:
	del *.o
	del *.exe
	

