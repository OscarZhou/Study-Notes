#include "puzzle.h"
#include <iostream>
#include <vector>

using namespace std;

struct Node {
	Puzzle* puzzle;
	Node* next;
};


/// Queue for BFS

class Queue
{
private:
	Node* head;
	Node* tail;
	int maxLength;
	int currentLength;
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
	
};


// Queue for PFS

class Stack
{
private:
	Node* top;
	int maxLength;
	int currentLength;
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
	
};





// Queue for Astar


class Heap{
private:
	std::vector<Puzzle> v;
	int last;
	int maxLength;
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
	void InsertOrReplace(const Puzzle &p, int &numOfDeletionsFromMiddleOfHeap);
	int MaxLength();

};