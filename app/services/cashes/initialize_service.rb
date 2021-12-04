# frozen_string_literal: true

module Cashes
  class InitializeService < ApplicationService
    def call
      Success.new(cash)
    end

    private

    def cash
      Cash.new
    end
  end
end
