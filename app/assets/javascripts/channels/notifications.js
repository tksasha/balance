(function(){
  App.notifications = App.cable.subscriptions.create('NotificationsChannel', {
    received: function(data) {
      switch(data.type) {
        case 'at_end':
          at_end.Update(data.value);
          break;
        case 'balance':
          balance.Update(data.value);
          break;
      };
    }
  });
}).call(this);
