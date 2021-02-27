# frozen_string_literal: true

class CalculateAtEndService
  attr_reader :currency

  def initialize(params)
    self.currency = params[:currency]
  end

  def calculate
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

  class << self
    def calculate(*args)
      new(*args).calculate
    end
  end
end
