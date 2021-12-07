# frozen_string_literal: true

module Categories
  class GetResourceService < ApplicationService
    def initialize(params)
      @id = params[:id]
    end

    def call
      Success.new(category)
    end

    private

    def category
      @category ||= Category.find(@id)
    end
  end
end
