# frozen_string_literal: true

module Categories
  class GetCollectionService < ApplicationService
    def initialize(params)
      @params = params
    end

    def call
      CategorySearcher.call(categories, @params)
    end

    private

    def categories
      Category.order(:income)
    end
  end
end
