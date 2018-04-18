#include "puzzle.h"
#include <iostream>

using namespace std;

struct Board {
	Puzzle* puzzle;
	Board* next;
};

class Stack
{
private:
	Board* top;
	int maxLength;
public:
	Stack();
	Stack(const Puzzle &p);
	~Stack();
	
	void Push(const Puzzle &p);
	Puzzle Top();
	void Pop();
	bool IsEmpty();
	int Length();
	int MaxLength();
	void Print();
};



