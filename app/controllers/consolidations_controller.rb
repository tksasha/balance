# frozen_string_literal: true

class ConsolidationsController < ApplicationController
  private

  def date
    DateFactory.build params
  end

  def collection
    @collection ||= \
      ConsolidationSearcher.
      search(Consolidation.includes(:category), date: date).
      decorate(context: { date: date })
  end
end
