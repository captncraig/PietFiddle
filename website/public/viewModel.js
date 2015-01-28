$(function(){
	function PietViewModel() {
    	this.size = ko.observable(0);
	}
	window.ViewModel = new PietViewModel();
	ko.applyBindings(ViewModel)
})


