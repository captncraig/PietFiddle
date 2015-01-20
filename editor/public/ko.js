
function Row(data) {
    this.data = data
}

function PietViewModel() {
    this.rows = ko.observableArray([]);
	for(var i = 0; i<H; i++){
		var d = DATA.substring(i*W,W+i*W)
		this.rows.push(new Row(d))
	}
}

// Activates knockout.js
ko.applyBindings(new PietViewModel());
