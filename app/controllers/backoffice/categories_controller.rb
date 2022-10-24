# frozen_string_literal: true

module Backoffice
  class CategoriesController < ApplicationController
    def create
      respond_to :js

      render :new unless resource.save
    end

    def update
      respond_to :js

      render :edit unless resource.update(resource_params)
    end

    private

    def scope
      Category.order(:name)
    end

    def collection
      @collection ||= ::CategorySearcher.search(scope, params)
    end

    def resource
      @resource ||= Category.find(params[:id])
    end

    def initialize_resource
      @resource = Category.new
    end

    def resource_params
      params
        .require(:category)
        .permit(:name, :supercategory, :income, :visible)
        .merge(currency: params[:currency])
    end

    def build_resource
      @resource = Category.new(resource_params)
    end
  end
end
