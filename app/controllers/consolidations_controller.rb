# frozen_string_literal: true

class ConsolidationsController < ApplicationController
  private

  def collection
    @collection ||= ::Consolidations::GetCollectionService.call(params)
  end
end
