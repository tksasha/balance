# frozen_string_literal: true

module HasCurrencyScopes
  class << self
    def included(base)
      CURRENCIES.each_key do |currency|
        default = currency == CURRENCIES.keys.first

        base.send(:scope, currency.upcase, default:, show_count: false) do |scope|
          scope.public_send(currency)
        end
      end
    end
  end
end
