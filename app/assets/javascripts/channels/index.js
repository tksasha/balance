//= require actioncable

//= require_self

//= require_tree .

(function(){
  this.App = {};

  App.cable = ActionCable.createConsumer();

  ActionCable.logger.enabled = false;
}).call(this);
