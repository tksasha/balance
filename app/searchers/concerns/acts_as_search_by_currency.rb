# frozen_string_literal: true

module ActsAsSearchByCurrency
  extend ActiveSupport::Concern

  def search_by_currency(currency)
    return unless currency.present?

    results.where(currency: currency)
  end
end
