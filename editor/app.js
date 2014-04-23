angular.module('piet',[])

angular.module('piet')
	.controller('EditorCtrl', function EditorCtrl($scope) {
	$scope.palette = makePalette()
	$scope.program = makeProgram(30,29,"BBBBBBBBBBBCQQQQQQQRNNNNNNNNNNBBBBBBBBKKKQQQQQQQQRNNNNNNNNNNBBBBBBBBKKQQQQQQQQQRNNNNNNNNNNBBBBBBBKKQQQQQQQQQQNNNNNNNNNNNBBBBBBBKKQQQQQQQQQQNNNNNNNNNNNBBBBBBQQQQQQQQQQQQQNNNNNNNNNNNBBBBTBBQQQQQQQQQQQQNNNNNNNNNNNBBBTBTBBQQQQQQQQQQQNNNNNNNNNNNBBTBBBTBQQQQQQQQQQQNNNNNNNNNNNBBBTFTBBBQQQQQQQQQQNNNNNNNNNNNEEEEEBBBBBHHHHHHHHHTMMMMMMMMMNEEEEEHHHHHHHHHHHHHHHHTMMMMMMOOEEEEHHHHHHHHHHHHHHHHMMMMMMMMIIEEEIHHHHHHHHHHHHHHHHMMMMMMMMCMEEEEEEEETHHHHHHHHHHHMMMMMMMMQMEEEEEEEEMLHHHHHHHHHHMMMMMMMMMMOOOOOOOOOHHHHHHHHHHHMMMMMMMMMMOOOOOOOOOHHHHHHHHHHHTMMMMMMMMMOOOOOOOOOTTAAAAAAAATAMMMMMMMMMOOOOOOOOOBAAAAAAAAAAAMMMMMMMMMOOOOOOOOOBAAAAAAAAAAAMMMMMMMMMOOOOOOOOOBAAAAAAAAAAAMMMMMMMMMOOOOOOOOOBAAAAAAAAAAAEEEEMMMMMOOOOOOOOOBAAAAAAAAAAAEEEENNNNNOOOOOOOOOBAAAAAAAAAAAEEEEHHHHHOOOOOOOOOBAAAAAAAAAAAEDDDDDDDDOOOOOOOOPBAAAAAAAAAAAEDDDDDDDDOOOOOOOOPBAAAAAAAAAAAEDDDDDDDDOOOOOOOOPBAAAAAAAAAAAEDDDDDDDD") 
	$scope.settings = {}
	$scope.editState = {selectedColor:'Q',painting:false}
	$scope.hover = {size:0}
	console.log($scope.program)
	
	$scope.getCellText = function(cell){
		return ""
	}
	$scope.getCellStyle = function(cell){
		obj = {"box-sizing":"border-box"}
		style = "1px solid black"
		if(cell.x == 0 || $scope.program.rows[cell.y].cells[cell.x-1].color != cell.color){
			obj["border-left"] = style
		}
		if(cell.x >= ($scope.program.w - 1) || $scope.program.rows[cell.y].cells[cell.x+1].color != cell.color){
			obj["border-right"] = style
		}
		if(cell.y == 0 || $scope.program.rows[cell.y-1].cells[cell.x].color != cell.color){
			obj["border-top"] = style
		}
		if(cell.y >= ($scope.program.h - 1) || $scope.program.rows[cell.y+1].cells[cell.x].color != cell.color){
			obj["border-bottom"] = style
		}
		return obj
		
	}
	$scope.mouseDown = function(cell,ev){
		if(ev.which == 1){
			cell.color = $scope.editState.selectedColor
			$scope.editState.painting = true
			$scope.hover.size = floodFill(cell,$scope.program)
		}
		else if(ev.which == 3){
			$scope.editState.selectedColor = cell.color
		}
	}
	$scope.mouseEnter = function(cell){
		if($scope.editState.painting){
			cell.color = $scope.editState.selectedColor
		}
		$scope.hover.size = floodFill(cell,$scope.program)
	}
	$scope.mouseUp = function(cell){
		$scope.editState.painting = false;
	}

	$scope.setColor = function(c){
		$scope.editState.selectedColor = c
	}
});

var mark = 0;
function floodFill(cell,program){
	mark++
	var count = 0
	var stack = [cell]
	while(stack.length){
		var target = stack.pop()
		if(target.mark == mark || target.color != cell.color) continue;
		//not marked, color match
		target.mark = mark
		count++
		if(target.x > 0) stack.push(program.rows[target.y].cells[target.x-1])
		if(target.x < program.w - 1) stack.push(program.rows[target.y].cells[target.x+1])
		if(target.y > 0) stack.push(program.rows[target.y-1].cells[target.x])
		if(target.y < program.h - 1) stack.push(program.rows[target.y+1].cells[target.x])
	}
	return count
}

function makeProgram(w,h,dat){
	var program = {rows:[],w:w,h:h}
	for(var y=0; y<h; y++){
		
		var row = {cells:[]}
		for(var x = 0; x<w; x++){
			row.cells.push({color:dat[x+y*w],x:x,y:y,mark:0})
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
