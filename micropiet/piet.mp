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


//input rem size x y .....
//put into program locationd
5 moveAfter  
//decrement remaining and break loop if necessary
1 - d
br readProgram doneReading

:doneReading

11- //dp
d //cc
d //x
d //y
d //tries
d //blocksize
d //stacksize

//MAIN LOOP STACK: x y dp cc stacksize blocksize tries progSize width height [program] [stack]

7 moveUp d 8 moveAfter > //EXIT IF GREATER

//FINDEXIT

//maxX = x
d
//maxy = y
3 1 2 - r d 4 1 r
//index = maxY * width + xd
d 12 1 2 - r d 13 1 r *
3 1 2 - r d 4 1 r +
//currentcolor = program[index]
d 15 + d 1 2 - r 
d 3 1 2 - r- p r
//blocksize = 0
10 1 2 - r pop 1 1 - 10 1 r


