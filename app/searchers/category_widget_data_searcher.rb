# frozen_string_literal: true

class CategoryWidgetDataSearcher
  attr_reader :currency

  def initialize(params)
    self.currency = params[:currency]
  end

  def search
    Category
      .visible
      .where(currency:)
      .order(:income)
      .pluck(:name, :id, :income)
      .group_by { |array| array[2] }
      .map do |key, value|
        value.map(&:pop)

        [I18n.t(key, scope: :category_widget_data), value]
      end
  end

  private

  def currency=(object)
    @currency = Currency.parse(object)
  end

  class << self
    def search(*)
      new(*).search
    end
  end
end
