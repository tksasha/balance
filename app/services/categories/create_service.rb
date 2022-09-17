# frozen_string_literal: true

module Categories
  class CreateService < ApplicationService
    def initialize(params)
      @params = params
    end

    def call
      return Success.new(category) if category.save

      Failure.new(category)
    end

    private

    def resource_params
      @params.require(:category).permit(:name, :supercategory, :income, :visible, :currency)
    end

    def category
      @category ||= Category.new(resource_params)
    end
  end
end
