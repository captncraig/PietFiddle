#moveUp{ //move nth cell to top. tip of stack is 0
	1 2 - r	
}
#moveAfter{ //move to after the nth cell
	1 r
}

//read width and height, calculate size
print "width:"
Id
print "height:"
Id
//hhww
3 moveUp 
*
d
//remaining size h w

:readProgram
//read char and put from A-Z range to 0-18 range
i 65 -
//put into program location
4 moveAfter  
//decrement remaining and break loop if necessary
1 - d
br readProgram doneReading

:doneReading
p
