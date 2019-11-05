# frozen_string_literal: true

class BalancesController < ApplicationController
  private

  def resource
    @resource ||= BalanceService.new params
  end
end
