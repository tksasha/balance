# frozen_string_literal: true

module Frontend
  class Dashboard
    def initialize(params = {})
      @currency = Currency.parse(params[:currency])

      @date = DateRange.parse(params)
    end

    def items
      scope
        .where(date:)
        .order(date: :desc)
    end

    def at_end
      income - expense
    end

    def balance
      cashes.sum(:sum) - at_end
    end

    def item
      Item.new(currency:)
    end

    private

    attr_reader :currency, :date

    def scope
      Item.where(currency:)
    end

    def cashes
      Cash.where(currency:)
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
