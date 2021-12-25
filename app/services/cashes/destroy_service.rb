# frozen_string_literal: true

module Cashes
  class DestroyService < ApplicationService
    include ActsAsUpdateBalanceViaWebsocketService

    def initialize(params)
      @id = params[:id]
    end

    def call
      return Failure.new(cash) unless cash.destroy

      update_balance_via_websocket

      Success.new(cash)
    end

    private

    def cash
      @cash ||= Cash.find(@id)
    end

    delegate :currency, to: :cash
  end
end
