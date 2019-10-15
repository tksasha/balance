# frozen_string_literal: true

class ItemSearcher < ApplicationSearcher
  def search_by_date_range(date_range)
    results.where date: date_range
  end

  def search_by_category(slug)
    return unless slug.present?

    results.where categories: { slug: slug }
  end

  private

  def date
    @date ||= DateFactory.build @params
  end

  def date_range
    @date_range ||= DateRange.month date
  end

  def params
    @params.tap { |params| params[:date_range] = date_range }
  end

  def results
    @results ||= @relation.includes(:category)
  end
end
