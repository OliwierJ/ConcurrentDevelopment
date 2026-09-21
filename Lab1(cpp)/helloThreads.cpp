#include "Semaphore.h"
#include <iostream>
#include <thread>
#include <chrono>

/*! displays a message first*/
void taskOne(std::shared_ptr<Semaphore> theSemaphore, int delay) {
  std::this_thread::sleep_for(std::chrono::seconds(delay));
  std::cout <<"I ";
  std::cout << "must ";
  std::cout << "print ";
  std::cout << "first. \n"<<std::endl;

  theSemaphore->Signal(); //signal the semaphore to allow taskTwo to continue
}

/*! displays a message second*/
void taskTwo(std::shared_ptr<Semaphore> theSemaphore) {
  //wait here until taskOne finishes...
  theSemaphore->Wait();

  std::cout <<"This ";
  std::cout << "will ";
  std::this_thread::sleep_for(std::chrono::seconds(5));
  std::cout << "appear ";
  std::cout << "second"<<std::endl;
}


int main(void) {
  std::thread threadOne, threadTwo;
  std::shared_ptr<Semaphore> sem(new Semaphore);
  
  /**< Launch the threads  */
  int taskOneDelay=5;
  std::cout << "Launched from the main\n";
  threadOne=std::thread(taskTwo,sem);
  threadTwo=std::thread(taskOne,sem,taskOneDelay);
  
  /**< Wait for the threads to finish */
  threadOne.join();
  threadTwo.join();
  return 0;
}
