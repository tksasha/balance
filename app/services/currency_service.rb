# frozen_string_literal: true

class CurrencyService
  DEFAULT = CURRENCIES.first

  def initialize(currency)
    @currency = currency
  end

  def currency
    return DEFAULT unless @currency.present?

    currency = @currency.downcase

    return DEFAULT unless CURRENCIES.include? currency

    currency
  end

  class << self
    def currency(*args)
      new(*args).currency
    end
  end
end
