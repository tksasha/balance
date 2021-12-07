# frozen_string_literal: true

module Backoffice
  class CategoriesController < ApplicationController
    def create
      render :new, status: :unprocessable_entity unless resource.save
    end

    def update
      render :edit, status: :unprocessable_entity unless resource.update resource_params
    end

    private

    def collection
      @collection ||= ::Categories::GetCollectionService.call(params)
    end

    def result
      @result ||= ::Categories::GetResultService.call(action_name, params)
    end

    def resource_params
      params.require(:category).permit(:name, :income, :visible, :currency)
    end
  end
end
