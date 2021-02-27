# frozen_string_literal: true

module ApplicationHelper
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
    @category_widget_data ||= CategoryWidgetDataSearcher.search params
  end

  def at_end
    @at_end ||= AtEndCalculatorService.calculate params
  end

  def balance
    @balance ||= BalanceCalculatorService.calculate params
  end

  def currency_from_params
    ParseCurrencyService.call(params[:currency])
  end

  def new_item_for_inline_form
    Item.new(currency: currency_from_params)
  end

  def current_year
    month.year
  end

  def current_month
    month.month
  end

  private

  def month
    @month ||= ParseMonthService.call(params)
  end
end
