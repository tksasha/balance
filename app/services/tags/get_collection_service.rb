# frozen_string_literal: true

module Tags
  class GetCollectionService < ApplicationService
    def initialize(params)
      @category_id = params[:category_id]
    end

    def call
      category.tags.order(:name)
    end

    private

    def category
      @category ||= Category.find(@category_id)
    end
  end
end
