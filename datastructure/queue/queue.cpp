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
	q.Print();
	q.Enqueue(3);
	q.Enqueue(4);
	q.Enqueue(5);
	q.Enqueue(6);
	q.Print();
	q.Dequeue();
	q.Dequeue();
	q.Print();

}