# frozen_string_literal: true

module Cashes
  class CreateService < ApplicationService
    def initialize(params)
      @params = params
    end

    def call
      return Success.new(cash) if cash.save

      Failure.new(cash)
    end

    private

    def resource_params
      @params.require(:cash).permit(:name, :formula, :currency)
    end

    def cash
      @cash ||= Cash.new(resource_params)
    end
  end
end
