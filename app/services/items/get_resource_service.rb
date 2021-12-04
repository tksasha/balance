# frozen_string_literal: true

module Items
  class GetResourceService < ApplicationService
    def initialize(params)
      @id = params[:id]
    end

    def call
      Success.new(item)
    end

    private

    def item
      @item ||= Item.find(@id)
    end
  end
end
