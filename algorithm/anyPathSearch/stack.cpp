#include "stack.h"

using namespace std;

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
		head = NULL;	
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

