// = require jquery
// = require jquery_ujs
// = require materialize
//
// = require_tree .
//
// = require_self

$(function() {
  // Global AJAX-events
  $(document).ajaxStart(function() { $('#ajax-loader').show() });
  $(document).ajaxStop(function() { $('#ajax-loader').hide() });

  $('.datepicker').datepicker(DatepickerDefaultOptions);

  $('select').formSelect(FormSelectDefaultOptions);

  $('body').on('created.category', CreateOrUpdateCategoryCallback);

  $('body').on('updated.category', CreateOrUpdateCategoryCallback);

  window.Modal = M.Modal.init(document.getElementById('modal'));

  $.getScript('/consolidations');
});
