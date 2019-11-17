# frozen_string_literal: true

module ApplicationHelper
  def current_date
    @current_date ||= DateFactory.build params
  end

  def months
    %w[Січень Лютий Березень Квітень Травень Червень Липень Серпень Вересень Жовтень Листопад Грудень]
  end

  def money(sum)
    number_with_delimiter format('%.2f', sum)
  end

  def decorated
    resource.decorate
  end

  def breadcrumbs
    content_tag :div, class: :breadcrumb do
      if block_given?
        concat link_to('Backoffice', :backoffice, class: 'breadcrumb-item', data: { remote: true })

        yield
      else
        concat content_tag(:span, 'Backoffice', class: 'breadcrumb-item active')
      end
    end
  end

  # TODO: spec me
  def category_widget_data
    @category_widget_data ||= CategoryWidgetDataSearcher.search params
  end
end
