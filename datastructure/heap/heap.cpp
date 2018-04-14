#include <iostream>
#include <vector>

using namespace std;


struct Puzzle
{
	int cost;
};

class Heap
{
private:
	std::vector<Puzzle> puzzle;
	int last;
public:
	Heap();
	~Heap();
	
	void InsertHeap(int newData);
	int Root();
	void DeleteRoot();
	int IsEmpty();
	int Length();
	void Print();
};

Heap::Heap(){
	last = -1;
}

Heap::~Heap(){
	puzzle.clear();
}

void Heap::InsertHeap(int newData){
	Puzzle *p = new Puzzle();
	p->cost = newData;
	puzzle.push_back(*p);
	last += 1;
	if(last == 0){
		return;
	}

	int childIndex = last;
	int parentIndex = 0;
	bool swapping = true;
	while(swapping){
		swapping = false;
		parentIndex = (childIndex-1)/2;
		if(parentIndex>=0){
			if(puzzle[childIndex].cost>puzzle[parentIndex].cost){
				swap(puzzle[childIndex], puzzle[parentIndex]);
				swapping = true;
				childIndex = parentIndex;
			}
		}
	}
	delete p;
}

int Heap::Root(){
	if (last >= 0){
		return puzzle[0].cost;
	}
	return -1;
}

void Heap::DeleteRoot(){
	if(last <0){
		return;
	}
	unsigned int deletedData = puzzle[0].cost;
	puzzle[0] = puzzle[last];
	puzzle[last].cost = 0;
	last -= 1;
	int parentIndex = 0;
	int leftChildIndex = parentIndex*2+1;
	int rightChildIndex = parentIndex*2+2;
	while(puzzle[parentIndex].cost<puzzle[leftChildIndex].cost || puzzle[parentIndex].cost< puzzle[rightChildIndex].cost){
		if(puzzle[leftChildIndex].cost<puzzle[rightChildIndex].cost){
			swap(puzzle[parentIndex],puzzle[rightChildIndex]);
			parentIndex = rightChildIndex;
		}else{
			swap(puzzle[parentIndex], puzzle[leftChildIndex]);
			parentIndex = leftChildIndex;
		}
		leftChildIndex = parentIndex*2+1;
		rightChildIndex = parentIndex*2+2;
		if(leftChildIndex>last){
			break;
		}else{
			if (rightChildIndex>=last){
				if(puzzle[parentIndex].cost<puzzle[leftChildIndex].cost){
					swap(puzzle[parentIndex],puzzle[leftChildIndex]);

				}
				break;
			}
		}
	}
}

bool Heap::IsEmpty(){
	return puzzle.empty();
}

int Heap::Length(){
	return puzzle.size();
}

void Heap::Print(){
	cout<<"last="<<last<<endl;
	for(std::vector<Puzzle>::iterator it=puzzle.begin();it!=puzzle.end();++it){
		cout<<' '<<(*it).cost;
	}
	cout<<endl;
}

int main(){
	Heap q;
	q.Print();
	q.InsertHeap(10);
	q.Print();
	q.InsertHeap(11);
	q.Print();
	q.InsertHeap(11);
	q.Print();
	q.InsertHeap(8);
	q.Print();
	q.InsertHeap(7);
	q.Print();
	q.InsertHeap(2);
	q.InsertHeap(11);
	q.Print();
	q.InsertHeap(9);
	q.Print();
	q.DeleteRoot();
	q.Print();

}