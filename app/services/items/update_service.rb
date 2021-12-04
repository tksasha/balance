# frozen_string_literal: true

module Items
  class UpdateService < ApplicationService
    def initialize(params)
      @params = params
    end

    def call
      return Success.new(item) if item.update(resource_params)

      Failure.new(item)
    end

    private

    def resource_params
      @params.require(:item).permit(:date, :formula, :category_id, :description, :currency)
    end

    def item
      @item ||= Item.find(@params[:id])
    end
  end
end
