# frozen_string_literal: true

module Websockets
  class UpdateBalanceService < ApplicationService
    def initialize(currency)
      @currency = currency
    end

    def call
      ActionCable.server.broadcast('NotificationsChannel', { type: :balance, value: balance })
    end

    private

    def balance
      balance = CalculateBalanceService.call(@currency)

      MoneyDecorator.new(balance).to_s
    end
  end
end
