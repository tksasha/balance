# frozen_string_literal: true

class CalculateBalanceService
  attr_reader :currency

  def initialize(params)
    self.currency = params[:currency]
  end

  def calculate
    (sum - at_end).round(2)
  end

  private

  def currency=(object)
    @currency = ParseCurrencyService.call(object)
  end

  def at_end
    CalculateAtEndService.calculate currency: @currency
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
