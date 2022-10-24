# frozen_string_literal: true

module Backoffice
  class CashesController < ApplicationController
    def create
      respond_to :js

      render :new unless resource.save
    end

    def update
      respond_to :js

      render :edit unless resource.update(resource_params)
    end

    def destroy
      respond_to :js

      resource.destroy
    end

    private

    def scope
      Cash.order(:name)
    end

    def collection
      @collection ||= ::CashSearcher.search(scope, params)
    end

    def resource
      @resource ||= Cash.find(params[:id])
    end

    def initialize_resource
      @resource = Cash.new
    end

    def resource_params
      params
        .require(:cash)
        .permit(:name, :formula, :supercategory, :favorite)
        .merge(currency: params[:currency])
    end

    def build_resource
      @resource = Cash.new(resource_params)
    end
  end
end
