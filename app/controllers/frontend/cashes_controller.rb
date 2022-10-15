# frozen_string_literal: true

module Frontend
  class CashesController < ApplicationController
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
  end
end
