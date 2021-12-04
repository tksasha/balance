# frozen_string_literal: true

module Backoffice
  class CashesController < ApplicationController
    private

    def collection
      @collection ||= ::Cashes::GetCollectionService.call(params)
    end

    def result
      @result ||= ::Cashes::GetResultService.call(action_name, params)
    end
  end
end
