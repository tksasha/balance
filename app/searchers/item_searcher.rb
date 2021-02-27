# frozen_string_literal: true

class ItemSearcher < ApplicationSearcher
  include ActsAsSearchByCurrency

  def search_by_month(month)
    results.where date: month.dates
  end

  def search_by_category(slug)
    return if slug.blank?

    results.joins(:category).where(categories: { slug: slug })
  end

  private

  def month
    @month ||= ParseMonthService.call(@params)
  end

  def params
    @params.tap { |params| params[:month] = month }
  end
end
