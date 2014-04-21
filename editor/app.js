function EditorCtrl($scope) {
	console.log("ABC")
	$scope.program = makeProgram(6,4,"ADGJMPBEHKNQCFILORSSSTTT") 
	$scope.settings = {cellSize:30}
	$scope.editState = {selectedColor:'Q',painting:false}
	console.log($scope.program)
	
	$scope.getCellText = function(cell){
		return $scope.settings.cellSize > 20? cell.color : ""
	}
	$scope.getCellStyle = function(cell){
		return {
			height: $scope.settings.cellSize + "px",
			width: $scope.settings.cellSize + "px",
			"line-height": $scope.settings.cellSize + "px",
			
		}
	}
	$scope.mouseDown = function(cell,ev){
		if(ev.which == 1){
			cell.color = $scope.editState.selectedColor
			$scope.editState.painting = true
		}
		else if(ev.which == 3){
			$scope.editState.selectedColor = cell.color
		}
	}
	$scope.mouseEnter = function(cell){
		if($scope.editState.painting){
			cell.color = $scope.editState.selectedColor
		}
	}
	$scope.mouseUp = function(cell){
		$scope.editState.painting = false;
	}
}

function makeProgram(w,h,dat){
	var program = {rows:[]}
	for(var y=0; y<h; y++){
		
		var row = {cells:[]}
		for(var x = 0; x<w; x++){
			row.cells.push({color:dat[x+y*w]})
		}
		program.rows.push(row)
	}
	return program
}
