# frozen_string_literal: true

module Cashes
  class UpdateService < ApplicationService
    include ActsAsUpdateBalanceViaWebsocketService

    def initialize(params)
      @params = params
    end

    def call
      return Failure.new(cash) unless cash.update(resource_params)

      update_balance_via_websocket

      Success.new(cash)
    end

    private

    def resource_params
      @params.require(:cash).permit(:formula, :name, :currency, :supercategory, :favorite)
    end

    def cash
      @cash ||= Cash.find(@params[:id])
    end

    delegate :currency, to: :cash
  end
end
