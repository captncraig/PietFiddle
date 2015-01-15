angular.module('piet',[])

angular.module('piet')
	.controller('EditorCtrl', function EditorCtrl($scope,$http) {
	$scope.palette = makePalette()
	$scope.program = makeProgram(10,10,null) 
	$scope.settings = {}
	$scope.editState = {selectedColor:'A',painting:false,rightDown:false,filled:false}
	$scope.hover = {size:0}
	console.log($scope.program)
	
	$scope.getCellText = function(cell){
		return ""
	}
	$scope.getCellStyle = function(cell){
		obj = {"box-sizing":"border-box"}
		style = "1px solid black"
		edgeStyle = "2px solid black"
		if(cell.x == 0){
			obj["border-left"] = edgeStyle
		}
		else if($scope.program.rows[cell.y].cells[cell.x-1].color != cell.color){
			obj["border-left"] = style
		}
		if(cell.x >= ($scope.program.w - 1)){
			obj["border-right"] = edgeStyle
		} 
		else if($scope.program.rows[cell.y].cells[cell.x+1].color != cell.color){
			obj["border-right"] = style
		}
		if(cell.y == 0){
			obj["border-top"] = edgeStyle
		}
		else if( $scope.program.rows[cell.y-1].cells[cell.x].color != cell.color){
			obj["border-top"] = style
		}
		if(cell.y >= ($scope.program.h - 1)){
			obj["border-bottom"] = edgeStyle
		} 
		else if($scope.program.rows[cell.y+1].cells[cell.x].color != cell.color){
			obj["border-bottom"] = style
		}
		return obj
	}
	$scope.mouseDown = function(cell,ev){
		//left click
		if(ev.which == 1){
			//right held down. Flood fill.
			if($scope.editState.rightDown){
				$scope.editState.filled = true;
				$scope.hover.size = floodFill(cell,$scope.program,$scope.editState.selectedColor)
				$scope.hover.size = floodFill(cell,$scope.program)
			}
			//regular click. start painting
			else{
				cell.color = $scope.editState.selectedColor
				$scope.editState.painting = true
				$scope.hover.size = floodFill(cell,$scope.program)
			}
		}
		else if(ev.which == 3){
			$scope.editState.filled = false;
			$scope.editState.rightDown = true;
		}
	}
	$scope.mouseEnter = function(cell){
		if($scope.editState.painting){
			cell.color = $scope.editState.selectedColor
		}
		$scope.hover.size = floodFill(cell,$scope.program)
	}
	$scope.mouseUp = function(cell,ev){
		$scope.editState.painting = false;
		if(ev.which == 3){
			$scope.editState.rightDown = false;
			if (!$scope.editState.filled)
				$scope.editState.selectedColor = cell.color;
			$scope.editState.filled = false;
		}
	}
	keys = [110,112,94,43,45,42,47,37,33,62,60,115,100,114,105,73,111,79]
	$scope.keypress = function(ev){
		console.log(ev.keyCode, ev.keyIdentifier)
		idx = keys.indexOf(ev.keyCode)
		if(idx != -1)
			$scope.editState.selectedColor = $scope.rotate($scope.editState.selectedColor,Math.floor(idx / 3),idx%3)
		else if(ev.keyCode == 125){ // } key to cycle colors for whole program
			rotateColors($scope.program)
		}
	}
	$scope.rotate = function(src,hue,light){
		x="ABCDEFGHIJKLMNOPQRST"
		idx = x.indexOf(src)
		light = ((idx % 3) + light) % 3
		hue = (Math.floor(idx / 3) + hue) % 6
		console.log(hue,light)
		return x[hue*3 + light]
	}
	$scope.setColor = function(c){
		$scope.editState.selectedColor = c
	}
});

var mark = 0;
function floodFill(cell,program,set){
	mark++
	var count = 0
	var stack = [cell]
	var targetColor = cell.color;
	while(stack.length){
		var target = stack.pop()
		if(target.mark == mark || target.color != targetColor) continue;
		//not marked, color match
		target.mark = mark
		if(set) target.color = set
		count++
		if(target.x > 0) stack.push(program.rows[target.y].cells[target.x-1])
		if(target.x < program.w - 1) stack.push(program.rows[target.y].cells[target.x+1])
		if(target.y > 0) stack.push(program.rows[target.y-1].cells[target.x])
		if(target.y < program.h - 1) stack.push(program.rows[target.y+1].cells[target.x])
	}
	return count
}

function rotateColors(program){
	letters = "ABCDEFGHIJKLMNOPQR"
	for(var y=0; y<program.h; y++){
		for(var x = 0; x<program.w; x++){
			cell = program.rows[y].cells[x]
			idx = letters.indexOf(cell.color)
			if(idx != -1){
				cell.color = letters[(idx+1) % 18]
			}
		}
	}
}

function makeProgram(w,h,dat){
	var program = {rows:[],w:w,h:h}
	for(var y=0; y<h; y++){
		var row = {cells:[]}
		for(var x = 0; x<w; x++){
			c = "S"
			if(dat) c = dat[x+y*w]
			row.cells.push({color:c,x:x,y:y,mark:0})
		}
		program.rows.push(row)
	}
	return program
}

function makePalette(){
	return {
		'A':{color:'A'},
		'B':{color:'B'},
		'C':{color:'C'},
		'D':{color:'D'},
		'E':{color:'E'},
		'F':{color:'F'},
		'G':{color:'G'},
		'H':{color:'H'},
		'I':{color:'I'},
		'J':{color:'J'},
		'K':{color:'K'},
		'L':{color:'L'},
		'M':{color:'M'},
		'N':{color:'N'},
		'O':{color:'O'},
		'P':{color:'P'},
		'Q':{color:'Q'},
		'R':{color:'R'},
	}
}
