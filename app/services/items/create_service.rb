# frozen_string_literal: true

module Items
  class CreateService < ApplicationService
    include ActsAsUpdateAtEndViaWebsocketService
    include ActsAsUpdateBalanceViaWebsocketService

    def initialize(params)
      @params = params
    end

    def call
      return Failure.new(item) unless item.save

      update_at_end_via_websocket
        .update_balance_via_websocket

      Success.new(item)
    end

    private

    def resource_params
      @params.require(:item).permit(:date, :formula, :category_id, :description, :currency)
    end

    def item
      @item ||= Item.new(resource_params)
    end

    delegate :currency, to: :item
  end
end
