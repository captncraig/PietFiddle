$(function(){
	function PietViewModel() {
		var self = this;
		this.uid = ko.observable(UID);
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
		this.loggedIn = ko.computed(function(){
			return self.uid() !== "";
		});
		this.loggedOut = ko.computed(function(){
			return self.uid() === "";
		});
		this.loginUn = ko.observable('');
		this.loginPw = ko.observable('');
		this.signupUn = ko.observable('');
		this.signupPw = ko.observable('');
	}
	window.ViewModel = new PietViewModel();
	ko.applyBindings(ViewModel)
	
	
})


