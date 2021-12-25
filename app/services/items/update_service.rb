# frozen_string_literal: true

module Items
  class UpdateService < ApplicationService
    include ActsAsUpdateAtEndViaWebsocketService
    include ActsAsUpdateBalanceViaWebsocketService

    def initialize(params)
      @params = params
    end

    def call
      return Failure.new(item) unless item.update(resource_params)

      update_at_end_via_websocket
        .update_balance_via_websocket

      Success.new(item)
    end

    private

    def resource_params
      @params.require(:item).permit(:date, :formula, :category_id, :description, :currency, tag_ids: [])
    end

    def item
      @item ||= Item.find(@params[:id])
    end

    delegate :currency, to: :item
  end
end
