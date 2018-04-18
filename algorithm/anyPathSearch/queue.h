#include "puzzle.h"
#include <iostream>

using namespace std;

struct Board {
	Puzzle* puzzle;
	Board* next;
};

class Queue
{
private:
	Board* head;
	Board* tail;
	int maxLength;
public:
	Queue();
	Queue(const Puzzle &p);
	~Queue();
	
	void Enqueue(const Puzzle &p);
	Puzzle Peek();
	void Dequeue();
	bool IsEmpty();
	int Length();
	int MaxLength();
	void Print();
};



