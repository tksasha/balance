$(function() {
  $('#currencies-widget').popover({
    placement: 'bottom',
    html: true,
    content: function() {
      return $('#currencies-widget-content').html();
    }
  });
});
