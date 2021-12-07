# frozen_string_literal: true

module Consolidations
  class GetCollectionService < ApplicationService
    def initialize(params)
      @params = params
    end

    def call
      ConsolidationSearcher.search(scope, @params)
    end

    private

    def scope
      Consolidation.includes(:category)
    end
  end
end
