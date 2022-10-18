# frozen_string_literal: true

module Cashes
  class CreateService < ApplicationService
    include ActsAsUpdateBalanceViaWebsocketService

    def initialize(params)
      @params = params
    end

    def call
      return Failure.new(cash) unless cash.save

      update_balance_via_websocket

      Success.new(cash)
    end

    private

    def resource_params
      @params.require(:cash).permit(:name, :formula, :currency, :supercategory, :favorite)
    end

    def cash
      @cash ||= Cash.new(resource_params)
    end

    delegate :currency, to: :cash
  end
end
