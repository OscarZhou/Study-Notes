#include <iostream>

using namespace std;


struct Puzzle
{
	int data;
	Puzzle* next;
};

class Stack
{
private:
	Puzzle* top;
public:
	Stack();
	~Stack();
	
	void Push(int newData);
	int Top();
	void Pop();
	int IsEmpty();
	int Length();
	void Print();
};

Stack::Stack(){
	top = NULL;
}

Stack::~Stack(){
	if(!IsEmpty()){
		delete top;
		top = NULL;	
	}
}

void Stack::Push(int newData){
	Puzzle *p = new Puzzle();
	p->data = newData;
	p->next = NULL;

	p->next = top;
	top = p;
}

int Stack::Top(){
	if (IsEmpty()){
		cout << "the Stack is empty"<<endl;
		return -1;
	}
	return top->data;
}

void Stack::Pop(){
	if(IsEmpty()){
		cout << "the Stack is empty"<<endl;
		return ;
	}
	Puzzle *p = top;
	top = top->next;
	delete p;
}

int Stack::IsEmpty(){
	if(top == NULL){
		return true;
	}
	return false;
}

int Stack::Length(){
	Puzzle *b = top;
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

void Stack::Print(){
	if(!IsEmpty()){
		Puzzle *cur = top;
		while (cur != NULL){
			cout<<cur->data<<" ";
			cur = cur->next;
		}
		cout<<endl;
	}
}

int main(){
	Stack q;
	q.Print();
	q.Push(2);
	cout<<"length="<<q.Length()<<endl;
	q.Print();
	q.Push(3);
	cout<<"length="<<q.Length()<<endl;
	q.Push(4);
	cout<<"length="<<q.Length()<<endl;
	q.Push(5);
	cout<<"length="<<q.Length()<<endl;
	q.Push(6);
	cout<<"length="<<q.Length()<<endl;
	q.Print();
	q.Pop();
	cout<<"length="<<q.Length()<<endl;
	q.Pop();
	cout<<"length="<<q.Length()<<endl;
	q.Print();
	cout<<"current data="<<q.Top()<<endl;
}