# frozen_string_literal: true

module Items
  class GetCollectionService < ApplicationService
    def initialize(params)
      @params = params
    end

    def call
      ItemSearcher.call(items, @params)
    end

    private

    def items
      Item.order(date: :desc).includes(:category, :tags)
    end
  end
end
