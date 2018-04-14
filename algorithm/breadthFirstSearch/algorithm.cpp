

#include "algorithm.h"

using namespace std;



///////////////////////////////////////////////////////////////////////////////////////////
//
// Search Algorithm:  Breadth-First Search 
//
// Move Generator:  
//
////////////////////////////////////////////////////////////////////////////////////////////
string breadthFirstSearch(string const initialState, string const goalState, int &numOfStateExpansions, int& maxQLength, float &actualRunningTime){
    string path;
	clock_t startTime;
    //add necessary variables here

	Puzzle startPuzzle(initialState, goalState);
	Queue Q(startPuzzle);//1
	

    //algorithm implementation
	// cout << "------------------------------" << endl;
 //    cout << "<<breadthFirstSearch>>" << endl;
 //    cout << "------------------------------" << endl;
    
	startTime = clock();

	while(true){
		if(Q.IsEmpty()){
			break;
		}

		Puzzle currentPuzzle = Q.Peek();
		if(currentPuzzle.goalMatch()){
			path = currentPuzzle.getPath();
			maxQLength = Q.MaxLength();
			break;
		}

		Q.Dequeue();
		numOfStateExpansions++;
		if(currentPuzzle.canMoveUp()){
			Puzzle *temPuzzle = currentPuzzle.moveUp();
			Q.Enqueue(*temPuzzle);
			delete temPuzzle;
		}

		if(currentPuzzle.canMoveRight()){
			Puzzle *temPuzzle = currentPuzzle.moveRight();
			Q.Enqueue(*temPuzzle);
			delete temPuzzle;
		}

		if(currentPuzzle.canMoveDown()){
			Puzzle *temPuzzle = currentPuzzle.moveDown();
			Q.Enqueue(*temPuzzle);
			delete temPuzzle;
		}

		if(currentPuzzle.canMoveLeft()){
			Puzzle *temPuzzle = currentPuzzle.moveLeft();
			Q.Enqueue(*temPuzzle);
			delete temPuzzle;
		}
		if (Q.MaxLength() > 5000000){
			path = currentPuzzle.getPath();
			break;
		}
	}
	
	//srand(time(NULL)); //RANDOM NUMBER GENERATOR - ONLY FOR THIS DEMO.  YOU REALLY DON'T NEED THIS! DISABLE THIS STATEMENT.
	// maxQLength= rand() % 1500; //AT THE MOMENT, THIS IS JUST GENERATING SOME DUMMY VALUE.  YOUR ALGORITHM IMPLEMENTATION SHOULD COMPUTE THIS PROPERLY.
	// numOfStateExpansions = rand() % 800; //AT THE MOMENT, THIS IS JUST GENERATING SOME DUMMY VALUE.  YOUR ALGORITHM IMPLEMENTATION SHOULD COMPUTE THIS PROPERLY


	//***********************************************************************************************************
	actualRunningTime = ((float)(clock() - startTime)/CLOCKS_PER_SEC);
	//path = "DDRRLLLUUU";  //this is just a dummy path for testing the function           
	
	return path;		
		
}

