# frozen_string_literal: true

class CashSearcher < ApplicationSearcher
  include ActsAsSearchByCurrency

  def search_by_supercategory(supercategory)
    return if supercategory.blank?

    results.where(supercategory:)
  end
end
