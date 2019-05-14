# Use this setup block to configure all options available in SimpleForm.
SimpleForm.setup do |config|
  config.error_notification_class = 'alert alert-danger'

  config.button_class = 'btn btn-default'

  config.boolean_label_class = nil

  config.wrappers :vertical_form, tag: 'div', class: 'form-group', error_class: 'has-error' do |b|
    b.use :html5

    b.use :placeholder

    b.optional :maxlength

    b.optional :pattern

    b.optional :min_max

    b.optional :readonly

    b.use :label

    b.use :input, class: 'form-control'

    b.use :error, wrap_with: { tag: 'span', class: 'error' }

    b.use :hint,  wrap_with: { tag: 'p', class: 'help-block' }
  end

  config.default_wrapper = :vertical_form

  config.wrapper_mappings = {
    check_boxes: :vertical_radio_and_checkboxes,
    radio_buttons: :vertical_radio_and_checkboxes,
    file: :vertical_file_input,
    boolean: :vertical_boolean,
    datetime: :multi_select,
    date: :multi_select,
    time: :multi_select
  }

  config.label_text = lambda { |label, required, explicit_label| label }
end
