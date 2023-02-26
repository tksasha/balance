# frozen_string_literal: true

class ConsolidationsController < BaseController
  private

  def collection
    @collection ||= dashboard.consolidations
  end
end
