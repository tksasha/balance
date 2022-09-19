# frozen_string_literal: true

module Cashes
  class InitializeService < ApplicationService
    def initialize(params)
      @currency = Currency.parse(params[:currency])
    end

    def call
      Success.new(cash)
    end

    private

    def cash
      Cash.new(currency: @currency)
    end
  end
end
