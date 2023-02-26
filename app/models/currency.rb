# frozen_string_literal: true

class Currency
  DEFAULT = CURRENCIES.keys.first

  class << self
    def parse(currency)
      return DEFAULT if currency.blank?

      currency = currency.downcase

      return DEFAULT unless CURRENCIES.key?(currency)

      currency
    end
  end
end
