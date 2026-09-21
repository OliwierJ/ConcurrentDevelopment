#Oliwier Jakubiec 17/9/2026
#Makefile
#Used to contain commands to compile projects
#See http://www.cs.colby.edu/maxwell/courses/tutorials/maketutor/ for info.


CC=g++
CFLAGS= -std=c++11 -pthread
ODIR=obj

PROGRAMS = helloThreads mutualEx

OBJS-HelloThreads_=helloThreads.o Semaphore.o
OBJS2-mutualEx_=mutualEx.o Semaphore.o
OBJ_HELLO = $(patsubst %,$(ODIR)/%,$(OBJS-HelloThreads_))
OBJ_MUTUAL = $(patsubst %,$(ODIR)/%,$(OBJS2-mutualEx_))

.PHONY: all run-hello run-mutual debug clean

$(ODIR): 
	mkdir -p $(ODIR)

#build and run the program
all: $(PROGRAMS)

#build helloThreads
helloThreads: $(OBJ_HELLO)
	$(CC) -o helloThreads $(OBJ_HELLO) $(CFLAGS) $(DEBUGFLAGS)

#build mutualEx
mutualEx: $(OBJ_MUTUAL)
	$(CC) -o mutualEx $(OBJ_MUTUAL) $(CFLAGS) $(DEBUGFLAGS)
	
#This rule says that each .o file depends on a .cpp file of the same name
$(ODIR)/%.o: %.cpp | $(ODIR)
	$(CC) -c -o $@ $< $(CFLAGS) $(DEBUGFLAGS)

#build and run helloThreads
run-hello: helloThreads
	./helloThreads

#build and run mutualEx
run-mutual: mutualEx
	./mutualEx
	
DEBUG: DEBUGFLAGS = -DDEBUG
DEBUG: all

#deletes all object files and the executable
clean:
	rm -rf $(ODIR) 
	rm -f $(PROGRAMS)
