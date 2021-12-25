# frozen_string_literal: true

class CalculateAtEndService < ApplicationService
  def initialize(currency)
    @currency = ParseCurrencyService.call(currency)
  end

  def call
    income - expense
  end

  private

  def search_by_currency
    Item.where(currency: @currency)
  end

  def income
    search_by_currency.income.sum(:sum)
  end

  def expense
    search_by_currency.expense.sum(:sum)
  end
end
