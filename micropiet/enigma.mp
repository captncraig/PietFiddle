//reflector
19 0 9 21 22 2 25 5 1 4 8 12 10 14 6 13 23 15 3 11 18 16 7 20 17 24 
3
rotorLoop:
	d
	br settings chooseRotor
	chooseRotor:
		n 49 - //input and reduce to proper range (0-4)
		d br rot1 cr1
		cr1:
			1 - d br rot2 cr2
		cr2:
			1 - d br rot3 cr3
		cr3:
			1 - br rot4 rot5
		rot1: pop
			//push z-a and then rollover position. rollover will be on top, followed by a-z wirings.
			9 2 17 1 8 0 15 18 20 23 7 24 22 14 19 13 25 21 16 3 6 11 5 12 10 4 16
			jmp rotorDone
		rot2: pop
			4 14 21 5 24 15 13 25 6 16 2 12 19 22 7 11 1 23 20 17 8 18 10 3 9 0 4
			jmp rotorDone
		rot3: pop
			14 16 18 20 12 10 0 6 22 8 4 24 13 25 21 23 19 17 15 2 11 9 7 5 3 1 21
			jmp rotorDone
		rot4:
			1 22 12 2 3 10 6 19 5 13 11 23 7 17 8 20 16 24 0 9 25 15 21 14 18 4 9
			jmp rotorDone
		rot5:
			10 2 4 5 14 16 9 12 22 0 23 11 7 13 3 18 15 20 24 19 8 6 17 1 25 21 25
		rotorDone:
			28 1 2 - r //bring up rotor counter
			1 -
			jmp rotorLoop
settings:
	n 65 - n 65 - n 65 -
	//RML
	n 65 - 4 1 2 - r 
	//L Loff RM
	+ 26 %
	//L' RM
	n 65 - 4 1 2 - r + 26 %
	//M' L' R
	n 65 - 4 1 2 - r + 26 %
	// R' M' L'















p