///////////////////////////////////////////////////////////////////////////////////////////
//
// Search Algorithm:  Breadth-First Search with VisitedList
//
// Move Generator:  
//
////////////////////////////////////////////////////////////////////////////////////////////
string breadthFirstSearch_with_VisitedList(string const initialState, string const goalState, int &numOfStateExpansions, int& maxQLength, float &actualRunningTime){
    string path;
	clock_t startTime;
    //add necessary variables here
	Puzzle startPuzzle(initialState, goalState);
	Queue Q(startPuzzle);
	map<string, bool> visitedList;

    //algorithm implementation
	// cout << "------------------------------" << endl;
 //    cout << "<<breadthFirstSearch_with_VisitedList>>" << endl;
 //    cout << "------------------------------" << endl;


	startTime = clock();
	

	while(true){
		if(Q.IsEmpty()){
			break;
		}

		Puzzle currentPuzzle = Q.Peek();
		if(visitedList[currentPuzzle.getString()]){
			Q.Dequeue();
			continue;
		}	
		visitedList[currentPuzzle.getString()] = true;
		if(currentPuzzle.goalMatch()){
			path = currentPuzzle.getPath();
			visitedList.clear();
			break;
		}

		Q.Dequeue();
		numOfStateExpansions++;
		if(currentPuzzle.canMoveUp()){
			Puzzle *temPuzzle = currentPuzzle.moveUp();
			Q.Enqueue(*temPuzzle);
			delete temPuzzle;
		}

		if(currentPuzzle.canMoveRight()){
			Puzzle *temPuzzle = currentPuzzle.moveRight();
			Q.Enqueue(*temPuzzle);
			delete temPuzzle;
		}

		if(currentPuzzle.canMoveDown()){
			Puzzle *temPuzzle = currentPuzzle.moveDown();
			Q.Enqueue(*temPuzzle);
			delete temPuzzle;
		}

		if(currentPuzzle.canMoveLeft()){
			Puzzle *temPuzzle = currentPuzzle.moveLeft();
			Q.Enqueue(*temPuzzle);
			delete temPuzzle;
		}
		if (Q.MaxLength() > 5000000){
			path = currentPuzzle.getPath();
			break;
		}
	}
	srand(time(NULL)); //RANDOM NUMBER GENERATOR - ONLY FOR THIS DEMO.  YOU REALLY DON'T NEED THIS! DISABLE THIS STATEMENT.
	//maxQLength= rand() % 800; //AT THE MOMENT, THIS IS JUST GENERATING SOME DUMMY VALUE.  YOUR ALGORITHM IMPLEMENTATION SHOULD COMPUTE THIS PROPERLY.
	// numOfStateExpansions = rand() % 600; //AT THE MOMENT, THIS IS JUST GENERATING SOME DUMMY VALUE.  YOUR ALGORITHM IMPLEMENTATION SHOULD COMPUTE THIS PROPERLY


	
//***********************************************************************************************************
	actualRunningTime = ((float)(clock() - startTime)/CLOCKS_PER_SEC);
	//path = "DDRRLLLUUU";  //this is just a dummy path for testing the function           
	return path;		
		
}

///////////////////////////////////////////////////////////////////////////////////////////
//
// Search Algorithm:  
//
// Move Generator:  
//
////////////////////////////////////////////////////////////////////////////////////////////
string progressiveDeepeningSearch_No_VisitedList(string const initialState, string const goalState, int &numOfStateExpansions, int& maxQLength, float &actualRunningTime, int ultimateMaxDepth){
    string path;
	clock_t startTime;
    //add necessary variables here
    int currentThreshold = 1;
	Puzzle startPuzzle(initialState, goalState);
	startPuzzle.setDepth(0);
	Stack Q(startPuzzle);	
	

    //algorithm implementation
	// cout << "------------------------------" << endl;
 //    cout << "<<progressiveDeepeningSearch_No_VisitedList>>" << endl;
 //    cout << "------------------------------" << endl;

	startTime = clock();

	while(true){
		if(Q.IsEmpty()){
			if(currentThreshold < ultimateMaxDepth){
				currentThreshold++;
				Q.Push(startPuzzle);
				continue;
			}
			break;
		}
		
		Puzzle currentPuzzle = Q.Top();
		if(currentPuzzle.getDepth() <= currentThreshold){
			if(currentPuzzle.goalMatch()){
				path = currentPuzzle.getPath();
				maxQLength = Q.MaxLength();
				break;
			}

			Q.Pop();
			numOfStateExpansions++;

			if(currentPuzzle.canMoveLeft(currentThreshold)){
				Puzzle *temPuzzle = currentPuzzle.moveLeft();
				Q.Push(*temPuzzle);
				delete temPuzzle;
			}

			if(currentPuzzle.canMoveDown(currentThreshold)){
				Puzzle *temPuzzle = currentPuzzle.moveDown();
				Q.Push(*temPuzzle);
				delete temPuzzle;
			}

			if(currentPuzzle.canMoveRight(currentThreshold)){
				Puzzle *temPuzzle = currentPuzzle.moveRight();
				Q.Push(*temPuzzle);
				delete temPuzzle;
			}

			if(currentPuzzle.canMoveUp(currentThreshold)){
				Puzzle *temPuzzle = currentPuzzle.moveUp();
				Q.Push(*temPuzzle);
				delete temPuzzle;
			}

			if (Q.MaxLength() > 5000000){
				path = currentPuzzle.getPath();
				break;
			}
		}else{
			break;
		}
	}


	srand(time(NULL)); //RANDOM NUMBER GENERATOR - ONLY FOR THIS DEMO.  YOU REALLY DON'T NEED THIS! DISABLE THIS STATEMENT.
	// maxQLength= rand() % 500; //AT THE MOMENT, THIS IS JUST GENERATING SOME DUMMY VALUE.  YOUR ALGORITHM IMPLEMENTATION SHOULD COMPUTE THIS PROPERLY.
	// numOfStateExpansions = rand() % 600; //AT THE MOMENT, THIS IS JUST GENERATING SOME DUMMY VALUE.  YOUR ALGORITHM IMPLEMENTATION SHOULD COMPUTE THIS PROPERLY

	
	
//***********************************************************************************************************
	actualRunningTime = ((float)(clock() - startTime)/CLOCKS_PER_SEC);
	//path = "DDRRLLLUUU";  //this is just a dummy path for testing the function           
	return path;		
		
}
	



