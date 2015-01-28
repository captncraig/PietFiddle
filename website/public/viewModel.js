$(function(){
	function PietViewModel() {
		var self = this;
    	this.size = ko.observable(0);
		this.newW = ko.observable(W);
		this.newH = ko.observable(H);
		this.dirty = ko.observable(false);
		this.id = ko.observable(ID);
		this.link = ko.computed(function() {
        	return "/img/"+self.id()+".png?cs=15";
    	});
		this.tinyLink = ko.computed(function() {
        	return "/img/"+self.id()+".png?cs=1";
    	});
	}
	window.ViewModel = new PietViewModel();
	ko.applyBindings(ViewModel)
	
	
})


