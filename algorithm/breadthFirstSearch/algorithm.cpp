

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
			// delete temPuzzle;
		}

		if(currentPuzzle.canMoveRight()){
			Puzzle *temPuzzle = currentPuzzle.moveRight();
			Q.Enqueue(*temPuzzle);
			// delete temPuzzle;
		}

		if(currentPuzzle.canMoveDown()){
			Puzzle *temPuzzle = currentPuzzle.moveDown();
			Q.Enqueue(*temPuzzle);
			// delete temPuzzle;
		}

		if(currentPuzzle.canMoveLeft()){
			Puzzle *temPuzzle = currentPuzzle.moveLeft();
			Q.Enqueue(*temPuzzle);
			// delete temPuzzle;
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
	Queue expandedQueue(startPuzzle);

    //algorithm implementation
	// cout << "------------------------------" << endl;
 //    cout << "<<breadthFirstSearch_with_VisitedList>>" << endl;
 //    cout << "------------------------------" << endl;


	startTime = clock();
	

	while(true){
		if(expandedQueue.IsEmpty()){
			break;
		}
		numOfStateExpansions++;

		Puzzle currentPuzzle = expandedQueue.Peek();
	
		if(currentPuzzle.goalMatch()){
			path = currentPuzzle.getPath();
			break;
		}

		expandedQueue.Dequeue();
		if(currentPuzzle.canMoveLeft()){
			Puzzle *temPuzzle = currentPuzzle.moveLeft();
			expandedQueue.Enqueue(*temPuzzle);
			delete temPuzzle;
		}

		if(currentPuzzle.canMoveRight()){
			Puzzle *temPuzzle = currentPuzzle.moveLeft();
			expandedQueue.Enqueue(*temPuzzle);
			delete temPuzzle;
		}

		if(currentPuzzle.canMoveUp()){
			Puzzle *temPuzzle = currentPuzzle.moveLeft();
			expandedQueue.Enqueue(*temPuzzle);
			delete temPuzzle;
		}

		if(currentPuzzle.canMoveDown()){
			Puzzle *temPuzzle = currentPuzzle.moveLeft();
			expandedQueue.Enqueue(*temPuzzle);
			delete temPuzzle;
		}
	}
	srand(time(NULL)); //RANDOM NUMBER GENERATOR - ONLY FOR THIS DEMO.  YOU REALLY DON'T NEED THIS! DISABLE THIS STATEMENT.
	maxQLength= rand() % 800; //AT THE MOMENT, THIS IS JUST GENERATING SOME DUMMY VALUE.  YOUR ALGORITHM IMPLEMENTATION SHOULD COMPUTE THIS PROPERLY.
	// numOfStateExpansions = rand() % 600; //AT THE MOMENT, THIS IS JUST GENERATING SOME DUMMY VALUE.  YOUR ALGORITHM IMPLEMENTATION SHOULD COMPUTE THIS PROPERLY


	
//***********************************************************************************************************
	actualRunningTime = ((float)(clock() - startTime)/CLOCKS_PER_SEC);
	path = "DDRRLLLUUU";  //this is just a dummy path for testing the function           
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


    //algorithm implementation
	// cout << "------------------------------" << endl;
 //    cout << "<<progressiveDeepeningSearch_No_VisitedList>>" << endl;
 //    cout << "------------------------------" << endl;

	startTime = clock();
	srand(time(NULL)); //RANDOM NUMBER GENERATOR - ONLY FOR THIS DEMO.  YOU REALLY DON'T NEED THIS! DISABLE THIS STATEMENT.
	maxQLength= rand() % 500; //AT THE MOMENT, THIS IS JUST GENERATING SOME DUMMY VALUE.  YOUR ALGORITHM IMPLEMENTATION SHOULD COMPUTE THIS PROPERLY.
	numOfStateExpansions = rand() % 600; //AT THE MOMENT, THIS IS JUST GENERATING SOME DUMMY VALUE.  YOUR ALGORITHM IMPLEMENTATION SHOULD COMPUTE THIS PROPERLY

	
	
//***********************************************************************************************************
	actualRunningTime = ((float)(clock() - startTime)/CLOCKS_PER_SEC);
	path = "DDRRLLLUUU";  //this is just a dummy path for testing the function           
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


    //algorithm implementation
	// cout << "------------------------------" << endl;
 //    cout << "<<progressiveDeepeningSearch_with_NonStrict_VisitedList>>" << endl;
 //    cout << "------------------------------" << endl;

	startTime = clock();
	srand(time(NULL)); //RANDOM NUMBER GENERATOR - ONLY FOR THIS DEMO.  YOU REALLY DON'T NEED THIS! DISABLE THIS STATEMENT.
	maxQLength= rand() % 300; //AT THE MOMENT, THIS IS JUST GENERATING SOME DUMMY VALUE.  YOUR ALGORITHM IMPLEMENTATION SHOULD COMPUTE THIS PROPERLY.
	numOfStateExpansions = rand() % 300; //AT THE MOMENT, THIS IS JUST GENERATING SOME DUMMY VALUE.  YOUR ALGORITHM IMPLEMENTATION SHOULD COMPUTE THIS PROPERLY

	
	
//***********************************************************************************************************
    actualRunningTime = ((float)(clock() - startTime)/CLOCKS_PER_SEC);
	path = "DDRRLLLUUU"; //this is just a dummy path for testing the function
	
	return path;		
		
}
	

string aStar_ExpandedList(string const initialState, string const goalState, int &numOfStateExpansions, int& maxQLength, 
                               float &actualRunningTime, int &numOfDeletionsFromMiddleOfHeap, int &numOfLocalLoopsAvoided, int &numOfAttemptedNodeReExpansions, heuristicFunction heuristic){
											 
   string path;
   clock_t startTime;
   
   numOfDeletionsFromMiddleOfHeap=0;
   numOfLocalLoopsAvoided=0;
   numOfAttemptedNodeReExpansions=0;


    // cout << "------------------------------" << endl;
    // cout << "<<aStar_ExpandedList>>" << endl;
    // cout << "------------------------------" << endl;
	actualRunningTime=0.0;	
	startTime = clock();
	srand(time(NULL)); //RANDOM NUMBER GENERATOR - ONLY FOR THIS DEMO.  YOU REALLY DON'T NEED THIS! DISABLE THIS STATEMENT.
	maxQLength= rand() % 200; //AT THE MOMENT, THIS IS JUST GENERATING SOME DUMMY VALUE.  YOUR ALGORITHM IMPLEMENTATION SHOULD COMPUTE THIS PROPERLY.
	numOfStateExpansions = rand() % 200; //AT THE MOMENT, THIS IS JUST GENERATING SOME DUMMY VALUE.  YOUR ALGORITHM IMPLEMENTATION SHOULD COMPUTE THIS PROPERLY


	
	
//***********************************************************************************************************
	actualRunningTime = ((float)(clock() - startTime)/CLOCKS_PER_SEC);
	path = "DDRRLLLUUU"; //this is just a dummy path for testing the function
	             
	return path;		
		
}




