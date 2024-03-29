# frozen_string_literal: true

class ConsolidationSearcher
  attr_reader :currency, :month

  def initialize(relation, params)
    @relation = relation

    self.currency = params[:currency]

    self.month = params[:month]
  end

  def search
    @relation
      .where(date: dates, currency:)
      .select('SUM(sum) AS sum, category_id')
      .group(:category_id)
      .tap do |items|
        ConsolidationExpensesSum.sum = items.select { |item| item.income? == false }.sum(&:sum)
      end
  end

  private

  def currency=(object)
    @currency = Currency.parse(object)
  end

  def month=(object)
    @month = Month.parse(object)
  end

  def dates
    @dates ||= month.dates
  end

  class << self
    def search(*)
      new(*).search
    end
  end
end
