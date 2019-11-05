# frozen_string_literal: true

class CashSearcher < ApplicationSearcher
  def search_by_currency(currency)
    return unless currency.present?

    results.where(currency: currency)
  end
end
