# frozen_string_literal: true

module ApplicationHelper
  def months
    %w[Січень Лютий Березень Квітень Травень Червень Липень Серпень Вересень Жовтень Листопад Грудень]
  end

  def decorated
    resource.decorate
  end

  def breadcrumbs
    tag.div(class: :breadcrumb) do
      if block_given?
        concat link_to('Backoffice', :backoffice, class: 'breadcrumb-item', data: { remote: true })

        yield
      else
        concat tag.span('Backoffice', class: 'breadcrumb-item active')
      end
    end
  end

  def category_widget_data
    @category_widget_data ||= CategoryWidgetDataSearcher.search(params)
  end

  def currency_from_params
    ParseCurrencyService.call(params[:currency])
  end

  def new_item_for_inline_form
    Item.new(currency: currency_from_params)
  end

  def month
    @month ||= ParseMonthService.call(params)
  end

  def current_year
    month.year
  end

  def current_month
    month.month
  end

  # rubocop:disable Metrics/AbcSize
  # rubocop:disable Metrics/MethodLength
  # rubocop:disable Naming/MethodParameterName
  def input(f, attribute, options = {})
    input_html_class = ['form-control']

    input_html_class.push(options.fetch(:input_html, {}).fetch(:class, ''))

    if f.object.errors[attribute].present?
      input_html_class.push('is-invalid')

      concat f.label(attribute, class: 'text-danger')

      concat f.text_field(attribute, class: input_html_class.join(' '))

      content_tag :div, class: 'invalid-feedback' do
        f.object.errors[attribute].each do |error|
          concat content_tag :div, error
        end
      end
    else
      concat f.label(attribute)

      f.text_field(attribute, class: input_html_class.join(' '))
    end
  end
  # rubocop:enable Metrics/AbcSize
  # rubocop:enable Metrics/MethodLength
  # rubocop:enable Naming/MethodParameterName

  def serialized_collection
    collection.map { |item| serializer(item) }
  end
end
