$(function() {
  $('#cashes').on('cashes.changed', function() {
    $('.cash table.summarize').each(function(_, element) {
      const table = new Table(element);

      table.summarize();
    });
  });
});
