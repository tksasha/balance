$(function() {
  if((location.pathname != '/admin') && (location.pathname != '/admin/dashboard')) { return false }

  $('table').each(function(_, element) {
    const table = new Table(element);

    table.summarize();
  });
});
