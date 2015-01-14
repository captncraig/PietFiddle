#pullUp{
	//pull the nth item up to the top
	1 2 - r
}
#pushDown{
	//push down to under nth
	1 r 
}
#copyUp{
	//pull nth item up. copy it. Put it back
	d 1 + pullUp d
	3 pullUp 1 + pushDown
}


1 2 3 4 5
3 pullUp
@@ 3 5 4 2 1
2 pullUp
@@ 5 3 4 2 1

4 pushDown
@@ 3 4 2 5 1

4 copyUp
@@ 5 3 4 2 5 1
6 copyUp
@@ 1 5 3 4 2 5 1
1 copyUp
@@ 1 1 5 3 4 2 5 1
2 1 pullUp
@@ 2 1 1 5 3 4 2 5 1
