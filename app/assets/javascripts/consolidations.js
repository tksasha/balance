$(function() {
  $('#consolidations').on('consolidations.changed', function() {
    $('#consolidations table').each(function(_, element) {
      const table = new Table(element);

      table.summarize();
    });
  });
});
