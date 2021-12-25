# frozen_string_literal: true

module Websockets
  class UpdateAtEndService < ApplicationService
    def initialize(currency)
      @currency = currency
    end

    def call
      ActionCable.server.broadcast('NotificationsChannel', { type: :at_end, value: at_end })
    end

    private

    def at_end
      at_end = CalculateAtEndService.call(@currency)

      MoneyDecorator.new(at_end).to_s
    end
  end
end
