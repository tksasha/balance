# frozen_string_literal: true

class ItemSearcher < ApplicationSearcher
  include ActsAsSearchByCurrency

  def search_by_month(month)
    results.where date: month.dates
  end

  def search_by_category_id(category_id)
    return if category_id.blank?

    results.where(category_id:)
  end

  private

  def month
    @month ||= ParseMonthService.call(@params)
  end

  def params
    @params.tap { |params| params[:month] = month }
  end
end
