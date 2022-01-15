# frozen_string_literal: true

module ActsAsSearchByCurrency
  extend ActiveSupport::Concern

  def search_by_currency(currency)
    return if currency.blank?

    results.where(currency:)
  end
end
