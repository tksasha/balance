# frozen_string_literal: true

module Items
  class CreateService < ApplicationService
    def initialize(params)
      @params = params
    end

    def call
      return Success.new(item) if item.save

      Failure.new(item)
    end

    private

    def resource_params
      @params.require(:item).permit(:date, :formula, :category_id, :description, :currency)
    end

    def item
      @item ||= Item.new(resource_params)
    end
  end
end
