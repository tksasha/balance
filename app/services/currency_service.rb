# frozen_string_literal: true

class CurrencyService
  DEFAULT = CURRENCIES.first

  def initialize(currency)
    @currency = currency
  end

  def call
    currency
  end

  private

  def currency
    return DEFAULT if @currency.blank?

    currency = @currency.downcase

    return DEFAULT unless CURRENCIES.include? currency

    currency
  end

  class << self
    def call(*args)
      new(*args).call
    end
  end
end
