# frozen_string_literal: true

class ConsolidationsController < ApplicationController
  private

  def relation
    Consolidation.includes(:category)
  end

  def collection
    @collection ||= ConsolidationSearcher.search(relation, params)
  end
end
