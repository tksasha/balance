# frozen_string_literal: true

class ItemSearcher < ApplicationSearcher
  include ActsAsSearchByCurrency

  def search_by_month(month)
    dates = Month.parse(month).dates

    results.where date: dates
  rescue ArgumentError
    nil
  end

  def search_by_category(slug)
    return if slug.blank?

    results.joins(:category).where(categories: { slug: slug })
  end
end
