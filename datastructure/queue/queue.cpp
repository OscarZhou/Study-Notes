#include <iostream>

using namespace std;


struct Puzzle
{
	int data;
	Puzzle* next;
};

class Queue
{
private:
	Puzzle* head;
	Puzzle* tail;
public:
	Queue();
	~Queue();
	
	void Enqueue(int newData);
	int Peek();
	void Dequeue();
	int IsEmpty();
	int Length();
	void Print();
};

Queue::Queue(){
	head = tail = NULL;
}

Queue::~Queue(){
	if(!IsEmpty()){
		delete head;
		head = tail = NULL;	
	}
}

void Queue::Enqueue(int newData){
	Puzzle *p = new Puzzle();
	p->data = newData;
	p->next = NULL;

	if(IsEmpty()){
		head = p;
		tail = p;
	}else{
		tail->next = p;
		tail = p;		
	}
}

int Queue::Peek(){
	if (IsEmpty()){
		cout << "the Queue is empty"<<endl;
		return -1;
	}
	return head->data;
}

void Queue::Dequeue(){
	if(IsEmpty()){
		cout << "the Queue is empty"<<endl;
		return ;
	}
	Puzzle *p = head;
	head = head->next;
	delete p;
}

int Queue::IsEmpty(){
	if(head == tail && head == NULL){
		return true;
	}
	return false;
}

int Queue::Length(){
	Puzzle *b = head;
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

void Queue::Print(){
	if(!IsEmpty()){
		Puzzle *cur = head;
		while (cur != NULL){
			cout<<cur->data<<" ";
			cur = cur->next;
		}
		cout<<endl;
	}
}

int main(){
	Queue q;
	q.Print();
	q.Enqueue(2);
	cout<<"length="<<q.Length()<<endl;
	q.Print();
	q.Enqueue(3);
	cout<<"length="<<q.Length()<<endl;
	q.Enqueue(4);
	cout<<"length="<<q.Length()<<endl;
	q.Enqueue(5);
	cout<<"length="<<q.Length()<<endl;
	q.Enqueue(6);
	cout<<"length="<<q.Length()<<endl;
	q.Print();
	q.Dequeue();
	cout<<"length="<<q.Length()<<endl;
	q.Dequeue();
	cout<<"length="<<q.Length()<<endl;
	q.Print();

}