#include "customQueue.h"

using namespace std;

Queue::Queue(){
	head = tail = NULL;
	maxLength = 0;
}

Queue::Queue(const Puzzle &p){
	Board *b = new Board();
	b->puzzle = new Puzzle(p);
	b->next = NULL;
	head = b;
	tail = b;
	maxLength = 0;
}

Queue::~Queue(){
	if(!IsEmpty()){
		delete head;
		head = tail = NULL;	
	}
}

void Queue::Enqueue(const Puzzle &p){
	Board *b = new Board();
	b->puzzle = new Puzzle(p);
	b->next = NULL;

	if(IsEmpty()){
		head = b;
		tail = b;
	}else{
		tail->next = b;
		tail = b;		
	}
	maxLength++;
}

Puzzle Queue::Peek(){
	if (!IsEmpty()){
		return *(head->puzzle);
	}
}

void Queue::Dequeue(){
	Board *b = head;
	if(head == tail){
		head = tail = NULL;
	}else{
		head = head->next;	
	}
	delete b;
	maxLength--;
}

bool Queue::IsEmpty(){
	if(head == tail && head == NULL){
		return true;
	}
	return false;
}

int Queue::Length(){
	Board *b = head;
	int length = 0;
	if (b != NULL){
		length = 1;
	}
	while(b->next != NULL){
		length++;
		b = b->next;
	}
	return length;
}

int Queue::MaxLength(){
	return maxLength;
}




Stack::Stack(){
	top = NULL;
	maxLength = 0;
}

Stack::Stack(const Puzzle &p){
	Board *b = new Board();
	b->puzzle = new Puzzle(p);
	b->next = NULL;

	top = b;
	maxLength = 0;
}

Stack::~Stack(){
	if(!IsEmpty()){
		delete top;
		top = NULL;	
	}
}

void Stack::Push(const Puzzle &p){
	Board *b = new Board();
	b->puzzle = new Puzzle(p);
	b->next = NULL;

	b->next = top;
	top = b;
	maxLength++;
}

Puzzle Stack::Top(){
	if (!IsEmpty()){
		return *(top->puzzle);
	}
}

void Stack::Pop(){
	Board *b = top;
	if (!IsEmpty()){
		top = top->next;
	}
	delete b;
	maxLength--;
}

bool Stack::IsEmpty(){
	if(top == NULL){
		return true;
	}
	return false;
}

int Stack::Length(){
	Board *b = top;
	int length = 0;
	if (b != NULL){
		length = 1;
	}
	while(b->next != NULL){
		length++;
		b = b->next;
	}
	return length;
}

int Stack::MaxLength(){
	return maxLength;
}

