$(function() {
  $('#cashes').on('cashes.changed', function() {
    $('.cash table.summarize').each(function(_, element) {
      const table = new Table(element);

      table.summarize();

      if($('#cashes [class|="col"]').size() == 2) {
        const col = $('#cashes [class|="col"]:last');

        if(col.find('.summary .sum').text().length > 9) {
          col
            .removeClass('col-3')
            .addClass('col-4');
        }
      }
    });
  });
});
