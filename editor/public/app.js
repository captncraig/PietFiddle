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
	
	var cellSize = 20;
	
	function init(){
		canvas.width = cellSize * W;
		canvas.height = cellSize * H;
		drawAll();
	}
	
	function drawAll(){
		var ctx = canvas.getContext('2d');
		for (var y = 0; y< H; y++){
			for(var x = 0; x < W; x++){
				var color = programText[y*W + x];
				var px = x * cellSize;
				var py = y *cellSize;
				ctx.fillStyle = pallette[color];
				ctx.fillRect(px, py, cellSize,cellSize)
				if(x == 0 || programText[y*W + x - 1] != color){
					line(px,py,px,py + cellSize,ctx);
				}
				if(x == W-1 || programText[y*W + x + 1] != color){
					line(px + cellSize,py,px+cellSize,py + cellSize,ctx);
				}
				if(y == 0 || programText[(y-1)*W + x] != color){
					line(px,py,px + cellSize,py,ctx);
				}
				if(y == H-1 || programText[(y+1)*W + x] != color){
					line(px,py+cellSize,px+cellSize,py + cellSize,ctx);
				}
			}
		}
	}
	function line(x0,y0,x1,y1,ctx){
		ctx.beginPath();
    	ctx.moveTo(x0,y0);
		ctx.lineTo(x1,y1);
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
	function enterCell(x,y){
		currentX = x;
		currentY = y;
		console.log(x,y);
	}
	init();
})