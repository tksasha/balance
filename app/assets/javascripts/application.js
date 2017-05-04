// = require jquery
// = require jquery_ujs
// = require bootstrap/util
// = require bootstrap/modal
// = require bootstrap/dropdown
// = require bootstrap-datepicker/core
// = require bootstrap-datepicker/locales/bootstrap-datepicker.ua
// = require bootstrap-select.min
//
// = require_self

var BOOTSTRAP_DATEPICKER_DEFAULTS = { format: 'dd.mm.yyyy', autoclose: true, language: 'ua', todayHighlight: true };

$(function() {
  // Global AJAX-events
  $(document).ajaxStart(function() { $('#ajax-loader').show() });
  $(document).ajaxStop(function() { $('#ajax-loader').hide() });

  $('.bootstrap-datepicker').datepicker(BOOTSTRAP_DATEPICKER_DEFAULTS);

  //$('select').selectpicker();
});
