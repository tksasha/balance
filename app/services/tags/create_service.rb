# frozen_string_literal: true

module Tags
  class CreateService < ApplicationService
    def initialize(category, params)
      @category = category

      @params = params
    end

    def call
      return Success.new(tag) if tag.save

      Failure.new(tag)
    end

    private

    def resource_params
      @params.require(:tag).permit(:name)
    end

    def tag
      @tag ||= @category.tags.new(resource_params)
    end
  end
end
