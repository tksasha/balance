# frozen_string_literal: true

class Currency
  DEFAULT = CURRENCIES.first

  class << self
    def parse(currency)
      return DEFAULT if currency.blank?

      currency = currency.downcase

      return DEFAULT unless CURRENCIES.include? currency

      currency
    end
  end
end

def Currency(*args)
  Currency.parse(*args)
end