///////////////////////////////////////////////////////////////////////////////////////////
//
// Search Algorithm:  
//
// Move Generator:  
//
////////////////////////////////////////////////////////////////////////////////////////////
string progressiveDeepeningSearch_with_NonStrict_VisitedList(string const initialState, string const goalState, int &numOfStateExpansions, int& maxQLength, float &actualRunningTime, int ultimateMaxDepth){
    string path;
	clock_t startTime;
    //add necessary variables here
    int currentThreshold = 1;
	Puzzle startPuzzle(initialState, goalState);
	startPuzzle.setDepth(0);
	Stack Q(startPuzzle);	
	map<string, bool> nonStrictList;


    //algorithm implementation
	// cout << "------------------------------" << endl;
 //    cout << "<<progressiveDeepeningSearch_with_NonStrict_VisitedList>>" << endl;
 //    cout << "------------------------------" << endl;

	startTime = clock();
	while(true){
		if(Q.IsEmpty()){
			if(currentThreshold < ultimateMaxDepth){
				currentThreshold++;
				Q.Push(startPuzzle);
				nonStrictList.clear();
				continue;
			}
			break;
		}
		
		Puzzle currentPuzzle = Q.Top();

		if(currentPuzzle.getDepth() <= currentThreshold){
			string nonStrictListKey = currentPuzzle.getString()+to_string(currentPuzzle.getDepth());
			if(nonStrictList[nonStrictListKey]){
				Q.Pop();
				continue;
			}
			nonStrictList[nonStrictListKey] = true;

			if(currentPuzzle.goalMatch()){
				path = currentPuzzle.getPath();
				maxQLength = Q.MaxLength();
				nonStrictList.clear();
				break;
			}

			Q.Pop();
			numOfStateExpansions++;

			if(currentPuzzle.canMoveLeft(currentThreshold)){
				Puzzle *temPuzzle = currentPuzzle.moveLeft();
				Q.Push(*temPuzzle);
				delete temPuzzle;
			}

			if(currentPuzzle.canMoveDown(currentThreshold)){
				Puzzle *temPuzzle = currentPuzzle.moveDown();
				Q.Push(*temPuzzle);
				delete temPuzzle;
			}

			if(currentPuzzle.canMoveRight(currentThreshold)){
				Puzzle *temPuzzle = currentPuzzle.moveRight();
				Q.Push(*temPuzzle);
				delete temPuzzle;
			}

			if(currentPuzzle.canMoveUp(currentThreshold)){
				Puzzle *temPuzzle = currentPuzzle.moveUp();
				Q.Push(*temPuzzle);
				delete temPuzzle;
			}

			if (Q.MaxLength() > 5000000){
				path = currentPuzzle.getPath();
				break;
			}
		}else{
			break;
		}
	}

	srand(time(NULL)); //RANDOM NUMBER GENERATOR - ONLY FOR THIS DEMO.  YOU REALLY DON'T NEED THIS! DISABLE THIS STATEMENT.
	// maxQLength= rand() % 300; //AT THE MOMENT, THIS IS JUST GENERATING SOME DUMMY VALUE.  YOUR ALGORITHM IMPLEMENTATION SHOULD COMPUTE THIS PROPERLY.
	// numOfStateExpansions = rand() % 300; //AT THE MOMENT, THIS IS JUST GENERATING SOME DUMMY VALUE.  YOUR ALGORITHM IMPLEMENTATION SHOULD COMPUTE THIS PROPERLY

	
	
//***********************************************************************************************************
    actualRunningTime = ((float)(clock() - startTime)/CLOCKS_PER_SEC);
	// path = "DDRRLLLUUU"; //this is just a dummy path for testing the function
	
	return path;		
		
}
	

