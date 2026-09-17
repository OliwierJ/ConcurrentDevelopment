#Oliwier Jakubiec 17/9/2026
#Makefile
#Used to contain commands to compile projects
#See http://www.cs.colby.edu/maxwell/courses/tutorials/maketutor/ for info.


CC=g++
CFLAGS= -std=c++11 -pthread
EXEC=helloThreads
OBJS=helloThreads.o Semaphore.o

#build and run the program
ALL: BUILD
	./$(EXEC)

#This rule says that each .o file depends on a .cpp file of the same name
%.o: %.cpp
	$(CC) -c -o $@ $< $(CFLAGS) $(DEBUGFLAGS)

BUILD: $(OBJS)
	$(CC) -o $(EXEC) $(OBJS) $(CFLAGS) $(DEBUGFLAGS)

DEBUG: DEBUGFLAGS = -DDEBUG
DEBUG: BUILD

#deletes all object files and the executable
CLEAN:
	rm *.o
	rm $(EXEC)
