# frozen_string_literal: true

class CalculateBalanceService < ApplicationService
  attr_reader :currency

  def initialize(params)
    self.currency = params[:currency]
  end

  def call
    (sum - at_end).round(2)
  end

  private

  def currency=(object)
    @currency = ParseCurrencyService.call(object)
  end

  def at_end
    CalculateAtEndService.call(currency: @currency)
  end

  def sum
    Cash.where(currency: @currency).sum(:sum)
  end
end
