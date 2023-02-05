# frozen_string_literal: true

class ConsolidationsController < BaseController
  private

  def collection
    @collection ||= ::Consolidations::GetCollectionService.call(params)
  end
end
