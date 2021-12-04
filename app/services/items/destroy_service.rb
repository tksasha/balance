# frozen_string_literal: true

module Items
  class DestroyService < ApplicationService
    def initialize(params)
      @id = params[:id]
    end

    def call
      return Success.new(item) if item.destroy

      Failure.new(item)
    end

    private

    def item
      @item ||= Item.find(@id)
    end
  end
end
