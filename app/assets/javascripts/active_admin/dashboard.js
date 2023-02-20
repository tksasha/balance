$(function() {
  if((location.pathname != '/backoffice') && (location.pathname != '/backoffice/dashboard')) { return false }

  $('table').each(function(_, element) {
    const table = new Table(element);

    table.summarize();
  });
});
