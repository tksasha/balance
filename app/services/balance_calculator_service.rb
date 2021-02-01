# frozen_string_literal: true

class BalanceCalculatorService
  attr_reader :currency

  def initialize(params)
    self.currency = params[:currency]
  end

  def calculate
    (sum - at_end).round(2)
  end

  private

  def currency=(object)
    @currency = CurrencyService.call(object)
  end

  def at_end
    AtEndCalculatorService.calculate currency: @currency
  end

  def sum
    Cash.where(currency: @currency).sum(:sum)
  end

  class << self
    def calculate(*args)
      new(*args).calculate
    end
  end
end
