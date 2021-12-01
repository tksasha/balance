$(function() {
  const el = document.querySelector('#currencies-widget');

  bootstrap.Popover.getOrCreateInstance(el, {
    placement: 'bottom',
    html: true,
    content: function() {
      return $('#currencies-widget-content').html();
    }
  });
});
