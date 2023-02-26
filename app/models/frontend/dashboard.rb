# frozen_string_literal: true

module Frontend
  class Dashboard
    def initialize(params = {})
      self.currency = params[:currency]

      self.month = params[:month]
    end

    def cashes
      CashSearcher.search(Cash.all, currency:)
    end

    def at_end
      income - expense
    end

    def balance
      cashes_sum - at_end
    end

    def item
      Item.new(currency:)
    end

    def consolidations
      Frontend::Reports::Consolidations.call(currency:, month:)
    end

    private

    attr_reader :currency, :month

    def currency=(currency)
      @currency = Currency.parse(currency)
    end

    def month=(month)
      @month = Month.parse(month)
    end

    def scope
      Item.where(currency:)
    end

    def cashes_sum
      cashes.sum(:sum)
    end

    def income
      scope
        .income
        .sum(:sum)
    end

    def expense
      scope
        .expense
        .sum(:sum)
    end
  end
end
