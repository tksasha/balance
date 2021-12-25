# frozen_string_literal: true

module BalanceHelper
  def balance
    CalculateBalanceService.call(params[:currency])
  end
end
