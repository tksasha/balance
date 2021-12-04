# frozen_string_literal: true

module Cashes
  class InitializeResourceService < ApplicationService
    def call
      Success.new(cash)
    end

    private

    def cash
      Cash.new
    end
  end
end
