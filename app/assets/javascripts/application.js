// = require jquery
// = require jquery_ujs
// = require popper
// = require bootstrap-sprockets
//
// = require bootstrap-select
// = require bootstrap-datepicker
// = require bootstrap-datepicker.uk.min
//
// = require_tree .
//
// = require_self

var BOOTSTRAP_DATEPICKER_DEFAULTS = {
  format: 'dd.mm.yyyy', autoclose: true, language: 'uk', todayHighlight: true
};

$(function() {
  $(document).ajaxStart(function() {
    $('#ajax-loader').show()
  });

  $(document).ajaxStop(function() {
    $('#ajax-loader').hide()
  });

  $('.datepicker').datepicker(BOOTSTRAP_DATEPICKER_DEFAULTS);

  $('select').selectpicker();

  $('body').on('created.category', CreateOrUpdateCategoryCallback);

  $('body').on('updated.category', CreateOrUpdateCategoryCallback);
});
