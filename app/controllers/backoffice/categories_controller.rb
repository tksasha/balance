# frozen_string_literal: true

module Backoffice
  class CategoriesController < ApplicationController
    private

    def collection
      @collection ||= ::Categories::GetCollectionService.call(params)
    end

    def result
      @result ||= ::Categories::GetResultService.call(action_name, params)
    end
  end
end
