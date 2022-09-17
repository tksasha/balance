# frozen_string_literal: true

module Categories
  class UpdateService < ApplicationService
    def initialize(params)
      @params = params
    end

    def call
      return Success.new(category) if category.update(resource_params)

      Failure.new(category)
    end

    private

    def resource_params
      @params.require(:category).permit(:name, :supercategory, :income, :visible, :currency)
    end

    def category
      @category ||= Category.find(@params[:id])
    end
  end
end
