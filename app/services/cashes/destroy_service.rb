# frozen_string_literal: true

module Cashes
  class DestroyService < ApplicationService
    def initialize(params)
      @id = params[:id]
    end

    def call
      return Success.new(cash) if cash.destroy

      Failure.new(cash)
    end

    private

    def cash
      @cash ||= Cash.find(@id)
    end
  end
end
