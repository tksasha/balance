// = require jquery
// = require jquery_ujs
// = require popper
// = require bootstrap-sprockets
//
// vendor/assets/javascripts
// = require bootstrap-datepicker
// = require bootstrap-datepicker.uk.min
//
// app/assets/javascripts
// = require bootstrap/modal
// = require currencies-widget
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
});
