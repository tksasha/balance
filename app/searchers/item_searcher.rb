# frozen_string_literal: true

class ItemSearcher < ApplicationSearcher
  include ActsAsSearchByCurrency

  def search_by_month(month)
    month = Month.parse(month)

    results.where date: month.dates
  end

  def search_by_category_id(category_id)
    return if category_id.blank?

    results.where(category_id:)
  end

  private

  def relation
    return @relation if params[:month].present?

    date = Month.now.dates

    @relation.where(date:)
  end
end
