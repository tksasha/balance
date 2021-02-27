# frozen_string_literal: true

class CalculateAtEndService < ApplicationService
  attr_reader :currency

  def initialize(params)
    self.currency = params[:currency]
  end

  def call
    income - expense
  end

  private

  def currency=(object)
    @currency = ParseCurrencyService.call(object)
  end

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
