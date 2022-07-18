# frozen_string_literal: true

module Categories
  class InitializeService < ApplicationService
    def initialize(params)
      @currency = params[:currency]
    end

    def call
      Success.new(category)
    end

    private

    def category
      Category.new(currency: @currency)
    end
  end
end
