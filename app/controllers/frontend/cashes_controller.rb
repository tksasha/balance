# frozen_string_literal: true

module Frontend
  class CashesController < ApplicationController
    def update
      render :edit unless resource.update(resource_params)
    end

    private

    def cashes
      Cash.order(:name)
    end

    def collection
      @collection ||= ::CashSearcher.search(cashes, params)
    end

    def resource
      @resource ||= Cash.find(params[:id])
    end

    def resource_params
      params.require(:cash).permit(:name, :formula)
    end

    def dashboard
      ::Frontend::Dashboard.new(params)
    end

    helper_method :dashboard
  end
end