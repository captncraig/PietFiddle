$(function(){
	
	var canvas = $("#programCanvas")[0];
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
	
	var cellSize = 30;
	
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
				ctx.fillStyle = pallette[color];
				ctx.fillRect(x*cellSize, y*cellSize, cellSize,cellSize)
			}
		}
	}
	
	init();
})