#include "customQueue.h"

using namespace std;


/// Queue for BFS
Queue::Queue(){
	head = tail = NULL;
	currentLength = maxLength = 0;
}

Queue::Queue(const Puzzle &p){
	Board *b = new Board();
	b->puzzle = new Puzzle(p);
	b->next = NULL;
	head = b;
	tail = b;
	currentLength = maxLength = 1;
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
	currentLength++;
	if (currentLength>maxLength){
		maxLength = currentLength;
	}
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
	currentLength--;
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


// Queue for PFS

Stack::Stack(){
	top = NULL;
	currentLength = maxLength = 0;
}

Stack::Stack(const Puzzle &p){
	Board *b = new Board();
	b->puzzle = new Puzzle(p);
	b->next = NULL;

	top = b;
	currentLength = maxLength = 0;
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
	currentLength++;
	if(currentLength>maxLength){
		maxLength = currentLength;
	}
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
	currentLength--;
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

// Queue for Astar


Heap::Heap(){
	last = maxLength = -1;
}

Heap::Heap(const Puzzle &p){
	last = maxLength = -1;
	v.push_back(p);
	last += 1;
	t= 0;
}

Heap::~Heap(){
	v.clear();
	last = -1;
}


void Heap::InsertHeap(const Puzzle &p){
	v.push_back(p);
	last += 1;
	if(last>maxLength){
		maxLength = last;
	}
	if(last == 0){
		return;
	}

	int childIndex = last;
	int parentIndex = 0;
	bool swapping = true;
	while(swapping){
		swapping = false;
		parentIndex = (childIndex-1)/2;
		if (parentIndex>=0){
			if(v[childIndex].getFCost()<v[parentIndex].getFCost()){
				swap(v[childIndex], v[parentIndex]);
				swapping = true;
				childIndex = parentIndex;
			}
		}
	}

}

Puzzle Heap::Root(){
	if(last >=0){
		return v[0];
	}
}

void Heap::DeleteRoot(){
	if (last <0){
		return;
	}
	Puzzle deletePuzzle = v[0];
	v[0] = v[last];
	v.pop_back();
	last -= 1;
	// There are at least two elements in the Q
	// Then the Q can be fixed
	if (last <1){ 
		return;
	}
	int parentIndex = 0;
	int leftChildIndex = parentIndex*2+1;
	int rightChildIndex = parentIndex*2+2;
	while(v[parentIndex].getFCost()>v[leftChildIndex].getFCost() || v[parentIndex].getFCost()>v[rightChildIndex].getFCost()){
		if(v[leftChildIndex].getFCost()<v[rightChildIndex].getFCost()){
			swap(v[parentIndex],v[leftChildIndex]);
			parentIndex = leftChildIndex;
		}else{
			swap(v[parentIndex],v[rightChildIndex]);
			parentIndex = rightChildIndex;
		}

		leftChildIndex = parentIndex*2+1;
		rightChildIndex = parentIndex*2+2;

		if(leftChildIndex>last){
			break;
		}else{
			if (rightChildIndex > last ){
				if(v[parentIndex].getFCost()>v[leftChildIndex].getFCost()){
					swap(v[parentIndex], v[leftChildIndex]);
				}
				break;
			}
		}
	}
}

void Heap::Delete(const Puzzle &p){
	if (last <0){
		return;
	}
	
	int parentIndex = 0;
	Puzzle deletePuzzle = v[0];

	for(std::vector<Puzzle>::iterator it = v.begin() ; it!= v.end();++it){
		if((*it).getString() == p.getString()){
			deletePuzzle = *it;
			break;
		}
		parentIndex++;
	}
	v[parentIndex] = v[last];
	v.pop_back();
	last -= 1;
	
	if (last <1){ 
		return;
	}
	int leftChildIndex = parentIndex*2+1;
	int rightChildIndex = parentIndex*2+2;
	if(leftChildIndex>last){
		// cout<<"left parentIndex="<<parentIndex<<endl;
		return;
	}else{
		if (rightChildIndex > last ){
			if(v[parentIndex].getFCost()>v[leftChildIndex].getFCost()){
				swap(v[parentIndex], v[leftChildIndex]);
			}
			// cout<<"right parentIndex="<<parentIndex<<endl;
			return;
		}
	}
	while(v[parentIndex].getFCost()>v[leftChildIndex].getFCost() || v[parentIndex].getFCost()>v[rightChildIndex].getFCost()){
		if(v[leftChildIndex].getFCost()<v[rightChildIndex].getFCost()){
			swap(v[parentIndex],v[leftChildIndex]);
			parentIndex = leftChildIndex;
		}else{
			swap(v[parentIndex],v[rightChildIndex]);
			parentIndex = rightChildIndex;
		}
		leftChildIndex = parentIndex*2+1;
		rightChildIndex = parentIndex*2+2;

		if(leftChildIndex>last){
			break;
		}else{
			if (rightChildIndex > last ){
				if(v[parentIndex].getFCost()>v[leftChildIndex].getFCost()){
					swap(v[parentIndex], v[leftChildIndex]);
				}
				break;
			}
		}
	}
}

bool Heap::IsEmpty(){
	return v.empty();
}

int Heap::Length(){
	return v.size();
}

void Heap::InsertOrReplace(const Puzzle &p, int &numOfDeletionsFromMiddleOfHeap){
	bool isFind = false;
	
	for(std::vector<Puzzle>::iterator it=v.begin(); it!=v.end(); ++it){
		if((*it).getString() == p.getString() && (*it).getFCost()>p.getFCost()){
			numOfDeletionsFromMiddleOfHeap++;
			Delete(*it);
			InsertHeap(p);
			isFind = true;
			break;
		}
	}
	if (!isFind){
		InsertHeap(p);
	}
}

int Heap::MaxLength(){
	return maxLength;
}

void Heap::Print(){
	// for(std::vector<Puzzle>::iterator it=v.begin(); it!=v.end(); ++it){
	// 	cout<<"t="<<t<<" "<<(*it).getString()<<", hcost="<<(*it).getHCost()<<", fcost="<<(*it).getFCost()<<endl;
	// }
	t++;
	cout<<"t="<<t<<endl;
	// cout<<"0="<<v[0].getFCost()<< " 1="<<v[1].getFCost()<<" 2="<<v[2].getFCost()<<endl;
}

void Heap::PrintLast(){
	cout<<"last="<<last<<endl;

}