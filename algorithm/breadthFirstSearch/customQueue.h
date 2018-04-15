#include "puzzle.h"
#include <iostream>
#include <vector>

using namespace std;

struct Board {
	Puzzle* puzzle;
	Board* next;
};


/// Queue for BFS

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


// Queue for PFS

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





// Queue for Astar


class Heap{
private:
	std::vector<Puzzle> v;
	int last;
	int t;
public:
	Heap();
	Heap(const Puzzle &p);
	~Heap();

	void InsertHeap(const Puzzle &p);
	Puzzle Root();
	void DeleteRoot();
	void Delete(const Puzzle &p);
	bool IsEmpty();
	int Length();
	void InsertOrReplace(const Puzzle &p);
	int MaxLength();

	void Print();
};