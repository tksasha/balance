# frozen_string_literal: true

module Items
  class DestroyService < ApplicationService
    include ActsAsUpdateAtEndViaWebsocketService
    include ActsAsUpdateBalanceViaWebsocketService

    def initialize(params)
      @id = params[:id]
    end

    def call
      return Failure.new(item) unless item.destroy

      update_at_end_via_websocket
        .update_balance_via_websocket

      Success.new(item)
    end

    private

    def item
      @item ||= Item.find(@id)
    end

    delegate :currency, to: :item
  end
end
