# frozen_string_literal: true

module Backoffice
  class CashesController < ApplicationController
    delegate :destroy, to: :resource

    def create
      render :new, status: :unprocessable_entity unless resource.save
    end

    def update
      render :edit, status: :unprocessable_entity unless resource.update(resource_params)
    end

    private

    def collection
      @collection ||= ::Cashes::GetCollectionService.call(params)
    end

    def resource
      @resource ||= Cash.find(params[:id])
    end

    def initialize_resource
      @resource = Cash.new(currency: params[:currency])
    end

    def resource_params
      params.require(:cash).permit(:formula, :name, :currency)
    end

    def build_resource
      @resource = Cash.new(resource_params)
    end
  end
end
