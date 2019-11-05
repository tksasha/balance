# frozen_string_literal: true

class AtEndService
  def initialize(params)
    @currency = CurrencyService.currency params[:currency]
  end

  def at_end
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

  class << self
    def at_end(*args)
      new(*args).at_end
    end
  end
end
