# frozen_string_literal: true

module Backoffice
  class TagsController < ApplicationController
    private

    def category
      @category ||= Category.find(params[:category_id])
    end

    helper_method :category

    def collection
      @collection ||= ::Tags::GetCollectionService.call(category)
    end

    def result
      @result ||= ::Tags::GetResultService.call(action_name, category, params)
    end
  end
end
