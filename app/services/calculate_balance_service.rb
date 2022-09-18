# frozen_string_literal: true

class CalculateBalanceService < ApplicationService
  def initialize(currency)
    @currency = Currency(currency)
  end

  def call
    (sum - at_end).round(2)
  end

  private

  def at_end
    CalculateAtEndService.call(@currency)
  end

  def sum
    Cash.where(currency: @currency).sum(:sum)
  end
end
