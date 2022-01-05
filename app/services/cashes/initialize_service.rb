# frozen_string_literal: true

module Cashes
  class InitializeService < ApplicationService
    def initialize(params)
      @currency = ParseCurrencyService.call(params.fetch(:currency, nil))
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
