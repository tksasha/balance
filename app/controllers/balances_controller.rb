class BalancesController < ApplicationController
  private

  def resource
    @resource ||= BalanceService.new params
  end
end
