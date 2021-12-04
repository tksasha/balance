# frozen_string_literal: true

module Cashes
  class GetResourceService < ApplicationService
    def initialize(params)
      @cash_id = params[:id]
    end

    def call
      Success.new(cash)
    end

    private

    def cash
      Cash.find(@cash_id)
    end
  end
end
