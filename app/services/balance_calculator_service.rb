# frozen_string_literal: true

class BalanceCalculatorService
  def initialize(params)
    @currency = CurrencyService.currency params[:currency]
  end

  def calculate
    (sum - at_end).round(2)
  end

  private

  def at_end
    AtEndCalculatorService.calculate currency: @currency
  end

  def sum
    Cash.where(currency: @currency).sum :sum
  end

  class << self
    def calculate(*args)
      new(*args).calculate
    end
  end
end
