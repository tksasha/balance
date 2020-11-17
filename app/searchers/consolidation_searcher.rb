# frozen_string_literal: true

class ConsolidationSearcher
  def initialize(relation, params)
    @relation = relation

    date = DateFactory.build(params)

    @date_range = DateRange.month(date)

    @currency = CurrencyService.currency(params[:currency])
  end

  def search
    @relation
      .where(date: @date_range, currency: @currency)
      .select('SUM(sum) AS sum, category_id')
      .group(:category_id)
      .tap do |items|
        ConsolidationExpensesSum.sum = items.select { |item| item.income? == false }.sum(&:sum)
      end
  end

  class << self
    def search(*args)
      new(*args).search
    end
  end
end
