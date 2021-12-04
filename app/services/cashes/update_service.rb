# frozen_string_literal: true

module Cashes
  class UpdateService < ApplicationService
    def initialize(params)
      @params = params
    end

    def call
      return Success.new(cash) if cash.update(resource_params)

      Failure.new(cash)
    end

    private

    def resource_params
      @params.require(:cash).permit(:formula, :name, :currency)
    end

    def cash
      @cash ||= Cash.find(@params[:id])
    end
  end
end
