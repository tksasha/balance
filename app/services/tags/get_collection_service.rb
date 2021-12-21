# frozen_string_literal: true

module Tags
  class GetCollectionService < ApplicationService
    def initialize(category)
      @category = category
    end

    def call
      @category.tags.order(:name)
    end
  end
end
