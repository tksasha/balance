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
      @collection ||= Categories::GetCollectionService.call(params)
    end

    def resource_params
      params.require(:category).permit(:name, :income, :visible, :currency)
    end

    def resource
      @resource ||= Category.find params[:id]
    end

    def initialize_resource
      @resource = Category.new currency: params[:currency]
    end

    def build_resource
      @resource = Category.new resource_params
    end
  end
end
