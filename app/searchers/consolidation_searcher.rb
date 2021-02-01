# frozen_string_literal: true

class ConsolidationSearcher
  attr_reader :currency, :date

  def initialize(relation, params)
    @relation = relation

    self.currency = params[:currency]

    self.date = params
  end

  def search
    @relation
      .where(date: date_range, currency: currency)
      .select('SUM(sum) AS sum, category_id')
      .group(:category_id)
      .tap do |items|
        ConsolidationExpensesSum.sum = items.select { |item| item.income? == false }.sum(&:sum)
      end
  end

  private

  def currency=(object)
    @currency = CurrencyService.call(object)
  end

  def date=(object)
    @date = DateFactory.build(object)
  end

  def date_range
    @date_range ||= DateRange.month(date)
  end

  class << self
    def search(*args)
      new(*args).search
    end
  end
end
