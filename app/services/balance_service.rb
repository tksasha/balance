# frozen_string_literal: true

class BalanceService
  def initialize(params)
    @currency = CurrencyService.currency params[:currency]
  end

  def balance
    (sum - at_end).round(2)
  end

  private

  def at_end
    AtEndService.at_end currency: @currency
  end

  def sum
    Cash.where(currency: @currency).sum :sum
  end
end
