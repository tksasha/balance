# frozen_string_literal: true

module Cashes
  class GetCollectionService < ApplicationService
    def initialize(params)
      @params = params
    end

    def call
      CashSearcher.call(cashes, @params)
    end

    private

    def cashes
      Cash.order(:name)
    end
  end
end
