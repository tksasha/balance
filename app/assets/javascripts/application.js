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

  $('body').on('created.category', CreateOrUpdateCategoryCallback);

  $('body').on('updated.category', CreateOrUpdateCategoryCallback);
});
