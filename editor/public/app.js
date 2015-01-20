String.prototype.replaceAt=function(index, character) {
    return this.substr(0, index) + character + this.substr(index+character.length);
}
$(function(){
	var jCanvas = $("#programCanvas");
	var canvas = jCanvas[0];
	var programText = DATA;
	while(programText.length < W*H){
		programText += "A";
	}
	var pallette = [];
	pallette['A'] = "#FFC0C0";
	pallette['B'] = "#FF0000";
	pallette['C'] = "#C00000";
	pallette['D'] = "#FFFFC0";
	pallette['E'] = "#FFFF00";
	pallette['F'] = "#C0C000";
	pallette['G'] = "#C0FFC0";
	pallette['H'] = "#00FF00";
	pallette['I'] = "#00C000";
	pallette['J'] = "#C0FFFF";
	pallette['K'] = "#00FFFF";
	pallette['L'] = "#00C0C0";
	pallette['M'] = "#C0C0FF";
	pallette['N'] = "#0000FF";
	pallette['O'] = "#0000C0";
	pallette['P'] = "#FFC0FF";
	pallette['Q'] = "#FF00FF";
	pallette['R'] = "#C000C0";
	pallette['S'] = "#FFFFFF";
	pallette['T'] = "#000000";
	
	var cellSize = 25;
	
	function init(){
		canvas.width = cellSize * W;
		canvas.height = cellSize * H;
		drawAll();
	}
	var ctx = canvas.getContext('2d');
	function drawAll(){
		for (var y = 0; y< H; y++){
			for(var x = 0; x < W; x++){
				drawCell(x,y,false);
			}
		}
	}
	function drawCell(x,y,updateNeighbor){
		var color = programText[y*W + x];
		var px = x * cellSize;
		var py = y *cellSize;
		var tl = {x:x*cellSize,y:y*cellSize};
		var bl = {x:x*cellSize,y:y*cellSize+cellSize};
		var tr = {x:x*cellSize+cellSize,y:y*cellSize};
		var br = {x:x*cellSize+cellSize,y:y*cellSize+cellSize};
		ctx.fillStyle = pallette[color];
		ctx.fillRect(px, py, cellSize,cellSize)
		if(x == 0 || programText[y*W + x - 1] != color){
			line(tl,bl);
		}
		if(x == W-1 || programText[y*W + x + 1] != color){
			line(tr,br);
		}
		if(y == 0 || programText[(y-1)*W + x] != color){
			line(tl,tr);
		}
		if(y == H-1 || programText[(y+1)*W + x] != color){
			line(bl,br);
		}
		if(updateNeighbor){
			if(x != 0)drawCell(x-1,y,false);
			if(x != W-1)drawCell(x+1,y,false);
			if(y != 0)drawCell(x,y-1,false);
			if(y != H-1)drawCell(x,y+1,false);
		}
	}
	function line(a,b){
		ctx.beginPath();
		ctx.lineWidth = 2;
    	ctx.moveTo(a.x,a.y);
		ctx.lineTo(b.x,b.y);
    	ctx.stroke();
	}
	var currentX = -1,currentY = -1;
	jCanvas.bind('mousemove',function(ev){
		var x = Math.floor(ev.offsetX / cellSize);
		var y = Math.floor(ev.offsetY / cellSize);
		if(x != currentX || y != currentY){
			enterCell(x,y)
		}
	});
	jCanvas.bind('mouseleave',function(){
		currentX = currentY = -1;
	});
	jCanvas.bind('mouseenter',function(ev){
		var x = Math.floor(ev.offsetX / cellSize);
		var y = Math.floor(ev.offsetY / cellSize);
		enterCell(x,y);
	});
	jCanvas.bind('mousedown', function(ev){
		var x = Math.floor(ev.offsetX / cellSize);
		var y = Math.floor(ev.offsetY / cellSize);
		var isRight = ev.button == 2;
		mousedown(x,y,isRight);
	});
	jCanvas.bind('mouseup', function(ev){
		var x = Math.floor(ev.offsetX / cellSize);
		var y = Math.floor(ev.offsetY / cellSize);
		var isRight = ev.button == 2;
		mouseup(x,y,isRight);
	});
	jCanvas.bind('contextmenu', function(){return false;}); 
	
	var editState = {selectedColor:'A',painting:false,rightDown:false,filled:false}
	
	function enterCell(x,y){
		currentX = x;
		currentY = y;
		if(editState.painting){
			setCell(x,y);
		}
	}
	function setCell(x,y){
		programText = programText.replaceAt(y*W + x, editState.selectedColor);
		drawCell(x,y,true);
	}
	function mousedown(x,y,isRight){
		if(!isRight){
			if(editState.rightDown){
				editState.filled = true;
			}
			else{
				setCell(x,y);
				editState.painting = true
			}
		}
		else if(isRight){
			editState.filled = false;
			editState.rightDown = true;
		}
	}
	function mouseup(x,y,isRight){
		editState.painting = false;
		console.log("UP",x,y,isRight)
		if(isRight){
			editState.rightDown = false;
			if (!editState.filled)
				editState.selectedColor = programText[y*W + x];
			editState.filled = false;
		}
	}
	init();
})