string aStar_ExpandedList(string const initialState, string const goalState, int &numOfStateExpansions, int& maxQLength, 
                               float &actualRunningTime, int &numOfDeletionsFromMiddleOfHeap, int &numOfLocalLoopsAvoided, int &numOfAttemptedNodeReExpansions, heuristicFunction heuristic){										 
	string path;
	clock_t startTime;

	numOfDeletionsFromMiddleOfHeap=0;
	numOfLocalLoopsAvoided=0;
	numOfAttemptedNodeReExpansions=0;

   	Puzzle startPuzzle(initialState, goalState);
   	startPuzzle.updateHCost(heuristic);
   	startPuzzle.updateFCost();
	Heap PriorityQ(startPuzzle);	
	map<string, bool> expandedList;


    // cout << "------------------------------" << endl;
    // cout << "<<aStar_ExpandedList>>" << endl;
    // cout << "------------------------------" << endl;
	actualRunningTime=0.0;	
	startTime = clock();

	while(true){
		if(PriorityQ.IsEmpty()){
			break;
		}

		Puzzle currentPuzzle = PriorityQ.Root();
		
		if(currentPuzzle.goalMatch()){
			cout<<"1"<<endl;
			path = currentPuzzle.getPath();
			expandedList.clear();
			break;
		}
		PriorityQ.Print();

		if(expandedList[currentPuzzle.getString()]){
			PriorityQ.Print();
			PriorityQ.DeleteRoot();
			continue;
		}	
		PriorityQ.Print();

		expandedList[currentPuzzle.getString()] = true;
		PriorityQ.DeleteRoot();

		PriorityQ.Print();

		numOfStateExpansions++;
		if(currentPuzzle.canMoveUp()){
			Puzzle *temPuzzle = currentPuzzle.moveUp();
			if(!expandedList[temPuzzle->getString()]){
				temPuzzle->updateHCost(heuristic);
   				temPuzzle->updateFCost();
   				PriorityQ.InsertOrReplace(*temPuzzle);

			}
			delete temPuzzle;
		}
PriorityQ.Print();
		if(currentPuzzle.canMoveRight()){
			Puzzle *temPuzzle = currentPuzzle.moveRight();
			if(!expandedList[temPuzzle->getString()]){
				temPuzzle->updateHCost(heuristic);
   				temPuzzle->updateFCost();
   				PriorityQ.InsertOrReplace(*temPuzzle);
			}
			delete temPuzzle;
		}
PriorityQ.Print();
		if(currentPuzzle.canMoveDown()){
			Puzzle *temPuzzle = currentPuzzle.moveDown();
			if(!expandedList[temPuzzle->getString()]){
				temPuzzle->updateHCost(heuristic);
   				temPuzzle->updateFCost();
   				PriorityQ.InsertOrReplace(*temPuzzle);
			}
			delete temPuzzle;
		}
PriorityQ.Print();
		if(currentPuzzle.canMoveLeft()){
			Puzzle *temPuzzle = currentPuzzle.moveLeft();
			if(!expandedList[temPuzzle->getString()]){
				temPuzzle->updateHCost(heuristic);
   				temPuzzle->updateFCost();
   				PriorityQ.InsertOrReplace(*temPuzzle);
			}
			delete temPuzzle;
		}
		if (PriorityQ.MaxLength() > 5000000){
			path = currentPuzzle.getPath();
			break;
		}
		PriorityQ.Print();
	}
	srand(time(NULL)); //RANDOM NUMBER GENERATOR - ONLY FOR THIS DEMO.  YOU REALLY DON'T NEED THIS! DISABLE THIS STATEMENT.
	// maxQLength= rand() % 200; //AT THE MOMENT, THIS IS JUST GENERATING SOME DUMMY VALUE.  YOUR ALGORITHM IMPLEMENTATION SHOULD COMPUTE THIS PROPERLY.
	// numOfStateExpansions = rand() % 200; //AT THE MOMENT, THIS IS JUST GENERATING SOME DUMMY VALUE.  YOUR ALGORITHM IMPLEMENTATION SHOULD COMPUTE THIS PROPERLY


	
	
//***********************************************************************************************************
	actualRunningTime = ((float)(clock() - startTime)/CLOCKS_PER_SEC);
	// path = "DDRRLLLUUU"; //this is just a dummy path for testing the function
	             
	return path;		
		
}




