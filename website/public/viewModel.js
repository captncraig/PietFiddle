$(function(){
	function PietViewModel() {
    	this.size = ko.observable(0);
		this.newW = ko.observable(W);
		this.newH = ko.observable(H);
		this.dirty = ko.observable(false);
		this.id = ko.observable("");
	}
	window.ViewModel = new PietViewModel();
	ko.applyBindings(ViewModel)
})


