# frozen_string_literal: true

module Cashes
  class GetResourceService < ApplicationService
    def initialize(params)
      @cash_id = params[:id]
    end

    def call
      Cash.find(@cash_id)
    end
  end
end
