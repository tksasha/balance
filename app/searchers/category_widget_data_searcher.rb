# frozen_string_literal: true

class CategoryWidgetDataSearcher
  def initialize(params)
    @currency = params[:currency]
  end

  def search
    return [] unless CURRENCIES.include?(@currency)

    Category
      .where(currency: @currency)
      .order(:income)
      .pluck(:name, :id, :income)
      .group_by { |array| array[2] }
      .map do |key, value|
        value.map(&:pop)

        [I18n.t(key, scope: :category_widget_data), value]
      end
  end

  class << self
    def search(*args)
      new(*args).search
    end
  end
end
